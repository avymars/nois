// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/genesis_validator.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type GenesisValidator struct {
	LaunchID       uint64     `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
	Address        string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	GenTx          []byte     `protobuf:"bytes,3,opt,name=genTx,proto3" json:"genTx,omitempty"`
	ConsPubKey     []byte     `protobuf:"bytes,4,opt,name=consPubKey,proto3" json:"consPubKey,omitempty"`
	SelfDelegation types.Coin `protobuf:"bytes,5,opt,name=selfDelegation,proto3" json:"selfDelegation"`
	Peer           Peer       `protobuf:"bytes,6,opt,name=peer,proto3" json:"peer"`
}

func (m *GenesisValidator) Reset()         { *m = GenesisValidator{} }
func (m *GenesisValidator) String() string { return proto.CompactTextString(m) }
func (*GenesisValidator) ProtoMessage()    {}
func (*GenesisValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_da7a6c7c23ebd0c1, []int{0}
}
func (m *GenesisValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisValidator.Merge(m, src)
}
func (m *GenesisValidator) XXX_Size() int {
	return m.Size()
}
func (m *GenesisValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisValidator.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisValidator proto.InternalMessageInfo

func (m *GenesisValidator) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *GenesisValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GenesisValidator) GetGenTx() []byte {
	if m != nil {
		return m.GenTx
	}
	return nil
}

func (m *GenesisValidator) GetConsPubKey() []byte {
	if m != nil {
		return m.ConsPubKey
	}
	return nil
}

func (m *GenesisValidator) GetSelfDelegation() types.Coin {
	if m != nil {
		return m.SelfDelegation
	}
	return types.Coin{}
}

func (m *GenesisValidator) GetPeer() Peer {
	if m != nil {
		return m.Peer
	}
	return Peer{}
}

type Peer struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Connection:
	//	*Peer_TcpAddress
	//	*Peer_HttpTunnel
	//	*Peer_None
	Connection isPeer_Connection `protobuf_oneof:"connection"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_da7a6c7c23ebd0c1, []int{1}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

type isPeer_Connection interface {
	isPeer_Connection()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Peer_TcpAddress struct {
	TcpAddress string `protobuf:"bytes,2,opt,name=tcpAddress,proto3,oneof" json:"tcpAddress,omitempty"`
}
type Peer_HttpTunnel struct {
	HttpTunnel *Peer_HTTPTunnel `protobuf:"bytes,3,opt,name=httpTunnel,proto3,oneof" json:"httpTunnel,omitempty"`
}
type Peer_None struct {
	None *Peer_EmptyConnection `protobuf:"bytes,4,opt,name=none,proto3,oneof" json:"none,omitempty"`
}

func (*Peer_TcpAddress) isPeer_Connection() {}
func (*Peer_HttpTunnel) isPeer_Connection() {}
func (*Peer_None) isPeer_Connection()       {}

func (m *Peer) GetConnection() isPeer_Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *Peer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Peer) GetTcpAddress() string {
	if x, ok := m.GetConnection().(*Peer_TcpAddress); ok {
		return x.TcpAddress
	}
	return ""
}

func (m *Peer) GetHttpTunnel() *Peer_HTTPTunnel {
	if x, ok := m.GetConnection().(*Peer_HttpTunnel); ok {
		return x.HttpTunnel
	}
	return nil
}

func (m *Peer) GetNone() *Peer_EmptyConnection {
	if x, ok := m.GetConnection().(*Peer_None); ok {
		return x.None
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Peer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Peer_TcpAddress)(nil),
		(*Peer_HttpTunnel)(nil),
		(*Peer_None)(nil),
	}
}

type Peer_HTTPTunnel struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Peer_HTTPTunnel) Reset()         { *m = Peer_HTTPTunnel{} }
func (m *Peer_HTTPTunnel) String() string { return proto.CompactTextString(m) }
func (*Peer_HTTPTunnel) ProtoMessage()    {}
func (*Peer_HTTPTunnel) Descriptor() ([]byte, []int) {
	return fileDescriptor_da7a6c7c23ebd0c1, []int{1, 0}
}
func (m *Peer_HTTPTunnel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer_HTTPTunnel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer_HTTPTunnel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer_HTTPTunnel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer_HTTPTunnel.Merge(m, src)
}
func (m *Peer_HTTPTunnel) XXX_Size() int {
	return m.Size()
}
func (m *Peer_HTTPTunnel) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer_HTTPTunnel.DiscardUnknown(m)
}

var xxx_messageInfo_Peer_HTTPTunnel proto.InternalMessageInfo

func (m *Peer_HTTPTunnel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Peer_HTTPTunnel) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Peer_EmptyConnection struct {
}

func (m *Peer_EmptyConnection) Reset()         { *m = Peer_EmptyConnection{} }
func (m *Peer_EmptyConnection) String() string { return proto.CompactTextString(m) }
func (*Peer_EmptyConnection) ProtoMessage()    {}
func (*Peer_EmptyConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_da7a6c7c23ebd0c1, []int{1, 1}
}
func (m *Peer_EmptyConnection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer_EmptyConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer_EmptyConnection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer_EmptyConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer_EmptyConnection.Merge(m, src)
}
func (m *Peer_EmptyConnection) XXX_Size() int {
	return m.Size()
}
func (m *Peer_EmptyConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer_EmptyConnection.DiscardUnknown(m)
}

var xxx_messageInfo_Peer_EmptyConnection proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisValidator)(nil), "tendermint.spn.launch.GenesisValidator")
	proto.RegisterType((*Peer)(nil), "tendermint.spn.launch.Peer")
	proto.RegisterType((*Peer_HTTPTunnel)(nil), "tendermint.spn.launch.Peer.HTTPTunnel")
	proto.RegisterType((*Peer_EmptyConnection)(nil), "tendermint.spn.launch.Peer.EmptyConnection")
}

func init() { proto.RegisterFile("launch/genesis_validator.proto", fileDescriptor_da7a6c7c23ebd0c1) }

var fileDescriptor_da7a6c7c23ebd0c1 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x83, 0x5b, 0xe8, 0xa4, 0x2a, 0xb0, 0x0a, 0x92, 0x1b, 0x24, 0x63, 0xf5, 0x80, 0x2c,
	0x21, 0xd6, 0x6a, 0x10, 0x17, 0x6e, 0x49, 0x8b, 0x1a, 0xc4, 0x25, 0x32, 0x11, 0x07, 0x2e, 0x95,
	0x7f, 0x06, 0x67, 0x25, 0x7b, 0xd7, 0xf2, 0x6e, 0xaa, 0xe6, 0x2d, 0x38, 0x72, 0xe6, 0x19, 0x78,
	0x88, 0x1e, 0x2b, 0x4e, 0x9c, 0x10, 0x4a, 0x5e, 0x04, 0x79, 0xd7, 0x69, 0xa2, 0x08, 0xf5, 0xb6,
	0x33, 0xf3, 0x7d, 0xdf, 0xcc, 0x37, 0xb3, 0xe0, 0x15, 0xf1, 0x9c, 0xa7, 0xb3, 0x30, 0x47, 0x8e,
	0x92, 0xc9, 0xcb, 0xab, 0xb8, 0x60, 0x59, 0xac, 0x44, 0x4d, 0xab, 0x5a, 0x28, 0x41, 0x9e, 0x29,
	0xe4, 0x19, 0xd6, 0x25, 0xe3, 0x8a, 0xca, 0x8a, 0x53, 0x03, 0xef, 0xf7, 0x72, 0x91, 0x0b, 0x8d,
	0x08, 0x9b, 0x97, 0x01, 0xf7, 0xbd, 0x54, 0xc8, 0x52, 0xc8, 0x30, 0x89, 0x25, 0x86, 0x57, 0xa7,
	0x09, 0xaa, 0xf8, 0x34, 0x4c, 0x05, 0xe3, 0x6d, 0xfd, 0xd8, 0xd4, 0x2f, 0x0d, 0xd1, 0x04, 0xa6,
	0x74, 0xf2, 0xbd, 0x03, 0x4f, 0x2e, 0xcc, 0x0c, 0x9f, 0xd7, 0x23, 0x90, 0x3e, 0x3c, 0x32, 0xfd,
	0x3e, 0x9c, 0xbb, 0xb6, 0x6f, 0x07, 0x4e, 0x74, 0x17, 0x93, 0x01, 0x3c, 0x8c, 0xb3, 0xac, 0x46,
	0x29, 0xdd, 0x8e, 0x6f, 0x07, 0x07, 0x23, 0xf7, 0xd7, 0xcf, 0xd7, 0xbd, 0x56, 0x73, 0x68, 0x2a,
	0x9f, 0x54, 0xcd, 0x78, 0x1e, 0xad, 0x81, 0xa4, 0x07, 0x7b, 0x39, 0xf2, 0xe9, 0xb5, 0xfb, 0xc0,
	0xb7, 0x83, 0xc3, 0xc8, 0x04, 0xc4, 0x03, 0x48, 0x05, 0x97, 0x93, 0x79, 0xf2, 0x11, 0x17, 0xae,
	0xa3, 0x4b, 0x5b, 0x19, 0x72, 0x01, 0x47, 0x12, 0x8b, 0xaf, 0xe7, 0x58, 0x60, 0x1e, 0x2b, 0x26,
	0xb8, 0xbb, 0xe7, 0xdb, 0x41, 0x77, 0x70, 0x4c, 0xdb, 0x6e, 0x8d, 0x5d, 0xda, 0xda, 0xa5, 0x67,
	0x82, 0xf1, 0x91, 0x73, 0xf3, 0xe7, 0x85, 0x15, 0xed, 0xd0, 0xc8, 0x5b, 0x70, 0x2a, 0xc4, 0xda,
	0xdd, 0xd7, 0xf4, 0xe7, 0xf4, 0xbf, 0xab, 0xa5, 0x13, 0xc4, 0xba, 0x15, 0xd0, 0xf0, 0x93, 0x1f,
	0x1d, 0x70, 0x9a, 0x24, 0x39, 0x82, 0x0e, 0xcb, 0xf4, 0x22, 0x0e, 0xa2, 0x0e, 0xcb, 0x88, 0x0f,
	0xa0, 0xd2, 0x6a, 0xb8, 0xbd, 0x85, 0xb1, 0x15, 0x6d, 0xe5, 0xc8, 0x18, 0x60, 0xa6, 0x54, 0x35,
	0x9d, 0x73, 0x8e, 0x85, 0x76, 0xdd, 0x1d, 0xbc, 0xbc, 0xa7, 0x2f, 0x1d, 0x4f, 0xa7, 0x13, 0x83,
	0x6e, 0x94, 0x36, 0x5c, 0x32, 0x04, 0x87, 0x0b, 0x8e, 0x7a, 0x3d, 0xdd, 0xc1, 0xab, 0xfb, 0x34,
	0xde, 0x97, 0x95, 0x5a, 0x9c, 0x09, 0xce, 0x31, 0x6d, 0x6c, 0x8f, 0xad, 0x48, 0x53, 0xfb, 0xef,
	0x00, 0x36, 0xf2, 0x84, 0x80, 0xc3, 0xe3, 0x12, 0x5b, 0x3b, 0xfa, 0x4d, 0xdc, 0x9d, 0x9b, 0xde,
	0x5d, 0xae, 0xff, 0x14, 0x1e, 0xef, 0xc8, 0x8e, 0x0e, 0xf5, 0xd9, 0xd6, 0xd1, 0xe8, 0x66, 0xe9,
	0xd9, 0xb7, 0x4b, 0xcf, 0xfe, 0xbb, 0xf4, 0xec, 0x6f, 0x2b, 0xcf, 0xba, 0x5d, 0x79, 0xd6, 0xef,
	0x95, 0x67, 0x7d, 0x09, 0x72, 0xa6, 0x66, 0xf3, 0x84, 0xa6, 0xa2, 0x0c, 0x37, 0x53, 0x87, 0xb2,
	0xe2, 0xe1, 0x75, 0xd8, 0xfe, 0x7e, 0xb5, 0xa8, 0x50, 0x26, 0xfb, 0xfa, 0x2b, 0xbe, 0xf9, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xef, 0xc9, 0x08, 0x5b, 0x14, 0x03, 0x00, 0x00,
}

func (m *GenesisValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Peer.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesisValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.SelfDelegation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesisValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.ConsPubKey) > 0 {
		i -= len(m.ConsPubKey)
		copy(dAtA[i:], m.ConsPubKey)
		i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.ConsPubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.GenTx) > 0 {
		i -= len(m.GenTx)
		copy(dAtA[i:], m.GenTx)
		i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.GenTx)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchID != 0 {
		i = encodeVarintGenesisValidator(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Connection != nil {
		{
			size := m.Connection.Size()
			i -= size
			if _, err := m.Connection.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Peer_TcpAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer_TcpAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.TcpAddress)
	copy(dAtA[i:], m.TcpAddress)
	i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.TcpAddress)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *Peer_HttpTunnel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer_HttpTunnel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HttpTunnel != nil {
		{
			size, err := m.HttpTunnel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesisValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Peer_None) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer_None) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.None != nil {
		{
			size, err := m.None.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesisValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Peer_HTTPTunnel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer_HTTPTunnel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer_HTTPTunnel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGenesisValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Peer_EmptyConnection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer_EmptyConnection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer_EmptyConnection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintGenesisValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesisValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovGenesisValidator(uint64(m.LaunchID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	l = len(m.GenTx)
	if l > 0 {
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	l = len(m.ConsPubKey)
	if l > 0 {
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	l = m.SelfDelegation.Size()
	n += 1 + l + sovGenesisValidator(uint64(l))
	l = m.Peer.Size()
	n += 1 + l + sovGenesisValidator(uint64(l))
	return n
}

func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	if m.Connection != nil {
		n += m.Connection.Size()
	}
	return n
}

func (m *Peer_TcpAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TcpAddress)
	n += 1 + l + sovGenesisValidator(uint64(l))
	return n
}
func (m *Peer_HttpTunnel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpTunnel != nil {
		l = m.HttpTunnel.Size()
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	return n
}
func (m *Peer_None) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.None != nil {
		l = m.None.Size()
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	return n
}
func (m *Peer_HTTPTunnel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesisValidator(uint64(l))
	}
	return n
}

func (m *Peer_EmptyConnection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovGenesisValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesisValidator(x uint64) (n int) {
	return sovGenesisValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisValidator
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
			return fmt.Errorf("proto: GenesisValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenTx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenTx = append(m.GenTx[:0], dAtA[iNdEx:postIndex]...)
			if m.GenTx == nil {
				m.GenTx = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsPubKey = append(m.ConsPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsPubKey == nil {
				m.ConsPubKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfDelegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SelfDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Peer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisValidator
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
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisValidator
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Connection = &Peer_TcpAddress{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpTunnel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Peer_HTTPTunnel{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Connection = &Peer_HttpTunnel{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field None", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Peer_EmptyConnection{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Connection = &Peer_None{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisValidator
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
func (m *Peer_HTTPTunnel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisValidator
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
			return fmt.Errorf("proto: HTTPTunnel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPTunnel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisValidator
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
				return ErrInvalidLengthGenesisValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisValidator
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
func (m *Peer_EmptyConnection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisValidator
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
			return fmt.Errorf("proto: EmptyConnection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyConnection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisValidator
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
func skipGenesisValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesisValidator
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
					return 0, ErrIntOverflowGenesisValidator
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
					return 0, ErrIntOverflowGenesisValidator
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
				return 0, ErrInvalidLengthGenesisValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesisValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesisValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesisValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesisValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesisValidator = fmt.Errorf("proto: unexpected end of group")
)
