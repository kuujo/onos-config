// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/types/snapshot/snapshot.proto

package snapshot

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	change "github.com/onosproject/onos-config/pkg/types/change"
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

// DeviceSnapshot is a snapshot of a single device
type DeviceSnapshot struct {
	// the request identifier
	ID ID `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	// device_id is the device to which the snapshot applies
	DeviceID github_com_onosproject_onos_topo_pkg_northbound_device.ID `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3,casttype=github.com/onosproject/onos-topo/pkg/northbound/device.ID" json:"device_id,omitempty"`
	// values is a list of values to set
	Values []*change.Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	// timestamp is the time at which to take the snapshot
	Timestamp time.Time `protobuf:"bytes,4,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *DeviceSnapshot) Reset()         { *m = DeviceSnapshot{} }
func (m *DeviceSnapshot) String() string { return proto.CompactTextString(m) }
func (*DeviceSnapshot) ProtoMessage()    {}
func (*DeviceSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_2795dc4c070a10ed, []int{0}
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

func (m *DeviceSnapshot) GetValues() []*change.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *DeviceSnapshot) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*DeviceSnapshot)(nil), "onos.config.snapshot.DeviceSnapshot")
}

func init() { proto.RegisterFile("pkg/types/snapshot/snapshot.proto", fileDescriptor_2795dc4c070a10ed) }

var fileDescriptor_2795dc4c070a10ed = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xe3, 0xf4, 0xaa, 0x6a, 0xd3, 0x2b, 0x86, 0xa8, 0x43, 0xa8, 0x50, 0x52, 0x98, 0xba,
	0x60, 0x8b, 0x32, 0xb1, 0xa1, 0xa8, 0x4b, 0xd6, 0x80, 0x58, 0x51, 0xfe, 0xb8, 0x8e, 0xa1, 0xcd,
	0x67, 0xd5, 0x4e, 0x25, 0x1e, 0x02, 0xa9, 0x8f, 0xd5, 0xb1, 0x23, 0x53, 0x41, 0xe9, 0x5b, 0x30,
	0xa1, 0xda, 0x69, 0xcb, 0xc4, 0x62, 0x1f, 0xd9, 0xbf, 0xef, 0xf8, 0xf8, 0x38, 0x97, 0xe2, 0x95,
	0x11, 0xf5, 0x26, 0xa8, 0x24, 0xb2, 0x4c, 0x84, 0x2c, 0x40, 0x1d, 0x05, 0x16, 0x0b, 0x50, 0xe0,
	0xf6, 0xa1, 0x04, 0x89, 0x33, 0x28, 0xa7, 0x9c, 0xe1, 0xc3, 0xdd, 0x20, 0x60, 0x00, 0x6c, 0x46,
	0x89, 0x66, 0xd2, 0x6a, 0x4a, 0x14, 0x9f, 0x53, 0xa9, 0x92, 0xb9, 0x30, 0x63, 0x83, 0x3e, 0x03,
	0x06, 0x5a, 0x92, 0xbd, 0x6a, 0x4e, 0xef, 0x19, 0x57, 0x45, 0x95, 0xe2, 0x0c, 0xe6, 0x64, 0xef,
	0x2b, 0x16, 0xf0, 0x42, 0x33, 0xa5, 0xf5, 0xb5, 0x79, 0x83, 0x9c, 0x22, 0x65, 0x45, 0x52, 0x32,
	0xda, 0x6c, 0xc6, 0xe1, 0xea, 0xdd, 0x76, 0xce, 0x26, 0x74, 0xc9, 0x33, 0xfa, 0xd0, 0x64, 0x71,
	0x2f, 0x1c, 0x9b, 0xe7, 0x1e, 0x1a, 0xa2, 0x51, 0x37, 0xfc, 0x5f, 0x6f, 0x03, 0x3b, 0x9a, 0x7c,
	0xeb, 0x35, 0xb6, 0x79, 0xee, 0x4e, 0x9d, 0x6e, 0xae, 0xf9, 0x67, 0x9e, 0x7b, 0xb6, 0x86, 0xa2,
	0x7a, 0x1b, 0x74, 0x8c, 0x89, 0x46, 0xef, 0xfe, 0x4a, 0xa5, 0x40, 0x80, 0xce, 0x54, 0xc2, 0x42,
	0x15, 0x29, 0x54, 0x65, 0x4e, 0x8c, 0x21, 0x8e, 0x26, 0x71, 0xc7, 0xc8, 0x28, 0x77, 0x6f, 0x9c,
	0xf6, 0x32, 0x99, 0x55, 0x54, 0x7a, 0xad, 0x61, 0x6b, 0xd4, 0x1b, 0x9f, 0xe3, 0xdf, 0xc5, 0x35,
	0x7f, 0x78, 0xda, 0x13, 0x71, 0x03, 0xba, 0xa1, 0xd3, 0x3d, 0xd6, 0xe6, 0xfd, 0x1b, 0xa2, 0x51,
	0x6f, 0x3c, 0xc0, 0xa6, 0x58, 0x7c, 0x28, 0x16, 0x3f, 0x1e, 0x88, 0xb0, 0xb3, 0xde, 0x06, 0xd6,
	0xea, 0x33, 0x40, 0xf1, 0x69, 0x2c, 0xf4, 0xd6, 0xb5, 0x8f, 0x36, 0xb5, 0x8f, 0xbe, 0x6a, 0x1f,
	0xad, 0x76, 0xbe, 0xb5, 0xd9, 0xf9, 0xd6, 0xc7, 0xce, 0xb7, 0xd2, 0xb6, 0xb6, 0xb8, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x49, 0x2d, 0xe2, 0x44, 0xe4, 0x01, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSnapshot(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Values[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSnapshot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DeviceID) > 0 {
		i -= len(m.DeviceID)
		copy(dAtA[i:], m.DeviceID)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.DeviceID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
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
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
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
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = github_com_onosproject_onos_topo_pkg_northbound_device.ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &change.Value{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
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
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSnapshot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSnapshot
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
				next, err := skipSnapshot(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSnapshot
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
	ErrInvalidLengthSnapshot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot   = fmt.Errorf("proto: integer overflow")
)
