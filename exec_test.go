package testing_exec_sample

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var lsOutPutTestString = "test_file\ntest_file2\n"

func helperCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]
	switch cmd {
	case "ls":
		fmt.Fprint(os.Stdout, lsOutPutTestString)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
}

func TestRunLs(t *testing.T) {
	cmd := helperCommand("ls")
	out, err := runLs(cmd)
	if err != nil {
		t.Error(err)
	}

	if out != lsOutPutTestString {
		t.Fatalf("Error: expected %s, got %s", lsOutPutTestString, out)
	}
}
