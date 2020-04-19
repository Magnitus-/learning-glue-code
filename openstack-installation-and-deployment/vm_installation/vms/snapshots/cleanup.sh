virsh snapshot-delete --domain controller --snapshotname controller-snapshot 2>/dev/null || true;
virsh snapshot-delete --domain compute1 --snapshotname compute-snapshot 2>/dev/null || true;
virsh snapshot-delete --domain storage1 --snapshotname storage-snapshot 2>/dev/null || true;