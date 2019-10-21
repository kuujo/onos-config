// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/types/snapshot/device/types.proto

package device

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	github_com_onosproject_onos_config_pkg_types "github.com/onosproject/onos-config/pkg/types"
	device "github.com/onosproject/onos-config/pkg/types/change/device"
	github_com_onosproject_onos_config_pkg_types_change_device "github.com/onosproject/onos-config/pkg/types/change/device"
	snapshot "github.com/onosproject/onos-config/pkg/types/snapshot"
	github_com_onosproject_onos_topo_pkg_northbound_device "github.com/onosproject/onos-topo/pkg/northbound/device"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// DeviceSnapshot is a device snapshot
type DeviceSnapshot struct {
	// 'id' is the unique snapshot identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// 'device_id' is the device to which the snapshot applies
	DeviceID github_com_onosproject_onos_topo_pkg_northbound_device.ID `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3,casttype=github.com/onosproject/onos-topo/pkg/northbound/device.ID" json:"device_id,omitempty"`
	// 'revision' is the request revision number
	Revision Revision `protobuf:"varint,3,opt,name=revision,proto3,casttype=Revision" json:"revision,omitempty"`
	// 'network_snapshot' is a reference to the network snapshot from which this snapshot was created
	NetworkSnapshot NetworkSnapshotRef `protobuf:"bytes,4,opt,name=network_snapshot,json=networkSnapshot,proto3" json:"network_snapshot"`
	// 'max_network_change_index' is the maximum network change index to be snapshotted for the device
	MaxNetworkChangeIndex github_com_onosproject_onos_config_pkg_types.Index `protobuf:"varint,5,opt,name=max_network_change_index,json=maxNetworkChangeIndex,proto3,casttype=github.com/onosproject/onos-config/pkg/types.Index" json:"max_network_change_index,omitempty"`
	// 'status' is the snapshot status
	Status snapshot.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	// 'created' is the time at which the configuration was created
	Created time.Time `protobuf:"bytes,7,opt,name=created,proto3,stdtime" json:"created"`
	// 'updated' is the time at which the configuration was last updated
	Updated time.Time `protobuf:"bytes,8,opt,name=updated,proto3,stdtime" json:"updated"`
}

func (m *DeviceSnapshot) Reset()         { *m = DeviceSnapshot{} }
func (m *DeviceSnapshot) String() string { return proto.CompactTextString(m) }
func (*DeviceSnapshot) ProtoMessage()    {}
func (*DeviceSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f618bedd50a8adb9, []int{0}
}
func (m *DeviceSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSnapshot.Merge(m, src)
}
func (m *DeviceSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *DeviceSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSnapshot proto.InternalMessageInfo

func (m *DeviceSnapshot) GetID() ID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DeviceSnapshot) GetDeviceID() github_com_onosproject_onos_topo_pkg_northbound_device.ID {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DeviceSnapshot) GetRevision() Revision {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *DeviceSnapshot) GetNetworkSnapshot() NetworkSnapshotRef {
	if m != nil {
		return m.NetworkSnapshot
	}
	return NetworkSnapshotRef{}
}

func (m *DeviceSnapshot) GetMaxNetworkChangeIndex() github_com_onosproject_onos_config_pkg_types.Index {
	if m != nil {
		return m.MaxNetworkChangeIndex
	}
	return 0
}

func (m *DeviceSnapshot) GetStatus() snapshot.Status {
	if m != nil {
		return m.Status
	}
	return snapshot.Status{}
}

func (m *DeviceSnapshot) GetCreated() time.Time {
	if m != nil {
		return m.Created
	}
	return time.Time{}
}

func (m *DeviceSnapshot) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

// NetworkSnapshotRef is a back reference to the NetworkSnapshot that created a DeviceSnapshot
type NetworkSnapshotRef struct {
	// 'id' is the identifier of the network snapshot from which this snapshot was created
	ID github_com_onosproject_onos_config_pkg_types.ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=github.com/onosproject/onos-config/pkg/types.ID" json:"id,omitempty"`
	// 'index' is the index of the network snapshot from which this snapshot was created
	Index github_com_onosproject_onos_config_pkg_types.Index `protobuf:"varint,2,opt,name=index,proto3,casttype=github.com/onosproject/onos-config/pkg/types.Index" json:"index,omitempty"`
}

func (m *NetworkSnapshotRef) Reset()         { *m = NetworkSnapshotRef{} }
func (m *NetworkSnapshotRef) String() string { return proto.CompactTextString(m) }
func (*NetworkSnapshotRef) ProtoMessage()    {}
func (*NetworkSnapshotRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_f618bedd50a8adb9, []int{1}
}
func (m *NetworkSnapshotRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkSnapshotRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkSnapshotRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkSnapshotRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkSnapshotRef.Merge(m, src)
}
func (m *NetworkSnapshotRef) XXX_Size() int {
	return m.Size()
}
func (m *NetworkSnapshotRef) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkSnapshotRef.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkSnapshotRef proto.InternalMessageInfo

func (m *NetworkSnapshotRef) GetID() github_com_onosproject_onos_config_pkg_types.ID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NetworkSnapshotRef) GetIndex() github_com_onosproject_onos_config_pkg_types.Index {
	if m != nil {
		return m.Index
	}
	return 0
}

// Snapshot is a snapshot of the state of a single device
type Snapshot struct {
	// 'id' is a unique snapshot identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// 'device_id' is the device to which the snapshot applies
	DeviceID github_com_onosproject_onos_topo_pkg_northbound_device.ID `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3,casttype=github.com/onosproject/onos-topo/pkg/northbound/device.ID" json:"device_id,omitempty"`
	// 'snapshot_id' is the ID of the snapshot
	SnapshotID ID `protobuf:"bytes,3,opt,name=snapshot_id,json=snapshotId,proto3,casttype=ID" json:"snapshot_id,omitempty"`
	// 'change_index' is the change index at which the snapshot ended
	ChangeIndex github_com_onosproject_onos_config_pkg_types_change_device.Index `protobuf:"varint,4,opt,name=change_index,json=changeIndex,proto3,casttype=github.com/onosproject/onos-config/pkg/types/change/device.Index" json:"change_index,omitempty"`
	// 'values' is a list of values to set
	Values []*device.PathValue `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f618bedd50a8adb9, []int{2}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetID() ID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Snapshot) GetDeviceID() github_com_onosproject_onos_topo_pkg_northbound_device.ID {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *Snapshot) GetSnapshotID() ID {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

func (m *Snapshot) GetChangeIndex() github_com_onosproject_onos_config_pkg_types_change_device.Index {
	if m != nil {
		return m.ChangeIndex
	}
	return 0
}

func (m *Snapshot) GetValues() []*device.PathValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceSnapshot)(nil), "onos.config.snapshot.device.DeviceSnapshot")
	proto.RegisterType((*NetworkSnapshotRef)(nil), "onos.config.snapshot.device.NetworkSnapshotRef")
	proto.RegisterType((*Snapshot)(nil), "onos.config.snapshot.device.Snapshot")
}

func init() {
	proto.RegisterFile("pkg/types/snapshot/device/types.proto", fileDescriptor_f618bedd50a8adb9)
}

var fileDescriptor_f618bedd50a8adb9 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0x6d, 0xb2, 0x6d, 0x37, 0x3b, 0x2d, 0x2a, 0xc3, 0x0a, 0xa1, 0x2e, 0x49, 0x29, 0x0a, 0xbd,
	0x38, 0x81, 0x8a, 0xc2, 0x8a, 0x88, 0xd6, 0x22, 0x04, 0x44, 0x24, 0x2b, 0x5e, 0x6b, 0x9a, 0x4c,
	0xd3, 0xb8, 0x9b, 0x4c, 0x48, 0x26, 0xb5, 0x7e, 0x8b, 0xfd, 0x22, 0x7e, 0x8f, 0xbd, 0xb9, 0x47,
	0x4f, 0x51, 0xd2, 0xbb, 0x1f, 0xa0, 0x27, 0xc9, 0xfc, 0x29, 0x2d, 0x96, 0x85, 0xea, 0xc5, 0x4b,
	0x19, 0x7e, 0xf3, 0xde, 0xef, 0xf7, 0x7e, 0x6f, 0x5e, 0x03, 0x1e, 0x24, 0xe7, 0x81, 0x45, 0xbf,
	0x24, 0x38, 0xb3, 0xb2, 0xd8, 0x4d, 0xb2, 0x19, 0xa1, 0x96, 0x8f, 0xe7, 0xa1, 0x87, 0x79, 0x15,
	0x25, 0x29, 0xa1, 0x04, 0xde, 0x23, 0x31, 0xc9, 0x90, 0x47, 0xe2, 0x69, 0x18, 0x20, 0x09, 0x44,
	0x1c, 0xd8, 0x31, 0x03, 0x42, 0x82, 0x0b, 0x6c, 0x31, 0xe8, 0x24, 0x9f, 0x5a, 0x34, 0x8c, 0x70,
	0x46, 0xdd, 0x28, 0xe1, 0xec, 0xce, 0x71, 0x40, 0x02, 0xc2, 0x8e, 0x56, 0x75, 0x12, 0xd5, 0x97,
	0x41, 0x48, 0x67, 0xf9, 0x04, 0x79, 0x24, 0xb2, 0xaa, 0xf6, 0x49, 0x4a, 0x3e, 0x61, 0x8f, 0xb2,
	0xf3, 0x43, 0x3e, 0xca, 0xda, 0xa1, 0x6e, 0x43, 0x56, 0xe7, 0xf5, 0x5e, 0x2d, 0xbc, 0x99, 0x1b,
	0x07, 0x78, 0xc7, 0x7a, 0xbd, 0x6f, 0x75, 0x70, 0x6b, 0xc4, 0xca, 0x67, 0x62, 0x0c, 0x3c, 0x01,
	0x6a, 0xe8, 0xeb, 0x4a, 0x57, 0xe9, 0x1f, 0x0d, 0xdb, 0x65, 0x61, 0xaa, 0xf6, 0x68, 0xc5, 0x7e,
	0x1d, 0x35, 0xf4, 0xe1, 0x14, 0x1c, 0xf1, 0x36, 0xe3, 0xd0, 0xd7, 0x55, 0x06, 0xb2, 0xcb, 0xc2,
	0xd4, 0x78, 0x13, 0x06, 0x3d, 0xbd, 0x49, 0x1b, 0x25, 0x09, 0x61, 0xca, 0x62, 0x92, 0xd2, 0xd9,
	0x84, 0xe4, 0xb1, 0x2f, 0x74, 0x21, 0x7b, 0xe4, 0x68, 0xfc, 0x68, 0xfb, 0xb0, 0x0f, 0xb4, 0x14,
	0xcf, 0xc3, 0x2c, 0x24, 0xb1, 0x7e, 0xd0, 0x55, 0xfa, 0xf5, 0x61, 0x7b, 0x55, 0x98, 0x9a, 0x23,
	0x6a, 0xce, 0xfa, 0x16, 0x7e, 0x04, 0x77, 0x62, 0x4c, 0x3f, 0x93, 0xf4, 0x7c, 0x2c, 0xad, 0xd2,
	0xeb, 0x5d, 0xa5, 0xdf, 0x1a, 0x58, 0xe8, 0x86, 0xc7, 0x43, 0x6f, 0x39, 0x49, 0xee, 0xed, 0xe0,
	0xe9, 0xb0, 0x7e, 0x55, 0x98, 0x35, 0xe7, 0x76, 0xbc, 0x7d, 0x03, 0x09, 0xd0, 0x23, 0x77, 0x31,
	0x96, 0x53, 0xb8, 0x9b, 0xe3, 0x30, 0xf6, 0xf1, 0x42, 0x6f, 0x30, 0x6d, 0x4f, 0x56, 0x85, 0x39,
	0xd8, 0xe7, 0x49, 0x90, 0x5d, 0xb1, 0x9d, 0xbb, 0x91, 0xbb, 0x10, 0x3a, 0x5e, 0xb1, 0xae, 0xac,
	0x0c, 0x9f, 0x82, 0x66, 0x46, 0x5d, 0x9a, 0x67, 0x7a, 0x93, 0x2d, 0x72, 0xb2, 0x7b, 0x91, 0x33,
	0x86, 0x11, 0xaa, 0x05, 0x03, 0x3e, 0x07, 0x87, 0x5e, 0x8a, 0x5d, 0x8a, 0x7d, 0xfd, 0x90, 0x91,
	0x3b, 0x88, 0xa7, 0x14, 0xc9, 0x94, 0xa2, 0xf7, 0x32, 0xa5, 0x43, 0xad, 0xa2, 0x5e, 0xfe, 0x30,
	0x15, 0x47, 0x92, 0x2a, 0x7e, 0x9e, 0xf8, 0x8c, 0xaf, 0xed, 0xc3, 0x17, 0xa4, 0xde, 0x57, 0x05,
	0xc0, 0x3f, 0xad, 0x85, 0xf6, 0x46, 0xaa, 0x4e, 0xd7, 0xa9, 0xb2, 0xf6, 0xf3, 0x8c, 0x47, 0xf0,
	0x0d, 0x68, 0x70, 0xef, 0xd5, 0x7f, 0xf2, 0x9e, 0x37, 0xe9, 0xfd, 0x52, 0x81, 0xf6, 0x9f, 0x65,
	0xff, 0x31, 0x68, 0xc9, 0x37, 0xae, 0x26, 0x1d, 0xb0, 0x49, 0xc7, 0x65, 0x61, 0x02, 0x29, 0x74,
	0x2d, 0x0b, 0x48, 0xa0, 0xed, 0xc3, 0x00, 0xb4, 0xb7, 0xa2, 0x59, 0x67, 0xf6, 0x54, 0xc8, 0x17,
	0x7f, 0xff, 0xb5, 0x10, 0x66, 0xb5, 0xbc, 0x8d, 0x78, 0x3e, 0x03, 0xcd, 0xb9, 0x7b, 0x91, 0xe3,
	0x4c, 0x6f, 0x74, 0x0f, 0xfa, 0xad, 0xc1, 0xfd, 0xad, 0x78, 0x72, 0xa4, 0xfc, 0x97, 0xbd, 0x73,
	0xe9, 0xec, 0x43, 0x05, 0x76, 0x04, 0x67, 0xa8, 0x5f, 0x95, 0x86, 0x72, 0x5d, 0x1a, 0xca, 0xcf,
	0xd2, 0x50, 0x2e, 0x97, 0x46, 0xed, 0x7a, 0x69, 0xd4, 0xbe, 0x2f, 0x8d, 0xda, 0xa4, 0xc9, 0x12,
	0xf6, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x7b, 0x4e, 0xd6, 0x9b, 0x05, 0x00, 0x00,
}

func (m *DeviceSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Updated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Created, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Created):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.MaxNetworkChangeIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxNetworkChangeIndex))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.NetworkSnapshot.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Revision != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DeviceID) > 0 {
		i -= len(m.DeviceID)
		copy(dAtA[i:], m.DeviceID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DeviceID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NetworkSnapshotRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkSnapshotRef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkSnapshotRef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Values[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ChangeIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ChangeIndex))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SnapshotID) > 0 {
		i -= len(m.SnapshotID)
		copy(dAtA[i:], m.SnapshotID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SnapshotID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DeviceID) > 0 {
		i -= len(m.DeviceID)
		copy(dAtA[i:], m.DeviceID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DeviceID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Revision != 0 {
		n += 1 + sovTypes(uint64(m.Revision))
	}
	l = m.NetworkSnapshot.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.MaxNetworkChangeIndex != 0 {
		n += 1 + sovTypes(uint64(m.MaxNetworkChangeIndex))
	}
	l = m.Status.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *NetworkSnapshotRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SnapshotID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ChangeIndex != 0 {
		n += 1 + sovTypes(uint64(m.ChangeIndex))
	}
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = github_com_onosproject_onos_topo_pkg_northbound_device.ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= Revision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkSnapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetworkSnapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNetworkChangeIndex", wireType)
			}
			m.MaxNetworkChangeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNetworkChangeIndex |= github_com_onosproject_onos_config_pkg_types.Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Created, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Updated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NetworkSnapshotRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkSnapshotRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkSnapshotRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = github_com_onosproject_onos_config_pkg_types.ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= github_com_onosproject_onos_config_pkg_types.Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = github_com_onosproject_onos_topo_pkg_northbound_device.ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotID = ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeIndex", wireType)
			}
			m.ChangeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChangeIndex |= github_com_onosproject_onos_config_pkg_types_change_device.Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &device.PathValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTypes
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)
