package utils

import (
	stderrors "errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)


func SimpleCmdExec(cmdName string, cmdArgs []string, workingDir string, environ []string, silent bool) (string, error) {

	if !silent {
		fmt.Printf("DEBUG: EXTERNAL COMMAND: %v %v\n", cmdName, strings.Join( cmdArgs, " "))
	}
	cmd := exec.Command(cmdName, cmdArgs...)
	//cmd.Stdout = logStreamerOut
	//cmd.Stderr = logStreamerErr
	if environ != nil {
		cmd.Env = environ
	}
	if workingDir != "" && path.IsAbs(workingDir) {
		cmd.Dir = workingDir
	} else if workingDir != "" {
		return "", stderrors.New("Working Directory must be an absolute path")
	}


	stdoutBytes, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error running Cmd", err)
		if eerr, ok := err.(*exec.ExitError); ok && !silent {
			fmt.Printf("DEBUG: STDERR: %v\n", string(eerr.Stderr))
		}

		return "", err
	}

	stdoutStr := strings.TrimSpace(string(stdoutBytes))
	if !silent {
		fmt.Printf("DEBUG: STDOUT: %v\n", stdoutStr)
	}
	return stdoutStr, nil
}