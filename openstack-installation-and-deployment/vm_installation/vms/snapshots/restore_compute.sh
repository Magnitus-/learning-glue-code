virsh shutdown --domain compute1;
sleep 3;
virsh snapshot-revert --domain compute1 --snapshotname compute-snapshot --running; 