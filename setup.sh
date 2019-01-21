#!/usr/bin/env sh
# -*- coding: utf-8 -*-

sudo sh -c "mkdir -p /opt/qemu/images/macos10.d"

if [[ ! -f /opt/qemu/images/macos10.qcow2 ]]; then
    sudo sh -c ". $HOME/.dev-openstack/search-openrc.sh && openstack image save --file /opt/qemu/images/macos10.qcow2 macos10"
fi

if [[ ! -f /opt/qemu/images/macos10.d/macos10kernel ]]; then
    sudo sh -c ". $HOME/.dev-openstack/search-openrc.sh && openstack image save --file /opt/qemu/images/macos10.d/macos10kernel macos10kernel"
fi
