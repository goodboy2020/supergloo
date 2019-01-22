// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/encryption.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Encryption struct {
	// If set to true, TLS is enabled across the entire mesh.
	TlsEnabled bool `protobuf:"varint,1,opt,name=tlsEnabled,proto3" json:"tlsEnabled,omitempty"`
	// This is a ref to a secret that should have at least ca-cert.pem and ca-key.pem fields.
	// The expected format is the same as defined in
	// github.com/solo-io/supergloo/pkg/api/external/istio/encryption/v1/secret.proto
	// If deploying to Consul, Consul Connect requires that the cert and key are generated using ec, not rsa.
	// If tlsEnabled is not true, this won't be used.
	Secret               *core.ResourceRef `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Encryption) Reset()         { *m = Encryption{} }
func (m *Encryption) String() string { return proto.CompactTextString(m) }
func (*Encryption) ProtoMessage()    {}
func (*Encryption) Descriptor() ([]byte, []int) {
	return fileDescriptor_encryption_0b9726061104b91a, []int{0}
}
func (m *Encryption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Encryption.Unmarshal(m, b)
}
func (m *Encryption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Encryption.Marshal(b, m, deterministic)
}
func (dst *Encryption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Encryption.Merge(dst, src)
}
func (m *Encryption) XXX_Size() int {
	return xxx_messageInfo_Encryption.Size(m)
}
func (m *Encryption) XXX_DiscardUnknown() {
	xxx_messageInfo_Encryption.DiscardUnknown(m)
}

var xxx_messageInfo_Encryption proto.InternalMessageInfo

func (m *Encryption) GetTlsEnabled() bool {
	if m != nil {
		return m.TlsEnabled
	}
	return false
}

func (m *Encryption) GetSecret() *core.ResourceRef {
	if m != nil {
		return m.Secret
	}
	return nil
}

func init() {
	proto.RegisterType((*Encryption)(nil), "supergloo.solo.io.Encryption")
}
func (this *Encryption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Encryption)
	if !ok {
		that2, ok := that.(Encryption)
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
	if this.TlsEnabled != that1.TlsEnabled {
		return false
	}
	if !this.Secret.Equal(that1.Secret) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/encryption.proto", fileDescriptor_encryption_0b9726061104b91a)
}

var fileDescriptor_encryption_0b9726061104b91a = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x2f,
	0x2e, 0x2d, 0x48, 0x2d, 0x4a, 0xcf, 0xc9, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd4,
	0x4f, 0xcd, 0x4b, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x84, 0x2b, 0xd1, 0x03, 0x69, 0xd2, 0xcb, 0xcc, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0xcb, 0xea, 0x83, 0x58, 0x10, 0x85, 0x52, 0x3a, 0xd8, 0x8c, 0x07, 0xd1, 0xd9, 0x99, 0x25,
	0x30, 0xd3, 0x8b, 0x52, 0xd3, 0x20, 0xaa, 0x95, 0xe2, 0xb9, 0xb8, 0x5c, 0xe1, 0x56, 0x09, 0xc9,
	0x71, 0x71, 0x95, 0xe4, 0x14, 0xbb, 0xe6, 0x25, 0x26, 0xe5, 0xa4, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x04, 0x21, 0x89, 0x08, 0x19, 0x72, 0xb1, 0x15, 0xa7, 0x26, 0x17, 0xa5, 0x96, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xea, 0x25, 0xe7, 0x17, 0xa5, 0xc2, 0x1c, 0xa4, 0x17,
	0x94, 0x5a, 0x9c, 0x5f, 0x5a, 0x94, 0x9c, 0x1a, 0x94, 0x9a, 0x16, 0x04, 0x55, 0xe8, 0xa4, 0xbb,
	0xe2, 0x91, 0x1c, 0x63, 0x94, 0x3a, 0x5e, 0x3f, 0x17, 0x64, 0xa7, 0x43, 0x5d, 0x96, 0xc4, 0x06,
	0x76, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xbf, 0x9e, 0x94, 0x25, 0x01, 0x00, 0x00,
}
