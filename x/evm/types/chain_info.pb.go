// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/chain_info.proto

package types

import (
	fmt "fmt"
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

type ChainInfo_Status int32

const (
	ChainInfo_IN_PROPOSAL          ChainInfo_Status = 0
	ChainInfo_ACTIVE               ChainInfo_Status = 1
	ChainInfo_WAITING_FOR_EVIDENCE ChainInfo_Status = 2
)

var ChainInfo_Status_name = map[int32]string{
	0: "IN_PROPOSAL",
	1: "ACTIVE",
	2: "WAITING_FOR_EVIDENCE",
}

var ChainInfo_Status_value = map[string]int32{
	"IN_PROPOSAL":          0,
	"ACTIVE":               1,
	"WAITING_FOR_EVIDENCE": 2,
}

func (x ChainInfo_Status) String() string {
	return proto.EnumName(ChainInfo_Status_name, int32(x))
}

func (ChainInfo_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_439fc96864abe15f, []int{0, 0}
}

type ChainInfo struct {
	Id                    uint64 `protobuf:"varint,9999,opt,name=id,proto3" json:"id,omitempty"`
	ChainReferenceID      string `protobuf:"bytes,1,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	ChainID               uint64 `protobuf:"varint,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	SmartContractUniqueID []byte `protobuf:"bytes,3,opt,name=smartContractUniqueID,proto3" json:"smartContractUniqueID,omitempty"`
	SmartContractAddr     string `protobuf:"bytes,4,opt,name=smartContractAddr,proto3" json:"smartContractAddr,omitempty"`
	// used to verify by pigeons if they are at the correct chain
	ReferenceBlockHeight          uint64           `protobuf:"varint,5,opt,name=referenceBlockHeight,proto3" json:"referenceBlockHeight,omitempty"`
	ReferenceBlockHash            string           `protobuf:"bytes,6,opt,name=referenceBlockHash,proto3" json:"referenceBlockHash,omitempty"`
	Abi                           string           `protobuf:"bytes,7,opt,name=abi,proto3" json:"abi,omitempty"`
	Bytecode                      []byte           `protobuf:"bytes,8,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	ConstructorInput              []byte           `protobuf:"bytes,9,opt,name=constructorInput,proto3" json:"constructorInput,omitempty"`
	Status                        ChainInfo_Status `protobuf:"varint,10,opt,name=status,proto3,enum=palomachain.paloma.evm.ChainInfo_Status" json:"status,omitempty"`
	ActiveSmartContractID         uint64           `protobuf:"varint,11,opt,name=activeSmartContractID,proto3" json:"activeSmartContractID,omitempty"`
	SmartContractDeployingVersion uint64           `protobuf:"varint,12,opt,name=smartContractDeployingVersion,proto3" json:"smartContractDeployingVersion,omitempty"`
}

func (m *ChainInfo) Reset()         { *m = ChainInfo{} }
func (m *ChainInfo) String() string { return proto.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()    {}
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_439fc96864abe15f, []int{0}
}
func (m *ChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainInfo.Merge(m, src)
}
func (m *ChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainInfo proto.InternalMessageInfo

func (m *ChainInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChainInfo) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *ChainInfo) GetChainID() uint64 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *ChainInfo) GetSmartContractUniqueID() []byte {
	if m != nil {
		return m.SmartContractUniqueID
	}
	return nil
}

func (m *ChainInfo) GetSmartContractAddr() string {
	if m != nil {
		return m.SmartContractAddr
	}
	return ""
}

func (m *ChainInfo) GetReferenceBlockHeight() uint64 {
	if m != nil {
		return m.ReferenceBlockHeight
	}
	return 0
}

func (m *ChainInfo) GetReferenceBlockHash() string {
	if m != nil {
		return m.ReferenceBlockHash
	}
	return ""
}

func (m *ChainInfo) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *ChainInfo) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

func (m *ChainInfo) GetConstructorInput() []byte {
	if m != nil {
		return m.ConstructorInput
	}
	return nil
}

func (m *ChainInfo) GetStatus() ChainInfo_Status {
	if m != nil {
		return m.Status
	}
	return ChainInfo_IN_PROPOSAL
}

func (m *ChainInfo) GetActiveSmartContractID() uint64 {
	if m != nil {
		return m.ActiveSmartContractID
	}
	return 0
}

func (m *ChainInfo) GetSmartContractDeployingVersion() uint64 {
	if m != nil {
		return m.SmartContractDeployingVersion
	}
	return 0
}

type SmartContract struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AbiJSON  string `protobuf:"bytes,2,opt,name=abiJSON,proto3" json:"abiJSON,omitempty"`
	Bytecode []byte `protobuf:"bytes,3,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
}

func (m *SmartContract) Reset()         { *m = SmartContract{} }
func (m *SmartContract) String() string { return proto.CompactTextString(m) }
func (*SmartContract) ProtoMessage()    {}
func (*SmartContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_439fc96864abe15f, []int{1}
}
func (m *SmartContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmartContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SmartContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContract.Merge(m, src)
}
func (m *SmartContract) XXX_Size() int {
	return m.Size()
}
func (m *SmartContract) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContract.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContract proto.InternalMessageInfo

func (m *SmartContract) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SmartContract) GetAbiJSON() string {
	if m != nil {
		return m.AbiJSON
	}
	return ""
}

func (m *SmartContract) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

type SmartContractDeployment struct {
	// which smart contract is getting deployed
	SmartContractID uint64 `protobuf:"varint,1,opt,name=smartContractID,proto3" json:"smartContractID,omitempty"`
	// to which chain info the smart contract is getting deployed to
	ChainReferenceID string `protobuf:"bytes,2,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	// application level ID which uniquely identifies a deployed smart contract.
	// It's used if we have multiple smart contracts deployed on a same EVM so that
	// the contract can figure out if the message was actually sent for them.
	// (message includes the unique id and smart contract has simple logic to
	// disallow those that to not match)
	UniqueID []byte `protobuf:"bytes,3,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`
}

func (m *SmartContractDeployment) Reset()         { *m = SmartContractDeployment{} }
func (m *SmartContractDeployment) String() string { return proto.CompactTextString(m) }
func (*SmartContractDeployment) ProtoMessage()    {}
func (*SmartContractDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_439fc96864abe15f, []int{2}
}
func (m *SmartContractDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartContractDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmartContractDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SmartContractDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContractDeployment.Merge(m, src)
}
func (m *SmartContractDeployment) XXX_Size() int {
	return m.Size()
}
func (m *SmartContractDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContractDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContractDeployment proto.InternalMessageInfo

func (m *SmartContractDeployment) GetSmartContractID() uint64 {
	if m != nil {
		return m.SmartContractID
	}
	return 0
}

func (m *SmartContractDeployment) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *SmartContractDeployment) GetUniqueID() []byte {
	if m != nil {
		return m.UniqueID
	}
	return nil
}

func init() {
	proto.RegisterEnum("palomachain.paloma.evm.ChainInfo_Status", ChainInfo_Status_name, ChainInfo_Status_value)
	proto.RegisterType((*ChainInfo)(nil), "palomachain.paloma.evm.ChainInfo")
	proto.RegisterType((*SmartContract)(nil), "palomachain.paloma.evm.SmartContract")
	proto.RegisterType((*SmartContractDeployment)(nil), "palomachain.paloma.evm.SmartContractDeployment")
}

func init() { proto.RegisterFile("evm/chain_info.proto", fileDescriptor_439fc96864abe15f) }

var fileDescriptor_439fc96864abe15f = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4f, 0x8f, 0x12, 0x4f,
	0x10, 0xa5, 0x61, 0x97, 0x3f, 0xb5, 0xfb, 0x5b, 0xf8, 0x75, 0x50, 0x3b, 0x9b, 0x38, 0x21, 0x9c,
	0x46, 0x63, 0x86, 0x64, 0xf5, 0xae, 0x2c, 0x83, 0x3a, 0xc6, 0xc0, 0x66, 0xd8, 0xc5, 0xc4, 0x0b,
	0x69, 0x86, 0x06, 0x3a, 0x32, 0xdd, 0x38, 0xd3, 0x43, 0xe4, 0xec, 0xd1, 0x83, 0x7e, 0x2c, 0x8f,
	0x7b, 0xf4, 0x68, 0xe0, 0x8b, 0x98, 0x69, 0x16, 0xb2, 0x03, 0xa3, 0xb7, 0xaa, 0xf7, 0xaa, 0xba,
	0xaa, 0x5f, 0x55, 0x41, 0x95, 0x2d, 0xfc, 0x86, 0x37, 0xa5, 0x5c, 0x0c, 0xb8, 0x18, 0x4b, 0x6b,
	0x1e, 0x48, 0x25, 0xf1, 0xc3, 0x39, 0x9d, 0x49, 0x9f, 0x6a, 0xdc, 0xda, 0xd8, 0x16, 0x5b, 0xf8,
	0xf5, 0xaf, 0xc7, 0x50, 0x6a, 0xc5, 0xa0, 0x23, 0xc6, 0x12, 0x97, 0x21, 0xcb, 0x47, 0xe4, 0x7b,
	0xa7, 0x86, 0xcc, 0x23, 0x37, 0xcb, 0x47, 0xf8, 0x29, 0x54, 0x74, 0x8a, 0xcb, 0xc6, 0x2c, 0x60,
	0xc2, 0x63, 0x8e, 0x4d, 0x50, 0x0d, 0x99, 0x25, 0xf7, 0x00, 0xc7, 0x04, 0x0a, 0x1a, 0x73, 0x6c,
	0x92, 0xd5, 0x0f, 0x6c, 0x5d, 0xfc, 0x02, 0x1e, 0x84, 0x3e, 0x0d, 0x54, 0x4b, 0x0a, 0x15, 0x50,
	0x4f, 0xdd, 0x08, 0xfe, 0x39, 0x8a, 0x9f, 0xca, 0xd5, 0x90, 0x79, 0xea, 0xa6, 0x93, 0xf8, 0x19,
	0xfc, 0x9f, 0x20, 0x9a, 0xa3, 0x51, 0x40, 0x8e, 0x74, 0xf1, 0x43, 0x02, 0x5f, 0x40, 0x35, 0xd8,
	0x36, 0x73, 0x39, 0x93, 0xde, 0xa7, 0xb7, 0x8c, 0x4f, 0xa6, 0x8a, 0x1c, 0xeb, 0x56, 0x52, 0x39,
	0x6c, 0x01, 0xde, 0xc3, 0x69, 0x38, 0x25, 0x79, 0x5d, 0x22, 0x85, 0xc1, 0x15, 0xc8, 0xd1, 0x21,
	0x27, 0x05, 0x1d, 0x10, 0x9b, 0xf8, 0x1c, 0x8a, 0xc3, 0xa5, 0x62, 0x9e, 0x1c, 0x31, 0x52, 0xd4,
	0x9f, 0xd9, 0xf9, 0x5a, 0x3b, 0x29, 0x42, 0x15, 0x44, 0x9e, 0x92, 0x81, 0x23, 0xe6, 0x91, 0x22,
	0x25, 0x1d, 0x73, 0x80, 0xe3, 0x57, 0x90, 0x0f, 0x15, 0x55, 0x51, 0x48, 0xa0, 0x86, 0xcc, 0xb3,
	0x0b, 0xd3, 0x4a, 0x9f, 0x97, 0xb5, 0x9b, 0x95, 0xd5, 0xd3, 0xf1, 0xee, 0x5d, 0x5e, 0xac, 0x31,
	0xf5, 0x14, 0x5f, 0xb0, 0xde, 0x7d, 0x69, 0x1c, 0x9b, 0x9c, 0x68, 0x01, 0xd2, 0x49, 0x6c, 0xc3,
	0xe3, 0x84, 0x94, 0x36, 0x9b, 0xcf, 0xe4, 0x92, 0x8b, 0x49, 0x9f, 0x05, 0x21, 0x97, 0x82, 0x9c,
	0xea, 0xec, 0x7f, 0x07, 0xd5, 0x5f, 0x42, 0x7e, 0xd3, 0x0d, 0x2e, 0xc3, 0x89, 0xd3, 0x19, 0x5c,
	0xb9, 0xdd, 0xab, 0x6e, 0xaf, 0xf9, 0xbe, 0x92, 0xc1, 0x00, 0xf9, 0x66, 0xeb, 0xda, 0xe9, 0xb7,
	0x2b, 0x08, 0x13, 0xa8, 0x7e, 0x68, 0x3a, 0xd7, 0x4e, 0xe7, 0xcd, 0xe0, 0x75, 0xd7, 0x1d, 0xb4,
	0xfb, 0x8e, 0xdd, 0xee, 0xb4, 0xda, 0x95, 0x6c, 0xfd, 0x06, 0xfe, 0x4b, 0x74, 0x86, 0xcf, 0xf4,
	0x22, 0xa2, 0xdd, 0x1e, 0x12, 0x28, 0xd0, 0x21, 0x7f, 0xd7, 0xeb, 0x76, 0xf4, 0x6e, 0x95, 0xdc,
	0xad, 0x9b, 0x98, 0x40, 0x2e, 0x39, 0x81, 0xfa, 0x37, 0x04, 0x8f, 0x7a, 0x87, 0x9d, 0xfb, 0x4c,
	0x28, 0x6c, 0x42, 0x39, 0xdc, 0x53, 0x6a, 0x53, 0x6e, 0x1f, 0x4e, 0xbd, 0x81, 0xec, 0x5f, 0x6e,
	0xe0, 0x1c, 0x8a, 0x51, 0x72, 0xb9, 0x77, 0xfe, 0x65, 0xeb, 0xe7, 0xca, 0x40, 0xb7, 0x2b, 0x03,
	0xfd, 0x5e, 0x19, 0xe8, 0xc7, 0xda, 0xc8, 0xdc, 0xae, 0x8d, 0xcc, 0xaf, 0xb5, 0x91, 0xf9, 0xf8,
	0x64, 0xc2, 0xd5, 0x34, 0x1a, 0x5a, 0x9e, 0xf4, 0x1b, 0xf7, 0xe6, 0x7e, 0x67, 0x37, 0xbe, 0x34,
	0xe2, 0x93, 0x56, 0xcb, 0x39, 0x0b, 0x87, 0x79, 0x7d, 0xce, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0x76, 0xff, 0x8d, 0xe6, 0x03, 0x00, 0x00,
}

func (m *ChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x4
		i--
		dAtA[i] = 0xf0
		i--
		dAtA[i] = 0xf8
	}
	if m.SmartContractDeployingVersion != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.SmartContractDeployingVersion))
		i--
		dAtA[i] = 0x60
	}
	if m.ActiveSmartContractID != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.ActiveSmartContractID))
		i--
		dAtA[i] = 0x58
	}
	if m.Status != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x50
	}
	if len(m.ConstructorInput) > 0 {
		i -= len(m.ConstructorInput)
		copy(dAtA[i:], m.ConstructorInput)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.ConstructorInput)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ReferenceBlockHash) > 0 {
		i -= len(m.ReferenceBlockHash)
		copy(dAtA[i:], m.ReferenceBlockHash)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.ReferenceBlockHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.ReferenceBlockHeight != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.ReferenceBlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SmartContractAddr) > 0 {
		i -= len(m.SmartContractAddr)
		copy(dAtA[i:], m.SmartContractAddr)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.SmartContractAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SmartContractUniqueID) > 0 {
		i -= len(m.SmartContractUniqueID)
		copy(dAtA[i:], m.SmartContractUniqueID)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.SmartContractUniqueID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainID != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.ChainID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SmartContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AbiJSON) > 0 {
		i -= len(m.AbiJSON)
		copy(dAtA[i:], m.AbiJSON)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.AbiJSON)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SmartContractDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartContractDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartContractDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UniqueID) > 0 {
		i -= len(m.UniqueID)
		copy(dAtA[i:], m.UniqueID)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.UniqueID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintChainInfo(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x12
	}
	if m.SmartContractID != 0 {
		i = encodeVarintChainInfo(dAtA, i, uint64(m.SmartContractID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChainInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovChainInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	if m.ChainID != 0 {
		n += 1 + sovChainInfo(uint64(m.ChainID))
	}
	l = len(m.SmartContractUniqueID)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.SmartContractAddr)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	if m.ReferenceBlockHeight != 0 {
		n += 1 + sovChainInfo(uint64(m.ReferenceBlockHeight))
	}
	l = len(m.ReferenceBlockHash)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.ConstructorInput)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovChainInfo(uint64(m.Status))
	}
	if m.ActiveSmartContractID != 0 {
		n += 1 + sovChainInfo(uint64(m.ActiveSmartContractID))
	}
	if m.SmartContractDeployingVersion != 0 {
		n += 1 + sovChainInfo(uint64(m.SmartContractDeployingVersion))
	}
	if m.Id != 0 {
		n += 3 + sovChainInfo(uint64(m.Id))
	}
	return n
}

func (m *SmartContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovChainInfo(uint64(m.Id))
	}
	l = len(m.AbiJSON)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	return n
}

func (m *SmartContractDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SmartContractID != 0 {
		n += 1 + sovChainInfo(uint64(m.SmartContractID))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	l = len(m.UniqueID)
	if l > 0 {
		n += 1 + l + sovChainInfo(uint64(l))
	}
	return n
}

func sovChainInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainInfo(x uint64) (n int) {
	return sovChainInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainInfo
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
			return fmt.Errorf("proto: ChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			m.ChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContractUniqueID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SmartContractUniqueID = append(m.SmartContractUniqueID[:0], dAtA[iNdEx:postIndex]...)
			if m.SmartContractUniqueID == nil {
				m.SmartContractUniqueID = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContractAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SmartContractAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceBlockHeight", wireType)
			}
			m.ReferenceBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReferenceBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceBlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReferenceBlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstructorInput", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConstructorInput = append(m.ConstructorInput[:0], dAtA[iNdEx:postIndex]...)
			if m.ConstructorInput == nil {
				m.ConstructorInput = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ChainInfo_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveSmartContractID", wireType)
			}
			m.ActiveSmartContractID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveSmartContractID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContractDeployingVersion", wireType)
			}
			m.SmartContractDeployingVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SmartContractDeployingVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9999:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChainInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainInfo
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
func (m *SmartContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainInfo
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
			return fmt.Errorf("proto: SmartContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbiJSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbiJSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChainInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainInfo
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
func (m *SmartContractDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainInfo
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
			return fmt.Errorf("proto: SmartContractDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartContractDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContractID", wireType)
			}
			m.SmartContractID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SmartContractID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniqueID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainInfo
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
				return ErrInvalidLengthChainInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniqueID = append(m.UniqueID[:0], dAtA[iNdEx:postIndex]...)
			if m.UniqueID == nil {
				m.UniqueID = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChainInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainInfo
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
func skipChainInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChainInfo
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
					return 0, ErrIntOverflowChainInfo
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
					return 0, ErrIntOverflowChainInfo
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
				return 0, ErrInvalidLengthChainInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChainInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChainInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChainInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChainInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChainInfo = fmt.Errorf("proto: unexpected end of group")
)
