// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: internal/pkg/state/state.proto

package state

import (
	messages "github.com/Reallife/pbuf-example/internal/pkg/messages"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts          map[string]*ListOfPosts   `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // key = account
	ScheduledPosts []*messages.Post          `protobuf:"bytes,2,rep,name=scheduledPosts,proto3" json:"scheduledPosts,omitempty"`
	DirectMessages map[string]*ListOfDirects `protobuf:"bytes,3,rep,name=directMessages,proto3" json:"directMessages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //key = toUserID
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_state_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_state_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_internal_pkg_state_state_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetPosts() map[string]*ListOfPosts {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *State) GetScheduledPosts() []*messages.Post {
	if x != nil {
		return x.ScheduledPosts
	}
	return nil
}

func (x *State) GetDirectMessages() map[string]*ListOfDirects {
	if x != nil {
		return x.DirectMessages
	}
	return nil
}

type ListOfPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*messages.Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *ListOfPosts) Reset() {
	*x = ListOfPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_state_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfPosts) ProtoMessage() {}

func (x *ListOfPosts) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_state_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfPosts.ProtoReflect.Descriptor instead.
func (*ListOfPosts) Descriptor() ([]byte, []int) {
	return file_internal_pkg_state_state_proto_rawDescGZIP(), []int{1}
}

func (x *ListOfPosts) GetPosts() []*messages.Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type ListOfDirects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Directs []*messages.DirectMessage `protobuf:"bytes,1,rep,name=directs,proto3" json:"directs,omitempty"`
}

func (x *ListOfDirects) Reset() {
	*x = ListOfDirects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_state_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfDirects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfDirects) ProtoMessage() {}

func (x *ListOfDirects) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_state_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfDirects.ProtoReflect.Descriptor instead.
func (*ListOfDirects) Descriptor() ([]byte, []int) {
	return file_internal_pkg_state_state_proto_rawDescGZIP(), []int{2}
}

func (x *ListOfDirects) GetDirects() []*messages.DirectMessage {
	if x != nil {
		return x.Directs
	}
	return nil
}

var File_internal_pkg_state_state_proto protoreflect.FileDescriptor

var file_internal_pkg_state_state_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12,
	0x48, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x4c, 0x0a, 0x0a, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x57, 0x0a, 0x13, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x33, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12,
	0x24, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61, 0x6c, 0x6c, 0x69, 0x66, 0x65,
	0x2f, 0x70, 0x62, 0x75, 0x66, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x3b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_state_state_proto_rawDescOnce sync.Once
	file_internal_pkg_state_state_proto_rawDescData = file_internal_pkg_state_state_proto_rawDesc
)

func file_internal_pkg_state_state_proto_rawDescGZIP() []byte {
	file_internal_pkg_state_state_proto_rawDescOnce.Do(func() {
		file_internal_pkg_state_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_state_state_proto_rawDescData)
	})
	return file_internal_pkg_state_state_proto_rawDescData
}

var file_internal_pkg_state_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pkg_state_state_proto_goTypes = []interface{}{
	(*State)(nil),                  // 0: state.State
	(*ListOfPosts)(nil),            // 1: state.ListOfPosts
	(*ListOfDirects)(nil),          // 2: state.ListOfDirects
	nil,                            // 3: state.State.PostsEntry
	nil,                            // 4: state.State.DirectMessagesEntry
	(*messages.Post)(nil),          // 5: messages.Post
	(*messages.DirectMessage)(nil), // 6: messages.DirectMessage
}
var file_internal_pkg_state_state_proto_depIdxs = []int32{
	3, // 0: state.State.posts:type_name -> state.State.PostsEntry
	5, // 1: state.State.scheduledPosts:type_name -> messages.Post
	4, // 2: state.State.directMessages:type_name -> state.State.DirectMessagesEntry
	5, // 3: state.ListOfPosts.posts:type_name -> messages.Post
	6, // 4: state.ListOfDirects.directs:type_name -> messages.DirectMessage
	1, // 5: state.State.PostsEntry.value:type_name -> state.ListOfPosts
	2, // 6: state.State.DirectMessagesEntry.value:type_name -> state.ListOfDirects
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_internal_pkg_state_state_proto_init() }
func file_internal_pkg_state_state_proto_init() {
	if File_internal_pkg_state_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_state_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_internal_pkg_state_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfPosts); i {
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
		file_internal_pkg_state_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfDirects); i {
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
			RawDescriptor: file_internal_pkg_state_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_state_state_proto_goTypes,
		DependencyIndexes: file_internal_pkg_state_state_proto_depIdxs,
		MessageInfos:      file_internal_pkg_state_state_proto_msgTypes,
	}.Build()
	File_internal_pkg_state_state_proto = out.File
	file_internal_pkg_state_state_proto_rawDesc = nil
	file_internal_pkg_state_state_proto_goTypes = nil
	file_internal_pkg_state_state_proto_depIdxs = nil
}