// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package datasyncimport (
	"github.com/aws/aws-sdk-go/service/datasync"
)func expandProtocol(l []interface{}) *datasync.FsxProtocol {
	if len(l) == 0 || l[0] == nil {
return nil
	}	m := l[0].(map[string]interface{})
	protocol := &datasync.FsxProtocol{}	if v, ok := m["nfs"].([]interface{}); ok {
otocol.NFS = expandNFS(v)
	}
	if v, ok := m["smb"].([]interface{}); ok {
otocol.SMB = expandSMB(v)
	}	return protocol
}func flattenProtocol(protocol *datasync.FsxProtocol) []interface{} {
	if protocol == nil {
turn []interface{}{}
	}	m := map[string]interface{}{}	if protocol.NFS != nil {
"nfs"] = flattenNFS(protocol.NFS)
	}
	if protocol.SMB != nil {
"smb"] = flattenSMB(protocol.SMB)
	}	return []interface{}{m}
}func expandNFS(l []interface{}) *datasync.FsxProtocolNfs {
	if len(l) == 0 || l[0] == nil {
turn nil
	}	m := l[0].(map[string]interface{})	protocol := &datasync.FsxProtocolNfs{
untOptions: expandNFSMountOptions(m["mount_options"].([]interface{})),
	}	return protocol
}func expandSMB(l []interface{}) *datasync.FsxProtocolSmb {
	if len(l) == 0 || l[0] == nil {
turn nil
	}	m := l[0].(map[string]interface{})	protocol := &datasync.FsxProtocolSmb{
untOptions: expandSMBMountOptions(m["mount_options"].([]interface{})),
	}	return protocol
}// todo: go another level down?
func flattenNFS(nfs *datasync.FsxProtocolNfs) []interface{} {
	if nfs == nil {
turn []interface{}{}
	}	m := map[string]interface{}{
ount_options": flattenNFSMountOptions(nfs.MountOptions),
	}	return []interface{}{m}
}func flattenSMB(smb *datasync.FsxProtocolSmb) []interface{} {
	if smb == nil {
turn []interface{}{}
	}	m := map[string]interface{}{
ount_options": flattenSMBMountOptions(smb.MountOptions),
	}	return []interface{}{m}
}
