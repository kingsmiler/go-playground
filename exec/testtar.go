package main

import "syscall"
import "os"
import (
	"os/exec"
	"fmt"
)

func main() {

	gits := os.Getenv("GITS_HOME")
	err := os.Chdir(gits)
	checkError(err)

	pwd, err := os.Getwd()
	checkError(err)

	fmt.Println(pwd)

	tarTarget()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func tarTarget() {
	binary, lookErr := exec.LookPath("tar")
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()

	args := []string{"tar", "czf", "code.tar.gz", "docker-client"}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

