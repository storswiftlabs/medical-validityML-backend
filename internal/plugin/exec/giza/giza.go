package giza

import (
	"bytes"
	"fmt"
	expect "github.com/google/goexpect"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

type GizaRunner struct {
	user   string
	passwd string
	email  string
}

func NewGizaRunner(viper *viper.Viper) GizaRunner {
	user := viper.GetString("giza.user")
	email := viper.GetString("giza.email")
	passwd := viper.GetString("giza.passwd")

	if user == "" || email == "" || passwd == "" {
		fmt.Println("Please enter the complete information of Giza")
		os.Exit(5)
	}

	return GizaRunner{user: user, email: email, passwd: passwd}
}

func (g *GizaRunner) CheckVersion() string {
	cmd := exec.Command("giza", "--version")
	var output bytes.Buffer
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		fmt.Println("The GIza program cannot be executed, please install it")
		os.Exit(3)
	}
	section := strings.Split(output.String(), " ")
	return section[len(section)-1]
}

func (g *GizaRunner) checkLogin() bool {
	cmd := exec.Command("giza", "users", "me")
	var output bytes.Buffer
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		fmt.Println("The GIza program cannot be executed, please install it")
		os.Exit(3)
	}

	return output.String() == fmt.Sprintf("{\n  \"username\": \"%s\",\n  \"email\": \"%s\",\n  \"is_active\": true\n}\n", g.user, g.email)
}

func (g *GizaRunner) UserLogin() error {
	e, _, err := expect.Spawn("giza users login --renew", -1)
	if err != nil {
		return err
	}
	defer e.Close()

	e.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: "Enter your username*:"},
		&expect.BSnd{S: g.user + "\n"},
		&expect.BExp{R: "Enter your password*:"},
		&expect.BSnd{S: g.passwd + "\n"},
	}, -1)

	return nil
}

func (g *GizaRunner) GenerateProof(path string, disease string) error {
	if !g.checkLogin() {
		if err := g.UserLogin(); err != nil {
			return err
		}
	}

	cmd := exec.Command("giza", "prove", "--size", "M", "--output-path", fmt.Sprintf("%s/%s.%s", path, disease, "cairo"), fmt.Sprintf("%s/%s.%s", path, disease, "proof"))
	return cmd.Run()
}
