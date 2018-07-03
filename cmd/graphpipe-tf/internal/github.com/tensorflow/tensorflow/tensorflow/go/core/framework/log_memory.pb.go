// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/log_memory.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MemoryLogStep struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Handle describing the feeds and fetches of the step.
	Handle string `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"`
}

func (m *MemoryLogStep) Reset()                    { *m = MemoryLogStep{} }
func (m *MemoryLogStep) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogStep) ProtoMessage()               {}
func (*MemoryLogStep) Descriptor() ([]byte, []int) { return fileDescriptorLogMemory, []int{0} }

func (m *MemoryLogStep) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogStep) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

type MemoryLogTensorAllocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the kernel making the allocation as set in GraphDef,
	// e.g., "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName,proto3" json:"kernel_name,omitempty"`
	// Allocated tensor details.
	Tensor *TensorDescription `protobuf:"bytes,3,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *MemoryLogTensorAllocation) Reset()         { *m = MemoryLogTensorAllocation{} }
func (m *MemoryLogTensorAllocation) String() string { return proto.CompactTextString(m) }
func (*MemoryLogTensorAllocation) ProtoMessage()    {}
func (*MemoryLogTensorAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptorLogMemory, []int{1}
}

func (m *MemoryLogTensorAllocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogTensorAllocation) GetKernelName() string {
	if m != nil {
		return m.KernelName
	}
	return ""
}

func (m *MemoryLogTensorAllocation) GetTensor() *TensorDescription {
	if m != nil {
		return m.Tensor
	}
	return nil
}

type MemoryLogTensorDeallocation struct {
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,1,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,2,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
}

func (m *MemoryLogTensorDeallocation) Reset()         { *m = MemoryLogTensorDeallocation{} }
func (m *MemoryLogTensorDeallocation) String() string { return proto.CompactTextString(m) }
func (*MemoryLogTensorDeallocation) ProtoMessage()    {}
func (*MemoryLogTensorDeallocation) Descriptor() ([]byte, []int) {
	return fileDescriptorLogMemory, []int{2}
}

func (m *MemoryLogTensorDeallocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogTensorDeallocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

type MemoryLogTensorOutput struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the kernel producing an output as set in GraphDef, e.g.,
	// "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName,proto3" json:"kernel_name,omitempty"`
	// Index of the output being set.
	Index int32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// Output tensor details.
	Tensor *TensorDescription `protobuf:"bytes,4,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *MemoryLogTensorOutput) Reset()                    { *m = MemoryLogTensorOutput{} }
func (m *MemoryLogTensorOutput) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorOutput) ProtoMessage()               {}
func (*MemoryLogTensorOutput) Descriptor() ([]byte, []int) { return fileDescriptorLogMemory, []int{3} }

func (m *MemoryLogTensorOutput) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogTensorOutput) GetKernelName() string {
	if m != nil {
		return m.KernelName
	}
	return ""
}

func (m *MemoryLogTensorOutput) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MemoryLogTensorOutput) GetTensor() *TensorDescription {
	if m != nil {
		return m.Tensor
	}
	return nil
}

type MemoryLogRawAllocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the operation making the allocation.
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Number of bytes in the allocation.
	NumBytes int64 `protobuf:"varint,3,opt,name=num_bytes,json=numBytes,proto3" json:"num_bytes,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,4,opt,name=ptr,proto3" json:"ptr,omitempty"`
	// Id of the tensor buffer being allocated, used to match to a
	// corresponding deallocation.
	AllocationId int64 `protobuf:"varint,5,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,6,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
}

func (m *MemoryLogRawAllocation) Reset()                    { *m = MemoryLogRawAllocation{} }
func (m *MemoryLogRawAllocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogRawAllocation) ProtoMessage()               {}
func (*MemoryLogRawAllocation) Descriptor() ([]byte, []int) { return fileDescriptorLogMemory, []int{4} }

func (m *MemoryLogRawAllocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MemoryLogRawAllocation) GetNumBytes() int64 {
	if m != nil {
		return m.NumBytes
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

type MemoryLogRawDeallocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the operation making the deallocation.
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,3,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,4,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// True if the deallocation is queued and will be performed later,
	// e.g. for GPU lazy freeing of buffers.
	Deferred bool `protobuf:"varint,5,opt,name=deferred,proto3" json:"deferred,omitempty"`
}

func (m *MemoryLogRawDeallocation) Reset()         { *m = MemoryLogRawDeallocation{} }
func (m *MemoryLogRawDeallocation) String() string { return proto.CompactTextString(m) }
func (*MemoryLogRawDeallocation) ProtoMessage()    {}
func (*MemoryLogRawDeallocation) Descriptor() ([]byte, []int) {
	return fileDescriptorLogMemory, []int{5}
}

func (m *MemoryLogRawDeallocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogRawDeallocation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MemoryLogRawDeallocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogRawDeallocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *MemoryLogRawDeallocation) GetDeferred() bool {
	if m != nil {
		return m.Deferred
	}
	return false
}

func init() {
	proto.RegisterType((*MemoryLogStep)(nil), "tensorflow.MemoryLogStep")
	proto.RegisterType((*MemoryLogTensorAllocation)(nil), "tensorflow.MemoryLogTensorAllocation")
	proto.RegisterType((*MemoryLogTensorDeallocation)(nil), "tensorflow.MemoryLogTensorDeallocation")
	proto.RegisterType((*MemoryLogTensorOutput)(nil), "tensorflow.MemoryLogTensorOutput")
	proto.RegisterType((*MemoryLogRawAllocation)(nil), "tensorflow.MemoryLogRawAllocation")
	proto.RegisterType((*MemoryLogRawDeallocation)(nil), "tensorflow.MemoryLogRawDeallocation")
}
func (m *MemoryLogStep) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemoryLogStep) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StepId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.StepId))
	}
	if len(m.Handle) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.Handle)))
		i += copy(dAtA[i:], m.Handle)
	}
	return i, nil
}

func (m *MemoryLogTensorAllocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemoryLogTensorAllocation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StepId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.StepId))
	}
	if len(m.KernelName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.KernelName)))
		i += copy(dAtA[i:], m.KernelName)
	}
	if m.Tensor != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.Tensor.Size()))
		n1, err := m.Tensor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *MemoryLogTensorDeallocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemoryLogTensorDeallocation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AllocationId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.AllocationId))
	}
	if len(m.AllocatorName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.AllocatorName)))
		i += copy(dAtA[i:], m.AllocatorName)
	}
	return i, nil
}

func (m *MemoryLogTensorOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemoryLogTensorOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StepId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.StepId))
	}
	if len(m.KernelName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.KernelName)))
		i += copy(dAtA[i:], m.KernelName)
	}
	if m.Index != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.Index))
	}
	if m.Tensor != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.Tensor.Size()))
		n2, err := m.Tensor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *MemoryLogRawAllocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemoryLogRawAllocation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StepId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.StepId))
	}
	if len(m.Operation) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.Operation)))
		i += copy(dAtA[i:], m.Operation)
	}
	if m.NumBytes != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.NumBytes))
	}
	if m.Ptr != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.Ptr))
	}
	if m.AllocationId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.AllocationId))
	}
	if len(m.AllocatorName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.AllocatorName)))
		i += copy(dAtA[i:], m.AllocatorName)
	}
	return i, nil
}

func (m *MemoryLogRawDeallocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemoryLogRawDeallocation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StepId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.StepId))
	}
	if len(m.Operation) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.Operation)))
		i += copy(dAtA[i:], m.Operation)
	}
	if m.AllocationId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(m.AllocationId))
	}
	if len(m.AllocatorName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogMemory(dAtA, i, uint64(len(m.AllocatorName)))
		i += copy(dAtA[i:], m.AllocatorName)
	}
	if m.Deferred {
		dAtA[i] = 0x28
		i++
		if m.Deferred {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintLogMemory(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MemoryLogStep) Size() (n int) {
	var l int
	_ = l
	if m.StepId != 0 {
		n += 1 + sovLogMemory(uint64(m.StepId))
	}
	l = len(m.Handle)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	return n
}

func (m *MemoryLogTensorAllocation) Size() (n int) {
	var l int
	_ = l
	if m.StepId != 0 {
		n += 1 + sovLogMemory(uint64(m.StepId))
	}
	l = len(m.KernelName)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	if m.Tensor != nil {
		l = m.Tensor.Size()
		n += 1 + l + sovLogMemory(uint64(l))
	}
	return n
}

func (m *MemoryLogTensorDeallocation) Size() (n int) {
	var l int
	_ = l
	if m.AllocationId != 0 {
		n += 1 + sovLogMemory(uint64(m.AllocationId))
	}
	l = len(m.AllocatorName)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	return n
}

func (m *MemoryLogTensorOutput) Size() (n int) {
	var l int
	_ = l
	if m.StepId != 0 {
		n += 1 + sovLogMemory(uint64(m.StepId))
	}
	l = len(m.KernelName)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovLogMemory(uint64(m.Index))
	}
	if m.Tensor != nil {
		l = m.Tensor.Size()
		n += 1 + l + sovLogMemory(uint64(l))
	}
	return n
}

func (m *MemoryLogRawAllocation) Size() (n int) {
	var l int
	_ = l
	if m.StepId != 0 {
		n += 1 + sovLogMemory(uint64(m.StepId))
	}
	l = len(m.Operation)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	if m.NumBytes != 0 {
		n += 1 + sovLogMemory(uint64(m.NumBytes))
	}
	if m.Ptr != 0 {
		n += 1 + sovLogMemory(uint64(m.Ptr))
	}
	if m.AllocationId != 0 {
		n += 1 + sovLogMemory(uint64(m.AllocationId))
	}
	l = len(m.AllocatorName)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	return n
}

func (m *MemoryLogRawDeallocation) Size() (n int) {
	var l int
	_ = l
	if m.StepId != 0 {
		n += 1 + sovLogMemory(uint64(m.StepId))
	}
	l = len(m.Operation)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	if m.AllocationId != 0 {
		n += 1 + sovLogMemory(uint64(m.AllocationId))
	}
	l = len(m.AllocatorName)
	if l > 0 {
		n += 1 + l + sovLogMemory(uint64(l))
	}
	if m.Deferred {
		n += 2
	}
	return n
}

func sovLogMemory(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogMemory(x uint64) (n int) {
	return sovLogMemory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MemoryLogStep) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogMemory
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
			return fmt.Errorf("proto: MemoryLogStep: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemoryLogStep: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepId", wireType)
			}
			m.StepId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogMemory
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
func (m *MemoryLogTensorAllocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogMemory
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
			return fmt.Errorf("proto: MemoryLogTensorAllocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemoryLogTensorAllocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepId", wireType)
			}
			m.StepId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tensor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tensor == nil {
				m.Tensor = &TensorDescription{}
			}
			if err := m.Tensor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogMemory
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
func (m *MemoryLogTensorDeallocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogMemory
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
			return fmt.Errorf("proto: MemoryLogTensorDeallocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemoryLogTensorDeallocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationId", wireType)
			}
			m.AllocationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocationId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogMemory
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
func (m *MemoryLogTensorOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogMemory
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
			return fmt.Errorf("proto: MemoryLogTensorOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemoryLogTensorOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepId", wireType)
			}
			m.StepId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tensor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tensor == nil {
				m.Tensor = &TensorDescription{}
			}
			if err := m.Tensor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogMemory
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
func (m *MemoryLogRawAllocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogMemory
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
			return fmt.Errorf("proto: MemoryLogRawAllocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemoryLogRawAllocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepId", wireType)
			}
			m.StepId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBytes", wireType)
			}
			m.NumBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ptr", wireType)
			}
			m.Ptr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ptr |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationId", wireType)
			}
			m.AllocationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocationId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogMemory
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
func (m *MemoryLogRawDeallocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogMemory
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
			return fmt.Errorf("proto: MemoryLogRawDeallocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemoryLogRawDeallocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepId", wireType)
			}
			m.StepId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationId", wireType)
			}
			m.AllocationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocationId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
				return ErrInvalidLengthLogMemory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deferred", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogMemory
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
			m.Deferred = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLogMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogMemory
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
func skipLogMemory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogMemory
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
					return 0, ErrIntOverflowLogMemory
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
					return 0, ErrIntOverflowLogMemory
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
				return 0, ErrInvalidLengthLogMemory
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogMemory
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
				next, err := skipLogMemory(dAtA[start:])
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
	ErrInvalidLengthLogMemory = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogMemory   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensorflow/core/framework/log_memory.proto", fileDescriptorLogMemory) }

var fileDescriptorLogMemory = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x38, 0x31, 0xc9, 0x94, 0x00, 0xb2, 0xa0, 0x98, 0x16, 0x42, 0x14, 0x84, 0x14,
	0x71, 0x48, 0xa4, 0x22, 0x8e, 0x48, 0x10, 0xf5, 0x52, 0xa9, 0x40, 0x65, 0x38, 0x71, 0xb1, 0x9c,
	0x78, 0xe2, 0x5a, 0xf5, 0xee, 0x58, 0xeb, 0xb5, 0x42, 0x4f, 0xbc, 0x00, 0xcf, 0xc0, 0x7b, 0xf0,
	0x06, 0x88, 0x13, 0x47, 0x8e, 0x28, 0x4f, 0xc1, 0x11, 0xd9, 0x6b, 0xbc, 0x26, 0x4d, 0xa5, 0x00,
	0x37, 0xff, 0xe3, 0xd9, 0x99, 0xef, 0x9f, 0xd1, 0xc0, 0x63, 0x85, 0x22, 0x23, 0xb9, 0x48, 0x68,
	0x39, 0x99, 0x93, 0xc4, 0xc9, 0x42, 0x06, 0x1c, 0x97, 0x24, 0xcf, 0x26, 0x09, 0x45, 0x3e, 0x47,
	0x4e, 0xf2, 0x7c, 0x9c, 0x4a, 0x52, 0xe4, 0x80, 0xc9, 0xdd, 0x3b, 0xb8, 0xfc, 0x9d, 0xfe, 0xe3,
	0x87, 0x98, 0xcd, 0x65, 0x9c, 0xaa, 0x98, 0x84, 0x7e, 0x3f, 0x7c, 0x0e, 0xbd, 0x97, 0x65, 0xbd,
	0x63, 0x8a, 0xde, 0x28, 0x4c, 0x9d, 0x3b, 0x70, 0x35, 0x53, 0x98, 0xfa, 0x71, 0xe8, 0xb2, 0x01,
	0x1b, 0x59, 0x9e, 0x5d, 0xc8, 0xa3, 0xd0, 0xd9, 0x05, 0xfb, 0x34, 0x10, 0x61, 0x82, 0xee, 0x95,
	0x01, 0x1b, 0x75, 0xbd, 0x4a, 0x0d, 0x3f, 0x32, 0xb8, 0x5b, 0x97, 0x78, 0x5b, 0xf6, 0x79, 0x91,
	0x24, 0x34, 0x0f, 0x8a, 0x2e, 0x97, 0x97, 0x7b, 0x00, 0x3b, 0x67, 0x28, 0x05, 0x26, 0xbe, 0x08,
	0xf8, 0xef, 0x9a, 0xa0, 0x43, 0xaf, 0x02, 0x8e, 0xce, 0x53, 0xb0, 0x35, 0xb5, 0x6b, 0x0d, 0xd8,
	0x68, 0xe7, 0xe0, 0xfe, 0xd8, 0xd8, 0x1b, 0xeb, 0x3e, 0x87, 0xc6, 0x8e, 0x57, 0x25, 0x0f, 0x63,
	0xd8, 0x5f, 0xa3, 0x39, 0xc4, 0xc0, 0xf0, 0x3c, 0x84, 0x9e, 0x51, 0x86, 0xea, 0x9a, 0x09, 0x1e,
	0x85, 0xce, 0x23, 0xb8, 0x5e, 0x69, 0x92, 0x4d, 0xbc, 0x5e, 0x1d, 0x2d, 0x08, 0x87, 0x9f, 0x18,
	0xdc, 0x5e, 0xeb, 0xf5, 0x3a, 0x57, 0x69, 0xae, 0xfe, 0xc3, 0xf5, 0x2d, 0x68, 0xc7, 0x22, 0xc4,
	0xf7, 0xa5, 0xe9, 0xb6, 0xa7, 0x45, 0x63, 0x16, 0xad, 0xbf, 0x99, 0xc5, 0x57, 0x06, 0xbb, 0x35,
	0xa0, 0x17, 0x2c, 0xb7, 0xd9, 0xcb, 0x3d, 0xe8, 0x52, 0x8a, 0xb2, 0xcc, 0xaa, 0xf8, 0x4c, 0xc0,
	0xd9, 0x87, 0xae, 0xc8, 0xb9, 0x3f, 0x3b, 0x57, 0x98, 0x95, 0x88, 0x96, 0xd7, 0x11, 0x39, 0x9f,
	0x16, 0xda, 0xb9, 0x09, 0x56, 0xaa, 0x34, 0x62, 0xcb, 0x2b, 0x3e, 0x2f, 0x4e, 0xbb, 0xbd, 0xd5,
	0xb4, 0xed, 0x4d, 0xd3, 0xfe, 0xcc, 0xc0, 0x6d, 0x9a, 0xf9, 0x63, 0xad, 0xff, 0x68, 0xe7, 0x02,
	0x9f, 0xb5, 0x15, 0x5f, 0x6b, 0x03, 0x9f, 0xb3, 0x07, 0x9d, 0x10, 0x17, 0x28, 0x25, 0x6a, 0x9b,
	0x1d, 0xaf, 0xd6, 0xd3, 0x0f, 0x5f, 0x56, 0x7d, 0xf6, 0x6d, 0xd5, 0x67, 0xdf, 0x57, 0x7d, 0xf6,
	0x63, 0xd5, 0x67, 0xe0, 0x92, 0x8c, 0x9a, 0x4b, 0xac, 0x4f, 0x75, 0x7a, 0xe3, 0x98, 0x22, 0xed,
	0xf3, 0xa4, 0xb8, 0xd0, 0xec, 0x84, 0xbd, 0x7b, 0x16, 0xc5, 0xea, 0x34, 0x9f, 0x8d, 0xe7, 0xc4,
	0x27, 0x8d, 0x1b, 0xdf, 0xfc, 0x19, 0xd1, 0xda, 0xf1, 0xff, 0x64, 0x6c, 0x66, 0x97, 0xd7, 0xfe,
	0xe4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x52, 0x02, 0x07, 0x5b, 0x04, 0x00, 0x00,
}
