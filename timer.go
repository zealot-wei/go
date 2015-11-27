package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func ExecTimeout(command string, args []string, timeout int) (string, error, bool) {
	// Create *Cmd
	cmd := exec.Command(command, args...)
	cmdstring := strings.Join(cmd.Args, " ")

	// Stdout buffer
	buffer := &bytes.Buffer{}
	result := make(chan error)

	// Attach buffer to command stdout
	cmd.Stdout = buffer

	// Start command
	if err := cmd.Start(); err != nil {
		fmt.Printf("Command: %s error when start\n", cmdstring)
		return "", errors.New("Start command errors"), false
	}

	go func() {
		result <- cmd.Wait()
	}()

	// Create a timer than will kill the process
	select {
	case <-time.After(time.Second * time.Duration(timeout)):
		fmt.Printf("timeout, process: %s will be killed\n", cmdstring)

		go func() {
			t := <-result
			fmt.Println("timeout return: ", t)
		}()

		// timeout, kill the process
		if err := cmd.Process.Kill(); err != nil {
			fmt.Printf("failed to kill: %s, error: %s", cmdstring, err)
		}
		return "", errors.New("process timeout"), true
	case <-result:
		return buffer.String(), nil, false
	}
}

func main() {
	out, err, t := ExecTimeout("ls", []string{"/dev/random"}, 2)
	fmt.Printf("%V %T %s\n", out, out, out)
	fmt.Printf("%V %T %s\n", err, err, err)
	fmt.Println(t)

}
