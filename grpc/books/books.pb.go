// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: grpc/books/books.proto

package books

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Pages    uint32 `protobuf:"varint,3,opt,name=Pages,proto3" json:"Pages,omitempty"`
	AuthorID uint32 `protobuf:"varint,4,opt,name=AuthorID,proto3" json:"AuthorID,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_books_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_books_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_grpc_books_books_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetPages() uint32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *Book) GetAuthorID() uint32 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

type BookAndAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Pages     uint32 `protobuf:"varint,3,opt,name=Pages,proto3" json:"Pages,omitempty"`
	AuthorID  uint32 `protobuf:"varint,4,opt,name=AuthorID,proto3" json:"AuthorID,omitempty"`
	FirstName string `protobuf:"bytes,5,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,6,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func (x *BookAndAuthor) Reset() {
	*x = BookAndAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_books_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookAndAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookAndAuthor) ProtoMessage() {}

func (x *BookAndAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_books_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookAndAuthor.ProtoReflect.Descriptor instead.
func (*BookAndAuthor) Descriptor() ([]byte, []int) {
	return file_grpc_books_books_proto_rawDescGZIP(), []int{1}
}

func (x *BookAndAuthor) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *BookAndAuthor) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookAndAuthor) GetPages() uint32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *BookAndAuthor) GetAuthorID() uint32 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

func (x *BookAndAuthor) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *BookAndAuthor) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type BooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BooksResponse) Reset() {
	*x = BooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_books_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksResponse) ProtoMessage() {}

func (x *BooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_books_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksResponse.ProtoReflect.Descriptor instead.
func (*BooksResponse) Descriptor() ([]byte, []int) {
	return file_grpc_books_books_proto_rawDescGZIP(), []int{2}
}

func (x *BooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type FindBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *FindBookRequest) Reset() {
	*x = FindBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_books_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBookRequest) ProtoMessage() {}

func (x *FindBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_books_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBookRequest.ProtoReflect.Descriptor instead.
func (*FindBookRequest) Descriptor() ([]byte, []int) {
	return file_grpc_books_books_proto_rawDescGZIP(), []int{3}
}

func (x *FindBookRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_grpc_books_books_proto protoreflect.FileDescriptor

var file_grpc_books_books_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x44, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x6e,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x0d, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0x9c, 0x01, 0x0a, 0x05, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3a, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x0e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6c, 0x61, 0x73, 0x6f, 0x76, 0x41, 0x72,
	0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_books_books_proto_rawDescOnce sync.Once
	file_grpc_books_books_proto_rawDescData = file_grpc_books_books_proto_rawDesc
)

func file_grpc_books_books_proto_rawDescGZIP() []byte {
	file_grpc_books_books_proto_rawDescOnce.Do(func() {
		file_grpc_books_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_books_books_proto_rawDescData)
	})
	return file_grpc_books_books_proto_rawDescData
}

var file_grpc_books_books_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_books_books_proto_goTypes = []interface{}{
	(*Book)(nil),            // 0: Book
	(*BookAndAuthor)(nil),   // 1: BookAndAuthor
	(*BooksResponse)(nil),   // 2: BooksResponse
	(*FindBookRequest)(nil), // 3: FindBookRequest
	(*emptypb.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_grpc_books_books_proto_depIdxs = []int32{
	0, // 0: BooksResponse.books:type_name -> Book
	4, // 1: Books.GetBooks:input_type -> google.protobuf.Empty
	3, // 2: Books.FindBook:input_type -> FindBookRequest
	1, // 3: Books.AddBookAndAuthor:input_type -> BookAndAuthor
	2, // 4: Books.GetBooks:output_type -> BooksResponse
	0, // 5: Books.FindBook:output_type -> Book
	4, // 6: Books.AddBookAndAuthor:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_books_books_proto_init() }
func file_grpc_books_books_proto_init() {
	if File_grpc_books_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_books_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_grpc_books_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookAndAuthor); i {
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
		file_grpc_books_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksResponse); i {
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
		file_grpc_books_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBookRequest); i {
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
			RawDescriptor: file_grpc_books_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_books_books_proto_goTypes,
		DependencyIndexes: file_grpc_books_books_proto_depIdxs,
		MessageInfos:      file_grpc_books_books_proto_msgTypes,
	}.Build()
	File_grpc_books_books_proto = out.File
	file_grpc_books_books_proto_rawDesc = nil
	file_grpc_books_books_proto_goTypes = nil
	file_grpc_books_books_proto_depIdxs = nil
}
