pkg_name=doha
pkg_origin=guskovd
pkg_version='1.0.12'
pkg_description="doha"
pkg_maintainer='guskovd'
pkg_upstream_url="https://github.com/guskovd/doha"

pkg_hab_shell_interpreter="bash"

RUBY_VERSION=2.5.1

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
    core/ruby/$RUBY_VERSION
    core/rsync
    core/sshpass
    core/gcc
    core/gcc-libs
    core/gox
    core/make
    guskovd/python-openstackclient
)

do_shell() {
    ruby_bundle_path=$HOME/.hab-shell/ruby/bundle/$RUBY_VERSION

    mkdir -p $ruby_bundle_path
    export BUNDLE_PATH=$ruby_bundle_path

    pushd "$( builtin cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" > /dev/null
    bundle install --binstubs 
    popd > /dev/null

    export PATH="$( builtin cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.hab-shell/bin:$PATH"
    export PATH="$( builtin cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/bin:$PATH"
    
    . ~/.bashrc
}

do_build() {
    return 0
}

do_install() {
    return 0
}

