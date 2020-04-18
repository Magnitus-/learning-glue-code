# About

At the beginning of the class, Kris installs openstack inside 3 vms (controller, compute1, storage1).

He uses Virtualbox. I translated the infrastructure setup in scripts using libvirt and kvm.

# Usage

1. Get the Ubuntu server amd64 18.04.4 live iso and put it in an **ubuntu** directory under the **vms** directory

2. go to **network** and run:

```
./setup.sh
```

3. go to **vms** and run the following scripts:

- launch_controller.sh (when you get the instruction to eject the media, run the eject_controller_installer.sh script)
- launch_compute.sh (when you get the instruction to eject the media, run the eject_compute_installer.sh script)
- launch_storage.sh (when you get the instruction to eject the media, run the eject_storage_installer.sh script)