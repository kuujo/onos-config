// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/types/network/snapshot/types.proto

package snapshot

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// Status is the status of a snapshot object
type Status int32

const (
	// PENDING indicates the snapshot is waiting to be applied
	Status_PENDING Status = 0
	// APPLYING indicates the snapshot is being applied
	Status_APPLYING Status = 1
	// SUCCEEDED indicates the snapshot was applied successfully
	Status_SUCCEEDED Status = 2
	// FAILED indicates the snapshot failed
	Status_FAILED Status = 3
)

var Status_name = map[int32]string{
	0: "PENDING",
	1: "APPLYING",
	2: "SUCCEEDED",
	3: "FAILED",
}

var Status_value = map[string]int32{
	"PENDING":   0,
	"APPLYING":  1,
	"SUCCEEDED": 2,
	"FAILED":    3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de0254b3b791dba5, []int{0}
}

// Reason is a failure reason
type Reason int32

const (
	// ERROR indicates an error occurred during snapshot
	Reason_ERROR Reason = 0
)

var Reason_name = map[int32]string{
	0: "ERROR",
}

var Reason_value = map[string]int32{
	"ERROR": 0,
}

func (x Reason) String() string {
	return proto.EnumName(Reason_name, int32(x))
}

func (Reason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de0254b3b791dba5, []int{1}
}

// NetworkSnapshot is a network snapshot request
// The snapshot request indicates a set of devices to snapshot. Once the request is created, snapshots
// will be taken of the state of each device in the 'devices' list or all devices if no devices are specified.
type NetworkSnapshot struct {
	// 'id' is the the snapshot request identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// 'index' is the globally unique snapshot request index
	// Snapshot requests are assigned a global total order which is used to apply snapshots in FIFO
	// order. Each request is assigned a monotonically increasing snapshot index, and snapshots are
	// applied in the order in which they are created.
	Index Index `protobuf:"varint,2,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
	// 'revision' is the request revision number
	// Each time the snapshot state is updated, a new revision number will be assigned. The revision
	// is guaranteed to be unique and monotonically increasing and is thus suitable for optimistic
	// concurrency control.
	Revision Revision `protobuf:"varint,3,opt,name=revision,proto3,casttype=Revision" json:"revision,omitempty"`
	// 'status' is the current status of the snapshot
	// The snapshot will be initialized with the 'PENDING' status and will progress through the 'APPLYING'
	// status and ultimately marked 'SUCCEEDED' or 'FAILED' once the snapshot is complete. If the snapshot
	// fails, a 'reason' and optional 'message' will be provided.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=onos.config.network.snapshot.Status" json:"status,omitempty"`
	// 'reason' is the reason the snapshot failed
	// The 'reason' field will be set to 'ERROR' by default and is invalid unless the 'status' is 'FAILED'.
	Reason Reason `protobuf:"varint,5,opt,name=reason,proto3,enum=onos.config.network.snapshot.Reason" json:"reason,omitempty"`
	// 'message' is an optional status message
	// The message will be empty unless the snapshot is failed with a 'FAILED' status, in which case the
	// 'message' field may be used to provide a failure message.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// 'devices' is a list of device/version pairs to snapshot
	// If the 'devices' list is empty, snapshots will be taken of all devices for which changes exist.
	// If a list of devices is provided, snapshots will taken only of the devices present in the list.
	// For each device in the list, an optional version may be specified. If no version is specified,
	// the snapshot will be applied to all known versions of the device.
	Devices []*DeviceSnapshotId `protobuf:"bytes,7,rep,name=devices,proto3" json:"devices,omitempty"`
	// 'timestamp' is the wall clock time at which to take the snapshot
	// The timestamp will dictate how much of the history is retained following the snapshot. Changes
	// submitted after the configured timestamp will be preserved, while changes submitted prior to the
	// timestamp will be aggregated into a single Snapshot and the individual changes will be deleted.
	Timestamp time.Time `protobuf:"bytes,8,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// 'created' is the time at which the snapshot was created
	Created time.Time `protobuf:"bytes,9,opt,name=created,proto3,stdtime" json:"created"`
	// 'updated' is the time at which the snapshot was last updated
	Updated time.Time `protobuf:"bytes,10,opt,name=updated,proto3,stdtime" json:"updated"`
}

func (m *NetworkSnapshot) Reset()         { *m = NetworkSnapshot{} }
func (m *NetworkSnapshot) String() string { return proto.CompactTextString(m) }
func (*NetworkSnapshot) ProtoMessage()    {}
func (*NetworkSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_de0254b3b791dba5, []int{0}
}
func (m *NetworkSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkSnapshot.Merge(m, src)
}
func (m *NetworkSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *NetworkSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkSnapshot proto.InternalMessageInfo

func (m *NetworkSnapshot) GetID() ID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NetworkSnapshot) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *NetworkSnapshot) GetRevision() Revision {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *NetworkSnapshot) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PENDING
}

func (m *NetworkSnapshot) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_ERROR
}

func (m *NetworkSnapshot) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NetworkSnapshot) GetDevices() []*DeviceSnapshotId {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *NetworkSnapshot) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *NetworkSnapshot) GetCreated() time.Time {
	if m != nil {
		return m.Created
	}
	return time.Time{}
}

func (m *NetworkSnapshot) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

// DeviceSnapshotId is an identifier for a device/version pair
type DeviceSnapshotId struct {
	// 'device_id' is the device identifier
	DeviceID github_com_onosproject_onos_topo_pkg_northbound_device.ID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3,casttype=github.com/onosproject/onos-topo/pkg/northbound/device.ID" json:"device_id,omitempty"`
	// 'device_version' is an optional device version
	// If no version is provided, the snapshot will be applied to all versions of the device.
	DeviceVersion string `protobuf:"bytes,2,opt,name=device_version,json=deviceVersion,proto3" json:"device_version,omitempty"`
}

func (m *DeviceSnapshotId) Reset()         { *m = DeviceSnapshotId{} }
func (m *DeviceSnapshotId) String() string { return proto.CompactTextString(m) }
func (*DeviceSnapshotId) ProtoMessage()    {}
func (*DeviceSnapshotId) Descriptor() ([]byte, []int) {
	return fileDescriptor_de0254b3b791dba5, []int{1}
}
func (m *DeviceSnapshotId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceSnapshotId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceSnapshotId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceSnapshotId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSnapshotId.Merge(m, src)
}
func (m *DeviceSnapshotId) XXX_Size() int {
	return m.Size()
}
func (m *DeviceSnapshotId) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSnapshotId.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSnapshotId proto.InternalMessageInfo

func (m *DeviceSnapshotId) GetDeviceID() github_com_onosproject_onos_topo_pkg_northbound_device.ID {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DeviceSnapshotId) GetDeviceVersion() string {
	if m != nil {
		return m.DeviceVersion
	}
	return ""
}

func init() {
	proto.RegisterEnum("onos.config.network.snapshot.Status", Status_name, Status_value)
	proto.RegisterEnum("onos.config.network.snapshot.Reason", Reason_name, Reason_value)
	proto.RegisterType((*NetworkSnapshot)(nil), "onos.config.network.snapshot.NetworkSnapshot")
	proto.RegisterType((*DeviceSnapshotId)(nil), "onos.config.network.snapshot.DeviceSnapshotId")
}

func init() {
	proto.RegisterFile("pkg/types/network/snapshot/types.proto", fileDescriptor_de0254b3b791dba5)
}

var fileDescriptor_de0254b3b791dba5 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0xdb, 0x4c,
	0x14, 0x85, 0x25, 0x25, 0x96, 0xa5, 0x89, 0x93, 0xdf, 0xcc, 0xdf, 0x85, 0x30, 0x41, 0x32, 0xa1,
	0x2d, 0x22, 0xd0, 0x11, 0xb8, 0xab, 0x42, 0x09, 0xc4, 0x91, 0xda, 0x0a, 0x82, 0x6b, 0xc6, 0x6d,
	0xa1, 0xab, 0x22, 0x5b, 0x63, 0x59, 0x4d, 0xad, 0x11, 0x9a, 0xb1, 0xdb, 0xbe, 0x45, 0x9e, 0xa0,
	0xd0, 0xb7, 0xc9, 0x32, 0xcb, 0xae, 0xdc, 0x62, 0xbf, 0x85, 0x57, 0x45, 0x33, 0x92, 0x0b, 0x5d,
	0x84, 0x66, 0x63, 0xee, 0x1c, 0x9f, 0xef, 0x8a, 0x7b, 0xef, 0x01, 0x8f, 0xf3, 0xab, 0xc4, 0xe3,
	0x5f, 0x73, 0xc2, 0xbc, 0x8c, 0xf0, 0xcf, 0xb4, 0xb8, 0xf2, 0x58, 0x16, 0xe5, 0x6c, 0x46, 0xb9,
	0x94, 0x51, 0x5e, 0x50, 0x4e, 0xe1, 0x31, 0xcd, 0x28, 0x43, 0x13, 0x9a, 0x4d, 0xd3, 0x04, 0x55,
	0x4e, 0x54, 0x3b, 0x3b, 0x4e, 0x42, 0x69, 0xf2, 0x89, 0x78, 0xc2, 0x3b, 0x5e, 0x4c, 0x3d, 0x9e,
	0xce, 0x09, 0xe3, 0xd1, 0x3c, 0x97, 0x78, 0xe7, 0x41, 0x42, 0x13, 0x2a, 0x4a, 0xaf, 0xac, 0xa4,
	0x7a, 0xf2, 0x6d, 0x1f, 0xfc, 0x37, 0x90, 0xbd, 0x46, 0x55, 0x2b, 0x78, 0x0c, 0xb4, 0x34, 0xb6,
	0xd4, 0xae, 0xea, 0x9a, 0xfd, 0xd6, 0x7a, 0xe5, 0x68, 0xa1, 0xbf, 0x15, 0xbf, 0x58, 0x4b, 0x63,
	0xe8, 0x80, 0x46, 0x9a, 0xc5, 0xe4, 0x8b, 0xa5, 0x75, 0x55, 0x77, 0xbf, 0x6f, 0x6e, 0x57, 0x4e,
	0x23, 0x2c, 0x05, 0x2c, 0x75, 0xe8, 0x02, 0xa3, 0x20, 0xcb, 0x94, 0xa5, 0x34, 0xb3, 0xf6, 0x84,
	0xa7, 0xb5, 0x5d, 0x39, 0x06, 0xae, 0x34, 0xbc, 0xfb, 0x17, 0x3e, 0x07, 0x3a, 0xe3, 0x11, 0x5f,
	0x30, 0x6b, 0xbf, 0xab, 0xba, 0x47, 0xbd, 0x87, 0xe8, 0xae, 0x11, 0xd1, 0x48, 0x78, 0x71, 0xc5,
	0x94, 0x74, 0x41, 0x22, 0x46, 0x33, 0xab, 0xf1, 0x2f, 0x34, 0x16, 0x5e, 0x5c, 0x31, 0xd0, 0x02,
	0xcd, 0x39, 0x61, 0x2c, 0x4a, 0x88, 0xa5, 0x97, 0x93, 0xe2, 0xfa, 0x09, 0x5f, 0x81, 0x66, 0x4c,
	0x96, 0xe9, 0x84, 0x30, 0xab, 0xd9, 0xdd, 0x73, 0x0f, 0x7a, 0xe8, 0xee, 0xc6, 0xbe, 0x30, 0xd7,
	0xdb, 0x0b, 0x63, 0x5c, 0xe3, 0xb0, 0x0f, 0xcc, 0xdd, 0x15, 0x2c, 0xa3, 0xab, 0xba, 0x07, 0xbd,
	0x0e, 0x92, 0x77, 0x42, 0xf5, 0x9d, 0xd0, 0x9b, 0xda, 0xd1, 0x37, 0x6e, 0x56, 0x8e, 0x72, 0xfd,
	0xd3, 0x51, 0xf1, 0x1f, 0x0c, 0x9e, 0x81, 0xe6, 0xa4, 0x20, 0x11, 0x27, 0xb1, 0x65, 0xde, 0xa3,
	0x43, 0x0d, 0x95, 0xfc, 0x22, 0x8f, 0x05, 0x0f, 0xee, 0xc3, 0x57, 0xd0, 0xc9, 0x77, 0x15, 0xb4,
	0xff, 0x9e, 0x10, 0x4e, 0x81, 0x29, 0x67, 0xfc, 0xb0, 0x0b, 0x4a, 0xb8, 0x5e, 0x39, 0x86, 0x34,
	0x8a, 0xb8, 0x3c, 0x4b, 0x52, 0x3e, 0x5b, 0x8c, 0xd1, 0x84, 0xce, 0xbd, 0x72, 0x7d, 0x79, 0x41,
	0x3f, 0x92, 0x09, 0x17, 0xf5, 0x13, 0x4e, 0x73, 0xea, 0x95, 0xb1, 0xcf, 0x68, 0xc1, 0x67, 0x63,
	0xba, 0xc8, 0x62, 0x4f, 0x36, 0x44, 0xa1, 0x8f, 0x0d, 0x59, 0x86, 0x31, 0x7c, 0x04, 0x8e, 0xaa,
	0xef, 0x2c, 0x49, 0x21, 0x02, 0xa5, 0x89, 0x5b, 0x1d, 0x4a, 0xf5, 0x9d, 0x14, 0x4f, 0xcf, 0x80,
	0x2e, 0xb3, 0x01, 0x0f, 0x40, 0x73, 0x18, 0x0c, 0xfc, 0x70, 0xf0, 0xb2, 0xad, 0xc0, 0x16, 0x30,
	0xce, 0x87, 0xc3, 0xcb, 0xf7, 0xe5, 0x4b, 0x85, 0x87, 0xc0, 0x1c, 0xbd, 0xbd, 0xb8, 0x08, 0x02,
	0x3f, 0xf0, 0xdb, 0x1a, 0x04, 0x40, 0x7f, 0x71, 0x1e, 0x5e, 0x06, 0x7e, 0x7b, 0xef, 0xf4, 0x7f,
	0xa0, 0xcb, 0x74, 0x40, 0x13, 0x34, 0x02, 0x8c, 0x5f, 0xe3, 0xb6, 0xd2, 0xb7, 0x6e, 0xd6, 0xb6,
	0x7a, 0xbb, 0xb6, 0xd5, 0x5f, 0x6b, 0x5b, 0xbd, 0xde, 0xd8, 0xca, 0xed, 0xc6, 0x56, 0x7e, 0x6c,
	0x6c, 0x65, 0xac, 0x8b, 0xcd, 0x3d, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x30, 0x24, 0xd3, 0x85,
	0xb9, 0x03, 0x00, 0x00,
}

func (m *NetworkSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	dAtA[i] = 0x52
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Created, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Created):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTypes(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x42
	if len(m.Devices) > 0 {
		for iNdEx := len(m.Devices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Devices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x32
	}
	if m.Reason != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Reason))
		i--
		dAtA[i] = 0x28
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Revision != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x18
	}
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

func (m *DeviceSnapshotId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceSnapshotId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceSnapshotId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeviceVersion) > 0 {
		i -= len(m.DeviceVersion)
		copy(dAtA[i:], m.DeviceVersion)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DeviceVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeviceID) > 0 {
		i -= len(m.DeviceID)
		copy(dAtA[i:], m.DeviceID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DeviceID)))
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
func (m *NetworkSnapshot) Size() (n int) {
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
	if m.Revision != 0 {
		n += 1 + sovTypes(uint64(m.Revision))
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	if m.Reason != 0 {
		n += 1 + sovTypes(uint64(m.Reason))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *DeviceSnapshotId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DeviceVersion)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkSnapshot) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NetworkSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
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
				m.Index |= Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			m.Reason = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reason |= Reason(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
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
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devices", wireType)
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
			m.Devices = append(m.Devices, &DeviceSnapshotId{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
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
		case 10:
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
func (m *DeviceSnapshotId) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeviceSnapshotId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceSnapshotId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceVersion", wireType)
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
			m.DeviceVersion = string(dAtA[iNdEx:postIndex])
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
