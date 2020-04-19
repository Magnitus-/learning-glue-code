virsh snapshot-delete --domain storage1 --snapshotname storage-snapshot 2>/dev/null || true;
virsh snapshot-create-as --name "storage-snapshot" \
--description "storage node backup" \
--domain storage1;