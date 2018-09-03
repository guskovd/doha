package docker

import (
	"syscall"
	"os"
	"os/exec"
)

func ContainerExec(args []string) {
	docker_binary, lookErr := exec.LookPath("docker")
	if lookErr != nil {
		panic(lookErr)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	syscall.Exec(
		docker_binary,
		append([]string{
			"docker", "exec",
			"-w", pwd,
			"-ti", "doha",
		}, args...),
		os.Environ(),
	)
}
