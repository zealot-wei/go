package main

import (
  "fmt"
  "os/exec"
)

func exec_without_output(cmd string, args []string) (bool, error) {
  if err := exec.Command(cmd, args...).Run(); err != nil {
   return false, err
  }
  return true, nil
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
