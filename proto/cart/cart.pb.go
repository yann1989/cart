// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: proto/cart/cart.proto

package cart

import (
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

type CartFindAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CartFindAll) Reset() {
	*x = CartFindAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartFindAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartFindAll) ProtoMessage() {}

func (x *CartFindAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartFindAll.ProtoReflect.Descriptor instead.
func (*CartFindAll) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartFindAll) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CartAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*CartInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *CartAll) Reset() {
	*x = CartAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartAll) ProtoMessage() {}

func (x *CartAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartAll.ProtoReflect.Descriptor instead.
func (*CartAll) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CartAll) GetInfos() []*CartInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CartID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CartID) Reset() {
	*x = CartID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartID) ProtoMessage() {}

func (x *CartID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartID.ProtoReflect.Descriptor instead.
func (*CartID) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CartID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChangeNum int64 `protobuf:"varint,2,opt,name=change_num,json=changeNum,proto3" json:"change_num,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetChangeNum() int64 {
	if x != nil {
		return x.ChangeNum
	}
	return 0
}

type Clean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Clean) Reset() {
	*x = Clean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clean) ProtoMessage() {}

func (x *Clean) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clean.ProtoReflect.Descriptor instead.
func (*Clean) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{4}
}

func (x *Clean) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CartInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,4,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *CartInfo) Reset() {
	*x = CartInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartInfo) ProtoMessage() {}

func (x *CartInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartInfo.ProtoReflect.Descriptor instead.
func (*CartInfo) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{6}
}

func (x *CartInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CartInfo) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartInfo) GetSizeId() int64 {
	if x != nil {
		return x.SizeId
	}
	return 0
}

func (x *CartInfo) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ResponseAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId int64  `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResponseAdd) Reset() {
	*x = ResponseAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAdd) ProtoMessage() {}

func (x *ResponseAdd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAdd.ProtoReflect.Descriptor instead.
func (*ResponseAdd) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseAdd) GetCartId() int64 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *ResponseAdd) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_cart_cart_proto protoreflect.FileDescriptor

var file_proto_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x2a, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x20, 0x0a, 0x05,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1c,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7d, 0x0a, 0x08,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x38, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xd2, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x24,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x12, 0x09, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41,
	0x64, 0x64, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x72,
	0x74, 0x12, 0x06, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04, 0x49, 0x6e, 0x63, 0x72, 0x12, 0x05,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04, 0x44, 0x65, 0x63, 0x72, 0x12, 0x05, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x07, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x0c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x1a, 0x08,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x6c, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cart_cart_proto_rawDescOnce sync.Once
	file_proto_cart_cart_proto_rawDescData = file_proto_cart_cart_proto_rawDesc
)

func file_proto_cart_cart_proto_rawDescGZIP() []byte {
	file_proto_cart_cart_proto_rawDescOnce.Do(func() {
		file_proto_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cart_cart_proto_rawDescData)
	})
	return file_proto_cart_cart_proto_rawDescData
}

var file_proto_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_cart_cart_proto_goTypes = []interface{}{
	(*CartFindAll)(nil), // 0: CartFindAll
	(*CartAll)(nil),     // 1: CartAll
	(*CartID)(nil),      // 2: CartID
	(*Item)(nil),        // 3: Item
	(*Clean)(nil),       // 4: Clean
	(*Response)(nil),    // 5: Response
	(*CartInfo)(nil),    // 6: CartInfo
	(*ResponseAdd)(nil), // 7: ResponseAdd
}
var file_proto_cart_cart_proto_depIdxs = []int32{
	6, // 0: CartAll.infos:type_name -> CartInfo
	6, // 1: Cart.AddCart:input_type -> CartInfo
	4, // 2: Cart.CleanCart:input_type -> Clean
	3, // 3: Cart.Incr:input_type -> Item
	3, // 4: Cart.Decr:input_type -> Item
	2, // 5: Cart.DeleteItemByID:input_type -> CartID
	0, // 6: Cart.GetAll:input_type -> CartFindAll
	7, // 7: Cart.AddCart:output_type -> ResponseAdd
	5, // 8: Cart.CleanCart:output_type -> Response
	5, // 9: Cart.Incr:output_type -> Response
	5, // 10: Cart.Decr:output_type -> Response
	5, // 11: Cart.DeleteItemByID:output_type -> Response
	1, // 12: Cart.GetAll:output_type -> CartAll
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cart_cart_proto_init() }
func file_proto_cart_cart_proto_init() {
	if File_proto_cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartFindAll); i {
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
		file_proto_cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartAll); i {
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
		file_proto_cart_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartID); i {
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
		file_proto_cart_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_proto_cart_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clean); i {
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
		file_proto_cart_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_cart_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartInfo); i {
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
		file_proto_cart_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAdd); i {
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
			RawDescriptor: file_proto_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cart_cart_proto_goTypes,
		DependencyIndexes: file_proto_cart_cart_proto_depIdxs,
		MessageInfos:      file_proto_cart_cart_proto_msgTypes,
	}.Build()
	File_proto_cart_cart_proto = out.File
	file_proto_cart_cart_proto_rawDesc = nil
	file_proto_cart_cart_proto_goTypes = nil
	file_proto_cart_cart_proto_depIdxs = nil
}
