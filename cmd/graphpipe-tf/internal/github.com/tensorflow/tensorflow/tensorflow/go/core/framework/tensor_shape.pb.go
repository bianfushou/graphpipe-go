// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/tensor_shape.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Dimensions of a tensor.
type TensorShapeProto struct {
	// Dimensions of the tensor, such as {"input", 30}, {"output", 40}
	// for a 30 x 40 2D tensor.  If an entry has size -1, this
	// corresponds to a dimension of unknown size. The names are
	// optional.
	//
	// The order of entries in "dim" matters: It indicates the layout of the
	// values in the tensor in-memory representation.
	//
	// The first entry in "dim" is the outermost dimension used to layout the
	// values, the last entry is the innermost dimension.  This matches the
	// in-memory layout of RowMajor Eigen tensors.
	//
	// If "dim.size()" > 0, "unknown_rank" must be false.
	Dim []*TensorShapeProto_Dim `protobuf:"bytes,2,rep,name=dim" json:"dim,omitempty"`
	// If true, the number of dimensions in the shape is unknown.
	//
	// If true, "dim.size()" must be 0.
	UnknownRank bool `protobuf:"varint,3,opt,name=unknown_rank,json=unknownRank,proto3" json:"unknown_rank,omitempty"`
}

func (m *TensorShapeProto) Reset()                    { *m = TensorShapeProto{} }
func (m *TensorShapeProto) String() string            { return proto.CompactTextString(m) }
func (*TensorShapeProto) ProtoMessage()               {}
func (*TensorShapeProto) Descriptor() ([]byte, []int) { return fileDescriptorTensorShape, []int{0} }

func (m *TensorShapeProto) GetDim() []*TensorShapeProto_Dim {
	if m != nil {
		return m.Dim
	}
	return nil
}

func (m *TensorShapeProto) GetUnknownRank() bool {
	if m != nil {
		return m.UnknownRank
	}
	return false
}

// One dimension of the tensor.
type TensorShapeProto_Dim struct {
	// Size of the tensor in that dimension.
	// This value must be >= -1, but values of -1 are reserved for "unknown"
	// shapes (values of -1 mean "unknown" dimension).  Certain wrappers
	// that work with TensorShapeProto may fail at runtime when deserializing
	// a TensorShapeProto containing a dim value of -1.
	Size_ int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Optional name of the tensor dimension.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *TensorShapeProto_Dim) Reset()         { *m = TensorShapeProto_Dim{} }
func (m *TensorShapeProto_Dim) String() string { return proto.CompactTextString(m) }
func (*TensorShapeProto_Dim) ProtoMessage()    {}
func (*TensorShapeProto_Dim) Descriptor() ([]byte, []int) {
	return fileDescriptorTensorShape, []int{0, 0}
}

func (m *TensorShapeProto_Dim) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *TensorShapeProto_Dim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*TensorShapeProto)(nil), "tensorflow.TensorShapeProto")
	proto.RegisterType((*TensorShapeProto_Dim)(nil), "tensorflow.TensorShapeProto.Dim")
}
func (m *TensorShapeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TensorShapeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dim) > 0 {
		for _, msg := range m.Dim {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTensorShape(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.UnknownRank {
		dAtA[i] = 0x18
		i++
		if m.UnknownRank {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *TensorShapeProto_Dim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TensorShapeProto_Dim) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Size_ != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensorShape(dAtA, i, uint64(m.Size_))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTensorShape(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func encodeVarintTensorShape(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TensorShapeProto) Size() (n int) {
	var l int
	_ = l
	if len(m.Dim) > 0 {
		for _, e := range m.Dim {
			l = e.Size()
			n += 1 + l + sovTensorShape(uint64(l))
		}
	}
	if m.UnknownRank {
		n += 2
	}
	return n
}

func (m *TensorShapeProto_Dim) Size() (n int) {
	var l int
	_ = l
	if m.Size_ != 0 {
		n += 1 + sovTensorShape(uint64(m.Size_))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTensorShape(uint64(l))
	}
	return n
}

func sovTensorShape(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTensorShape(x uint64) (n int) {
	return sovTensorShape(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TensorShapeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorShape
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
			return fmt.Errorf("proto: TensorShapeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TensorShapeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
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
				return ErrInvalidLengthTensorShape
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dim = append(m.Dim, &TensorShapeProto_Dim{})
			if err := m.Dim[len(m.Dim)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnknownRank", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UnknownRank = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTensorShape(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorShape
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
func (m *TensorShapeProto_Dim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorShape
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
			return fmt.Errorf("proto: Dim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
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
				return ErrInvalidLengthTensorShape
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorShape(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorShape
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
func skipTensorShape(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensorShape
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
					return 0, ErrIntOverflowTensorShape
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
					return 0, ErrIntOverflowTensorShape
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
				return 0, ErrInvalidLengthTensorShape
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTensorShape
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
				next, err := skipTensorShape(dAtA[start:])
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
	ErrInvalidLengthTensorShape = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensorShape   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/tensor_shape.proto", fileDescriptorTensorShape)
}

var fileDescriptorTensorShape = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0xd6, 0x87, 0xc8, 0xc4, 0x17, 0x67, 0x24, 0x16, 0xa4, 0xea, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x54, 0x2b, 0xcd, 0x60, 0xe4, 0x12, 0x08, 0x01, 0x73,
	0x83, 0x41, 0x2a, 0x02, 0xc0, 0x0a, 0x8c, 0xb8, 0x98, 0x53, 0x32, 0x73, 0x25, 0x98, 0x14, 0x98,
	0x35, 0xb8, 0x8d, 0x14, 0xf4, 0x10, 0xca, 0xf5, 0xd0, 0x95, 0xea, 0xb9, 0x64, 0xe6, 0x06, 0x81,
	0x14, 0x0b, 0x29, 0x72, 0xf1, 0x94, 0xe6, 0x65, 0xe7, 0xe5, 0x97, 0xe7, 0xc5, 0x17, 0x25, 0xe6,
	0x65, 0x4b, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0x71, 0x43, 0xc5, 0x82, 0x12, 0xf3, 0xb2, 0xa5,
	0x74, 0xb9, 0x98, 0x5d, 0x32, 0x73, 0x85, 0x84, 0xb8, 0x58, 0x8a, 0x33, 0xab, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x98, 0xed, 0xd4, 0xc8, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x37, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xc8, 0x25, 0x91, 0x5f, 0x94, 0x8e, 0xec,
	0x1c, 0xb8, 0x37, 0x9d, 0x04, 0xd1, 0x5d, 0x56, 0x1c, 0xc0, 0x18, 0x65, 0x9b, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0x14, 0x42, 0xd8, 0x99, 0xe9, 0xf9, 0x68, 0x41,
	0xf7, 0x83, 0x91, 0x31, 0x89, 0x0d, 0x1c, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x18, 0xf8, 0xcd, 0x61, 0x01, 0x00, 0x00,
}
