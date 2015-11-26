package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

func exec_without_output(cmd string, args []string) (bool, error) {
  if err := exec.Command(cmd, args...).Run(); err != nil {
   return false, err
  }
  return true, nil
}

func ExecBash(command string, args []string) (int, string, error) {
	cmd := exec.Command(command, args...)
	var waitStatus syscall.WaitStatus

	//Stdout buffer
	cmdOutput := &bytes.Buffer{}

	//Attach buffer to command stdout
	cmd.Stdout = cmdOutput

	//Execute command
	err := cmd.Run()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Command:%s execute error\n", command)
		}
	}()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			return waitStatus.ExitStatus(), "", err
		} else {
			fmt.Println("something wrong")
		}
	}
	waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	return waitStatus.ExitStatus(), cmdOutput.String(), nil

}






func main() {
  cmd := "cat"
  args := []string{"/home/work/go/AUTHORS"}
  a, b := exec_without_output(cmd, args)
  if a {
   fmt.Println("ok")
  } else if b != nil {
   fmt.Println(b)
  }

}
