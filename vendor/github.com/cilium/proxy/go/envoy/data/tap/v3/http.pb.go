// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/data/tap/v3/http.proto

package tapv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A fully buffered HTTP trace message.
type HttpBufferedTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request message.
	Request *HttpBufferedTrace_Message `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Response message.
	Response *HttpBufferedTrace_Message `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	// downstream connection
	DownstreamConnection *Connection `protobuf:"bytes,3,opt,name=downstream_connection,json=downstreamConnection,proto3" json:"downstream_connection,omitempty"`
}

func (x *HttpBufferedTrace) Reset() {
	*x = HttpBufferedTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v3_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpBufferedTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpBufferedTrace) ProtoMessage() {}

func (x *HttpBufferedTrace) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v3_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpBufferedTrace.ProtoReflect.Descriptor instead.
func (*HttpBufferedTrace) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v3_http_proto_rawDescGZIP(), []int{0}
}

func (x *HttpBufferedTrace) GetRequest() *HttpBufferedTrace_Message {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *HttpBufferedTrace) GetResponse() *HttpBufferedTrace_Message {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *HttpBufferedTrace) GetDownstreamConnection() *Connection {
	if x != nil {
		return x.DownstreamConnection
	}
	return nil
}

// A streamed HTTP trace segment. Multiple segments make up a full trace.
// [#next-free-field: 8]
type HttpStreamedTraceSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Trace ID unique to the originating Envoy only. Trace IDs can repeat and should not be used
	// for long term stable uniqueness.
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are assignable to MessagePiece:
	//
	//	*HttpStreamedTraceSegment_RequestHeaders
	//	*HttpStreamedTraceSegment_RequestBodyChunk
	//	*HttpStreamedTraceSegment_RequestTrailers
	//	*HttpStreamedTraceSegment_ResponseHeaders
	//	*HttpStreamedTraceSegment_ResponseBodyChunk
	//	*HttpStreamedTraceSegment_ResponseTrailers
	MessagePiece isHttpStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
}

func (x *HttpStreamedTraceSegment) Reset() {
	*x = HttpStreamedTraceSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v3_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpStreamedTraceSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpStreamedTraceSegment) ProtoMessage() {}

func (x *HttpStreamedTraceSegment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v3_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpStreamedTraceSegment.ProtoReflect.Descriptor instead.
func (*HttpStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v3_http_proto_rawDescGZIP(), []int{1}
}

func (x *HttpStreamedTraceSegment) GetTraceId() uint64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (m *HttpStreamedTraceSegment) GetMessagePiece() isHttpStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (x *HttpStreamedTraceSegment) GetRequestHeaders() *v3.HeaderMap {
	if x, ok := x.GetMessagePiece().(*HttpStreamedTraceSegment_RequestHeaders); ok {
		return x.RequestHeaders
	}
	return nil
}

func (x *HttpStreamedTraceSegment) GetRequestBodyChunk() *Body {
	if x, ok := x.GetMessagePiece().(*HttpStreamedTraceSegment_RequestBodyChunk); ok {
		return x.RequestBodyChunk
	}
	return nil
}

func (x *HttpStreamedTraceSegment) GetRequestTrailers() *v3.HeaderMap {
	if x, ok := x.GetMessagePiece().(*HttpStreamedTraceSegment_RequestTrailers); ok {
		return x.RequestTrailers
	}
	return nil
}

func (x *HttpStreamedTraceSegment) GetResponseHeaders() *v3.HeaderMap {
	if x, ok := x.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseHeaders); ok {
		return x.ResponseHeaders
	}
	return nil
}

func (x *HttpStreamedTraceSegment) GetResponseBodyChunk() *Body {
	if x, ok := x.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseBodyChunk); ok {
		return x.ResponseBodyChunk
	}
	return nil
}

func (x *HttpStreamedTraceSegment) GetResponseTrailers() *v3.HeaderMap {
	if x, ok := x.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseTrailers); ok {
		return x.ResponseTrailers
	}
	return nil
}

type isHttpStreamedTraceSegment_MessagePiece interface {
	isHttpStreamedTraceSegment_MessagePiece()
}

type HttpStreamedTraceSegment_RequestHeaders struct {
	// Request headers.
	RequestHeaders *v3.HeaderMap `protobuf:"bytes,2,opt,name=request_headers,json=requestHeaders,proto3,oneof"`
}

type HttpStreamedTraceSegment_RequestBodyChunk struct {
	// Request body chunk.
	RequestBodyChunk *Body `protobuf:"bytes,3,opt,name=request_body_chunk,json=requestBodyChunk,proto3,oneof"`
}

type HttpStreamedTraceSegment_RequestTrailers struct {
	// Request trailers.
	RequestTrailers *v3.HeaderMap `protobuf:"bytes,4,opt,name=request_trailers,json=requestTrailers,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseHeaders struct {
	// Response headers.
	ResponseHeaders *v3.HeaderMap `protobuf:"bytes,5,opt,name=response_headers,json=responseHeaders,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseBodyChunk struct {
	// Response body chunk.
	ResponseBodyChunk *Body `protobuf:"bytes,6,opt,name=response_body_chunk,json=responseBodyChunk,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseTrailers struct {
	// Response trailers.
	ResponseTrailers *v3.HeaderMap `protobuf:"bytes,7,opt,name=response_trailers,json=responseTrailers,proto3,oneof"`
}

func (*HttpStreamedTraceSegment_RequestHeaders) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_RequestBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_RequestTrailers) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseHeaders) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseTrailers) isHttpStreamedTraceSegment_MessagePiece() {}

// HTTP message wrapper.
type HttpBufferedTrace_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message headers.
	Headers []*v3.HeaderValue `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// Message body.
	Body *Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Message trailers.
	Trailers []*v3.HeaderValue `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty"`
	// The timestamp after receiving the message headers.
	HeadersReceivedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=headers_received_time,json=headersReceivedTime,proto3" json:"headers_received_time,omitempty"`
}

func (x *HttpBufferedTrace_Message) Reset() {
	*x = HttpBufferedTrace_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v3_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpBufferedTrace_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpBufferedTrace_Message) ProtoMessage() {}

func (x *HttpBufferedTrace_Message) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v3_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpBufferedTrace_Message.ProtoReflect.Descriptor instead.
func (*HttpBufferedTrace_Message) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v3_http_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HttpBufferedTrace_Message) GetHeaders() []*v3.HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HttpBufferedTrace_Message) GetBody() *Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *HttpBufferedTrace_Message) GetTrailers() []*v3.HeaderValue {
	if x != nil {
		return x.Trailers
	}
	return nil
}

func (x *HttpBufferedTrace_Message) GetHeadersReceivedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.HeadersReceivedTime
	}
	return nil
}

var File_envoy_data_tap_v3_http_proto protoreflect.FileDescriptor

var file_envoy_data_tap_v3_http_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x70,
	0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76,
	0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74,
	0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x04, 0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x42, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x15, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0xbb, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x13, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a,
	0x2f, 0x9a, 0xc5, 0x88, 0x1e, 0x2a, 0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x22, 0xca, 0x04, 0x0a, 0x18, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d,
	0x61, 0x70, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61,
	0x70, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x4c, 0x0a,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x13, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x48,
	0x00, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x4e, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x70,
	0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x73, 0x3a, 0x36, 0x9a, 0xc5, 0x88, 0x1e, 0x31, 0x0a, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x69, 0x65, 0x63, 0x65, 0x42, 0x76, 0x0a,
	0x1f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33,
	0x42, 0x09, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x61, 0x70, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_tap_v3_http_proto_rawDescOnce sync.Once
	file_envoy_data_tap_v3_http_proto_rawDescData = file_envoy_data_tap_v3_http_proto_rawDesc
)

func file_envoy_data_tap_v3_http_proto_rawDescGZIP() []byte {
	file_envoy_data_tap_v3_http_proto_rawDescOnce.Do(func() {
		file_envoy_data_tap_v3_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_tap_v3_http_proto_rawDescData)
	})
	return file_envoy_data_tap_v3_http_proto_rawDescData
}

var file_envoy_data_tap_v3_http_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_data_tap_v3_http_proto_goTypes = []interface{}{
	(*HttpBufferedTrace)(nil),         // 0: envoy.data.tap.v3.HttpBufferedTrace
	(*HttpStreamedTraceSegment)(nil),  // 1: envoy.data.tap.v3.HttpStreamedTraceSegment
	(*HttpBufferedTrace_Message)(nil), // 2: envoy.data.tap.v3.HttpBufferedTrace.Message
	(*Connection)(nil),                // 3: envoy.data.tap.v3.Connection
	(*v3.HeaderMap)(nil),              // 4: envoy.config.core.v3.HeaderMap
	(*Body)(nil),                      // 5: envoy.data.tap.v3.Body
	(*v3.HeaderValue)(nil),            // 6: envoy.config.core.v3.HeaderValue
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_envoy_data_tap_v3_http_proto_depIdxs = []int32{
	2,  // 0: envoy.data.tap.v3.HttpBufferedTrace.request:type_name -> envoy.data.tap.v3.HttpBufferedTrace.Message
	2,  // 1: envoy.data.tap.v3.HttpBufferedTrace.response:type_name -> envoy.data.tap.v3.HttpBufferedTrace.Message
	3,  // 2: envoy.data.tap.v3.HttpBufferedTrace.downstream_connection:type_name -> envoy.data.tap.v3.Connection
	4,  // 3: envoy.data.tap.v3.HttpStreamedTraceSegment.request_headers:type_name -> envoy.config.core.v3.HeaderMap
	5,  // 4: envoy.data.tap.v3.HttpStreamedTraceSegment.request_body_chunk:type_name -> envoy.data.tap.v3.Body
	4,  // 5: envoy.data.tap.v3.HttpStreamedTraceSegment.request_trailers:type_name -> envoy.config.core.v3.HeaderMap
	4,  // 6: envoy.data.tap.v3.HttpStreamedTraceSegment.response_headers:type_name -> envoy.config.core.v3.HeaderMap
	5,  // 7: envoy.data.tap.v3.HttpStreamedTraceSegment.response_body_chunk:type_name -> envoy.data.tap.v3.Body
	4,  // 8: envoy.data.tap.v3.HttpStreamedTraceSegment.response_trailers:type_name -> envoy.config.core.v3.HeaderMap
	6,  // 9: envoy.data.tap.v3.HttpBufferedTrace.Message.headers:type_name -> envoy.config.core.v3.HeaderValue
	5,  // 10: envoy.data.tap.v3.HttpBufferedTrace.Message.body:type_name -> envoy.data.tap.v3.Body
	6,  // 11: envoy.data.tap.v3.HttpBufferedTrace.Message.trailers:type_name -> envoy.config.core.v3.HeaderValue
	7,  // 12: envoy.data.tap.v3.HttpBufferedTrace.Message.headers_received_time:type_name -> google.protobuf.Timestamp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_envoy_data_tap_v3_http_proto_init() }
func file_envoy_data_tap_v3_http_proto_init() {
	if File_envoy_data_tap_v3_http_proto != nil {
		return
	}
	file_envoy_data_tap_v3_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_tap_v3_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpBufferedTrace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v3_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpStreamedTraceSegment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_data_tap_v3_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpBufferedTrace_Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_data_tap_v3_http_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HttpStreamedTraceSegment_RequestHeaders)(nil),
		(*HttpStreamedTraceSegment_RequestBodyChunk)(nil),
		(*HttpStreamedTraceSegment_RequestTrailers)(nil),
		(*HttpStreamedTraceSegment_ResponseHeaders)(nil),
		(*HttpStreamedTraceSegment_ResponseBodyChunk)(nil),
		(*HttpStreamedTraceSegment_ResponseTrailers)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_tap_v3_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_tap_v3_http_proto_goTypes,
		DependencyIndexes: file_envoy_data_tap_v3_http_proto_depIdxs,
		MessageInfos:      file_envoy_data_tap_v3_http_proto_msgTypes,
	}.Build()
	File_envoy_data_tap_v3_http_proto = out.File
	file_envoy_data_tap_v3_http_proto_rawDesc = nil
	file_envoy_data_tap_v3_http_proto_goTypes = nil
	file_envoy_data_tap_v3_http_proto_depIdxs = nil
}
