package cairo

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/db"
	m "medical-zkml-backend/internal/module"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var rootPath string
var msg = "Cairo disease prediction"

func init() {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(path.Dir(path.Dir(filename)))))

}

type CairoRunner struct {
	root  string   // Command source file path
	dir   string   // Project Root Path
	tasks []string // task list
}

func InitCairoRunner(path, dir string) *CairoRunner {
	checkCairoRunner(path)
	PathCheck(dir)
	return &CairoRunner{
		root: path,
		dir:  dir,
	}
}

func checkCairoRunner(root string) {
	cmd := exec.Command(root, "--version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("The execution of the scarb cairo package management tool failed. Please check the path or install scarb :", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func PathCheck(path string) {
	circuitPath := fmt.Sprintf("%s/%s", rootPath, path)
	_, err := os.Stat(circuitPath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(circuitPath, os.ModePerm); err != nil {
			fmt.Printf("Could not create circuit: %s", err)
			os.Exit(2)
		}
	}
}

func (r *CairoRunner) prediction(path, disease, module string, quantized []int) (int, error) {

	if err := r.projectInject(path, disease, module, quantized); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", path))))
		return -1, err
	}

	// Project execution
	cmd := exec.Command(r.root, "cairo-run", "--available-gas", "999999999")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Dir = fmt.Sprintf("%s/internal/plugin/circuit/%s", rootPath, path)
	if err := cmd.Run(); err != nil {
		zap.L().Error(msg, zap.Error(errors.New(output.String())))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", path))))
		return -1, errors.New(stderr.String())
	}
	fmt.Println(output.String())
	match := regexp.MustCompile(`returning \[(\d+)\]`).FindStringSubmatch(output.String())

	result, _ := strconv.Atoi(match[1])
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s execution prediction", path), "success"))
	return result, nil
}

func (r *CairoRunner) DiseasePrediction(recordPrediction *m.RecordPrediction, disease, module string, inputs []int) {
	project, err := r.createProject(recordPrediction.ID, disease, module)
	if err != nil {
		// Update the status of the database
		recordPrediction.Status = m.CreateFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Failed to create Cairo project", err.Error())
		_ = db.RecordPredict(recordPrediction)
		return
	}
	// Successfully created, updating database status
	recordPrediction.Status = m.CreateSuccess
	_ = db.RecordPredict(recordPrediction)

	result, err := r.prediction(project, disease, module, inputs)
	if err != nil {
		// Update database status
		recordPrediction.Status = m.PredictionFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Cairo circuit failed to perform prediction", err.Error())
		_ = db.RecordPredict(recordPrediction)
		return
	}
	// Execution completed, update database status
	recordPrediction.Status = m.PredictionComplete
	_ = db.RecordPredict(recordPrediction)

	// end
	recordPrediction.Status = m.Complete
	recordPrediction.Output = result
	_ = db.RecordPredict(recordPrediction)

	//_ = r.cleanProject(project)

	return
}

func (r *CairoRunner) cleanProject(project string) error {
	if err := os.RemoveAll(fmt.Sprintf("%s/internal/plugin/circuit/%s", rootPath, project)); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s clean project", project), "success"))
	return nil
}

func (r *CairoRunner) createProject(id uint, disease, module string) (string, error) {
	zap.L().Info(msg, zap.String("Status", "start"))

	// scarb new module_id
	// scarb cannot create project names in uppercase letters, characters need to be converted to lowercase
	newProject := strings.ToLower(fmt.Sprintf("%s_%s_%d", strings.ReplaceAll(module, " ", "_"), strings.ReplaceAll(disease, " ", "_"), id))
	fmt.Printf("%s new %s/%s\n", r.root, r.dir, newProject)
	cmd := exec.Command(r.root, "new", newProject)
	cmd.Dir = r.dir
	if err := cmd.Run(); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", newProject))))
		return "", errors.New(fmt.Sprintf("Failed to create task: %s", err.Error()))
	}
	zap.L().Info(msg, zap.String("create newProject success", newProject))
	return newProject, nil
}

func (r *CairoRunner) projectInject(project, disease, module string, data []int) error {

	// Main function injection
	content, err := os.ReadFile(fmt.Sprintf("%s/internal/plugin/abi/%s/%s.cairo", rootPath, strings.ReplaceAll(module, " ", "_"), strings.ReplaceAll(disease, " ", "_")))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}

	if err := os.WriteFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/src/main.cairo", rootPath, project), content, 0644); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}

	// Mod reference injection
	if err := os.WriteFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/src/lib.cairo", rootPath, project), []byte("mod inputs;\nmod main;"), 0644); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s circuit injection", project), "success"))

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

	if err := os.WriteFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/src/inputs.cairo", rootPath, project), []byte(strings.Join(inputs, "\n")), 0755); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}

	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s inputs injection", project), "success"))

	// Scarb.toml management file injection
	manifest := []string{
		"[package]",
		"name = \"contract\"",
		"version = \"0.1.0\"",
		"[dependencies]",
		"orion = { git = \"https://github.com/gizatechxyz/orion.git\", rev = \"d392123\" }",
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/Scarb.toml", rootPath, project), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(strings.Join(manifest, "\n")); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}

	return nil
}
