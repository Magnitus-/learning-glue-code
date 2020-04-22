# About

Corrections to the course material due to newer versions of Ubuntu (18.04), Openstack (Train) or other dependencies.

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

# Openstack Client for Python 3

In lecture 11, there are instructions to install the openstack client.

With Ubuntu 18.04 and Train, the Python 3 implementation of the client should be installed:

```
apt install python3-openstackclient
```

Link: https://docs.openstack.org/install-guide/environment-packages-ubuntu.html#finalize-the-installation

# Getting Etcd Binaries

In the lecture 15, Kris gets a binary because the version in the Ubuntu repos is dated.

With Ubuntu 18.04, the version in the repo is the same as the only getting installed.

As such, you can skip the user create and the creation of the service file by running the following command:

```
apt install etcd
```

# Keystone Bootstrapping Adjustment


In the lecture 17, a separate port is given for the admin url in the keystone bootstrap command.

This is no longer required now that version 2 of the API (which required a separate port for the admin) is deprecated:

```
keystone-manage bootstrap --bootstrap-password openstack \
  --bootstrap-admin-url http://controller:5000/v3/ \
  --bootstrap-internal-url http://controller:5000/v3/ \
  --bootstrap-public-url http://controller:5000/v3/ \
  --bootstrap-region-id RegionOne
```

Also note that because Train runs on Python 3, the Python 3 version of mod wsgi should be installed:

```
apt install keystone apache2 libapache2-mod-wsgi-py3 crudini -y
```

Link: https://docs.openstack.org/keystone/train/install/keystone-install-ubuntu.html

# Keystone Environment Variables For Client

In the lecture 18, the port for **OS_AUTH_URL** should be 5000.

Link: https://docs.openstack.org/keystone/train/install/keystone-install-ubuntu.html

# Placement Service

Starting with the Stein release, an additional **placement** service needs to be deployed: https://docs.openstack.org/install-guide/openstack-services.html#minimal-deployment-for-train

Should be done between lectures 22 and 23.

Instructions can be found here: https://docs.openstack.org/placement/train/install/install-ubuntu.html

