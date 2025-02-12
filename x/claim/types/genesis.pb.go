// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: claim/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the claim module's genesis state.
type GenesisState struct {
	AirdropSupply types.Coin    `protobuf:"bytes,1,opt,name=airdropSupply,proto3" json:"airdropSupply"`
	ClaimRecords  []ClaimRecord `protobuf:"bytes,2,rep,name=claimRecords,proto3" json:"claimRecords"`
	Missions      []Mission     `protobuf:"bytes,3,rep,name=missions,proto3" json:"missions"`
	InitialClaim  InitialClaim  `protobuf:"bytes,4,opt,name=initialClaim,proto3" json:"initialClaim"`
	Params        Params        `protobuf:"bytes,5,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_53acb741836a34f8, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAirdropSupply() types.Coin {
	if m != nil {
		return m.AirdropSupply
	}
	return types.Coin{}
}

func (m *GenesisState) GetClaimRecords() []ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

func (m *GenesisState) GetMissions() []Mission {
	if m != nil {
		return m.Missions
	}
	return nil
}

func (m *GenesisState) GetInitialClaim() InitialClaim {
	if m != nil {
		return m.InitialClaim
	}
	return InitialClaim{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.claim.GenesisState")
}

func init() { proto.RegisterFile("claim/genesis.proto", fileDescriptor_53acb741836a34f8) }

var fileDescriptor_53acb741836a34f8 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x13, 0xe0, 0xa2, 0xab, 0xc0, 0x5d, 0x7c, 0x19, 0x02, 0xba, 0xd7, 0xa5, 0x2c, 0x65,
	0xb2, 0x05, 0xdd, 0xba, 0x54, 0x05, 0x55, 0x55, 0xd5, 0x56, 0xaa, 0x60, 0xeb, 0x82, 0x9c, 0x60,
	0xa5, 0x96, 0x88, 0x6d, 0xd9, 0xa6, 0x2a, 0x6f, 0xd1, 0xc7, 0x62, 0x64, 0xec, 0x54, 0x55, 0x30,
	0xf7, 0x1d, 0x2a, 0x6c, 0x53, 0x40, 0xca, 0x12, 0x25, 0xe7, 0xfc, 0xff, 0x77, 0xfe, 0x93, 0x13,
	0xfd, 0x4d, 0x67, 0x84, 0xe5, 0x38, 0xa3, 0x9c, 0x6a, 0xa6, 0x91, 0x54, 0xc2, 0x08, 0xd0, 0x30,
	0x94, 0x4f, 0xa9, 0xca, 0x19, 0x37, 0x48, 0x4b, 0x8e, 0xac, 0xa6, 0xd5, 0xc8, 0x44, 0x26, 0xac,
	0x00, 0x6f, 0xdf, 0x9c, 0xb6, 0x05, 0x53, 0xa1, 0x73, 0xa1, 0x71, 0x42, 0x34, 0xc5, 0x2f, 0xbd,
	0x84, 0x1a, 0xd2, 0xc3, 0xa9, 0x60, 0xdc, 0xf7, 0x81, 0x1b, 0x20, 0x89, 0x22, 0xb9, 0xe7, 0xb7,
	0x62, 0x57, 0xb3, 0xcf, 0x89, 0xa2, 0xa9, 0x50, 0x53, 0xdf, 0xf1, 0x71, 0x72, 0xa6, 0x35, 0x13,
	0x3b, 0x44, 0xd3, 0x15, 0x19, 0x67, 0x86, 0x91, 0xd9, 0xc4, 0x7e, 0xb9, 0x56, 0xe7, 0xab, 0x14,
	0xd5, 0x6f, 0x5c, 0xf6, 0xb1, 0x21, 0x86, 0x82, 0xeb, 0xe8, 0x0f, 0x61, 0x6a, 0xaa, 0x84, 0x1c,
	0xcf, 0xa5, 0x9c, 0x2d, 0xe2, 0xb0, 0x1d, 0x76, 0x6b, 0xfd, 0x26, 0x72, 0x31, 0xd1, 0x36, 0x26,
	0xf2, 0x31, 0xd1, 0x50, 0x30, 0x3e, 0xa8, 0x2c, 0x3f, 0x4e, 0x82, 0xd1, 0xb1, 0x0b, 0xdc, 0x45,
	0x75, 0x3b, 0x66, 0x64, 0xc3, 0xe9, 0xb8, 0xd4, 0x2e, 0x77, 0x6b, 0xfd, 0x53, 0x54, 0xf4, 0x63,
	0xd0, 0x70, 0xaf, 0xf4, 0xb4, 0x23, 0x33, 0xb8, 0x8c, 0x7e, 0xfb, 0x85, 0x74, 0x5c, 0xb6, 0xa0,
	0xff, 0xc5, 0xa0, 0x07, 0xa7, 0xf2, 0x90, 0x1f, 0x13, 0xb8, 0x8f, 0xea, 0x7e, 0x79, 0x3b, 0x2a,
	0xae, 0xd8, 0x9d, 0x3a, 0xc5, 0x90, 0xdb, 0x03, 0xe5, 0x2e, 0xce, 0xa1, 0x1b, 0x5c, 0x44, 0x55,
	0x77, 0x8d, 0xf8, 0x97, 0xe5, 0xfc, 0x2b, 0xe6, 0x3c, 0x5a, 0x8d, 0x27, 0x78, 0xc7, 0xe0, 0x6a,
	0xb9, 0x86, 0xe1, 0x6a, 0x0d, 0xc3, 0xcf, 0x35, 0x0c, 0xdf, 0x36, 0x30, 0x58, 0x6d, 0x60, 0xf0,
	0xbe, 0x81, 0xc1, 0xd3, 0x59, 0xc6, 0xcc, 0xf3, 0x3c, 0x41, 0xa9, 0xc8, 0xf1, 0x9e, 0x87, 0xb5,
	0xe4, 0xf8, 0xd5, 0x5d, 0x1a, 0x9b, 0x85, 0xa4, 0x3a, 0xa9, 0xda, 0xcb, 0x9d, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0xf4, 0x3a, 0x75, 0x7a, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.InitialClaim.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Missions) > 0 {
		for iNdEx := len(m.Missions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Missions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.AirdropSupply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AirdropSupply.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Missions) > 0 {
		for _, e := range m.Missions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.InitialClaim.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AirdropSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Missions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Missions = append(m.Missions, Mission{})
			if err := m.Missions[len(m.Missions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
