package cairo

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"go.uber.org/zap"
	"io"
	"medical-zkml-backend/internal/db"
	"medical-zkml-backend/internal/module"
	"os"
	"strconv"
	"strings"
)

type Runner struct {
	cli       *client.Client
	imageName string
	task      []string
	taskCount uint32
	taskMax   uint32
	dir       string
}

func NewDockerClient(imageName ...string) *Runner {
	var r Runner
	var err error

	r.dir = "internal/plugin/circuit"

	r.cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("connect docker server failed: ", err)
		os.Exit(7)
	}

	if len(imageName) == 0 {
		r.imageName = "cairo-runner:v2.3.1"
	} else {
		r.imageName = imageName[0]
	}

	exist, err := r.imageCheck()
	if err != nil {
		fmt.Println("get docker images failed: ", err)
		os.Exit(7)
	}

	if !exist {
		fmt.Printf("In the Docker image list, to obtain image %s, start building the image\n", r.imageName)
		if err := r.BuildCairoImage(r.imageName); err != nil {
			os.Exit(7)
		}
	}

	// TODO: Calculate taskCount by obtaining local resources
	r.taskMax = 30
	r.taskCount = 0

	return &r
}

// ImageCheck Check if there is a specified image in the image list
func (r *Runner) imageCheck() (bool, error) {
	images, err := r.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}
	for _, image := range images {
		if image.RepoTags[0] == r.imageName {
			return true, nil
		}
	}

	return false, nil
}

func (r *Runner) BuildCairoImage(name string) error {
	ctx := context.Background()
	buildCtx, buildCtxCancel := context.WithCancel(ctx)
	defer buildCtxCancel()

	options := types.ImageBuildOptions{
		Dockerfile: "build/",
		Tags:       []string{name},
	}
	buildResponse, err := r.cli.ImageBuild(buildCtx, nil, options)
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()

	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(name + " image construction completed")
	return nil
}

func (r *Runner) RunCairo(ctx context.Context, name, path string) (int, error) {

	config := &container.Config{
		Image: r.imageName,
	}

	// TODO: Place the generation of the cairo project in the code and inject it into the container
	hostConfig := &container.HostConfig{
		Binds: []string{fmt.Sprintf("%s:/app", path)},
	}

	// Create Container
	resp, err := r.cli.ContainerCreate(ctx, config, hostConfig, nil, nil, name)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("container create failed: %s", err))
	}

	// Start Container
	if err := r.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return -1, errors.New(fmt.Sprintf("container start failed: %s", err))
	}

	// Waiting for container execution to complete
	statusCh, errCh := r.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return -1, errors.New(fmt.Sprintf("container exec failed: %s", err))
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			return -1, errors.New("container exec failed: Container abnormal exit")
		}
	}

	// Get container logs
	out, err := r.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return -1, errors.New(fmt.Sprintf("container get logs failed: %s", err))
	}
	defer out.Close()
	logs, err := io.ReadAll(out)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("container get logs failed: %s", err))
	}

	// TODO: Get the output in a different way, this method is unstable
	output := strings.Split(string(logs), "\n")
	line := output[len(output)-3]
	predicted, err := strconv.Atoi(string(line[len(line)-2]))
	if err != nil {
		return -1, errors.New(fmt.Sprintf("output conversion failed: %s", err))
	}

	// Delete Container
	_ = r.cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})

	return predicted, nil
}

// DiseasePrediction TODO: concurrency control
func (r *Runner) DiseasePrediction(recordPrediction *module.RecordPrediction, disease, model string, inputs []int) {
	newProject := strings.ToLower(fmt.Sprintf("%s_%s_%d", strings.ReplaceAll(model, " ", "_"), strings.ReplaceAll(disease, " ", "_"), recordPrediction.ID))
	path := rootPath + "/" + r.dir + "/" + newProject
	if err := r.createProject(path, disease, model, inputs); err != nil {
		recordPrediction.Status = module.CreateFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Failed to create Cairo project", err.Error())
		_ = db.RecordPredict(recordPrediction)
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s create failed: ", err))))
		return
	}

	predicted, err := r.RunCairo(context.Background(), newProject, path)
	if err != nil {
		recordPrediction.Status = module.PredictionFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Cairo circuit exec failed: ", err.Error())
		_ = db.RecordPredict(recordPrediction)
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed: ", err))))
	}

	_ = r.cleanProject(path)

	recordPrediction.Status = module.Complete
	recordPrediction.Output = predicted
	_ = db.RecordPredict(recordPrediction)
}

func (r *Runner) createProject(path string, disease, model string, inputs []int) error {

	if err := os.MkdirAll(path+"/src", 0755); err != nil {
		return errors.New(fmt.Sprintf("Failed to create task: %s", err.Error()))
	}

	if err := r.projectInject(path, disease, model, inputs); err != nil {
		return errors.New(fmt.Sprintf("Failed to inject: %s", err.Error()))
	}

	return nil
}

func (r *Runner) projectInject(path, disease, model string, data []int) error {

	// Main function injection
	content, err := os.ReadFile(fmt.Sprintf("%s/internal/plugin/abi/%s/%s.cairo", rootPath, strings.ReplaceAll(model, " ", "_"), strings.ReplaceAll(disease, " ", "_")))
	if err != nil {
		return err
	}

	if err := os.WriteFile(path+"/src/main.cairo", content, 0755); err != nil {
		return err
	}

	// Mod reference injection
	if err := os.WriteFile(path+"/src/lib.cairo", []byte("mod inputs;\nmod main;"), 0755); err != nil {
		return err
	}

	// Inputs function injection
	inputs := []string{
		"use array::{SpanTrait, ArrayTrait};",
		"use orion::operators::tensor::{TensorTrait, FP16x16Tensor, Tensor};",
		"use orion::numbers::{FixedTrait, FP16x16};",
		"fn input() -> Tensor<FP16x16> {",
		"TensorTrait::<FP16x16>::new(",
	}
	inputs = append(inputs, fmt.Sprintf("array![1, %d].span(),", len(data)))
	inputs = append(inputs, "array![")
	for _, ele := range data {
		fmt.Println("ele: ", ele)
		inputs = append(inputs, fmt.Sprintf("FixedTrait::<FP16x16>::new(%d, false),", ele))
	}
	inputs = append(inputs, "].span()")
	inputs = append(inputs, ")")
	inputs = append(inputs, "}")

	if err := os.WriteFile(path+"/src/inputs.cairo", []byte(strings.Join(inputs, "\n")), 0755); err != nil {
		return err
	}

	// Scarb.toml management file injection
	manifest := []string{
		"[package]",
		"name = \"contract\"",
		"version = \"0.1.0\"",
		"[dependencies]",
		"orion = { path = \"/root/giza/orion\" }",
	}

	if err := os.WriteFile(path+"/Scarb.toml", []byte(strings.Join(manifest, "\n")), 0755); err != nil {
		return err
	}

	return nil
}

func (r *Runner) cleanProject(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}
