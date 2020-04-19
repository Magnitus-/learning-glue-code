virsh snapshot-delete --domain compute1 --snapshotname compute-snapshot 2>/dev/null || true;
virsh snapshot-create-as --name "compute-snapshot" \
--description "Compute node backup" \
--domain compute1;