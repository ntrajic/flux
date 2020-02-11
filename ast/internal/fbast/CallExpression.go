// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbast

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CallExpression struct {
	_tab flatbuffers.Table
}

func GetRootAsCallExpression(buf []byte, offset flatbuffers.UOffsetT) *CallExpression {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CallExpression{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *CallExpression) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CallExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CallExpression) BaseNode(obj *BaseNode) *BaseNode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BaseNode)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *CallExpression) CalleeType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CallExpression) MutateCalleeType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *CallExpression) Callee(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *CallExpression) Arguments(obj *ObjectExpression) *ObjectExpression {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ObjectExpression)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func CallExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func CallExpressionAddBaseNode(builder *flatbuffers.Builder, baseNode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(baseNode), 0)
}
func CallExpressionAddCalleeType(builder *flatbuffers.Builder, calleeType byte) {
	builder.PrependByteSlot(1, calleeType, 0)
}
func CallExpressionAddCallee(builder *flatbuffers.Builder, callee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(callee), 0)
}
func CallExpressionAddArguments(builder *flatbuffers.Builder, arguments flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(arguments), 0)
}
func CallExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}