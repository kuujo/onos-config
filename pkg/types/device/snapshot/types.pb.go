// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/types/device/snapshot/types.proto

package snapshot

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	change "github.com/onosproject/onos-config/pkg/types/device/change"
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
	return fileDescriptor_9a742975e59c0afb, []int{0}
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
	return fileDescriptor_9a742975e59c0afb, []int{1}
}

// DeviceSnapshot is a device snapshot request
// The DeviceSnapshot maintains the status of the snapshot request but does not store the snapshot
// state itself. During the snapshot process, the DeviceSnapshot will be updated with status
// information indicating when the snapshot is in progress and whether it succeeds or fails.
type DeviceSnapshot struct {
	// 'id' is the snapshot request identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// 'device_id' is the device to which the snapshot applies
	DeviceID github_com_onosproject_onos_topo_pkg_northbound_device.ID `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3,casttype=github.com/onosproject/onos-topo/pkg/northbound/device.ID" json:"device_id,omitempty"`
	// 'revision' is the request revision number
	// Each time the snapshot state is updated, a new revision number will be assigned. The revision
	// is guaranteed to be unique and monotonically increasing and is thus suitable for optimistic
	// concurrency control.
	Revision Revision `protobuf:"varint,3,opt,name=revision,proto3,casttype=Revision" json:"revision,omitempty"`
	// 'status' is the current status of the snapshot request
	// The snapshot will be initialized with the 'PENDING' status and will progress through the 'APPLYING'
	// status and ultimately marked 'SUCCEEDED' or 'FAILED' once the snapshot is complete. If the snapshot
	// fails, a 'reason' and optional 'message' will be provided.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=onos.config.device.snapshot.Status" json:"status,omitempty"`
	// 'reason' is the snapshot failure reason
	// The 'reason' field will be set to 'ERROR' by default and is invalid unless the 'status' is 'FAILED'.
	Reason Reason `protobuf:"varint,5,opt,name=reason,proto3,enum=onos.config.device.snapshot.Reason" json:"reason,omitempty"`
	// 'message' is an optional status message
	// The message will be empty unless the snapshot is failed with a 'FAILED' status, in which case the
	// 'message' field may be used to provide a failure message.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// 'timestamp' is the wall clock time at which to take the snapshot
	// The timestamp will dictate how much of the history is retained following the snapshot. Device changes
	// created after the configured timestamp will be preserved, while changes created prior to the snapshot
	// timestamp will be aggregated into a single Snapshot and the individual changes will be deleted.
	Timestamp time.Time `protobuf:"bytes,7,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// 'created' is the time at which the configuration was created
	Created time.Time `protobuf:"bytes,8,opt,name=created,proto3,stdtime" json:"created"`
	// 'updated' is the time at which the configuration was last updated
	Updated time.Time `protobuf:"bytes,9,opt,name=updated,proto3,stdtime" json:"updated"`
}

func (m *DeviceSnapshot) Reset()         { *m = DeviceSnapshot{} }
func (m *DeviceSnapshot) String() string { return proto.CompactTextString(m) }
func (*DeviceSnapshot) ProtoMessage()    {}
func (*DeviceSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a742975e59c0afb, []int{0}
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

func (m *DeviceSnapshot) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PENDING
}

func (m *DeviceSnapshot) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_ERROR
}

func (m *DeviceSnapshot) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeviceSnapshot) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
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

// Snapshot is the state of a device snapshot
// A single snapshot stores the aggregate of state changes applied to a device.
type Snapshot struct {
	// 'id' is the unique snapshot identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// 'device_id' is the device to which the snapshot applies
	DeviceID github_com_onosproject_onos_topo_pkg_northbound_device.ID `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3,casttype=github.com/onosproject/onos-topo/pkg/northbound/device.ID" json:"device_id,omitempty"`
	// 'device_version' is the version of the device
	DeviceVersion string `protobuf:"bytes,3,opt,name=device_version,json=deviceVersion,proto3" json:"device_version,omitempty"`
	// 'snapshot_id' is the ID of the snapshot request
	SnapshotID ID `protobuf:"bytes,4,opt,name=snapshot_id,json=snapshotId,proto3,casttype=ID" json:"snapshot_id,omitempty"`
	// 'index' is the index at which the snapshot was taken
	Index Index `protobuf:"varint,5,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
	// 'values' is a list of values that were snapshotted
	// The values are the state of the snapshot and represent an aggregation of all changes applied to the
	// device up to the configured 'index'.
	Values []*change.Value `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	// 'timestamp' is the wall clock time at which the snapshot was taken
	Timestamp time.Time `protobuf:"bytes,7,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a742975e59c0afb, []int{1}
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

func (m *Snapshot) GetDeviceVersion() string {
	if m != nil {
		return m.DeviceVersion
	}
	return ""
}

func (m *Snapshot) GetSnapshotID() ID {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

func (m *Snapshot) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Snapshot) GetValues() []*change.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Snapshot) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("onos.config.device.snapshot.Status", Status_name, Status_value)
	proto.RegisterEnum("onos.config.device.snapshot.Reason", Reason_name, Reason_value)
	proto.RegisterType((*DeviceSnapshot)(nil), "onos.config.device.snapshot.DeviceSnapshot")
	proto.RegisterType((*Snapshot)(nil), "onos.config.device.snapshot.Snapshot")
}

func init() {
	proto.RegisterFile("pkg/types/device/snapshot/types.proto", fileDescriptor_9a742975e59c0afb)
}

var fileDescriptor_9a742975e59c0afb = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xb5, 0xfc, 0x23, 0x4b, 0xd7, 0x49, 0x30, 0xf3, 0x65, 0x21, 0xf2, 0x15, 0xcb, 0xa4, 0x04,
	0x4c, 0xa0, 0x23, 0x48, 0x29, 0xb4, 0x14, 0x02, 0x71, 0xe4, 0x14, 0x41, 0x48, 0xc3, 0xa4, 0x0d,
	0x74, 0x55, 0x64, 0x6b, 0x22, 0xab, 0x8d, 0x35, 0x42, 0x33, 0x36, 0xed, 0x43, 0x14, 0xb2, 0xec,
	0x23, 0x65, 0x99, 0x65, 0x57, 0x6a, 0x51, 0xde, 0xc2, 0xab, 0xa2, 0x19, 0xc9, 0x2d, 0xb4, 0x84,
	0x06, 0xba, 0xe8, 0x46, 0xdc, 0x39, 0x73, 0xce, 0x9d, 0xa3, 0xb9, 0x67, 0x60, 0x27, 0x79, 0x1f,
	0x3a, 0xe2, 0x63, 0x42, 0xb9, 0x13, 0xd0, 0x45, 0x34, 0xa1, 0x0e, 0x8f, 0xfd, 0x84, 0x4f, 0x99,
	0x50, 0x28, 0x4e, 0x52, 0x26, 0x18, 0xfa, 0x9f, 0xc5, 0x8c, 0xe3, 0x09, 0x8b, 0x2f, 0xa2, 0x10,
	0x2b, 0x22, 0xae, 0x88, 0x5b, 0x76, 0xc8, 0x58, 0x78, 0x49, 0x1d, 0x49, 0x1d, 0xcf, 0x2f, 0x1c,
	0x11, 0xcd, 0x28, 0x17, 0xfe, 0x2c, 0x51, 0xea, 0xad, 0xcd, 0x90, 0x85, 0x4c, 0x96, 0x4e, 0x51,
	0x95, 0xe8, 0x51, 0x18, 0x89, 0xe9, 0x7c, 0x8c, 0x27, 0x6c, 0xe6, 0x14, 0xed, 0x93, 0x94, 0xbd,
	0xa3, 0x13, 0x21, 0xeb, 0x47, 0xea, 0x28, 0xe7, 0x17, 0x77, 0x93, 0xa9, 0x1f, 0x87, 0xf4, 0x67,
	0x6f, 0xdb, 0x9f, 0x9b, 0xb0, 0xe1, 0xca, 0xdd, 0xb3, 0xd2, 0x11, 0x7a, 0x00, 0xf5, 0x28, 0xb0,
	0xb4, 0xbe, 0x36, 0x30, 0x87, 0x6b, 0x79, 0x66, 0xd7, 0x3d, 0x77, 0x29, 0xbf, 0xa4, 0x1e, 0x05,
	0xe8, 0x02, 0x4c, 0xd5, 0xed, 0x6d, 0x14, 0x58, 0x75, 0x49, 0xf2, 0xf2, 0xcc, 0x36, 0x54, 0x13,
	0x49, 0x7d, 0x76, 0x97, 0x37, 0xc1, 0x12, 0x26, 0x9d, 0xc5, 0x2c, 0x15, 0xd3, 0x31, 0x9b, 0xc7,
	0x41, 0x69, 0x0f, 0x7b, 0x2e, 0x31, 0x54, 0xe9, 0x05, 0x68, 0x00, 0x46, 0x4a, 0x17, 0x11, 0x8f,
	0x58, 0x6c, 0x35, 0xfa, 0xda, 0xa0, 0x39, 0x5c, 0x5b, 0x66, 0xb6, 0x41, 0x4a, 0x8c, 0xac, 0x76,
	0xd1, 0x73, 0xd0, 0xb9, 0xf0, 0xc5, 0x9c, 0x5b, 0xcd, 0xbe, 0x36, 0xd8, 0xd8, 0x7b, 0x88, 0xef,
	0xb8, 0x6f, 0x7c, 0x26, 0xa9, 0xa4, 0x94, 0x14, 0xe2, 0x94, 0xfa, 0x9c, 0xc5, 0x56, 0xeb, 0x0f,
	0xc4, 0x44, 0x52, 0x49, 0x29, 0x41, 0x16, 0xb4, 0x67, 0x94, 0x73, 0x3f, 0xa4, 0x96, 0x5e, 0xdc,
	0x04, 0xa9, 0x96, 0x68, 0x08, 0xe6, 0x6a, 0x8e, 0x56, 0xbb, 0xaf, 0x0d, 0x3a, 0x7b, 0x5b, 0x58,
	0x4d, 0x1a, 0x57, 0x93, 0xc6, 0xaf, 0x2a, 0xc6, 0xd0, 0xb8, 0xce, 0xec, 0xda, 0xd5, 0x57, 0x5b,
	0x23, 0x3f, 0x64, 0x68, 0x1f, 0xda, 0x93, 0x94, 0xfa, 0x82, 0x06, 0x96, 0x71, 0x8f, 0x0e, 0x95,
	0xa8, 0xd0, 0xcf, 0x93, 0x40, 0xea, 0xcd, 0xfb, 0xe8, 0x4b, 0xd1, 0xf6, 0xa7, 0x06, 0x18, 0xff,
	0x58, 0x28, 0x76, 0x60, 0xa3, 0x3c, 0x67, 0x41, 0xd3, 0x55, 0x34, 0x4c, 0xb2, 0xae, 0xd0, 0x73,
	0x05, 0xa2, 0x27, 0xd0, 0xa9, 0x46, 0x56, 0x18, 0x6a, 0x4a, 0x43, 0x9b, 0x79, 0x66, 0x43, 0xf5,
	0x3f, 0x2b, 0xf7, 0x50, 0x11, 0xbd, 0x00, 0xd9, 0xd0, 0x8a, 0xe2, 0x80, 0x7e, 0x90, 0x51, 0x68,
	0x0e, 0xcd, 0x65, 0x66, 0xb7, 0xbc, 0x02, 0x20, 0x0a, 0x47, 0x4f, 0x41, 0x5f, 0xf8, 0x97, 0x73,
	0xca, 0x2d, 0xbd, 0xdf, 0x18, 0x74, 0xf6, 0xfa, 0xbf, 0x0b, 0x8b, 0x7a, 0x64, 0xf8, 0xbc, 0x20,
	0x92, 0x92, 0xff, 0x37, 0xf2, 0xb0, 0xbb, 0x0f, 0xba, 0x0a, 0x2f, 0xea, 0x40, 0xfb, 0x74, 0x74,
	0xe2, 0x7a, 0x27, 0x2f, 0xba, 0x35, 0xb4, 0x06, 0xc6, 0xc1, 0xe9, 0xe9, 0xf1, 0x9b, 0x62, 0xa5,
	0xa1, 0x75, 0x30, 0xcf, 0x5e, 0x1f, 0x1e, 0x8e, 0x46, 0xee, 0xc8, 0xed, 0xd6, 0x11, 0x80, 0x7e,
	0x74, 0xe0, 0x1d, 0x8f, 0xdc, 0x6e, 0x63, 0xf7, 0x3f, 0xd0, 0x55, 0x7e, 0x91, 0x09, 0xad, 0x11,
	0x21, 0x2f, 0x49, 0xb7, 0x36, 0xb4, 0xae, 0xf3, 0x9e, 0x76, 0x93, 0xf7, 0xb4, 0x6f, 0x79, 0x4f,
	0xbb, 0xba, 0xed, 0xd5, 0x6e, 0x6e, 0x7b, 0xb5, 0x2f, 0xb7, 0xbd, 0xda, 0x58, 0x97, 0xbe, 0x1e,
	0x7f, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x04, 0x96, 0xed, 0x18, 0xe5, 0x04, 0x00, 0x00,
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
	dAtA[i] = 0x4a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Created, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Created):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTypes(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x3a
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
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintTypes(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x3a
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
			dAtA[i] = 0x32
		}
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SnapshotID) > 0 {
		i -= len(m.SnapshotID)
		copy(dAtA[i:], m.SnapshotID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SnapshotID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DeviceVersion) > 0 {
		i -= len(m.DeviceVersion)
		copy(dAtA[i:], m.DeviceVersion)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DeviceVersion)))
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
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovTypes(uint64(l))
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
	l = len(m.DeviceVersion)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SnapshotID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTypes(uint64(l))
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
		case 8:
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
		case 9:
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
		case 4:
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
		case 5:
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
		case 6:
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
			m.Values = append(m.Values, &change.Value{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
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
