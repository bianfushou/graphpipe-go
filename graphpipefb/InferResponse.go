// automatically generated by the FlatBuffers compiler, do not modify

package graphpipe

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InferResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsInferResponse(buf []byte, offset flatbuffers.UOffsetT) *InferResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InferResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *InferResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InferResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *InferResponse) OutputTensors(obj *Tensor, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *InferResponse) OutputTensorsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *InferResponse) Errors(obj *Error, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *InferResponse) ErrorsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func InferResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func InferResponseAddOutputTensors(builder *flatbuffers.Builder, outputTensors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(outputTensors), 0)
}
func InferResponseStartOutputTensorsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func InferResponseAddErrors(builder *flatbuffers.Builder, errors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(errors), 0)
}
func InferResponseStartErrorsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func InferResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
