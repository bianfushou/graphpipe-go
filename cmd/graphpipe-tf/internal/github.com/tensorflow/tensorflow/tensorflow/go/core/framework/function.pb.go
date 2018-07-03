// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/function.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function []*FunctionDef `protobuf:"bytes,1,rep,name=function" json:"function,omitempty"`
	Gradient []*GradientDef `protobuf:"bytes,2,rep,name=gradient" json:"gradient,omitempty"`
}

func (m *FunctionDefLibrary) Reset()                    { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string            { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()               {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) { return fileDescriptorFunction, []int{0} }

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *FunctionDef) Reset()                    { *m = FunctionDef{} }
func (m *FunctionDef) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()               {}
func (*FunctionDef) Descriptor() ([]byte, []int) { return fileDescriptorFunction, []int{1} }

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc,proto3" json:"gradient_func,omitempty"`
}

func (m *GradientDef) Reset()                    { *m = GradientDef{} }
func (m *GradientDef) String() string            { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()               {}
func (*GradientDef) Descriptor() ([]byte, []int) { return fileDescriptorFunction, []int{2} }

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}
func (m *FunctionDefLibrary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionDefLibrary) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Function) > 0 {
		for _, msg := range m.Function {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFunction(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Gradient) > 0 {
		for _, msg := range m.Gradient {
			dAtA[i] = 0x12
			i++
			i = encodeVarintFunction(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *FunctionDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFunction(dAtA, i, uint64(m.Signature.Size()))
		n1, err := m.Signature.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.NodeDef) > 0 {
		for _, msg := range m.NodeDef {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintFunction(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Ret) > 0 {
		keysForRet := make([]string, 0, len(m.Ret))
		for k, _ := range m.Ret {
			keysForRet = append(keysForRet, string(k))
		}
		sortkeys.Strings(keysForRet)
		for _, k := range keysForRet {
			dAtA[i] = 0x22
			i++
			v := m.Ret[string(k)]
			mapSize := 1 + len(k) + sovFunction(uint64(len(k))) + 1 + len(v) + sovFunction(uint64(len(v)))
			i = encodeVarintFunction(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintFunction(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintFunction(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Attr) > 0 {
		keysForAttr := make([]string, 0, len(m.Attr))
		for k, _ := range m.Attr {
			keysForAttr = append(keysForAttr, string(k))
		}
		sortkeys.Strings(keysForAttr)
		for _, k := range keysForAttr {
			dAtA[i] = 0x2a
			i++
			v := m.Attr[string(k)]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovFunction(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovFunction(uint64(len(k))) + msgSize
			i = encodeVarintFunction(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintFunction(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintFunction(dAtA, i, uint64(v.Size()))
				n2, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n2
			}
		}
	}
	return i, nil
}

func (m *GradientDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GradientDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FunctionName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFunction(dAtA, i, uint64(len(m.FunctionName)))
		i += copy(dAtA[i:], m.FunctionName)
	}
	if len(m.GradientFunc) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFunction(dAtA, i, uint64(len(m.GradientFunc)))
		i += copy(dAtA[i:], m.GradientFunc)
	}
	return i, nil
}

func encodeVarintFunction(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FunctionDefLibrary) Size() (n int) {
	var l int
	_ = l
	if len(m.Function) > 0 {
		for _, e := range m.Function {
			l = e.Size()
			n += 1 + l + sovFunction(uint64(l))
		}
	}
	if len(m.Gradient) > 0 {
		for _, e := range m.Gradient {
			l = e.Size()
			n += 1 + l + sovFunction(uint64(l))
		}
	}
	return n
}

func (m *FunctionDef) Size() (n int) {
	var l int
	_ = l
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovFunction(uint64(l))
	}
	if len(m.NodeDef) > 0 {
		for _, e := range m.NodeDef {
			l = e.Size()
			n += 1 + l + sovFunction(uint64(l))
		}
	}
	if len(m.Ret) > 0 {
		for k, v := range m.Ret {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovFunction(uint64(len(k))) + 1 + len(v) + sovFunction(uint64(len(v)))
			n += mapEntrySize + 1 + sovFunction(uint64(mapEntrySize))
		}
	}
	if len(m.Attr) > 0 {
		for k, v := range m.Attr {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovFunction(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovFunction(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovFunction(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *GradientDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.FunctionName)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	l = len(m.GradientFunc)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	return n
}

func sovFunction(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFunction(x uint64) (n int) {
	return sovFunction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FunctionDefLibrary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: FunctionDefLibrary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionDefLibrary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Function", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Function = append(m.Function, &FunctionDef{})
			if err := m.Function[len(m.Function)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gradient", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gradient = append(m.Gradient, &GradientDef{})
			if err := m.Gradient[len(m.Gradient)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFunction
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
func (m *FunctionDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: FunctionDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &OpDef{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeDef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeDef = append(m.NodeDef, &NodeDef{})
			if err := m.NodeDef[len(m.NodeDef)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ret == nil {
				m.Ret = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFunction
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFunction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthFunction
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFunction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthFunction
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFunction(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFunction
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Ret[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attr == nil {
				m.Attr = make(map[string]*AttrValue)
			}
			var mapkey string
			var mapvalue *AttrValue
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFunction
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFunction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthFunction
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFunction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthFunction
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthFunction
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &AttrValue{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFunction(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFunction
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Attr[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFunction
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
func (m *GradientDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: GradientDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GradientDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunctionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunctionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GradientFunc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
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
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GradientFunc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFunction
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
func skipFunction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFunction
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
					return 0, ErrIntOverflowFunction
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
					return 0, ErrIntOverflowFunction
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
				return 0, ErrInvalidLengthFunction
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFunction
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
				next, err := skipFunction(dAtA[start:])
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
	ErrInvalidLengthFunction = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFunction   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensorflow/core/framework/function.proto", fileDescriptorFunction) }

var fileDescriptorFunction = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x99, 0xa4, 0xab, 0xed, 0xeb, 0x2a, 0xeb, 0xa8, 0x38, 0xf4, 0x10, 0x6a, 0x05, 0x29,
	0x0a, 0x09, 0x74, 0x51, 0x44, 0xf0, 0xe0, 0xe2, 0x0f, 0x10, 0xa9, 0x4b, 0x0e, 0x0a, 0x5e, 0xc2,
	0xb4, 0x99, 0xc4, 0xb0, 0x9b, 0x99, 0x32, 0x9d, 0xb8, 0xf4, 0xb2, 0x7f, 0x9f, 0x47, 0x8f, 0x1e,
	0x25, 0x17, 0xff, 0x05, 0x8f, 0x32, 0x93, 0x4c, 0x33, 0xe8, 0xc6, 0xdb, 0x30, 0xf3, 0xf9, 0x7c,
	0xf3, 0x5e, 0xde, 0x83, 0xb9, 0x62, 0x7c, 0x2b, 0x64, 0x76, 0x2e, 0x2e, 0xa2, 0xb5, 0x90, 0x2c,
	0xca, 0x24, 0x2d, 0xd9, 0x85, 0x90, 0x67, 0x51, 0x56, 0xf1, 0xb5, 0x2a, 0x04, 0x0f, 0x37, 0x52,
	0x28, 0x81, 0xa1, 0x23, 0x27, 0x8f, 0xfa, 0x2d, 0xaa, 0x94, 0x4c, 0xbe, 0xd2, 0xf3, 0x8a, 0x35,
	0xde, 0xe4, 0x3f, 0x5f, 0xe0, 0x22, 0x65, 0x49, 0xca, 0xb2, 0x96, 0x7c, 0xd8, 0x4f, 0x8a, 0x4d,
	0xc7, 0xcd, 0x2e, 0x01, 0xbf, 0x69, 0x6b, 0x7b, 0xc5, 0xb2, 0xf7, 0xc5, 0x4a, 0x52, 0xb9, 0xc3,
	0xc7, 0x30, 0xb4, 0x15, 0x13, 0x34, 0xf5, 0xe7, 0xe3, 0xc5, 0xbd, 0xb0, 0x0b, 0x0c, 0x1d, 0x23,
	0xde, 0x83, 0x5a, 0xca, 0x25, 0x4d, 0x0b, 0xc6, 0x15, 0xf1, 0xfe, 0x95, 0xde, 0xb6, 0x6f, 0x46,
	0xb2, 0xe0, 0xec, 0x97, 0x07, 0x63, 0x27, 0x0e, 0x47, 0x30, 0xda, 0x16, 0x39, 0xa7, 0xaa, 0x92,
	0x8c, 0xa0, 0x29, 0x9a, 0x8f, 0x17, 0xb7, 0xdc, 0x94, 0x0f, 0x1b, 0xed, 0x77, 0x0c, 0x0e, 0x61,
	0x68, 0x5b, 0x27, 0xbe, 0xf9, 0xea, 0x6d, 0x97, 0x5f, 0x8a, 0x94, 0x69, 0xe3, 0x3a, 0x6f, 0x0e,
	0x78, 0x01, 0xbe, 0x64, 0x8a, 0x0c, 0x0c, 0x3a, 0xed, 0xe9, 0x2a, 0x8c, 0x99, 0x7a, 0xcd, 0x95,
	0xdc, 0xc5, 0x1a, 0xc6, 0x4f, 0x60, 0xa0, 0x47, 0x41, 0x0e, 0x8c, 0x74, 0xbf, 0x4f, 0x7a, 0xa9,
	0x94, 0x6c, 0x2c, 0x83, 0x4f, 0x96, 0x30, 0xda, 0x5f, 0xe1, 0x23, 0xf0, 0xcf, 0xd8, 0xce, 0xb4,
	0x34, 0x8a, 0xf5, 0x11, 0x3f, 0x86, 0x03, 0x33, 0x5b, 0xe2, 0x99, 0x36, 0xef, 0xba, 0xb1, 0xda,
	0xfb, 0xa8, 0x1f, 0xe3, 0x86, 0x79, 0xee, 0x3d, 0x43, 0x93, 0xa7, 0x30, 0xb4, 0x75, 0x5d, 0x11,
	0x77, 0xc7, 0x8d, 0x1b, 0x39, 0xde, 0xbb, 0xc1, 0xd0, 0x3b, 0xf2, 0x67, 0x9f, 0x60, 0xec, 0x8c,
	0x00, 0x3f, 0x80, 0x1b, 0x76, 0x72, 0x09, 0xa7, 0x25, 0x6b, 0xa3, 0x0e, 0xed, 0xe5, 0x92, 0x96,
	0x4c, 0x43, 0x76, 0x52, 0x89, 0x7e, 0x68, 0xb3, 0x0f, 0xed, 0xa5, 0xee, 0xfe, 0xe4, 0xf2, 0x5b,
	0x1d, 0xa0, 0xef, 0x75, 0x80, 0x7e, 0xd4, 0x01, 0xfa, 0x59, 0x07, 0x08, 0x88, 0x90, 0xb9, 0xdb,
	0xcd, 0x7e, 0xf7, 0x4e, 0x6e, 0xda, 0xff, 0x75, 0xaa, 0xb7, 0x6f, 0x7b, 0x8a, 0x3e, 0xbf, 0xc8,
	0x0b, 0xf5, 0xa5, 0x5a, 0x85, 0x6b, 0x51, 0x46, 0xce, 0xce, 0x5e, 0x7d, 0xcc, 0xc5, 0x5f, 0xcb,
	0xfc, 0x1b, 0xa1, 0xd5, 0x35, 0xb3, 0xc9, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xd8,
	0xce, 0x37, 0x7f, 0x03, 0x00, 0x00,
}
