// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signature.proto

package verifier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MAC int32

const (
	MAC_HMAC_SHA2_256 MAC = 0
	MAC_HMAC_SHA2_512 MAC = 1
	MAC_HMAC_SHA3_256 MAC = 2
	MAC_HMAC_SHA3_512 MAC = 3
	MAC_POLY1305_AES  MAC = 4
)

var MAC_name = map[int32]string{
	0: "HMAC_SHA2_256",
	1: "HMAC_SHA2_512",
	2: "HMAC_SHA3_256",
	3: "HMAC_SHA3_512",
	4: "POLY1305_AES",
}
var MAC_value = map[string]int32{
	"HMAC_SHA2_256": 0,
	"HMAC_SHA2_512": 1,
	"HMAC_SHA3_256": 2,
	"HMAC_SHA3_512": 3,
	"POLY1305_AES":  4,
}

func (x MAC) String() string {
	return proto.EnumName(MAC_name, int32(x))
}
func (MAC) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{0}
}

type HashAlgorithm int32

const (
	HashAlgorithm_SHA2_256 HashAlgorithm = 0
	HashAlgorithm_SHA2_512 HashAlgorithm = 1
	HashAlgorithm_SHA3_256 HashAlgorithm = 2
	HashAlgorithm_SHA3_512 HashAlgorithm = 3
)

var HashAlgorithm_name = map[int32]string{
	0: "SHA2_256",
	1: "SHA2_512",
	2: "SHA3_256",
	3: "SHA3_512",
}
var HashAlgorithm_value = map[string]int32{
	"SHA2_256": 0,
	"SHA2_512": 1,
	"SHA3_256": 2,
	"SHA3_512": 3,
}

func (x HashAlgorithm) String() string {
	return proto.EnumName(HashAlgorithm_name, int32(x))
}
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{1}
}

type SignatureLevel int32

const (
	SignatureLevel_ADVANCED  SignatureLevel = 0
	SignatureLevel_QUALIFIED SignatureLevel = 1
)

var SignatureLevel_name = map[int32]string{
	0: "ADVANCED",
	1: "QUALIFIED",
}
var SignatureLevel_value = map[string]int32{
	"ADVANCED":  0,
	"QUALIFIED": 1,
}

func (x SignatureLevel) String() string {
	return proto.EnumName(SignatureLevel_name, int32(x))
}
func (SignatureLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{2}
}

type SignatureData struct {
	DocumentHash         []byte         `protobuf:"bytes,1,opt,name=document_hash,json=documentHash,proto3" json:"document_hash,omitempty"`
	HashAlgorithm        HashAlgorithm  `protobuf:"varint,2,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=HashAlgorithm" json:"hash_algorithm,omitempty"`
	MacKey               []byte         `protobuf:"bytes,3,opt,name=mac_key,json=macKey,proto3" json:"mac_key,omitempty"`
	Mac                  MAC            `protobuf:"varint,4,opt,name=mac,proto3,enum=MAC" json:"mac,omitempty"`
	OtherMacs            [][]byte       `protobuf:"bytes,5,rep,name=other_macs,json=otherMacs,proto3" json:"other_macs,omitempty"`
	SignatureLevel       SignatureLevel `protobuf:"varint,6,opt,name=signature_level,json=signatureLevel,proto3,enum=SignatureLevel" json:"signature_level,omitempty"`
	IdToken              []byte         `protobuf:"bytes,7,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	X509Idp              [][]byte       `protobuf:"bytes,8,rep,name=x509_idp,json=x509Idp,proto3" json:"x509_idp,omitempty"`
	LtvIdpCa             *LTV           `protobuf:"bytes,9,opt,name=ltv_idp_ca,json=ltvIdpCa,proto3" json:"ltv_idp_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignatureData) Reset()         { *m = SignatureData{} }
func (m *SignatureData) String() string { return proto.CompactTextString(m) }
func (*SignatureData) ProtoMessage()    {}
func (*SignatureData) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{0}
}
func (m *SignatureData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureData.Unmarshal(m, b)
}
func (m *SignatureData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureData.Marshal(b, m, deterministic)
}
func (dst *SignatureData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureData.Merge(dst, src)
}
func (m *SignatureData) XXX_Size() int {
	return xxx_messageInfo_SignatureData.Size(m)
}
func (m *SignatureData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureData.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureData proto.InternalMessageInfo

func (m *SignatureData) GetDocumentHash() []byte {
	if m != nil {
		return m.DocumentHash
	}
	return nil
}

func (m *SignatureData) GetHashAlgorithm() HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return HashAlgorithm_SHA2_256
}

func (m *SignatureData) GetMacKey() []byte {
	if m != nil {
		return m.MacKey
	}
	return nil
}

func (m *SignatureData) GetMac() MAC {
	if m != nil {
		return m.Mac
	}
	return MAC_HMAC_SHA2_256
}

func (m *SignatureData) GetOtherMacs() [][]byte {
	if m != nil {
		return m.OtherMacs
	}
	return nil
}

func (m *SignatureData) GetSignatureLevel() SignatureLevel {
	if m != nil {
		return m.SignatureLevel
	}
	return SignatureLevel_ADVANCED
}

func (m *SignatureData) GetIdToken() []byte {
	if m != nil {
		return m.IdToken
	}
	return nil
}

func (m *SignatureData) GetX509Idp() [][]byte {
	if m != nil {
		return m.X509Idp
	}
	return nil
}

func (m *SignatureData) GetLtvIdpCa() *LTV {
	if m != nil {
		return m.LtvIdpCa
	}
	return nil
}

type Timestamped struct {
	Rfc3161Timestamp     []byte   `protobuf:"bytes,1,opt,name=rfc3161_timestamp,json=rfc3161Timestamp,proto3" json:"rfc3161_timestamp,omitempty"`
	LtvTimestampCa       *LTV     `protobuf:"bytes,2,opt,name=ltv_timestamp_ca,json=ltvTimestampCa,proto3" json:"ltv_timestamp_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamped) Reset()         { *m = Timestamped{} }
func (m *Timestamped) String() string { return proto.CompactTextString(m) }
func (*Timestamped) ProtoMessage()    {}
func (*Timestamped) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{1}
}
func (m *Timestamped) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamped.Unmarshal(m, b)
}
func (m *Timestamped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamped.Marshal(b, m, deterministic)
}
func (dst *Timestamped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamped.Merge(dst, src)
}
func (m *Timestamped) XXX_Size() int {
	return xxx_messageInfo_Timestamped.Size(m)
}
func (m *Timestamped) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamped.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamped proto.InternalMessageInfo

func (m *Timestamped) GetRfc3161Timestamp() []byte {
	if m != nil {
		return m.Rfc3161Timestamp
	}
	return nil
}

func (m *Timestamped) GetLtvTimestampCa() *LTV {
	if m != nil {
		return m.LtvTimestampCa
	}
	return nil
}

type SignatureContainer struct {
	EnvelopedSignatureDataPkcs7 []byte   `protobuf:"bytes,1,opt,name=enveloped_signature_data_pkcs7,json=envelopedSignatureDataPkcs7,proto3" json:"enveloped_signature_data_pkcs7,omitempty"`
	LtvSigningCa                *LTV     `protobuf:"bytes,2,opt,name=ltv_signing_ca,json=ltvSigningCa,proto3" json:"ltv_signing_ca,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *SignatureContainer) Reset()         { *m = SignatureContainer{} }
func (m *SignatureContainer) String() string { return proto.CompactTextString(m) }
func (*SignatureContainer) ProtoMessage()    {}
func (*SignatureContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{2}
}
func (m *SignatureContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureContainer.Unmarshal(m, b)
}
func (m *SignatureContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureContainer.Marshal(b, m, deterministic)
}
func (dst *SignatureContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureContainer.Merge(dst, src)
}
func (m *SignatureContainer) XXX_Size() int {
	return xxx_messageInfo_SignatureContainer.Size(m)
}
func (m *SignatureContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureContainer.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureContainer proto.InternalMessageInfo

func (m *SignatureContainer) GetEnvelopedSignatureDataPkcs7() []byte {
	if m != nil {
		return m.EnvelopedSignatureDataPkcs7
	}
	return nil
}

func (m *SignatureContainer) GetLtvSigningCa() *LTV {
	if m != nil {
		return m.LtvSigningCa
	}
	return nil
}

type LTV struct {
	Ocsp                 []byte   `protobuf:"bytes,1,opt,name=ocsp,proto3" json:"ocsp,omitempty"`
	Crl                  []byte   `protobuf:"bytes,2,opt,name=crl,proto3" json:"crl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LTV) Reset()         { *m = LTV{} }
func (m *LTV) String() string { return proto.CompactTextString(m) }
func (*LTV) ProtoMessage()    {}
func (*LTV) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{3}
}
func (m *LTV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LTV.Unmarshal(m, b)
}
func (m *LTV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LTV.Marshal(b, m, deterministic)
}
func (dst *LTV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LTV.Merge(dst, src)
}
func (m *LTV) XXX_Size() int {
	return xxx_messageInfo_LTV.Size(m)
}
func (m *LTV) XXX_DiscardUnknown() {
	xxx_messageInfo_LTV.DiscardUnknown(m)
}

var xxx_messageInfo_LTV proto.InternalMessageInfo

func (m *LTV) GetOcsp() []byte {
	if m != nil {
		return m.Ocsp
	}
	return nil
}

func (m *LTV) GetCrl() []byte {
	if m != nil {
		return m.Crl
	}
	return nil
}

type SignatureFile struct {
	SignatureContainer   *SignatureContainer `protobuf:"bytes,1,opt,name=signature_container,json=signatureContainer,proto3" json:"signature_container,omitempty"`
	Timestamps           []*Timestamped      `protobuf:"bytes,2,rep,name=timestamps,proto3" json:"timestamps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SignatureFile) Reset()         { *m = SignatureFile{} }
func (m *SignatureFile) String() string { return proto.CompactTextString(m) }
func (*SignatureFile) ProtoMessage()    {}
func (*SignatureFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_fbfe2b1ca6703e24, []int{4}
}
func (m *SignatureFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureFile.Unmarshal(m, b)
}
func (m *SignatureFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureFile.Marshal(b, m, deterministic)
}
func (dst *SignatureFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureFile.Merge(dst, src)
}
func (m *SignatureFile) XXX_Size() int {
	return xxx_messageInfo_SignatureFile.Size(m)
}
func (m *SignatureFile) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureFile.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureFile proto.InternalMessageInfo

func (m *SignatureFile) GetSignatureContainer() *SignatureContainer {
	if m != nil {
		return m.SignatureContainer
	}
	return nil
}

func (m *SignatureFile) GetTimestamps() []*Timestamped {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func init() {
	proto.RegisterType((*SignatureData)(nil), "SignatureData")
	proto.RegisterType((*Timestamped)(nil), "Timestamped")
	proto.RegisterType((*SignatureContainer)(nil), "SignatureContainer")
	proto.RegisterType((*LTV)(nil), "LTV")
	proto.RegisterType((*SignatureFile)(nil), "SignatureFile")
	proto.RegisterEnum("MAC", MAC_name, MAC_value)
	proto.RegisterEnum("HashAlgorithm", HashAlgorithm_name, HashAlgorithm_value)
	proto.RegisterEnum("SignatureLevel", SignatureLevel_name, SignatureLevel_value)
}

func init() { proto.RegisterFile("signature.proto", fileDescriptor_signature_fbfe2b1ca6703e24) }

var fileDescriptor_signature_fbfe2b1ca6703e24 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0x9a, 0xb2, 0x76, 0xb7, 0x69, 0x97, 0x79, 0x12, 0x04, 0x21, 0x50, 0x55, 0x5e, 0xaa,
	0x0e, 0xa2, 0x35, 0x55, 0x07, 0x3c, 0x86, 0x74, 0x63, 0x15, 0x2d, 0x8c, 0xb4, 0x4c, 0x82, 0x17,
	0xcb, 0x38, 0xde, 0x1a, 0x96, 0x2f, 0x25, 0x5e, 0xc4, 0x9e, 0x11, 0xbf, 0x8f, 0xbf, 0x84, 0xec,
	0x2d, 0x59, 0xc2, 0xde, 0x72, 0x3e, 0x72, 0xee, 0x3d, 0x8e, 0x03, 0xbb, 0x99, 0x7f, 0x19, 0x11,
	0x7e, 0x9d, 0x32, 0x33, 0x49, 0x63, 0x1e, 0x0f, 0xfe, 0x36, 0xa0, 0xbb, 0x2a, 0xb8, 0x19, 0xe1,
	0x04, 0xbd, 0x84, 0xae, 0x17, 0xd3, 0xeb, 0x90, 0x45, 0x1c, 0x6f, 0x48, 0xb6, 0x31, 0x94, 0xbe,
	0x32, 0xd4, 0x5c, 0xad, 0x20, 0x4f, 0x49, 0xb6, 0x41, 0x53, 0xe8, 0x09, 0x0d, 0x93, 0xe0, 0x32,
	0x4e, 0x7d, 0xbe, 0x09, 0x8d, 0x46, 0x5f, 0x19, 0xf6, 0xac, 0x9e, 0x29, 0x64, 0xbb, 0x60, 0xdd,
	0xee, 0xa6, 0x0a, 0xd1, 0x13, 0x68, 0x85, 0x84, 0xe2, 0x2b, 0x76, 0x63, 0xa8, 0x32, 0x75, 0x3b,
	0x24, 0xf4, 0x23, 0xbb, 0x41, 0x8f, 0x41, 0x0d, 0x09, 0x35, 0x9a, 0x32, 0xa4, 0x69, 0x2e, 0x6d,
	0xc7, 0x15, 0x04, 0x7a, 0x0e, 0x10, 0xf3, 0x0d, 0x4b, 0x71, 0x48, 0x68, 0x66, 0x3c, 0xea, 0xab,
	0x43, 0xcd, 0xdd, 0x91, 0xcc, 0x92, 0xd0, 0x0c, 0xbd, 0xad, 0x14, 0xc2, 0x01, 0xcb, 0x59, 0x60,
	0x6c, 0xcb, 0x88, 0x5d, 0xb3, 0x2c, 0xb5, 0x10, 0xb4, 0xdb, 0xcb, 0x6a, 0x18, 0x3d, 0x85, 0xb6,
	0xef, 0x61, 0x1e, 0x5f, 0xb1, 0xc8, 0x68, 0xc9, 0x55, 0x5a, 0xbe, 0xb7, 0x16, 0x50, 0x48, 0xbf,
	0xa6, 0x87, 0xef, 0xb0, 0xef, 0x25, 0x46, 0x5b, 0x4e, 0x6c, 0x09, 0x3c, 0xf7, 0x12, 0x34, 0x00,
	0x08, 0x78, 0x2e, 0x14, 0x4c, 0x89, 0xb1, 0xd3, 0x57, 0x86, 0x1d, 0xab, 0x69, 0x2e, 0xd6, 0xe7,
	0x6e, 0x3b, 0xe0, 0xf9, 0xdc, 0x4b, 0x1c, 0x32, 0xf8, 0x09, 0x9d, 0xb5, 0x1f, 0xb2, 0x8c, 0x93,
	0x30, 0x61, 0x1e, 0x3a, 0x80, 0xbd, 0xf4, 0x82, 0x4e, 0xc6, 0x47, 0x63, 0xcc, 0x0b, 0xfa, 0xee,
	0x48, 0xf5, 0x3b, 0xa1, 0xb4, 0x23, 0x13, 0x74, 0x91, 0x5f, 0x1a, 0xc5, 0x94, 0x46, 0x65, 0x4a,
	0x2f, 0xe0, 0x79, 0xe9, 0x76, 0xc8, 0xe0, 0x8f, 0x02, 0xa8, 0x2c, 0xea, 0xc4, 0x11, 0x27, 0x7e,
	0xc4, 0x52, 0xe4, 0xc0, 0x0b, 0x16, 0xe5, 0x2c, 0x88, 0x13, 0xe6, 0xe1, 0xfb, 0x03, 0xf2, 0x08,
	0x27, 0x38, 0xb9, 0xa2, 0xd9, 0x9b, 0xbb, 0x05, 0x9e, 0x95, 0xae, 0xda, 0x15, 0x38, 0x13, 0x16,
	0x34, 0x02, 0x31, 0x4d, 0xbe, 0xee, 0x47, 0x97, 0xff, 0x6f, 0xa2, 0x05, 0x3c, 0x5f, 0xdd, 0x4a,
	0x0e, 0x19, 0x1c, 0x80, 0xba, 0x58, 0x9f, 0x23, 0x04, 0xcd, 0x98, 0x66, 0x45, 0x3d, 0xf9, 0x8c,
	0x74, 0x50, 0x69, 0x1a, 0xc8, 0x77, 0x35, 0x57, 0x3c, 0x0e, 0x7e, 0x2b, 0x95, 0x2b, 0x77, 0xe2,
	0x07, 0x0c, 0xcd, 0x60, 0xff, 0x7e, 0x4b, 0x5a, 0xd4, 0x90, 0x31, 0x1d, 0x6b, 0xdf, 0x7c, 0xd8,
	0xd0, 0x45, 0xd9, 0xc3, 0xd6, 0xaf, 0x00, 0xca, 0x83, 0xcb, 0x8c, 0x46, 0x5f, 0x1d, 0x76, 0x2c,
	0xcd, 0xac, 0x7c, 0x0b, 0xb7, 0xa2, 0x8f, 0x28, 0xa8, 0x4b, 0xdb, 0x41, 0x7b, 0xd0, 0x3d, 0x5d,
	0xda, 0x0e, 0x5e, 0x9d, 0xda, 0x16, 0xb6, 0xa6, 0x47, 0xfa, 0x56, 0x9d, 0x9a, 0x8e, 0x2d, 0x5d,
	0xa9, 0x52, 0x13, 0xe9, 0x6a, 0xd4, 0x29, 0xe1, 0x52, 0x91, 0x0e, 0xda, 0xd9, 0xe7, 0xc5, 0xb7,
	0xf1, 0xe4, 0x70, 0x8a, 0xed, 0xe3, 0x95, 0xde, 0x1c, 0x7d, 0x80, 0x6e, 0xed, 0x7f, 0x40, 0x1a,
	0xb4, 0x2b, 0x93, 0x0a, 0x74, 0x3b, 0xe4, 0x16, 0x15, 0xf9, 0x05, 0x92, 0xd1, 0xa3, 0xd7, 0xd0,
	0xab, 0x5f, 0x68, 0xa1, 0xdb, 0xb3, 0x73, 0xfb, 0x93, 0x73, 0x3c, 0xd3, 0xb7, 0x50, 0x17, 0x76,
	0xbe, 0x7c, 0xb5, 0x17, 0xf3, 0x93, 0xf9, 0xf1, 0x4c, 0x57, 0xde, 0xc3, 0xf7, 0x76, 0xce, 0x52,
	0xff, 0xc2, 0x67, 0xe9, 0x8f, 0x6d, 0xf9, 0xa3, 0x4f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x73, 0xba, 0xdb, 0xfb, 0x03, 0x00, 0x00,
}
