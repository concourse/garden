// Code generated by protoc-gen-gogo.
// source: attach.proto
// DO NOT EDIT!

/*
Package garden is a generated protocol buffer package.

It is generated from these files:
	attach.proto
	destroy.proto
	environment_variable.proto
	error.proto
	get_property.proto
	info.proto
	limit_bandwidth.proto
	limit_cpu.proto
	limit_disk.proto
	limit_memory.proto
	list.proto
	message.proto
	net_in.proto
	net_out.proto
	process_payload.proto
	property.proto
	remove_property.proto
	resource_limits.proto
	run.proto
	set_property.proto
	stop.proto
	stream_in.proto
	stream_out.proto
	tty.proto

It has these top-level messages:
	AttachRequest
*/
package garden

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type AttachRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	ProcessId        *uint32 `protobuf:"varint,2,req,name=process_id" json:"process_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AttachRequest) Reset()         { *m = AttachRequest{} }
func (m *AttachRequest) String() string { return proto.CompactTextString(m) }
func (*AttachRequest) ProtoMessage()    {}

func (m *AttachRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *AttachRequest) GetProcessId() uint32 {
	if m != nil && m.ProcessId != nil {
		return *m.ProcessId
	}
	return 0
}

func init() {
}
