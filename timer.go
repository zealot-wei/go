package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ExecTimeout(command string, args []string, timeout int) (int, string, error) {
	// Create *Cmd
	cmd := exec.Command(command, args...)
	cmdstring := strings.Join(cmd.Args, " ")

	// Stdout buffer
	buffer := &bytes.Buffer{}

	// Attach buffer to command stdout
	cmd.Stdout = buffer

	// Start command
	if err := cmd.Start(); err != nil {
		fmt.Println("Command: %s error when start", cmdstring)
		return 911, "", err
	}

	// Create a timer than will kill the process
	timer := time.NewTimer(time.Second * time.Duration(timeout))
	go func(timer *time.Timer, cmd *exec.Cmd) {
		for _ = range timer.C {
			err := cmd.Process.Signal(os.Kill)
			if err != nil {
				fmt.Println("process tiemout\n")
				fmt.Printf("%V %T %s\n", err, err, err)
			} else {
				fmt.Println("ok")
			}
		}
	}(timer, cmd)

	// Jobs after process has finished
	cmd.Wait()
	fmt.Printf("%d bytes generated", len(buffer.Bytes()))

	return 0, "", nil

}

func main() {
	retcode, out, err := ExecTimeout("cat", []string{"/dev/random"}, 4)
	fmt.Printf("%V %T %d\n", retcode, retcode, retcode)
	fmt.Printf("%V %T %s\n", out, out, out)
	fmt.Printf("%V %T %s\n", err, err, err)

}
