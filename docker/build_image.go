package docker

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
)

var dockerFileContent = []byte(`
FROM docker.io/dguskov/doha:base

ARG username=guest
ARG userid=1000
ARG groupname=guest
ARG groupid=1000
ARG homedir=/home/guest

ARG docker_groupid=995

RUN addgroup -g $groupid -S $groupname 
RUN adduser -D -g $groupid -G wheel -u $userid -h $homedir $username

RUN addgroup -g $docker_groupid -S docker
RUN adduser $username docker

CMD sudo chmod 666 /var/run/docker.sock && tail -f /dev/null
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

	// docker_group, err := user.LookupGroup("docker")
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
		"--build-arg", fmt.Sprintf("homedir=%s", os.Getenv("HOME")),
		// "--build-arg", fmt.Sprintf("groupid=%s", current_group.Gid),
		// "--build-arg", fmt.Sprintf("docker_groupid=%s", docker_group.Gid),
		"-t", DohaImageLocal,
		tmpdir,
	).CombinedOutput()

	if execErr != nil {
		log.Fatal(string(out))
		log.Fatal(lookErr)
	}
}
