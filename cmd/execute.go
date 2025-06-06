package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(command []string) (bool, error) {
	startGo := exec.Command(command[0], command[1:]...)
	startGo.Stdout = os.Stdout
	startGo.Stderr = os.Stderr
	err := startGo.Run()
	if err != nil {
		fmt.Println("Error executing command ", command, err)
		return false, err
	}
	return true, nil
}
func ExecuteShell(script string, command []string) (bool, error) {

	cmd := &exec.Cmd{
		Path:   script,
		Args:   command,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Println("Executing command ", cmd)

	err := cmd.Start()
	if err != nil {
		return false, err
	}

	err = cmd.Wait()
	if err != nil {
		return false, err
	}

	return true, nil
}
