// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: taxi.proto

package trade

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAssetsRequest) Reset() {
	*x = ListAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsRequest) ProtoMessage() {}

func (x *ListAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taxi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsRequest.ProtoReflect.Descriptor instead.
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return file_taxi_proto_rawDescGZIP(), []int{0}
}

type ListAssetsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetHash []string `protobuf:"bytes,1,rep,name=asset_hash,json=assetHash,proto3" json:"asset_hash,omitempty"` // asset hash accepted as payout
}

func (x *ListAssetsReply) Reset() {
	*x = ListAssetsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsReply) ProtoMessage() {}

func (x *ListAssetsReply) ProtoReflect() protoreflect.Message {
	mi := &file_taxi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsReply.ProtoReflect.Descriptor instead.
func (*ListAssetsReply) Descriptor() ([]byte, []int) {
	return file_taxi_proto_rawDescGZIP(), []int{1}
}

func (x *ListAssetsReply) GetAssetHash() []string {
	if x != nil {
		return x.AssetHash
	}
	return nil
}

type TopupWithAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetHash       string `protobuf:"bytes,1,opt,name=asset_hash,json=assetHash,proto3" json:"asset_hash,omitempty"`                      // asset hash to be used for payout
	FeeAmount       uint64 `protobuf:"varint,2,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty"`                     // amount in satoshis of bitcoin needed to cover the fees. It's up to the client to estimate and ask the precise amount
	MillisatPerByte uint64 `protobuf:"varint,3,opt,name=millisat_per_byte,json=millisatPerByte,proto3" json:"millisat_per_byte,omitempty"` // how many millisatoshi per byte we want to spend. ie. 0.1 sat/byte is 100
}

func (x *TopupWithAssetRequest) Reset() {
	*x = TopupWithAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopupWithAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopupWithAssetRequest) ProtoMessage() {}

func (x *TopupWithAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taxi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopupWithAssetRequest.ProtoReflect.Descriptor instead.
func (*TopupWithAssetRequest) Descriptor() ([]byte, []int) {
	return file_taxi_proto_rawDescGZIP(), []int{2}
}

func (x *TopupWithAssetRequest) GetAssetHash() string {
	if x != nil {
		return x.AssetHash
	}
	return ""
}

func (x *TopupWithAssetRequest) GetFeeAmount() uint64 {
	if x != nil {
		return x.FeeAmount
	}
	return 0
}

func (x *TopupWithAssetRequest) GetMillisatPerByte() uint64 {
	if x != nil {
		return x.MillisatPerByte
	}
	return 0
}

type TopupWithAssetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topup  *Topup `protobuf:"bytes,1,opt,name=topup,proto3" json:"topup,omitempty"`
	Expiry uint64 `protobuf:"varint,2,opt,name=expiry,proto3" json:"expiry,omitempty"` // the unix timestamp after wich the locked LBTC input will provably be double-spent
}

func (x *TopupWithAssetReply) Reset() {
	*x = TopupWithAssetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopupWithAssetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopupWithAssetReply) ProtoMessage() {}

func (x *TopupWithAssetReply) ProtoReflect() protoreflect.Message {
	mi := &file_taxi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopupWithAssetReply.ProtoReflect.Descriptor instead.
func (*TopupWithAssetReply) Descriptor() ([]byte, []int) {
	return file_taxi_proto_rawDescGZIP(), []int{3}
}

func (x *TopupWithAssetReply) GetTopup() *Topup {
	if x != nil {
		return x.Topup
	}
	return nil
}

func (x *TopupWithAssetReply) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

type Topup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopupId     string  `protobuf:"bytes,1,opt,name=topup_id,json=topupId,proto3" json:"topup_id,omitempty"`              //random identifier of the currer topup
	Partial     string  `protobuf:"bytes,2,opt,name=partial,proto3" json:"partial,omitempty"`                             // PSET signed with SIGHASH_SINGLE | ANYONECANPAY
	AssetHash   string  `protobuf:"bytes,3,opt,name=asset_hash,json=assetHash,proto3" json:"asset_hash,omitempty"`        // the asset hash used as payout for bitcoin fees
	AssetAmount uint64  `protobuf:"varint,4,opt,name=asset_amount,json=assetAmount,proto3" json:"asset_amount,omitempty"` // the asset denominated amount expressed in satoshis to be used as payout. It includes also the spread as taxi service fee
	AssetPrice  float32 `protobuf:"fixed32,5,opt,name=asset_price,json=assetPrice,proto3" json:"asset_price,omitempty"`   // the price of bitcoin expressed in asset
	BasisPoint  int32   `protobuf:"varint,6,opt,name=basis_point,json=basisPoint,proto3" json:"basis_point,omitempty"`    // the spread expressed in basis point on top the amount needed to repay for bitcoin fees
}

func (x *Topup) Reset() {
	*x = Topup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topup) ProtoMessage() {}

func (x *Topup) ProtoReflect() protoreflect.Message {
	mi := &file_taxi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topup.ProtoReflect.Descriptor instead.
func (*Topup) Descriptor() ([]byte, []int) {
	return file_taxi_proto_rawDescGZIP(), []int{4}
}

func (x *Topup) GetTopupId() string {
	if x != nil {
		return x.TopupId
	}
	return ""
}

func (x *Topup) GetPartial() string {
	if x != nil {
		return x.Partial
	}
	return ""
}

func (x *Topup) GetAssetHash() string {
	if x != nil {
		return x.AssetHash
	}
	return ""
}

func (x *Topup) GetAssetAmount() uint64 {
	if x != nil {
		return x.AssetAmount
	}
	return 0
}

func (x *Topup) GetAssetPrice() float32 {
	if x != nil {
		return x.AssetPrice
	}
	return 0
}

func (x *Topup) GetBasisPoint() int32 {
	if x != nil {
		return x.BasisPoint
	}
	return 0
}

var File_taxi_proto protoreflect.FileDescriptor

var file_taxi_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x30, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x54, 0x6f, 0x70, 0x75, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x65, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x66, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x61, 0x74,
	0x50, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x54, 0x6f, 0x70, 0x75, 0x70,
	0x57, 0x69, 0x74, 0x68, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x54, 0x6f, 0x70, 0x75, 0x70, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x22, 0xc0, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x75, 0x70, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x6f, 0x70, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x6f, 0x70, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x73, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x61, 0x73,
	0x69, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0x7a, 0x0a, 0x04, 0x54, 0x61, 0x78, 0x69, 0x12,
	0x32, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x12, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x0e, 0x54, 0x6f, 0x70, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x54, 0x6f, 0x70, 0x75, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x54, 0x6f, 0x70, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x75, 0x6c, 0x70, 0x65, 0x6d, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2f, 0x74, 0x61, 0x78, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taxi_proto_rawDescOnce sync.Once
	file_taxi_proto_rawDescData = file_taxi_proto_rawDesc
)

func file_taxi_proto_rawDescGZIP() []byte {
	file_taxi_proto_rawDescOnce.Do(func() {
		file_taxi_proto_rawDescData = protoimpl.X.CompressGZIP(file_taxi_proto_rawDescData)
	})
	return file_taxi_proto_rawDescData
}

var file_taxi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_taxi_proto_goTypes = []interface{}{
	(*ListAssetsRequest)(nil),     // 0: ListAssetsRequest
	(*ListAssetsReply)(nil),       // 1: ListAssetsReply
	(*TopupWithAssetRequest)(nil), // 2: TopupWithAssetRequest
	(*TopupWithAssetReply)(nil),   // 3: TopupWithAssetReply
	(*Topup)(nil),                 // 4: Topup
}
var file_taxi_proto_depIdxs = []int32{
	4, // 0: TopupWithAssetReply.topup:type_name -> Topup
	0, // 1: Taxi.ListAssets:input_type -> ListAssetsRequest
	2, // 2: Taxi.TopupWithAsset:input_type -> TopupWithAssetRequest
	1, // 3: Taxi.ListAssets:output_type -> ListAssetsReply
	3, // 4: Taxi.TopupWithAsset:output_type -> TopupWithAssetReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_taxi_proto_init() }
func file_taxi_proto_init() {
	if File_taxi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_taxi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsRequest); i {
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
		file_taxi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsReply); i {
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
		file_taxi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopupWithAssetRequest); i {
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
		file_taxi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopupWithAssetReply); i {
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
		file_taxi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topup); i {
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
			RawDescriptor: file_taxi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taxi_proto_goTypes,
		DependencyIndexes: file_taxi_proto_depIdxs,
		MessageInfos:      file_taxi_proto_msgTypes,
	}.Build()
	File_taxi_proto = out.File
	file_taxi_proto_rawDesc = nil
	file_taxi_proto_goTypes = nil
	file_taxi_proto_depIdxs = nil
}
