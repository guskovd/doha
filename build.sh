#!/usr/bin/env bash

set -e

HAB_VERSION=0.61.0-20180815171844
HAB_BINARY_PATH=artifacts/hab-$HAB_VERSION-x86_64-linux/hab
HAB_TGZ=hab-$HAB_VERSION-x86_64-linux.tar.gz

mkdir -p artifacts

err() {
    echo "Error: $*" >&2
    exit 1
}

if [[ $( id -u ) -eq 0 ]]; then
    err "Should be runned under user, not root"
fi


if [ ! -f artifacts/$HAB_TGZ ]; then
    wget https://dl.bintray.com/habitat/stable/linux/x86_64/$HAB_TGZ -O artifacts/$HAB_TGZ
fi

if [ ! -f $HAB_BINARY_PATH ]; then
    pushd artifacts > /dev/null
    tar -xzf $HAB_TGZ
    popd > /dev/null
fi

docker build \
    --build-arg HAB_BINARY_PATH=$HAB_BINARY_PATH \
    -t dguskov/doha:base .
