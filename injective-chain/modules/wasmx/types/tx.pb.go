// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/wasmx/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgExecuteContractCompat submits the given message data to a smart contract, compatible with EIP712
type MsgExecuteContractCompat struct {
	// Sender is the that actor that signed the messages
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// Msg json encoded message to be passed to the contract
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	// Funds coins that are transferred to the contract on execution
	Funds string `protobuf:"bytes,4,opt,name=funds,proto3" json:"funds,omitempty"`
}

func (m *MsgExecuteContractCompat) Reset()         { *m = MsgExecuteContractCompat{} }
func (m *MsgExecuteContractCompat) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContractCompat) ProtoMessage()    {}
func (*MsgExecuteContractCompat) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7afe23baa925f70, []int{0}
}
func (m *MsgExecuteContractCompat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContractCompat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContractCompat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContractCompat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContractCompat.Merge(m, src)
}
func (m *MsgExecuteContractCompat) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContractCompat) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContractCompat.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContractCompat proto.InternalMessageInfo

func (m *MsgExecuteContractCompat) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgExecuteContractCompat) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *MsgExecuteContractCompat) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MsgExecuteContractCompat) GetFunds() string {
	if m != nil {
		return m.Funds
	}
	return ""
}

// MsgExecuteContractCompatResponse returns execution result data.
type MsgExecuteContractCompatResponse struct {
	// Data contains bytes to returned from the contract
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgExecuteContractCompatResponse) Reset()         { *m = MsgExecuteContractCompatResponse{} }
func (m *MsgExecuteContractCompatResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContractCompatResponse) ProtoMessage()    {}
func (*MsgExecuteContractCompatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7afe23baa925f70, []int{1}
}
func (m *MsgExecuteContractCompatResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContractCompatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContractCompatResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContractCompatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContractCompatResponse.Merge(m, src)
}
func (m *MsgExecuteContractCompatResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContractCompatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContractCompatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContractCompatResponse proto.InternalMessageInfo

func (m *MsgExecuteContractCompatResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgExecuteContractCompat)(nil), "injective.wasmx.v1.MsgExecuteContractCompat")
	proto.RegisterType((*MsgExecuteContractCompatResponse)(nil), "injective.wasmx.v1.MsgExecuteContractCompatResponse")
}

func init() { proto.RegisterFile("injective/wasmx/v1/tx.proto", fileDescriptor_f7afe23baa925f70) }

var fileDescriptor_f7afe23baa925f70 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x6b, 0x5a, 0x2a, 0xb0, 0x18, 0x90, 0x55, 0x50, 0x54, 0x24, 0xab, 0xea, 0xc4, 0x00,
	0xb6, 0x0a, 0x88, 0x0f, 0xa0, 0x62, 0x40, 0xd0, 0xa5, 0x23, 0x9b, 0xeb, 0x3c, 0xdc, 0x20, 0x62,
	0x47, 0xb1, 0x13, 0x82, 0xd8, 0xf8, 0x02, 0x3e, 0x8b, 0xb1, 0x23, 0x23, 0x4a, 0x7e, 0x04, 0xd5,
	0x4d, 0x23, 0x24, 0xe8, 0xc0, 0x76, 0xef, 0xbb, 0x57, 0x3e, 0xcf, 0x7a, 0xf8, 0x28, 0xd2, 0x8f,
	0x20, 0x5d, 0x94, 0x03, 0x7f, 0x16, 0x36, 0x2e, 0x78, 0x3e, 0xe2, 0xae, 0x60, 0x49, 0x6a, 0x9c,
	0x21, 0xa4, 0x09, 0x99, 0x0f, 0x59, 0x3e, 0xea, 0xf7, 0x94, 0x51, 0xc6, 0xc7, 0x7c, 0xa9, 0x56,
	0xcd, 0x61, 0x8e, 0x83, 0x89, 0x55, 0xd7, 0x05, 0xc8, 0xcc, 0xc1, 0xd8, 0x68, 0x97, 0x0a, 0xe9,
	0xc6, 0x26, 0x4e, 0x84, 0x23, 0x87, 0xb8, 0x6b, 0x41, 0x87, 0x90, 0x06, 0x68, 0x80, 0x8e, 0x77,
	0xa7, 0xb5, 0x23, 0x7d, 0xbc, 0x23, 0xeb, 0x66, 0xb0, 0xe5, 0x93, 0xc6, 0x93, 0x7d, 0xdc, 0x8e,
	0xad, 0x0a, 0xda, 0x7e, 0xbc, 0x94, 0xa4, 0x87, 0xb7, 0x1f, 0x32, 0x1d, 0xda, 0xa0, 0xe3, 0x67,
	0x2b, 0x33, 0xbc, 0xc4, 0x83, 0x4d, 0xdc, 0x29, 0xd8, 0xc4, 0x68, 0x0b, 0x84, 0xe0, 0x4e, 0x28,
	0x9c, 0xf0, 0xf4, 0xbd, 0xa9, 0xd7, 0x67, 0x6f, 0x08, 0xb7, 0x27, 0x56, 0x91, 0x57, 0x7c, 0xf0,
	0xf7, 0xd2, 0x27, 0xec, 0xf7, 0xdf, 0xd9, 0x26, 0x54, 0xff, 0xe2, 0x3f, 0xed, 0xf5, 0x62, 0x57,
	0xf0, 0x51, 0x52, 0xb4, 0x28, 0x29, 0xfa, 0x2a, 0x29, 0x7a, 0xaf, 0x68, 0x6b, 0x51, 0xd1, 0xd6,
	0x67, 0x45, 0x5b, 0xf7, 0xb7, 0x2a, 0x72, 0xf3, 0x6c, 0xc6, 0xa4, 0x89, 0xf9, 0xcd, 0xfa, 0xe5,
	0x3b, 0x31, 0xb3, 0xbc, 0xe1, 0x9c, 0x4a, 0x93, 0xc2, 0x4f, 0x3b, 0x17, 0x91, 0xe6, 0xb1, 0x09,
	0xb3, 0x27, 0xb0, 0xf5, 0x2d, 0xdd, 0x4b, 0x02, 0x76, 0xd6, 0xf5, 0x27, 0x3a, 0xff, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x35, 0x56, 0xbb, 0xe7, 0xeb, 0x01, 0x00, 0x00,
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
	ExecuteContractCompat(ctx context.Context, in *MsgExecuteContractCompat, opts ...grpc.CallOption) (*MsgExecuteContractCompatResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ExecuteContractCompat(ctx context.Context, in *MsgExecuteContractCompat, opts ...grpc.CallOption) (*MsgExecuteContractCompatResponse, error) {
	out := new(MsgExecuteContractCompatResponse)
	err := c.cc.Invoke(ctx, "/injective.wasmx.v1.Msg/ExecuteContractCompat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ExecuteContractCompat(context.Context, *MsgExecuteContractCompat) (*MsgExecuteContractCompatResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ExecuteContractCompat(ctx context.Context, req *MsgExecuteContractCompat) (*MsgExecuteContractCompatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteContractCompat not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ExecuteContractCompat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteContractCompat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteContractCompat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective.wasmx.v1.Msg/ExecuteContractCompat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteContractCompat(ctx, req.(*MsgExecuteContractCompat))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "injective.wasmx.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteContractCompat",
			Handler:    _Msg_ExecuteContractCompat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "injective/wasmx/v1/tx.proto",
}

func (m *MsgExecuteContractCompat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContractCompat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContractCompat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		i -= len(m.Funds)
		copy(dAtA[i:], m.Funds)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Funds)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteContractCompatResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContractCompatResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContractCompatResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgExecuteContractCompat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Funds)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgExecuteContractCompatResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgExecuteContractCompat) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecuteContractCompat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContractCompat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
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
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
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
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
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
			m.Funds = string(dAtA[iNdEx:postIndex])
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
func (m *MsgExecuteContractCompatResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecuteContractCompatResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContractCompatResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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