#!/usr/bin/env sh
# -*- coding: utf-8 -*-

sudo sh -c "mkdir -p /opt/qemu/images/macos.d"
sudo sh -c ". $HOME/.dev-openstack/search-openrc.sh && openstack image save --file /opt/qemu/images/macos10.qcow2 macos10"
sudo sh -c ". $HOME/.dev-openstack/search-openrc.sh && openstack image save --file /opt/qemu/images/macos10.d/macos10kernel.qcow2 macos10kernel"
