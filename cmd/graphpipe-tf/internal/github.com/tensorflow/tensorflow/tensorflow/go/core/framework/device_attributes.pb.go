// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/device_attributes.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InterconnectLink struct {
	DeviceId int32  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Strength int32  `protobuf:"varint,3,opt,name=strength,proto3" json:"strength,omitempty"`
}

func (m *InterconnectLink) Reset()                    { *m = InterconnectLink{} }
func (m *InterconnectLink) String() string            { return proto.CompactTextString(m) }
func (*InterconnectLink) ProtoMessage()               {}
func (*InterconnectLink) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAttributes, []int{0} }

func (m *InterconnectLink) GetDeviceId() int32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *InterconnectLink) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InterconnectLink) GetStrength() int32 {
	if m != nil {
		return m.Strength
	}
	return 0
}

type LocalLinks struct {
	Link []*InterconnectLink `protobuf:"bytes,1,rep,name=link" json:"link,omitempty"`
}

func (m *LocalLinks) Reset()                    { *m = LocalLinks{} }
func (m *LocalLinks) String() string            { return proto.CompactTextString(m) }
func (*LocalLinks) ProtoMessage()               {}
func (*LocalLinks) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAttributes, []int{1} }

func (m *LocalLinks) GetLink() []*InterconnectLink {
	if m != nil {
		return m.Link
	}
	return nil
}

type DeviceLocality struct {
	// Optional bus locality of device.  Default value of 0 means
	// no specific locality.  Specific localities are indexed from 1.
	BusId int32 `protobuf:"varint,1,opt,name=bus_id,json=busId,proto3" json:"bus_id,omitempty"`
	// Optional NUMA locality of device.
	NumaNode int32 `protobuf:"varint,2,opt,name=numa_node,json=numaNode,proto3" json:"numa_node,omitempty"`
	// Optional local interconnect links to other devices.
	Links *LocalLinks `protobuf:"bytes,3,opt,name=links" json:"links,omitempty"`
}

func (m *DeviceLocality) Reset()                    { *m = DeviceLocality{} }
func (m *DeviceLocality) String() string            { return proto.CompactTextString(m) }
func (*DeviceLocality) ProtoMessage()               {}
func (*DeviceLocality) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAttributes, []int{2} }

func (m *DeviceLocality) GetBusId() int32 {
	if m != nil {
		return m.BusId
	}
	return 0
}

func (m *DeviceLocality) GetNumaNode() int32 {
	if m != nil {
		return m.NumaNode
	}
	return 0
}

func (m *DeviceLocality) GetLinks() *LocalLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type DeviceAttributes struct {
	// Fully specified name of the device within a cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit int64 `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// Platform-specific data about device that may be useful
	// for supporting efficient data transfers.
	Locality *DeviceLocality `protobuf:"bytes,5,opt,name=locality" json:"locality,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation,proto3" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc string `protobuf:"bytes,7,opt,name=physical_device_desc,json=physicalDeviceDesc,proto3" json:"physical_device_desc,omitempty"`
}

func (m *DeviceAttributes) Reset()                    { *m = DeviceAttributes{} }
func (m *DeviceAttributes) String() string            { return proto.CompactTextString(m) }
func (*DeviceAttributes) ProtoMessage()               {}
func (*DeviceAttributes) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAttributes, []int{3} }

func (m *DeviceAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceAttributes) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DeviceAttributes) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DeviceAttributes) GetLocality() *DeviceLocality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *DeviceAttributes) GetIncarnation() uint64 {
	if m != nil {
		return m.Incarnation
	}
	return 0
}

func (m *DeviceAttributes) GetPhysicalDeviceDesc() string {
	if m != nil {
		return m.PhysicalDeviceDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*InterconnectLink)(nil), "tensorflow.InterconnectLink")
	proto.RegisterType((*LocalLinks)(nil), "tensorflow.LocalLinks")
	proto.RegisterType((*DeviceLocality)(nil), "tensorflow.DeviceLocality")
	proto.RegisterType((*DeviceAttributes)(nil), "tensorflow.DeviceAttributes")
}
func (m *InterconnectLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterconnectLink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DeviceId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.DeviceId))
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Strength != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.Strength))
	}
	return i, nil
}

func (m *LocalLinks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalLinks) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Link) > 0 {
		for _, msg := range m.Link {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeviceAttributes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DeviceLocality) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceLocality) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BusId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.BusId))
	}
	if m.NumaNode != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.NumaNode))
	}
	if m.Links != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.Links.Size()))
		n1, err := m.Links.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *DeviceAttributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceAttributes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.DeviceType) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(len(m.DeviceType)))
		i += copy(dAtA[i:], m.DeviceType)
	}
	if m.MemoryLimit != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.MemoryLimit))
	}
	if m.Locality != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(m.Locality.Size()))
		n2, err := m.Locality.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Incarnation != 0 {
		dAtA[i] = 0x31
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Incarnation))
		i += 8
	}
	if len(m.PhysicalDeviceDesc) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDeviceAttributes(dAtA, i, uint64(len(m.PhysicalDeviceDesc)))
		i += copy(dAtA[i:], m.PhysicalDeviceDesc)
	}
	return i, nil
}

func encodeVarintDeviceAttributes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InterconnectLink) Size() (n int) {
	var l int
	_ = l
	if m.DeviceId != 0 {
		n += 1 + sovDeviceAttributes(uint64(m.DeviceId))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDeviceAttributes(uint64(l))
	}
	if m.Strength != 0 {
		n += 1 + sovDeviceAttributes(uint64(m.Strength))
	}
	return n
}

func (m *LocalLinks) Size() (n int) {
	var l int
	_ = l
	if len(m.Link) > 0 {
		for _, e := range m.Link {
			l = e.Size()
			n += 1 + l + sovDeviceAttributes(uint64(l))
		}
	}
	return n
}

func (m *DeviceLocality) Size() (n int) {
	var l int
	_ = l
	if m.BusId != 0 {
		n += 1 + sovDeviceAttributes(uint64(m.BusId))
	}
	if m.NumaNode != 0 {
		n += 1 + sovDeviceAttributes(uint64(m.NumaNode))
	}
	if m.Links != nil {
		l = m.Links.Size()
		n += 1 + l + sovDeviceAttributes(uint64(l))
	}
	return n
}

func (m *DeviceAttributes) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDeviceAttributes(uint64(l))
	}
	l = len(m.DeviceType)
	if l > 0 {
		n += 1 + l + sovDeviceAttributes(uint64(l))
	}
	if m.MemoryLimit != 0 {
		n += 1 + sovDeviceAttributes(uint64(m.MemoryLimit))
	}
	if m.Locality != nil {
		l = m.Locality.Size()
		n += 1 + l + sovDeviceAttributes(uint64(l))
	}
	if m.Incarnation != 0 {
		n += 9
	}
	l = len(m.PhysicalDeviceDesc)
	if l > 0 {
		n += 1 + l + sovDeviceAttributes(uint64(l))
	}
	return n
}

func sovDeviceAttributes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeviceAttributes(x uint64) (n int) {
	return sovDeviceAttributes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterconnectLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAttributes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InterconnectLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterconnectLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			m.DeviceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strength", wireType)
			}
			m.Strength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Strength |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAttributes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAttributes
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
func (m *LocalLinks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAttributes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocalLinks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalLinks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Link = append(m.Link, &InterconnectLink{})
			if err := m.Link[len(m.Link)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAttributes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAttributes
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
func (m *DeviceLocality) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAttributes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceLocality: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceLocality: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BusId", wireType)
			}
			m.BusId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BusId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumaNode", wireType)
			}
			m.NumaNode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumaNode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Links == nil {
				m.Links = &LocalLinks{}
			}
			if err := m.Links.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAttributes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAttributes
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
func (m *DeviceAttributes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAttributes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceAttributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceAttributes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryLimit", wireType)
			}
			m.MemoryLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryLimit |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locality", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Locality == nil {
				m.Locality = &DeviceLocality{}
			}
			if err := m.Locality.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Incarnation", wireType)
			}
			m.Incarnation = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Incarnation = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhysicalDeviceDesc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAttributes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAttributes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PhysicalDeviceDesc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAttributes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAttributes
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
func skipDeviceAttributes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceAttributes
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
					return 0, ErrIntOverflowDeviceAttributes
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
					return 0, ErrIntOverflowDeviceAttributes
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDeviceAttributes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceAttributes
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
				next, err := skipDeviceAttributes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthDeviceAttributes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceAttributes   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/device_attributes.proto", fileDescriptorDeviceAttributes)
}

var fileDescriptorDeviceAttributes = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0x89, 0xed, 0xd4, 0xf6, 0x8d, 0xc8, 0x12, 0x74, 0x09, 0xab, 0xd4, 0xb1, 0xa7, 0x39,
	0x48, 0xbb, 0xae, 0xe0, 0x4d, 0xc1, 0x65, 0x2f, 0x85, 0x22, 0x4b, 0xf0, 0xe4, 0x65, 0x98, 0xc9,
	0x64, 0xdb, 0xd0, 0x99, 0xa4, 0x24, 0x19, 0x97, 0x7e, 0x01, 0x3f, 0x9b, 0x47, 0x8f, 0x1e, 0xa5,
	0x1f, 0x42, 0x3c, 0x4a, 0x92, 0xb6, 0x33, 0x96, 0xbd, 0xbd, 0xf9, 0xcf, 0xcb, 0x7b, 0xbf, 0xf7,
	0xe7, 0x0f, 0x6f, 0x2d, 0x97, 0x46, 0xe9, 0xbb, 0x4a, 0xdd, 0xcf, 0x98, 0xd2, 0x7c, 0x76, 0xa7,
	0xf3, 0x9a, 0xdf, 0x2b, 0xbd, 0x9e, 0x95, 0xfc, 0x9b, 0x60, 0x3c, 0xcb, 0xad, 0xd5, 0xa2, 0x68,
	0x2c, 0x37, 0xd3, 0x8d, 0x56, 0x56, 0x61, 0x68, 0x9f, 0x4c, 0x32, 0x38, 0x9b, 0x4b, 0xcb, 0x35,
	0x53, 0x52, 0x72, 0x66, 0x17, 0x42, 0xae, 0xf1, 0x0b, 0x18, 0xed, 0x9f, 0x8a, 0x92, 0xa0, 0x04,
	0xa5, 0x11, 0x1d, 0x06, 0x61, 0x5e, 0x62, 0x0c, 0x7d, 0xbb, 0xdd, 0x70, 0xf2, 0x28, 0x41, 0xe9,
	0x88, 0xfa, 0x1a, 0x5f, 0xc0, 0xd0, 0x58, 0xcd, 0xe5, 0xd2, 0xae, 0x48, 0x2f, 0xf4, 0x1f, 0xbe,
	0x27, 0x1f, 0x01, 0x16, 0x8a, 0xe5, 0x95, 0x9b, 0x6c, 0xf0, 0x25, 0xf4, 0x2b, 0x21, 0xd7, 0x04,
	0x25, 0xbd, 0x34, 0xbe, 0x7a, 0x39, 0x6d, 0x49, 0xa6, 0xa7, 0x18, 0xd4, 0x77, 0x4e, 0x34, 0x3c,
	0xbd, 0xf1, 0xbb, 0xfd, 0x14, 0x61, 0xb7, 0xf8, 0x39, 0x0c, 0x8a, 0xc6, 0xb4, 0x6c, 0x51, 0xd1,
	0x98, 0x79, 0xe9, 0xa8, 0x65, 0x53, 0xe7, 0x99, 0x54, 0x65, 0xa0, 0x8b, 0xe8, 0xd0, 0x09, 0x9f,
	0x55, 0xc9, 0xf1, 0x1b, 0x88, 0xdc, 0x34, 0xe3, 0xf1, 0xe2, 0xab, 0xf3, 0xee, 0xe2, 0x16, 0x8f,
	0x86, 0xa6, 0xc9, 0x1f, 0x04, 0x67, 0x61, 0xe9, 0xa7, 0xa3, 0x77, 0xee, 0x70, 0x99, 0xd7, 0xdc,
	0x2f, 0x1d, 0x51, 0x5f, 0xe3, 0x57, 0x10, 0xef, 0x9d, 0xea, 0x78, 0x02, 0x41, 0xfa, 0xe2, 0x9c,
	0x79, 0x0d, 0x4f, 0x6a, 0x5e, 0x2b, 0xbd, 0xcd, 0x2a, 0x51, 0x0b, 0x4b, 0xfa, 0x09, 0x4a, 0x7b,
	0x34, 0x0e, 0xda, 0xc2, 0x49, 0xf8, 0x3d, 0x0c, 0xab, 0xfd, 0x69, 0x24, 0xf2, 0x74, 0x17, 0x5d,
	0xba, 0xff, 0x8f, 0xa7, 0xc7, 0x5e, 0x9c, 0x40, 0x2c, 0x24, 0xcb, 0xb5, 0xcc, 0xad, 0x50, 0x92,
	0x0c, 0x12, 0x94, 0x0e, 0x68, 0x57, 0xc2, 0x97, 0xf0, 0x6c, 0xb3, 0xda, 0x1a, 0xc1, 0xf2, 0x2a,
	0xdb, 0x63, 0x96, 0xdc, 0x30, 0xf2, 0xd8, 0x63, 0xe2, 0xc3, 0xbf, 0xb0, 0xe1, 0x86, 0x1b, 0x76,
	0xfd, 0x1d, 0xfd, 0xd8, 0x8d, 0xd1, 0xcf, 0xdd, 0x18, 0xfd, 0xda, 0x8d, 0xd1, 0xef, 0xdd, 0x18,
	0x01, 0x51, 0x7a, 0xd9, 0xe5, 0x39, 0xc6, 0xeb, 0xfa, 0xfc, 0xd4, 0xa2, 0x5b, 0x97, 0x2e, 0x73,
	0x8b, 0xbe, 0x7e, 0x58, 0x0a, 0xbb, 0x6a, 0x8a, 0x29, 0x53, 0xf5, 0xac, 0x13, 0xcf, 0x87, 0xcb,
	0xa5, 0x3a, 0xc9, 0xed, 0x5f, 0x84, 0x8a, 0x81, 0x4f, 0xea, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x87, 0xd2, 0xaf, 0xc4, 0xde, 0x02, 0x00, 0x00,
}
