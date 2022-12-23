// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: nameservice.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamResolveOption int32

const (
	// Default value. Shall never be used.
	StreamResolveOption_STREAM_RESOLVE_OPTION_UNSPECIFIED StreamResolveOption = 0
	// Resolve names in `names_subscribe` fields.
	StreamResolveOption_STREAM_RESOLVE_OPTION_DELTA StreamResolveOption = 1
	// Return resolution results for all names subscribed by this stream.
	StreamResolveOption_STREAM_RESOLVE_OPTION_ALL StreamResolveOption = 2
)

// Enum value maps for StreamResolveOption.
var (
	StreamResolveOption_name = map[int32]string{
		0: "STREAM_RESOLVE_OPTION_UNSPECIFIED",
		1: "STREAM_RESOLVE_OPTION_DELTA",
		2: "STREAM_RESOLVE_OPTION_ALL",
	}
	StreamResolveOption_value = map[string]int32{
		"STREAM_RESOLVE_OPTION_UNSPECIFIED": 0,
		"STREAM_RESOLVE_OPTION_DELTA":       1,
		"STREAM_RESOLVE_OPTION_ALL":         2,
	}
)

func (x StreamResolveOption) Enum() *StreamResolveOption {
	p := new(StreamResolveOption)
	*p = x
	return p
}

func (x StreamResolveOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamResolveOption) Descriptor() protoreflect.EnumDescriptor {
	return file_nameservice_proto_enumTypes[0].Descriptor()
}

func (StreamResolveOption) Type() protoreflect.EnumType {
	return &file_nameservice_proto_enumTypes[0]
}

func (x StreamResolveOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamResolveOption.Descriptor instead.
func (StreamResolveOption) EnumDescriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{0}
}

// Register method request.
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies id returned in `RegisterResponse`. It must be empty in the first request, and
	// non-empty in the subsequent requests.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// DNS name to register to the name service.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// IP address to back the DNS name.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Register method response.
type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An immutable request ID associated with this Register streaming call.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// TTL indicates to clients the deadline to send next RegisterRequest upon receiving this
	// RegisterResponse.
	Ttl *durationpb.Duration `protobuf:"bytes,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterResponse) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// Resolve method request.
type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DNS name to resolve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Resolve method response.
type ResolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identical to the name field in ResolveRequest.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resolved IP addresses.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// ttl indicates the maximum period the addresses are considered valid.
	Ttl *durationpb.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *ResolveResponse) Reset() {
	*x = ResolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveResponse) ProtoMessage() {}

func (x *ResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveResponse.ProtoReflect.Descriptor instead.
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResolveResponse) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *ResolveResponse) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// StreamingResolve method request.
type StreamingResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resolve option of this request.
	Option StreamResolveOption `protobuf:"varint,1,opt,name=option,proto3,enum=qeco.apis.kns.v1.StreamResolveOption" json:"option,omitempty"`
	// Names to subscribe. The callee must include the name resolution results for any name included
	// in this field.
	NamesSubscribe []string `protobuf:"bytes,2,rep,name=names_subscribe,json=namesSubscribe,proto3" json:"names_subscribe,omitempty"`
	// Names to unsubscribe.
	NamesUnsubscribes []string `protobuf:"bytes,3,rep,name=names_unsubscribes,json=namesUnsubscribes,proto3" json:"names_unsubscribes,omitempty"`
}

func (x *StreamingResolveRequest) Reset() {
	*x = StreamingResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingResolveRequest) ProtoMessage() {}

func (x *StreamingResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingResolveRequest.ProtoReflect.Descriptor instead.
func (*StreamingResolveRequest) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingResolveRequest) GetOption() StreamResolveOption {
	if x != nil {
		return x.Option
	}
	return StreamResolveOption_STREAM_RESOLVE_OPTION_UNSPECIFIED
}

func (x *StreamingResolveRequest) GetNamesSubscribe() []string {
	if x != nil {
		return x.NamesSubscribe
	}
	return nil
}

func (x *StreamingResolveRequest) GetNamesUnsubscribes() []string {
	if x != nil {
		return x.NamesUnsubscribes
	}
	return nil
}

// StreamResolve method Response.
type StreamingResolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates resolve option in corresponding StreamingResolveRequest.
	Option  StreamResolveOption `protobuf:"varint,1,opt,name=option,proto3,enum=qeco.apis.kns.v1.StreamResolveOption" json:"option,omitempty"`
	Updates []*ResolutionUpdate `protobuf:"bytes,2,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *StreamingResolveResponse) Reset() {
	*x = StreamingResolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingResolveResponse) ProtoMessage() {}

func (x *StreamingResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingResolveResponse.ProtoReflect.Descriptor instead.
func (*StreamingResolveResponse) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{5}
}

func (x *StreamingResolveResponse) GetOption() StreamResolveOption {
	if x != nil {
		return x.Option
	}
	return StreamResolveOption_STREAM_RESOLVE_OPTION_UNSPECIFIED
}

func (x *StreamingResolveResponse) GetUpdates() []*ResolutionUpdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

type ResolutionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addresses []string             `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Ttl       *durationpb.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *ResolutionUpdate) Reset() {
	*x = ResolutionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolutionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolutionUpdate) ProtoMessage() {}

func (x *ResolutionUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_nameservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolutionUpdate.ProtoReflect.Descriptor instead.
func (*ResolutionUpdate) Descriptor() ([]byte, []int) {
	return file_nameservice_proto_rawDescGZIP(), []int{6}
}

func (x *ResolutionUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResolutionUpdate) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *ResolutionUpdate) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

var File_nameservice_proto protoreflect.FileDescriptor

var file_nameservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x24, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22,
	0xb0, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x71, 0x65,
	0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x2a,
	0x7c, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a,
	0x1b, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f,
	0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x10, 0x01, 0x12, 0x1d,
	0x0a, 0x19, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45,
	0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x32, 0xa3, 0x02,
	0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x71, 0x65, 0x63, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71,
	0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12,
	0x20, 0x2e, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x29, 0x2e, 0x71, 0x65, 0x63, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6b, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x71, 0x65, 0x63, 0x6f, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nameservice_proto_rawDescOnce sync.Once
	file_nameservice_proto_rawDescData = file_nameservice_proto_rawDesc
)

func file_nameservice_proto_rawDescGZIP() []byte {
	file_nameservice_proto_rawDescOnce.Do(func() {
		file_nameservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_nameservice_proto_rawDescData)
	})
	return file_nameservice_proto_rawDescData
}

var file_nameservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nameservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nameservice_proto_goTypes = []interface{}{
	(StreamResolveOption)(0),         // 0: qeco.apis.kns.v1.StreamResolveOption
	(*RegisterRequest)(nil),          // 1: qeco.apis.kns.v1.RegisterRequest
	(*RegisterResponse)(nil),         // 2: qeco.apis.kns.v1.RegisterResponse
	(*ResolveRequest)(nil),           // 3: qeco.apis.kns.v1.ResolveRequest
	(*ResolveResponse)(nil),          // 4: qeco.apis.kns.v1.ResolveResponse
	(*StreamingResolveRequest)(nil),  // 5: qeco.apis.kns.v1.StreamingResolveRequest
	(*StreamingResolveResponse)(nil), // 6: qeco.apis.kns.v1.StreamingResolveResponse
	(*ResolutionUpdate)(nil),         // 7: qeco.apis.kns.v1.ResolutionUpdate
	(*durationpb.Duration)(nil),      // 8: google.protobuf.Duration
}
var file_nameservice_proto_depIdxs = []int32{
	8, // 0: qeco.apis.kns.v1.RegisterResponse.ttl:type_name -> google.protobuf.Duration
	8, // 1: qeco.apis.kns.v1.ResolveResponse.ttl:type_name -> google.protobuf.Duration
	0, // 2: qeco.apis.kns.v1.StreamingResolveRequest.option:type_name -> qeco.apis.kns.v1.StreamResolveOption
	0, // 3: qeco.apis.kns.v1.StreamingResolveResponse.option:type_name -> qeco.apis.kns.v1.StreamResolveOption
	7, // 4: qeco.apis.kns.v1.StreamingResolveResponse.updates:type_name -> qeco.apis.kns.v1.ResolutionUpdate
	8, // 5: qeco.apis.kns.v1.ResolutionUpdate.ttl:type_name -> google.protobuf.Duration
	1, // 6: qeco.apis.kns.v1.NameService.Register:input_type -> qeco.apis.kns.v1.RegisterRequest
	3, // 7: qeco.apis.kns.v1.NameService.Resolve:input_type -> qeco.apis.kns.v1.ResolveRequest
	5, // 8: qeco.apis.kns.v1.NameService.StreamingResolve:input_type -> qeco.apis.kns.v1.StreamingResolveRequest
	2, // 9: qeco.apis.kns.v1.NameService.Register:output_type -> qeco.apis.kns.v1.RegisterResponse
	4, // 10: qeco.apis.kns.v1.NameService.Resolve:output_type -> qeco.apis.kns.v1.ResolveResponse
	6, // 11: qeco.apis.kns.v1.NameService.StreamingResolve:output_type -> qeco.apis.kns.v1.StreamingResolveResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nameservice_proto_init() }
func file_nameservice_proto_init() {
	if File_nameservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nameservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_nameservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_nameservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
		file_nameservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveResponse); i {
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
		file_nameservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingResolveRequest); i {
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
		file_nameservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingResolveResponse); i {
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
		file_nameservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolutionUpdate); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nameservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nameservice_proto_goTypes,
		DependencyIndexes: file_nameservice_proto_depIdxs,
		EnumInfos:         file_nameservice_proto_enumTypes,
		MessageInfos:      file_nameservice_proto_msgTypes,
	}.Build()
	File_nameservice_proto = out.File
	file_nameservice_proto_rawDesc = nil
	file_nameservice_proto_goTypes = nil
	file_nameservice_proto_depIdxs = nil
}
