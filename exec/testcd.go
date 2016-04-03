
package main

import "syscall"
import "os"
import (
	"os/exec"
	"fmt"
)

func main() {

	binary, lookErr := exec.LookPath("cd")
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()
	home := os.Getenv("HOME")

	fmt.Println("path=", os.Getenv("PWD"))
	args := []string{"cd", home + "/Gits"}

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}