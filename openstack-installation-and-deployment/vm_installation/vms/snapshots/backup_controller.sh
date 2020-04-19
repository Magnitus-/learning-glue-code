virsh snapshot-delete --domain controller --snapshotname controller-snapshot 2>/dev/null || true;
virsh snapshot-create-as --name "controller-snapshot" \
--description "controller node backup" \
--domain controller;