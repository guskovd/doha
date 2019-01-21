pkg_name=doha
pkg_origin=guskovd
pkg_version='1.0.12'
pkg_description="doha"
pkg_maintainer='guskovd'
pkg_upstream_url="https://github.com/guskovd/doha"

pkg_hab_shell_interpreter="bash"

pkg_deps=(
    core/bash
    core/coreutils
    core/gawk
    core/hab
    core/sudo
    core/docker
    core/docker-compose
    core/go
    core/dep
    core/git
    core/wget
    core/tar
    core/gzip
    guskovd/python-openstackclient
)

do_shell() {
    . ~/.bashrc
}

do_build() {
    return 0
}

do_install() {
    return 0
}

