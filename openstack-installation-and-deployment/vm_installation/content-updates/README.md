# About

Corrections to the course material due to newer versions of Ubuntu (18.04), Openstack (Train) or other dependencies.

# Disk resizing

For some reason, the Ubuntu default installation didn't take advantage of the full disk size with logical volumes.

You'll notice the problem if you get an output like this:

```
lsblk 
NAME                      MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
loop0                       7:0    0 93.8M  1 loop /snap/core/8935
loop1                       7:1    0 93.9M  1 loop /snap/core/9066
sr0                        11:0    1 1024M  0 rom  
vda                       252:0    0   20G  0 disk 
├─vda1                    252:1    0    1M  0 part 
├─vda2                    252:2    0    1G  0 part /boot
└─vda3                    252:3    0   19G  0 part 
  └─ubuntu--vg-ubuntu--lv 253:0    0    4G  0 lvm  /
```

In the above case, the **ubuntu--vg-ubuntu--lv** logical volume only occupies 4GB of a 19GB partition.

You can correct the above by typing the following to extend the logical volume:

```
lvextend -l +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv
```

To make the filesystem report the new size, first find out the kind of filesystem you have by typing:

```
df -hT
```

If you have ext4, type:

```
resize2fs /dev/mapper/ubuntu--vg-ubuntu--lv
```

If you have xfs, type:

```
xfs_growfs /
```

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

# Glance Configuration

Use the instructions in the link for the configuration (don't forget to replace the passwords by openstack).

The configurations are a bit different and apparently, the glance registry has been deprecated since Queens.

Link: https://docs.openstack.org/glance/train/install/install-ubuntu.html

# Placement Service

Starting with the Stein release, an additional **placement** service needs to be deployed: https://docs.openstack.org/install-guide/openstack-services.html#minimal-deployment-for-train

Should be done between lectures 22 and 23.

Instructions can be found here: https://docs.openstack.org/placement/train/install/install-ubuntu.html

