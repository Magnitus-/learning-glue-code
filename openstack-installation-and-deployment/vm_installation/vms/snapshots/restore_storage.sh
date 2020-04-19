virsh shutdown --domain storage1;
sleep 3;
virsh snapshot-revert --domain storage1 --snapshotname storage-snapshot --running; 