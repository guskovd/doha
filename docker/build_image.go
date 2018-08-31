package docker

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"os/user"

	"github.com/jhoonb/archivex"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
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

	tarfile, err := ioutil.TempFile("", "dohaDockerfileTar")
	if err != nil {
		log.Fatal(err)
	}

	tar := new(archivex.TarFile)
	defer os.Remove(tarfile.Name())

	tar.Create(tarfile.Name())
	tar.AddAll(tmpdir, false)

	ctx := context.Background()
	dockerBuildContext, err := os.Open(tarfile.Name())

	cli, err := client.NewClient("unix:///var/run/docker.sock", "", nil, nil)

	current_user, err := user.Current()
	current_group, err := user.LookupGroupId(current_user.Gid)

	buildOptions := types.ImageBuildOptions{
		Tags: []string{"doha:local"},
		SuppressOutput: true, // need!
		BuildArgs: map[string]*string{
			"username": &current_user.Username,
			"userid": &current_user.Uid,
			"groupid": &current_user.Gid,
			"groupname": &current_group.Name,
		},
	}

	_, err1 := cli.ImageBuild(ctx, dockerBuildContext, buildOptions)

	if err1 != nil {
		log.Fatal(err1)
	}
}

