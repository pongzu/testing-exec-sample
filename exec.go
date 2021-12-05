package testing_exec_sample

import (
	"os/exec"
)

func RunLs() (string, error) {
	cmd := exec.Command("ls")
	return runLs(cmd)
}

func runLs(cmd *exec.Cmd) (string, error) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
