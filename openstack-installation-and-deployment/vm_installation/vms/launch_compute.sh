virt-install \
--hvm \
--name=compute1 \
--ram=4096 \
--vcpus=2 \
--cpu host-passthrough,cache.mode=passthrough \
--disk size=10,path=$(pwd)/compute1.qcow2,bus=virtio \
--disk path=./ubuntu/ubuntu-18.04.4-live-server-amd64.iso,device=cdrom \
--os-type linux \
--os-variant ubuntu18.04 \
--arch x86_64 \
--virt-type=kvm \
--graphics vnc \
--network network=openstack-provider \
--network network=openstack-management \
--network network=openstack-internet