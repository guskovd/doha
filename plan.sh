pkg_name=doha
pkg_origin=guskovd
pkg_version='1.0.12'
pkg_description="doha"
pkg_maintainer='guskovd'
pkg_upstream_url="https://github.com/guskovd/doha"

pkg_hab_shell_interpreter="bash"

RUBY_VERSION=2.5.1

pkg_deps=(
    core/bash/4.4.19/20180608092913
    core/coreutils/8.29/20180608092141
    core/dep/0.5.0/20180731174047
    core/docker-compose/1.21.0/20181212192853
    core/docker/18.03.0/20180608150948
    core/gawk/4.2.0/20180608093856
    core/gcc-libs/7.3.0/20180608091701
    core/gcc/7.3.0/20180608051919
    core/git/2.18.0/20181218161804
    core/go/1.11.3/20181214192430
    core/gox/master/20180608163536
    core/gzip/1.9/20180608100716
    core/hab/0.73.0/20190115004751
    core/make/4.2.1/20180608100733
    core/rsync/3.1.2/20180608145950
    core/ruby/2.5.1/20181212185250
    core/sshpass/1.06/20180608151129
    core/sudo/1.8.18p1/20181219210923
    core/tar/1.30/20180608093304
    core/wget/1.19.4/20181212185851
    guskovd/python-openstackclient/1.9.0/20190121115438
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

