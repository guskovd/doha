package docker

import (
	"syscall"
	"os"
	"os/exec"

	"fmt"
)

func ContainerExec() {
	docker_binary, lookErr := exec.LookPath("docker")
	if lookErr != nil {
		panic(lookErr)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(pwd)
	
	syscall.Exec(
		docker_binary,
		[]string{
			"docker", "exec",
			"-w", pwd,
			"-ti", "doha",
			"/bin/bash",
		},
		os.Environ())
}
