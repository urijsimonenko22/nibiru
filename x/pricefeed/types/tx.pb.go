// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pricefeed/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgPostPrice represents a method for creating a new post price
type MsgPostPrice struct {
	// From: address of oracle
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Token0: denominator unit of the price, a.k.a. quote asset
	Token0 string `protobuf:"bytes,2,opt,name=token0,proto3" json:"token0,omitempty"`
	// Token1: numerator unit of price, a.k.a. base asset
	Token1 string `protobuf:"bytes,3,opt,name=token1,proto3" json:"token1,omitempty"`
	// Price: Price of the trading pair in units of token1 / token0.
	Price  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Expiry time.Time                              `protobuf:"bytes,5,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *MsgPostPrice) Reset()         { *m = MsgPostPrice{} }
func (m *MsgPostPrice) String() string { return proto.CompactTextString(m) }
func (*MsgPostPrice) ProtoMessage()    {}
func (*MsgPostPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d54c954ce5f810, []int{0}
}
func (m *MsgPostPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostPrice.Merge(m, src)
}
func (m *MsgPostPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostPrice proto.InternalMessageInfo

// MsgPostPriceResponse defines the Msg/PostPrice response type.
type MsgPostPriceResponse struct {
}

func (m *MsgPostPriceResponse) Reset()         { *m = MsgPostPriceResponse{} }
func (m *MsgPostPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPostPriceResponse) ProtoMessage()    {}
func (*MsgPostPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d54c954ce5f810, []int{1}
}
func (m *MsgPostPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostPriceResponse.Merge(m, src)
}
func (m *MsgPostPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostPriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPostPrice)(nil), "NibiruChain.pricefeed.v1.MsgPostPrice")
	proto.RegisterType((*MsgPostPriceResponse)(nil), "NibiruChain.pricefeed.v1.MsgPostPriceResponse")
}

func init() { proto.RegisterFile("pricefeed/tx.proto", fileDescriptor_27d54c954ce5f810) }

var fileDescriptor_27d54c954ce5f810 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xb5, 0x69, 0x44, 0x0f, 0xa6, 0x53, 0x55, 0x59, 0x56, 0x75, 0xae, 0x2c, 0x54, 0x75,
	0xe9, 0x1d, 0x29, 0x1b, 0x62, 0x0a, 0x5d, 0x8b, 0x8a, 0xc5, 0xc4, 0x66, 0xbb, 0x17, 0xe7, 0x14,
	0xec, 0x77, 0xf2, 0x5d, 0x50, 0xb2, 0x32, 0xc1, 0x16, 0xc4, 0x3f, 0xc0, 0xc8, 0x9f, 0x92, 0x31,
	0x12, 0x0b, 0x42, 0x28, 0x04, 0x87, 0x3f, 0x04, 0xf9, 0x6c, 0x27, 0x5e, 0x90, 0x98, 0xee, 0xbd,
	0xef, 0x7d, 0xef, 0xc7, 0xf7, 0x1d, 0x26, 0xaa, 0x90, 0x89, 0x18, 0x09, 0x71, 0xcf, 0xcd, 0x8c,
	0xa9, 0x02, 0x0c, 0x10, 0xf7, 0xa5, 0x8c, 0x65, 0x31, 0x7d, 0x31, 0x8e, 0x64, 0xce, 0x76, 0x75,
	0xf6, 0x6e, 0xe0, 0x9d, 0xa4, 0x90, 0x82, 0x25, 0xf1, 0x2a, 0xaa, 0xf9, 0x9e, 0x9f, 0x02, 0xa4,
	0x6f, 0x05, 0xb7, 0x59, 0x3c, 0x1d, 0x71, 0x23, 0x33, 0xa1, 0x4d, 0x94, 0xa9, 0x86, 0x70, 0xd6,
	0x10, 0x22, 0x25, 0x79, 0x94, 0xe7, 0x60, 0x22, 0x23, 0x21, 0xd7, 0x75, 0x35, 0xf8, 0x89, 0xf0,
	0xa3, 0x5b, 0x9d, 0xde, 0x81, 0x36, 0x77, 0xd5, 0x32, 0x42, 0x70, 0x6f, 0x54, 0x40, 0xe6, 0xa2,
	0x73, 0x74, 0x79, 0x1c, 0xda, 0x98, 0x9c, 0xe2, 0xbe, 0x81, 0x89, 0xc8, 0x9f, 0xb8, 0x07, 0x16,
	0x6d, 0xb2, 0x1d, 0x3e, 0x70, 0x0f, 0x3b, 0xf8, 0x80, 0xdc, 0xe0, 0x23, 0x7b, 0xb9, 0xdb, 0xab,
	0xe0, 0x21, 0x5b, 0xae, 0x7d, 0xe7, 0xc7, 0xda, 0xbf, 0x48, 0xa5, 0x19, 0x4f, 0x63, 0x96, 0x40,
	0xc6, 0x13, 0xd0, 0x19, 0xe8, 0xe6, 0xb9, 0xd2, 0xf7, 0x13, 0x6e, 0xe6, 0x4a, 0x68, 0x76, 0x23,
	0x92, 0xb0, 0x6e, 0x26, 0xcf, 0x71, 0x5f, 0xcc, 0x94, 0x2c, 0xe6, 0xee, 0xd1, 0x39, 0xba, 0x7c,
	0x78, 0xed, 0xb1, 0x5a, 0x09, 0x6b, 0xa5, 0xb2, 0xd7, 0xad, 0xd4, 0xe1, 0x83, 0x6a, 0xc5, 0xe2,
	0x97, 0x8f, 0xc2, 0xa6, 0xe7, 0x59, 0xef, 0xc3, 0x17, 0xdf, 0x09, 0x4e, 0xf1, 0x49, 0x57, 0x5d,
	0x28, 0xb4, 0x82, 0x5c, 0x8b, 0xeb, 0x4f, 0x08, 0x1f, 0xde, 0xea, 0x94, 0x7c, 0x44, 0xf8, 0x78,
	0xaf, 0xfd, 0x82, 0xfd, 0xcb, 0x7c, 0xd6, 0x9d, 0xe2, 0xb1, 0xff, 0xe3, 0xb5, 0xdb, 0x82, 0xc7,
	0xef, 0xbf, 0xfd, 0xf9, 0x7c, 0x40, 0x83, 0x33, 0x9e, 0xdb, 0x3e, 0xbe, 0xff, 0x77, 0x05, 0xda,
	0x5c, 0xd9, 0x74, 0xf8, 0x6a, 0xf3, 0x9b, 0xa2, 0xaf, 0x25, 0x45, 0xcb, 0x92, 0xa2, 0x55, 0x49,
	0xd1, 0xa6, 0xa4, 0x68, 0xb1, 0xa5, 0xce, 0x6a, 0x4b, 0x9d, 0xef, 0x5b, 0xea, 0xbc, 0xe1, 0x1d,
	0x03, 0x3b, 0x17, 0xb4, 0x53, 0x67, 0x9d, 0xb9, 0xd6, 0xcd, 0xb8, 0x6f, 0xad, 0x7a, 0xfa, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x01, 0xb5, 0xef, 0xe4, 0x69, 0x02, 0x00, 0x00,
}

func (this *MsgPostPrice) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MsgPostPrice)
	if !ok {
		that2, ok := that.(MsgPostPrice)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MsgPostPrice")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MsgPostPrice but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MsgPostPrice but is not nil && this == nil")
	}
	if this.From != that1.From {
		return fmt.Errorf("From this(%v) Not Equal that(%v)", this.From, that1.From)
	}
	if this.Token0 != that1.Token0 {
		return fmt.Errorf("Token0 this(%v) Not Equal that(%v)", this.Token0, that1.Token0)
	}
	if this.Token1 != that1.Token1 {
		return fmt.Errorf("Token1 this(%v) Not Equal that(%v)", this.Token1, that1.Token1)
	}
	if !this.Price.Equal(that1.Price) {
		return fmt.Errorf("Price this(%v) Not Equal that(%v)", this.Price, that1.Price)
	}
	if !this.Expiry.Equal(that1.Expiry) {
		return fmt.Errorf("Expiry this(%v) Not Equal that(%v)", this.Expiry, that1.Expiry)
	}
	return nil
}
func (this *MsgPostPrice) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPostPrice)
	if !ok {
		that2, ok := that.(MsgPostPrice)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.From != that1.From {
		return false
	}
	if this.Token0 != that1.Token0 {
		return false
	}
	if this.Token1 != that1.Token1 {
		return false
	}
	if !this.Price.Equal(that1.Price) {
		return false
	}
	if !this.Expiry.Equal(that1.Expiry) {
		return false
	}
	return true
}
func (this *MsgPostPriceResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MsgPostPriceResponse)
	if !ok {
		that2, ok := that.(MsgPostPriceResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MsgPostPriceResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MsgPostPriceResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MsgPostPriceResponse but is not nil && this == nil")
	}
	return nil
}
func (this *MsgPostPriceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPostPriceResponse)
	if !ok {
		that2, ok := that.(MsgPostPriceResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PostPrice defines a method for creating a new post price
	PostPrice(ctx context.Context, in *MsgPostPrice, opts ...grpc.CallOption) (*MsgPostPriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PostPrice(ctx context.Context, in *MsgPostPrice, opts ...grpc.CallOption) (*MsgPostPriceResponse, error) {
	out := new(MsgPostPriceResponse)
	err := c.cc.Invoke(ctx, "/NibiruChain.pricefeed.v1.Msg/PostPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PostPrice defines a method for creating a new post price
	PostPrice(context.Context, *MsgPostPrice) (*MsgPostPriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PostPrice(ctx context.Context, req *MsgPostPrice) (*MsgPostPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PostPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPostPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PostPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NibiruChain.pricefeed.v1.Msg/PostPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PostPrice(ctx, req.(*MsgPostPrice))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NibiruChain.pricefeed.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostPrice",
			Handler:    _Msg_PostPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricefeed/tx.proto",
}

func (m *MsgPostPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPostPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPostPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgPostPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPostPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPostPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPostPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPostPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
