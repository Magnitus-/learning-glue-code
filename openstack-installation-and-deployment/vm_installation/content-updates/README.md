# About

Corrections to the course material due to newer versions of Ubuntu, Openstack or other dependencies.

# Network Setup

In the lecture 9, after disabling the Ubuntu default setup for the network interfaces, Kris edits the **/etc/network/interfaces** file.

This worked out of the box for **Ubuntu 16.04** and earlier, but for **Ubuntu 18.04** and onwards, custom network setup on machines is done using netplan configurations:

1. On all 3 machines, create the **/etc/cloud/cloud.cfg.d/99-disable-network-config.cfg** file with the following content:

```
network: {config: disabled}
```

2. On each host, delete the **/etc/netplan/50-cloud-init.yaml** file and create the **/etc/netplan/01-network-card.yaml** file on each host, with content relative to each host as shown below:

controller:
```
network:
    version: 2
    renderer: networkd
    ethernets:
        eth0:
            dhcp4: false
            addresses: [10.0.0.11/24]
            nameservers:
                addresses: [10.0.0.1]    
        eth1:
            dhcp4: false
        eth2:
            dhcp4: true
```

compute1:
```
network:
    version: 2
    renderer: networkd
    ethernets:
        eth0:
            dhcp4: false
            addresses: [10.0.0.31/24]
            nameservers:
                addresses: [10.0.0.1]    
        eth1:
            dhcp4: false
        eth2:
            dhcp4: true
```

storage1:
```
network:
    version: 2
    renderer: networkd
    ethernets:
        eth0:
            dhcp4: false
            addresses: [10.0.0.41/24]
            nameservers:
                addresses: [10.0.0.1]    
        eth1:
            dhcp4: false
        eth2:
            dhcp4: true
```

3. Reboot each machine as shown in the lecture and proceed with the course