package docker

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"os/user"
	"os/exec"

)

var dockerFileContent = []byte(`
FROM docker.io/dguskov/doha:base

ARG username=guest
ARG userid=1000
ARG groupname=guest
ARG groupid=1000

RUN addgroup -g $groupid -S $groupname 
RUN adduser -D -g $groupid -G wheel -u $userid -H $username
`)

// BuildImage common
func BuildImage() {
	tmpdir, err := ioutil.TempDir("", "dohaDockerfile")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	fo, err := os.Create(fmt.Sprintf("%s/Dockerfile", tmpdir))
	fo.Write(dockerFileContent)
	fo.Close()

	current_user, err := user.Current()
	current_group, err := user.LookupGroupId(current_user.Gid)

	docker_binary, lookErr := exec.LookPath("docker")
	
	if lookErr != nil {
		panic(lookErr)
	}

	out, execErr := exec.Command(
		docker_binary,
		"build",
		"--build-arg", fmt.Sprintf("username=%s", current_user.Username),
		"--build-arg", fmt.Sprintf("userid=%s", current_user.Uid),
		"--build-arg", fmt.Sprintf("groupname=%s", current_group.Name),
		// "--build-arg", fmt.Sprintf("groupid=%s", current_group.Gid),
		"-t", DohaImageLocal,
		tmpdir,
	).CombinedOutput()

	if execErr != nil {
		log.Fatal(string(out))
		log.Fatal(lookErr)
	}
}

