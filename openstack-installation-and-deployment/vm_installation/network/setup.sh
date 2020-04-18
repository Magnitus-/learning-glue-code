#!/usr/bin/env bash

virsh net-define --file provider.xml;
virsh net-autostart openstack-provider;
virsh net-start openstack-provider;

virsh net-define --file management.xml;
virsh net-autostart openstack-management;
virsh net-start openstack-management;

virsh net-define --file internet.xml;
virsh net-autostart openstack-internet;
virsh net-start openstack-internet;