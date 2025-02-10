// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/gravity/bridge_transfer_limit_proposal.proto

package gravity

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	cosmossdk_io_math "cosmossdk.io/math"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LimitPeriod int32

const (
	LimitPeriod_NONE    LimitPeriod = 0
	LimitPeriod_DAILY   LimitPeriod = 1
	LimitPeriod_WEEKLY  LimitPeriod = 2
	LimitPeriod_MONTHLY LimitPeriod = 3
	LimitPeriod_YEARLY  LimitPeriod = 4
)

var LimitPeriod_name = map[int32]string{
	0: "NONE",
	1: "DAILY",
	2: "WEEKLY",
	3: "MONTHLY",
	4: "YEARLY",
}

var LimitPeriod_value = map[string]int32{
	"NONE":    0,
	"DAILY":   1,
	"WEEKLY":  2,
	"MONTHLY": 3,
	"YEARLY":  4,
}

func (x LimitPeriod) String() string {
	return proto.EnumName(LimitPeriod_name, int32(x))
}

func (LimitPeriod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_decc4855205ef000, []int{0}
}

type SetBridgeTransferLimitProposal struct {
	Title           string                `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description     string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Token           string                `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Limit           cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=limit,proto3,customtype=cosmossdk.io/math.Int" json:"limit"`
	LimitPeriod     LimitPeriod           `protobuf:"varint,5,opt,name=limit_period,json=limitPeriod,proto3,enum=palomachain.paloma.gravity.LimitPeriod" json:"limit_period,omitempty"`
	ExemptAddresses []string              `protobuf:"bytes,6,rep,name=exempt_addresses,json=exemptAddresses,proto3" json:"exempt_addresses,omitempty"`
}

func (m *SetBridgeTransferLimitProposal) Reset()         { *m = SetBridgeTransferLimitProposal{} }
func (m *SetBridgeTransferLimitProposal) String() string { return proto.CompactTextString(m) }
func (*SetBridgeTransferLimitProposal) ProtoMessage()    {}
func (*SetBridgeTransferLimitProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_decc4855205ef000, []int{0}
}

func (m *SetBridgeTransferLimitProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SetBridgeTransferLimitProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetBridgeTransferLimitProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SetBridgeTransferLimitProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBridgeTransferLimitProposal.Merge(m, src)
}

func (m *SetBridgeTransferLimitProposal) XXX_Size() int {
	return m.Size()
}

func (m *SetBridgeTransferLimitProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBridgeTransferLimitProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetBridgeTransferLimitProposal proto.InternalMessageInfo

func (m *SetBridgeTransferLimitProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetBridgeTransferLimitProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SetBridgeTransferLimitProposal) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SetBridgeTransferLimitProposal) GetLimitPeriod() LimitPeriod {
	if m != nil {
		return m.LimitPeriod
	}
	return LimitPeriod_NONE
}

func (m *SetBridgeTransferLimitProposal) GetExemptAddresses() []string {
	if m != nil {
		return m.ExemptAddresses
	}
	return nil
}

func init() {
	proto.RegisterEnum("palomachain.paloma.gravity.LimitPeriod", LimitPeriod_name, LimitPeriod_value)
	proto.RegisterType((*SetBridgeTransferLimitProposal)(nil), "palomachain.paloma.gravity.SetBridgeTransferLimitProposal")
}

func init() {
	proto.RegisterFile("palomachain/paloma/gravity/bridge_transfer_limit_proposal.proto", fileDescriptor_decc4855205ef000)
}

var fileDescriptor_decc4855205ef000 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x93, 0xb6, 0xa9, 0x76, 0x22, 0x1a, 0x86, 0x15, 0xc2, 0x82, 0xd9, 0xe0, 0xc5, 0xea,
	0x61, 0x06, 0x76, 0x0f, 0x1e, 0xa5, 0xc5, 0xa0, 0xab, 0xb1, 0xab, 0x71, 0x41, 0xe2, 0x25, 0x4c,
	0x93, 0x31, 0x1d, 0x9a, 0x64, 0xc2, 0xcc, 0x58, 0x9b, 0x6f, 0xe1, 0xd9, 0x4f, 0xd4, 0x63, 0x8f,
	0xe2, 0xa1, 0x48, 0xfb, 0x45, 0x24, 0x7f, 0x8a, 0x3d, 0xe8, 0xed, 0x7d, 0x9f, 0xf7, 0x7d, 0x7e,
	0xcc, 0xc3, 0x3b, 0xe0, 0x45, 0x49, 0x32, 0x9e, 0x93, 0x78, 0x41, 0x58, 0x81, 0xdb, 0x1a, 0xa7,
	0x82, 0xac, 0x98, 0xaa, 0xf0, 0x5c, 0xb0, 0x24, 0xa5, 0x91, 0x12, 0xa4, 0x90, 0x5f, 0xa8, 0x88,
	0x32, 0x96, 0x33, 0x15, 0x95, 0x82, 0x97, 0x5c, 0x92, 0x0c, 0x95, 0x82, 0x2b, 0x0e, 0xcf, 0x4f,
	0x00, 0xa8, 0xad, 0x51, 0x07, 0x38, 0x3f, 0x4b, 0x79, 0xca, 0x9b, 0x35, 0x5c, 0x57, 0xad, 0xe3,
	0xf1, 0x8f, 0x1e, 0x70, 0x3e, 0x52, 0x35, 0x6d, 0xe8, 0xb7, 0x1d, 0xdc, 0xaf, 0xd9, 0xef, 0x3b,
	0x34, 0x3c, 0x03, 0x86, 0x62, 0x2a, 0xa3, 0xb6, 0xee, 0xea, 0xe3, 0x51, 0xd0, 0x36, 0xd0, 0x05,
	0x66, 0x42, 0x65, 0x2c, 0x58, 0xa9, 0x18, 0x2f, 0xec, 0x5e, 0x33, 0x3b, 0x95, 0x1a, 0x1f, 0x5f,
	0xd2, 0xc2, 0xee, 0x77, 0xbe, 0xba, 0x81, 0x57, 0xc0, 0x68, 0x9e, 0x6e, 0x0f, 0x6a, 0x75, 0xfa,
	0x68, 0xb3, 0xbb, 0xd0, 0x7e, 0xed, 0x2e, 0x1e, 0xc6, 0x5c, 0xe6, 0x5c, 0xca, 0x64, 0x89, 0x18,
	0xc7, 0x39, 0x51, 0x0b, 0x74, 0x5d, 0xa8, 0xa0, 0xdd, 0x85, 0x6f, 0xc0, 0xbd, 0x2e, 0x2f, 0x15,
	0x8c, 0x27, 0xb6, 0xe1, 0xea, 0xe3, 0xfb, 0x97, 0x4f, 0xd0, 0xff, 0xe3, 0xa2, 0x36, 0x43, 0xb3,
	0x1e, 0x98, 0xd9, 0xdf, 0x06, 0x3e, 0x05, 0x16, 0x5d, 0xd3, 0xbc, 0x54, 0x11, 0x49, 0x12, 0x41,
	0xa5, 0xa4, 0xd2, 0x1e, 0xba, 0xfd, 0xf1, 0x28, 0x78, 0xd0, 0xea, 0x93, 0xa3, 0xfc, 0xec, 0x15,
	0x30, 0x4f, 0x30, 0xf0, 0x2e, 0x18, 0xcc, 0x6e, 0x66, 0x9e, 0xa5, 0xc1, 0x11, 0x30, 0x5e, 0x4e,
	0xae, 0xfd, 0xd0, 0xd2, 0x21, 0x00, 0xc3, 0x4f, 0x9e, 0xf7, 0xd6, 0x0f, 0xad, 0x1e, 0x34, 0xc1,
	0x9d, 0x77, 0x37, 0xb3, 0xdb, 0xd7, 0x7e, 0x68, 0xf5, 0xeb, 0x41, 0xe8, 0x4d, 0x02, 0x3f, 0xb4,
	0x06, 0xd3, 0x0f, 0x9b, 0xbd, 0xa3, 0x6f, 0xf7, 0x8e, 0xfe, 0x7b, 0xef, 0xe8, 0xdf, 0x0f, 0x8e,
	0xb6, 0x3d, 0x38, 0xda, 0xcf, 0x83, 0xa3, 0x7d, 0x7e, 0x9e, 0x32, 0xb5, 0xf8, 0x3a, 0x47, 0x31,
	0xcf, 0xf1, 0x3f, 0xae, 0xbf, 0xba, 0xc4, 0x6b, 0x2c, 0x97, 0xd5, 0x37, 0x52, 0x61, 0x55, 0x95,
	0x54, 0x1e, 0xff, 0xc3, 0x7c, 0xd8, 0xdc, 0xef, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25,
	0xc9, 0x92, 0x96, 0x34, 0x02, 0x00, 0x00,
}

func (m *SetBridgeTransferLimitProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetBridgeTransferLimitProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetBridgeTransferLimitProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExemptAddresses) > 0 {
		for iNdEx := len(m.ExemptAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ExemptAddresses[iNdEx])
			copy(dAtA[i:], m.ExemptAddresses[iNdEx])
			i = encodeVarintBridgeTransferLimitProposal(dAtA, i, uint64(len(m.ExemptAddresses[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.LimitPeriod != 0 {
		i = encodeVarintBridgeTransferLimitProposal(dAtA, i, uint64(m.LimitPeriod))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Limit.Size()
		i -= size
		if _, err := m.Limit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBridgeTransferLimitProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintBridgeTransferLimitProposal(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintBridgeTransferLimitProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintBridgeTransferLimitProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBridgeTransferLimitProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovBridgeTransferLimitProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *SetBridgeTransferLimitProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovBridgeTransferLimitProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovBridgeTransferLimitProposal(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovBridgeTransferLimitProposal(uint64(l))
	}
	l = m.Limit.Size()
	n += 1 + l + sovBridgeTransferLimitProposal(uint64(l))
	if m.LimitPeriod != 0 {
		n += 1 + sovBridgeTransferLimitProposal(uint64(m.LimitPeriod))
	}
	if len(m.ExemptAddresses) > 0 {
		for _, s := range m.ExemptAddresses {
			l = len(s)
			n += 1 + l + sovBridgeTransferLimitProposal(uint64(l))
		}
	}
	return n
}

func sovBridgeTransferLimitProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozBridgeTransferLimitProposal(x uint64) (n int) {
	return sovBridgeTransferLimitProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *SetBridgeTransferLimitProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridgeTransferLimitProposal
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
			return fmt.Errorf("proto: SetBridgeTransferLimitProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetBridgeTransferLimitProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeTransferLimitProposal
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
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeTransferLimitProposal
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
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeTransferLimitProposal
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
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeTransferLimitProposal
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
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Limit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitPeriod", wireType)
			}
			m.LimitPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeTransferLimitProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitPeriod |= LimitPeriod(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExemptAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeTransferLimitProposal
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
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeTransferLimitProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExemptAddresses = append(m.ExemptAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBridgeTransferLimitProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridgeTransferLimitProposal
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

func skipBridgeTransferLimitProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBridgeTransferLimitProposal
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
					return 0, ErrIntOverflowBridgeTransferLimitProposal
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
					return 0, ErrIntOverflowBridgeTransferLimitProposal
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
				return 0, ErrInvalidLengthBridgeTransferLimitProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBridgeTransferLimitProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBridgeTransferLimitProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBridgeTransferLimitProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBridgeTransferLimitProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBridgeTransferLimitProposal = fmt.Errorf("proto: unexpected end of group")
)
