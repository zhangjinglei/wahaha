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
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
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
	proto.RegisterType((*HttpRule)(nil), "bidewu.HttpRule")
	proto.RegisterExtension(E_Http)
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xcd, 0x6a, 0xe3, 0x30,
	0x18, 0x8c, 0x9c, 0xec, 0x6e, 0xa2, 0x2c, 0x1b, 0x23, 0x76, 0xc1, 0xbb, 0x07, 0x63, 0x72, 0x0a,
	0x7b, 0xb0, 0x96, 0x1c, 0xb3, 0xa7, 0x98, 0xfe, 0x24, 0xd0, 0xa6, 0xc1, 0x34, 0x04, 0x7a, 0x73,
	0x62, 0x55, 0x52, 0x9b, 0x48, 0xc2, 0xfe, 0x4c, 0xa0, 0xd7, 0xd2, 0x77, 0xe8, 0xb9, 0xcf, 0xd1,
	0x07, 0x68, 0xa1, 0x87, 0x3e, 0x42, 0x49, 0x5f, 0xa4, 0xd8, 0xa9, 0xe3, 0xdb, 0xcc, 0xc7, 0xe8,
	0x9b, 0x19, 0x7d, 0xd8, 0x36, 0x2c, 0x59, 0xcb, 0x34, 0x95, 0x5a, 0xf9, 0x26, 0xd1, 0xa0, 0xc9,
	0xd7, 0x85, 0x8c, 0xd9, 0x26, 0xfb, 0xe3, 0x71, 0xad, 0xf9, 0x8a, 0xd1, 0x62, 0xba, 0xc8, 0x2e,
	0x69, 0xcc, 0xd2, 0x65, 0x22, 0x0d, 0xe8, 0x64, 0xa7, 0xec, 0xde, 0x22, 0xdc, 0x1c, 0x01, 0x98,
	0x30, 0x5b, 0x31, 0x42, 0x70, 0x9d, 0x33, 0x70, 0x90, 0x87, 0x7a, 0xad, 0x51, 0x2d, 0xcc, 0x09,
	0xf9, 0x89, 0x1b, 0x46, 0xa7, 0xe0, 0x58, 0x9f, 0xc3, 0x82, 0x91, 0x7f, 0xb8, 0x91, 0x9b, 0x3a,
	0x75, 0xcf, 0xea, 0xfd, 0xe8, 0x13, 0x7f, 0xe7, 0xe7, 0x4f, 0xf7, 0x41, 0x06, 0xcd, 0x09, 0x63,
	0x71, 0xce, 0xc3, 0x42, 0x49, 0x6c, 0x5c, 0x8f, 0x8c, 0x71, 0xbe, 0xe4, 0x6b, 0xc2, 0x1c, 0x06,
	0x2d, 0xfc, 0xcd, 0x44, 0x00, 0x2c, 0x51, 0x7f, 0x8f, 0x30, 0xae, 0x9e, 0x92, 0x0e, 0x6e, 0x8f,
	0xb9, 0xd2, 0x09, 0x3b, 0xd1, 0x5c, 0x2a, 0xdb, 0x22, 0xbf, 0xf1, 0xaf, 0x02, 0xce, 0x25, 0x88,
	0x89, 0xae, 0x94, 0x36, 0x22, 0xdf, 0xf1, 0xde, 0xc8, 0xae, 0x0d, 0x8e, 0x71, 0x43, 0x00, 0x18,
	0xe2, 0xfa, 0xbb, 0xe2, 0x7e, 0x59, 0xdc, 0x3f, 0x65, 0x20, 0x74, 0x7c, 0x66, 0x40, 0x6a, 0x95,
	0x3a, 0x2f, 0xcf, 0x8f, 0x5d, 0x0f, 0xf5, 0xda, 0x7d, 0xbb, 0x4c, 0x5e, 0xfe, 0x41, 0x58, 0x2c,
	0x08, 0xee, 0xd0, 0xd3, 0xd6, 0x45, 0xaf, 0x5b, 0x17, 0xbd, 0x6d, 0x5d, 0x74, 0xff, 0xee, 0xd6,
	0x70, 0x67, 0xa9, 0xd7, 0xa5, 0x3a, 0xca, 0x40, 0x04, 0x76, 0x50, 0x90, 0x2a, 0xce, 0x14, 0x5d,
	0x0c, 0xb9, 0x04, 0x91, 0x2d, 0xfc, 0xa5, 0x5e, 0xd3, 0x1b, 0x11, 0x29, 0x7e, 0x25, 0x15, 0x5f,
	0x31, 0x49, 0x37, 0x91, 0x88, 0x44, 0x44, 0xcd, 0x35, 0xa7, 0x8a, 0x01, 0xcd, 0x4d, 0x68, 0x75,
	0xb6, 0xff, 0x15, 0x7c, 0xb0, 0x70, 0x30, 0x3e, 0x38, 0x9c, 0xcf, 0x86, 0xb3, 0xf3, 0xd1, 0x47,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x9c, 0x56, 0xd0, 0xdb, 0x01, 0x00, 0x00,
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