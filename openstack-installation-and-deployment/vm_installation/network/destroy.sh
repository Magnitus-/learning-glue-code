#!/usr/bin/env bash

virsh net-destroy openstack-provider;
virsh net-undefine openstack-provider;

virsh net-destroy openstack-management;
virsh net-undefine openstack-management;

virsh net-destroy openstack-internet;
virsh net-undefine openstack-internet;