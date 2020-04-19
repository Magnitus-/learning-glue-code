virsh shutdown --domain controller;
sleep 3;
virsh snapshot-revert --domain controller --snapshotname controller-snapshot --running; 