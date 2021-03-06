// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission.proto

package permission

import (
	fmt "fmt"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Permission int32

const (
	Permission_IgnoreLogin           Permission = 2
	Permission_LoginWithNoPermission Permission = 1
	Permission_NeedPerm              Permission = 0
)

var Permission_name = map[int32]string{
	2: "IgnoreLogin",
	1: "LoginWithNoPermission",
	0: "NeedPerm",
}

var Permission_value = map[string]int32{
	"IgnoreLogin":           2,
	"LoginWithNoPermission": 1,
	"NeedPerm":              0,
}

func (x Permission) Enum() *Permission {
	p := new(Permission)
	*p = x
	return p
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (x *Permission) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Permission_value, data, "Permission")
	if err != nil {
		return err
	}
	*x = Permission(value)
	return nil
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

type DataRule struct {
	Code                 *string  `protobuf:"bytes,1,req,name=code" json:"code,omitempty"`
	Summary              *string  `protobuf:"bytes,2,req,name=summary" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRule) Reset()         { *m = DataRule{} }
func (m *DataRule) String() string { return proto.CompactTextString(m) }
func (*DataRule) ProtoMessage()    {}
func (*DataRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}
func (m *DataRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataRule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRule.Merge(m, src)
}
func (m *DataRule) XXX_Size() int {
	return m.Size()
}
func (m *DataRule) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRule.DiscardUnknown(m)
}

var xxx_messageInfo_DataRule proto.InternalMessageInfo

func (m *DataRule) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *DataRule) GetSummary() string {
	if m != nil && m.Summary != nil {
		return *m.Summary
	}
	return ""
}

type HttpRule struct {
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Post
	Pattern isHttpRule_Pattern `protobuf_oneof:"pattern"`
	//权限
	Perm *Permission `protobuf:"varint,3,req,name=perm,enum=bidewu.Permission,def=0" json:"perm,omitempty"`
	//  //权限码
	//  optional string permcode=4;
	//  //权限分类
	//  optional string permgroup=6;
	//所属系统
	App                  *string  `protobuf:"bytes,5,opt,name=app" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{1}
}
func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(m, src)
}
func (m *HttpRule) XXX_Size() int {
	return m.Size()
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

const Default_HttpRule_Perm Permission = Permission_NeedPerm

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,1,opt,name=get,oneof" json:"get,omitempty"`
}
type HttpRule_Post struct {
	Post string `protobuf:"bytes,2,opt,name=post,oneof" json:"post,omitempty"`
}

func (*HttpRule_Get) isHttpRule_Pattern()  {}
func (*HttpRule_Post) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetPerm() Permission {
	if m != nil && m.Perm != nil {
		return *m.Perm
	}
	return Default_HttpRule_Perm
}

func (m *HttpRule) GetApp() string {
	if m != nil && m.App != nil {
		return *m.App
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Post)(nil),
	}
}

var E_Datarules = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
	ExtensionType: ([]*DataRule)(nil),
	Field:         8000,
	Name:          "bidewu.datarules",
	Tag:           "bytes,8000,rep,name=datarules",
	Filename:      "permission.proto",
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295628,
	Name:          "bidewu.http",
	Tag:           "bytes,72295628,opt,name=http",
	Filename:      "permission.proto",
}

func init() {
	proto.RegisterEnum("bidewu.Permission", Permission_name, Permission_value)
	proto.RegisterType((*DataRule)(nil), "bidewu.DataRule")
	proto.RegisterType((*HttpRule)(nil), "bidewu.HttpRule")
	proto.RegisterExtension(E_Datarules)
	proto.RegisterExtension(E_Http)
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x9d, 0xce, 0x8c, 0xee, 0x4c, 0xad, 0xb8, 0xa1, 0x51, 0x88, 0x1e, 0x62, 0x98, 0xd3, 0xe0,
	0x21, 0x91, 0x3d, 0xc9, 0x08, 0xc2, 0x86, 0x55, 0x67, 0x41, 0xd7, 0x21, 0xba, 0x2c, 0x78, 0xeb,
	0x49, 0xca, 0x4e, 0xeb, 0x24, 0xdd, 0x74, 0x2a, 0x2e, 0x7a, 0x15, 0xff, 0xc1, 0xb3, 0xdf, 0xe1,
	0xc1, 0xa3, 0x82, 0x07, 0x3f, 0x41, 0xc6, 0x1f, 0x91, 0x64, 0x36, 0x13, 0xc1, 0xdb, 0x7b, 0xc5,
	0xeb, 0x57, 0xaf, 0x1e, 0x0d, 0xae, 0x41, 0x5b, 0xa8, 0xaa, 0x52, 0xba, 0x0c, 0x8d, 0xd5, 0xa4,
	0xf9, 0xd5, 0x95, 0xca, 0xf0, 0xa2, 0xbe, 0x1d, 0x48, 0xad, 0xe5, 0x1a, 0xa3, 0x76, 0xba, 0xaa,
	0x5f, 0x47, 0x19, 0x56, 0xa9, 0x55, 0x86, 0xb4, 0xdd, 0x2a, 0xa7, 0xf7, 0x61, 0x7c, 0x2c, 0x48,
	0x24, 0xf5, 0x1a, 0x39, 0x87, 0x51, 0xaa, 0x33, 0xf4, 0x58, 0xe0, 0xcc, 0x26, 0x49, 0x8b, 0xb9,
	0x07, 0x7b, 0x55, 0x5d, 0x14, 0xc2, 0xbe, 0xf7, 0x9c, 0x76, 0xdc, 0xd1, 0xe9, 0x47, 0x06, 0xe3,
	0x05, 0x91, 0xb9, 0x7c, 0x3a, 0x94, 0x48, 0x1e, 0x0b, 0xd8, 0x6c, 0xb2, 0x18, 0x24, 0x0d, 0xe1,
	0x37, 0x60, 0x64, 0x74, 0x45, 0x9e, 0x73, 0x39, 0x6c, 0x19, 0xbf, 0x07, 0xa3, 0x26, 0xae, 0x37,
	0x0c, 0x9c, 0xd9, 0xf5, 0x43, 0x1e, 0x6e, 0x93, 0x86, 0xcb, 0xdd, 0x09, 0xf3, 0xf1, 0x29, 0x62,
	0xd6, 0xf0, 0xa4, 0x55, 0x72, 0x17, 0x86, 0xc2, 0x18, 0xef, 0x4a, 0x63, 0x93, 0x34, 0x30, 0x9e,
	0xc0, 0x9e, 0x11, 0x44, 0x68, 0xcb, 0xbb, 0x8f, 0x01, 0xfa, 0xa7, 0xfc, 0x00, 0xf6, 0x4f, 0x64,
	0xa9, 0x2d, 0x3e, 0xd5, 0x52, 0x95, 0xae, 0xc3, 0x6f, 0xc1, 0xcd, 0x16, 0x9e, 0x2b, 0xca, 0x4f,
	0x75, 0xaf, 0x74, 0x19, 0xbf, 0x06, 0xbb, 0x45, 0xee, 0x60, 0xbe, 0x84, 0x49, 0x26, 0x48, 0xd8,
	0x7a, 0x8d, 0x15, 0xbf, 0x13, 0x6e, 0x7b, 0x0b, 0xbb, 0xde, 0xc2, 0x17, 0x68, 0xdf, 0xa9, 0x14,
	0x9f, 0x1b, 0x52, 0xba, 0xac, 0xbc, 0x6f, 0x0f, 0x83, 0xe1, 0x6c, 0xff, 0xd0, 0xed, 0xd2, 0x77,
	0x15, 0x26, 0xbd, 0xc9, 0xfc, 0x09, 0x8c, 0x72, 0x22, 0xc3, 0xfd, 0xff, 0xcc, 0x9e, 0x21, 0xe5,
	0x3a, 0xeb, 0xbc, 0x7e, 0xfe, 0xf8, 0x3a, 0x0d, 0xd8, 0xbf, 0x6e, 0x5d, 0xab, 0x49, 0x6b, 0x10,
	0x7f, 0x62, 0xdf, 0x37, 0x3e, 0xfb, 0xb5, 0xf1, 0xd9, 0xef, 0x8d, 0xcf, 0x3e, 0xff, 0xf1, 0x07,
	0x70, 0x90, 0xea, 0xa2, 0x53, 0x8b, 0x9a, 0xf2, 0xd8, 0x8d, 0x5b, 0xd2, 0x1f, 0xb8, 0x64, 0xaf,
	0x8e, 0xa4, 0xa2, 0xbc, 0x5e, 0x85, 0xa9, 0x2e, 0xa2, 0x0f, 0xb9, 0x28, 0xe5, 0x1b, 0x55, 0xca,
	0x35, 0xaa, 0xe8, 0x42, 0xe4, 0x22, 0x17, 0x91, 0x79, 0x2b, 0xa3, 0x12, 0x29, 0x6a, 0x96, 0x44,
	0xfd, 0x17, 0x7a, 0xd0, 0xc3, 0x2f, 0x0e, 0xc4, 0x27, 0xc7, 0x8f, 0xce, 0xcf, 0x8e, 0xce, 0x5e,
	0x2e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x46, 0x95, 0xac, 0xdd, 0x67, 0x02, 0x00, 0x00,
}

func (m *DataRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataRule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataRule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Summary == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.Summary)
		copy(dAtA[i:], *m.Summary)
		i = encodeVarintPermission(dAtA, i, uint64(len(*m.Summary)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.Code)
		copy(dAtA[i:], *m.Code)
		i = encodeVarintPermission(dAtA, i, uint64(len(*m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HttpRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpRule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpRule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.App != nil {
		i -= len(*m.App)
		copy(dAtA[i:], *m.App)
		i = encodeVarintPermission(dAtA, i, uint64(len(*m.App)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Perm == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintPermission(dAtA, i, uint64(*m.Perm))
		i--
		dAtA[i] = 0x18
	}
	if m.Pattern != nil {
		{
			size := m.Pattern.Size()
			i -= size
			if _, err := m.Pattern.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *HttpRule_Get) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpRule_Get) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Get)
	copy(dAtA[i:], m.Get)
	i = encodeVarintPermission(dAtA, i, uint64(len(m.Get)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *HttpRule_Post) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpRule_Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Post)
	copy(dAtA[i:], m.Post)
	i = encodeVarintPermission(dAtA, i, uint64(len(m.Post)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintPermission(dAtA []byte, offset int, v uint64) int {
	offset -= sovPermission(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != nil {
		l = len(*m.Code)
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.Summary != nil {
		l = len(*m.Summary)
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pattern != nil {
		n += m.Pattern.Size()
	}
	if m.Perm != nil {
		n += 1 + sovPermission(uint64(*m.Perm))
	}
	if m.App != nil {
		l = len(*m.App)
		n += 1 + l + sovPermission(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpRule_Get) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Get)
	n += 1 + l + sovPermission(uint64(l))
	return n
}
func (m *HttpRule_Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Post)
	n += 1 + l + sovPermission(uint64(l))
	return n
}

func sovPermission(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPermission(x uint64) (n int) {
	return sovPermission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataRule) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
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
			return fmt.Errorf("proto: DataRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Code = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Summary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Summary = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HttpRule) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
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
			return fmt.Errorf("proto: HttpRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Get", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pattern = &HttpRule_Get{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pattern = &HttpRule_Post{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			var v Permission
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= Permission(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Perm = &v
			hasFields[0] |= uint64(0x00000001)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
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
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.App = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermission
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPermission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
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
				return 0, ErrInvalidLengthPermission
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPermission
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPermission
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPermission        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermission          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPermission = fmt.Errorf("proto: unexpected end of group")
)
