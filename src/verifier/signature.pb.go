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
	return fileDescriptor_signature_2433630c08b6c1ea, []int{0}
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
	return fileDescriptor_signature_2433630c08b6c1ea, []int{1}
}

type SignatureData struct {
	Hash                 []byte         `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	SaltedHashes         [][]byte       `protobuf:"bytes,2,rep,name=salted_hashes,json=saltedHashes,proto3" json:"salted_hashes,omitempty"`
	HashAlgorithm        HashAlgorithm  `protobuf:"varint,3,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=HashAlgorithm" json:"hash_algorithm,omitempty"`
	Salt                 []byte         `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	SignatureLevel       SignatureLevel `protobuf:"varint,5,opt,name=signature_level,json=signatureLevel,proto3,enum=SignatureLevel" json:"signature_level,omitempty"`
	X509                 []byte         `protobuf:"bytes,6,opt,name=x509,proto3" json:"x509,omitempty"`
	X509Ca               []byte         `protobuf:"bytes,7,opt,name=x509_ca,json=x509Ca,proto3" json:"x509_ca,omitempty"`
	Ocsp                 []byte         `protobuf:"bytes,8,opt,name=ocsp,proto3" json:"ocsp,omitempty"`
	Crl                  []byte         `protobuf:"bytes,9,opt,name=crl,proto3" json:"crl,omitempty"`
	IdToken              []byte         `protobuf:"bytes,10,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	X509Idp              []byte         `protobuf:"bytes,11,opt,name=x509_idp,json=x509Idp,proto3" json:"x509_idp,omitempty"`
	X509IdpCa            []byte         `protobuf:"bytes,12,opt,name=x509_idp_ca,json=x509IdpCa,proto3" json:"x509_idp_ca,omitempty"`
	OcspIdp              []byte         `protobuf:"bytes,13,opt,name=ocsp_idp,json=ocspIdp,proto3" json:"ocsp_idp,omitempty"`
	CrlIdp               []byte         `protobuf:"bytes,14,opt,name=crl_idp,json=crlIdp,proto3" json:"crl_idp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignatureData) Reset()         { *m = SignatureData{} }
func (m *SignatureData) String() string { return proto.CompactTextString(m) }
func (*SignatureData) ProtoMessage()    {}
func (*SignatureData) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_2433630c08b6c1ea, []int{0}
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

func (m *SignatureData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SignatureData) GetSaltedHashes() [][]byte {
	if m != nil {
		return m.SaltedHashes
	}
	return nil
}

func (m *SignatureData) GetHashAlgorithm() HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return HashAlgorithm_SHA2_256
}

func (m *SignatureData) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *SignatureData) GetSignatureLevel() SignatureLevel {
	if m != nil {
		return m.SignatureLevel
	}
	return SignatureLevel_ADVANCED
}

func (m *SignatureData) GetX509() []byte {
	if m != nil {
		return m.X509
	}
	return nil
}

func (m *SignatureData) GetX509Ca() []byte {
	if m != nil {
		return m.X509Ca
	}
	return nil
}

func (m *SignatureData) GetOcsp() []byte {
	if m != nil {
		return m.Ocsp
	}
	return nil
}

func (m *SignatureData) GetCrl() []byte {
	if m != nil {
		return m.Crl
	}
	return nil
}

func (m *SignatureData) GetIdToken() []byte {
	if m != nil {
		return m.IdToken
	}
	return nil
}

func (m *SignatureData) GetX509Idp() []byte {
	if m != nil {
		return m.X509Idp
	}
	return nil
}

func (m *SignatureData) GetX509IdpCa() []byte {
	if m != nil {
		return m.X509IdpCa
	}
	return nil
}

func (m *SignatureData) GetOcspIdp() []byte {
	if m != nil {
		return m.OcspIdp
	}
	return nil
}

func (m *SignatureData) GetCrlIdp() []byte {
	if m != nil {
		return m.CrlIdp
	}
	return nil
}

type Timestamped struct {
	Rfc3161Timestamp     []byte   `protobuf:"bytes,1,opt,name=rfc3161_timestamp,json=rfc3161Timestamp,proto3" json:"rfc3161_timestamp,omitempty"`
	X509Tsa              []byte   `protobuf:"bytes,2,opt,name=x509_tsa,json=x509Tsa,proto3" json:"x509_tsa,omitempty"`
	X509TsaCa            []byte   `protobuf:"bytes,3,opt,name=x509_tsa_ca,json=x509TsaCa,proto3" json:"x509_tsa_ca,omitempty"`
	Ocsp                 []byte   `protobuf:"bytes,4,opt,name=ocsp,proto3" json:"ocsp,omitempty"`
	Crl                  []byte   `protobuf:"bytes,5,opt,name=crl,proto3" json:"crl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamped) Reset()         { *m = Timestamped{} }
func (m *Timestamped) String() string { return proto.CompactTextString(m) }
func (*Timestamped) ProtoMessage()    {}
func (*Timestamped) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_2433630c08b6c1ea, []int{1}
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

func (m *Timestamped) GetX509Tsa() []byte {
	if m != nil {
		return m.X509Tsa
	}
	return nil
}

func (m *Timestamped) GetX509TsaCa() []byte {
	if m != nil {
		return m.X509TsaCa
	}
	return nil
}

func (m *Timestamped) GetOcsp() []byte {
	if m != nil {
		return m.Ocsp
	}
	return nil
}

func (m *Timestamped) GetCrl() []byte {
	if m != nil {
		return m.Crl
	}
	return nil
}

type SignatureFile struct {
	SignatureData        *SignatureData `protobuf:"bytes,1,opt,name=signature_data,json=signatureData,proto3" json:"signature_data,omitempty"`
	Timestamps           []*Timestamped `protobuf:"bytes,2,rep,name=timestamps,proto3" json:"timestamps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignatureFile) Reset()         { *m = SignatureFile{} }
func (m *SignatureFile) String() string { return proto.CompactTextString(m) }
func (*SignatureFile) ProtoMessage()    {}
func (*SignatureFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_2433630c08b6c1ea, []int{2}
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

func (m *SignatureFile) GetSignatureData() *SignatureData {
	if m != nil {
		return m.SignatureData
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
	proto.RegisterType((*SignatureFile)(nil), "SignatureFile")
	proto.RegisterEnum("HashAlgorithm", HashAlgorithm_name, HashAlgorithm_value)
	proto.RegisterEnum("SignatureLevel", SignatureLevel_name, SignatureLevel_value)
}

func init() { proto.RegisterFile("signature.proto", fileDescriptor_signature_2433630c08b6c1ea) }

var fileDescriptor_signature_2433630c08b6c1ea = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0x34, 0x1f, 0x13, 0xdb, 0x35, 0x7b, 0xe9, 0x72, 0x41, 0x51, 0xb8, 0x44, 0x05,
	0x2c, 0xe2, 0x28, 0x15, 0x1c, 0x4d, 0xd2, 0xd2, 0x48, 0x15, 0x12, 0x6e, 0xe0, 0xc0, 0xc5, 0x5a,
	0xec, 0x6d, 0xb3, 0xc2, 0xa9, 0xad, 0xdd, 0xa5, 0xe2, 0xbf, 0xf0, 0x0f, 0xf8, 0x95, 0x68, 0xd6,
	0x89, 0x63, 0xab, 0xa7, 0xcc, 0x7b, 0x6f, 0x67, 0xe6, 0xed, 0xfa, 0x05, 0xce, 0x94, 0x78, 0x78,
	0x64, 0xfa, 0xb7, 0xe4, 0x41, 0x29, 0x0b, 0x5d, 0x4c, 0xfe, 0xd9, 0xe0, 0xde, 0x1d, 0xb8, 0x15,
	0xd3, 0x8c, 0x10, 0xe8, 0x6e, 0x99, 0xda, 0x52, 0x6b, 0x6c, 0x4d, 0x9d, 0xd8, 0xd4, 0xe4, 0x35,
	0xb8, 0x8a, 0xe5, 0x9a, 0x67, 0x09, 0x42, 0xae, 0x68, 0x67, 0x6c, 0x4f, 0x9d, 0xd8, 0xa9, 0xc8,
	0x1b, 0xc3, 0x91, 0x05, 0x78, 0xa8, 0x26, 0x2c, 0x7f, 0x28, 0xa4, 0xd0, 0xdb, 0x1d, 0xb5, 0xc7,
	0xd6, 0xd4, 0x0b, 0xbd, 0x00, 0x0f, 0x44, 0x07, 0x36, 0x76, 0xb7, 0x4d, 0x88, 0xfb, 0x70, 0x0c,
	0xed, 0x56, 0xfb, 0xb0, 0x26, 0x1f, 0x1a, 0x46, 0x93, 0x9c, 0x3f, 0xf1, 0x9c, 0x9e, 0x9a, 0x59,
	0x67, 0x41, 0x6d, 0xf6, 0x16, 0xe9, 0xd8, 0x53, 0x2d, 0x8c, 0xd3, 0xfe, 0x2c, 0xde, 0x7f, 0xa4,
	0xbd, 0x6a, 0x1a, 0xd6, 0xe4, 0x1c, 0xfa, 0xf8, 0x9b, 0xa4, 0x8c, 0xf6, 0x0d, 0xdd, 0x43, 0xb8,
	0x34, 0x57, 0x2d, 0x52, 0x55, 0xd2, 0x41, 0x75, 0x18, 0x6b, 0xe2, 0x83, 0x9d, 0xca, 0x9c, 0x0e,
	0x0d, 0x85, 0x25, 0x79, 0x09, 0x03, 0x91, 0x25, 0xba, 0xf8, 0xc5, 0x1f, 0x29, 0x18, 0xba, 0x2f,
	0xb2, 0x0d, 0x42, 0x94, 0xcc, 0x64, 0x91, 0x95, 0x74, 0x54, 0x49, 0x88, 0xd7, 0x59, 0x49, 0x5e,
	0xc1, 0xe8, 0x20, 0xe1, 0x62, 0xc7, 0xa8, 0xc3, 0xbd, 0xba, 0x64, 0xd8, 0x8a, 0xfb, 0x4c, 0xab,
	0x5b, 0xb5, 0x22, 0xc6, 0xd6, 0x73, 0xe8, 0xa7, 0x32, 0x37, 0x8a, 0x57, 0xf9, 0x4d, 0x65, 0xbe,
	0xce, 0xca, 0xc9, 0x5f, 0x0b, 0x46, 0x1b, 0xb1, 0xe3, 0x4a, 0xb3, 0x5d, 0xc9, 0x33, 0xf2, 0x06,
	0x5e, 0xc8, 0xfb, 0x74, 0x3e, 0xbb, 0x9c, 0x25, 0xfa, 0x40, 0xef, 0xbf, 0x9b, 0xbf, 0x17, 0xea,
	0xe3, 0xb5, 0x57, 0xad, 0x18, 0xed, 0x1c, 0xbd, 0x6e, 0x14, 0xab, 0xbd, 0x6a, 0xc5, 0xd0, 0xab,
	0x7d, 0xf4, 0xba, 0x51, 0xac, 0xf1, 0x4e, 0xdd, 0xe7, 0xef, 0x74, 0x5a, 0xbf, 0xd3, 0x44, 0x37,
	0x92, 0x74, 0x2d, 0x72, 0x8e, 0x81, 0x38, 0x7e, 0xc5, 0x8c, 0x69, 0x66, 0xbc, 0x8d, 0x42, 0x2f,
	0x68, 0x25, 0x2e, 0x76, 0x55, 0x2b, 0x80, 0x6f, 0x01, 0xea, 0xdb, 0x54, 0x49, 0x1b, 0x85, 0x4e,
	0xd0, 0xb8, 0x77, 0xdc, 0xd0, 0x2f, 0x3e, 0x83, 0xdb, 0x8a, 0x17, 0x71, 0x60, 0x70, 0x77, 0x13,
	0x85, 0x49, 0xb8, 0xb8, 0xf4, 0x4f, 0x6a, 0xb4, 0x98, 0x85, 0xbe, 0xb5, 0x47, 0x73, 0xa3, 0x75,
	0x6a, 0x84, 0x9a, 0x7d, 0xf1, 0x0e, 0xbc, 0x76, 0xb6, 0x50, 0x8f, 0x56, 0xdf, 0xa3, 0x2f, 0xcb,
	0xab, 0x95, 0x7f, 0x42, 0x5c, 0x18, 0x7e, 0xfd, 0x16, 0xdd, 0xae, 0xaf, 0xd7, 0x57, 0x2b, 0xdf,
	0xfa, 0x04, 0x3f, 0x06, 0x4f, 0x5c, 0x8a, 0x7b, 0xc1, 0xe5, 0xcf, 0x9e, 0xf9, 0x2f, 0xcd, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xee, 0xe9, 0xcc, 0xb8, 0x5e, 0x03, 0x00, 0x00,
}