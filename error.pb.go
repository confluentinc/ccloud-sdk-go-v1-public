// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/error.proto

package ccloud

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	// http status code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" db:"code,omitempty" url:"code,omitempty"`
	// a short, human-readable summary of the problem. It should not change from occurrence to occurrence of the problem
	Title string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty" db:"title,omitempty" url:"title,omitempty"`
	// human-readable explanation specific to this occurrence of the problem
	Message      string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" db:"message,omitempty" url:"message,omitempty"`
	NestedErrors map[string]string `protobuf:"bytes,3,rep,name=nested_errors,json=nestedErrors,proto3" json:"nested_errors,omitempty" db:"nested_errors,omitempty" url:"nested_errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Details      []string          `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty" db:"details,omitempty" url:"details,omitempty"`
	Stack        *Stack            `protobuf:"bytes,5,opt,name=stack,proto3,customtype=Stack" json:"stack,omitempty" db:"stack,omitempty" url:"stack,omitempty"`
	// an application-specific error code, expressed as a string value.
	ErrorCode string `protobuf:"bytes,6,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty" db:"error_code,omitempty" url:"error_code,omitempty"`
	// optional pointer to input request document
	Source *ErrorSource `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty" db:"source,omitempty" url:"source,omitempty"`
	// list of errors that are needed for public apis
	MultiErrors          []*Error `protobuf:"bytes,8,rep,name=multi_errors,json=multiErrors,proto3" json:"multi_errors,omitempty" db:"multi_errors,omitempty" url:"multi_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_04e40541982b606d, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetNestedErrors() map[string]string {
	if m != nil {
		return m.NestedErrors
	}
	return nil
}

func (m *Error) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Error) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Error) GetSource() *ErrorSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Error) GetMultiErrors() []*Error {
	if m != nil {
		return m.MultiErrors
	}
	return nil
}

type ErrorSource struct {
	// a JSON Pointer [RFC6901] to the associated entity in the request document [e.g. \"/data\" for a primary data object, or \"/data/attributes/title\" for a specific attribute].
	Pointer string `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty" db:"pointer,omitempty" url:"pointer,omitempty"`
	// a string indicating which query parameter caused the error.
	Parameter string `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty" db:"parameter,omitempty" url:"parameter,omitempty"`
	// a url indicating the location of the conflicting resource
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty" db:"location,omitempty" url:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorSource) Reset()         { *m = ErrorSource{} }
func (m *ErrorSource) String() string { return proto.CompactTextString(m) }
func (*ErrorSource) ProtoMessage()    {}
func (*ErrorSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_04e40541982b606d, []int{1}
}
func (m *ErrorSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorSource.Merge(m, src)
}
func (m *ErrorSource) XXX_Size() int {
	return m.Size()
}
func (m *ErrorSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorSource.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorSource proto.InternalMessageInfo

func (m *ErrorSource) GetPointer() string {
	if m != nil {
		return m.Pointer
	}
	return ""
}

func (m *ErrorSource) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

func (m *ErrorSource) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "ccloud.core.Error")
	proto.RegisterMapType((map[string]string)(nil), "ccloud.core.Error.NestedErrorsEntry")
	proto.RegisterType((*ErrorSource)(nil), "ccloud.core.ErrorSource")
}

func init() { proto.RegisterFile("core/error.proto", fileDescriptor_04e40541982b606d) }

var fileDescriptor_04e40541982b606d = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe5, 0x75, 0xd9, 0x56, 0x77, 0x48, 0x60, 0x71, 0x88, 0x76, 0x58, 0xa2, 0x0a, 0xa1,
	0x30, 0xad, 0xe9, 0x56, 0xa6, 0x31, 0x95, 0x09, 0xa1, 0xc0, 0x6e, 0x80, 0x26, 0x0f, 0x09, 0x89,
	0xcb, 0x94, 0x3f, 0x6e, 0x89, 0xf2, 0xc7, 0x25, 0x71, 0x2a, 0xf5, 0xcc, 0x89, 0xcf, 0xc0, 0x17,
	0xe2, 0xc8, 0x89, 0x03, 0x87, 0x08, 0xf5, 0x23, 0xe4, 0x13, 0x20, 0xbf, 0x69, 0xba, 0xae, 0x2e,
	0xb7, 0xf4, 0x79, 0xfc, 0xfe, 0xfc, 0xfa, 0xf5, 0xe3, 0xe2, 0x87, 0x3e, 0xcf, 0x58, 0x9f, 0x65,
	0x19, 0xcf, 0xec, 0x49, 0xc6, 0x05, 0x27, 0x1d, 0xdf, 0x8f, 0x79, 0x11, 0xd8, 0xd2, 0x38, 0x78,
	0x3c, 0xe6, 0x63, 0x0e, 0x7a, 0x5f, 0x7e, 0xd5, 0x4b, 0xba, 0xbf, 0x77, 0xb1, 0x76, 0x25, 0x4b,
	0xc8, 0x6b, 0xbc, 0xed, 0xf3, 0x80, 0xe9, 0xc8, 0x44, 0x96, 0xe6, 0x1c, 0x57, 0xa5, 0x61, 0x05,
	0xde, 0xb0, 0x2b, 0xb5, 0x63, 0x9e, 0x84, 0x82, 0x25, 0x13, 0x31, 0xeb, 0x9a, 0x45, 0x16, 0x2b,
	0x22, 0x85, 0x4a, 0xf2, 0x16, 0x6b, 0x22, 0x14, 0x31, 0xd3, 0xdb, 0x26, 0xb2, 0xda, 0x8e, 0x5d,
	0x95, 0xc6, 0x91, 0x44, 0x80, 0xa8, 0x30, 0xd6, 0x55, 0x5a, 0x17, 0x93, 0x77, 0x78, 0x37, 0x61,
	0x79, 0xee, 0x8e, 0x99, 0xbe, 0x05, 0x9c, 0x41, 0x55, 0x1a, 0xb6, 0xe4, 0x2c, 0x64, 0x85, 0xa4,
	0xea, 0xb4, 0x41, 0x90, 0xef, 0x08, 0x3f, 0x48, 0x59, 0x2e, 0x58, 0x70, 0x0b, 0x93, 0xc9, 0xf5,
	0x96, 0xd9, 0xb2, 0x3a, 0x83, 0x27, 0xf6, 0xca, 0x6c, 0x6c, 0x98, 0x80, 0xfd, 0x01, 0xd6, 0xc1,
	0x77, 0x7e, 0x95, 0x8a, 0x6c, 0xe6, 0xbc, 0xaa, 0x4a, 0x63, 0x28, 0xb7, 0xbe, 0x87, 0x50, 0x1a,
	0xf8, 0x9f, 0x4b, 0xf7, 0xd3, 0x15, 0xa4, 0x3c, 0x59, 0xc0, 0x84, 0x1b, 0xc6, 0xb9, 0xbe, 0x6d,
	0xb6, 0x56, 0x4f, 0xb6, 0x90, 0x15, 0xb0, 0xaa, 0xd3, 0x06, 0x41, 0xde, 0x63, 0x2d, 0x17, 0xae,
	0x1f, 0xe9, 0x9a, 0x89, 0xac, 0x7d, 0xe7, 0xc5, 0x9f, 0xd2, 0xd0, 0x6e, 0xa4, 0xd0, 0x8c, 0x1d,
	0x5c, 0x05, 0xb9, 0xae, 0xd2, 0x9a, 0x42, 0x3e, 0x61, 0x0c, 0xfd, 0xdf, 0x42, 0x08, 0x76, 0x60,
	0xf2, 0x17, 0x55, 0x69, 0x9c, 0x49, 0xd4, 0x9d, 0xa3, 0xf0, 0x36, 0x5a, 0xb4, 0x0d, 0xf2, 0x1b,
	0x99, 0x0a, 0x0f, 0xef, 0xe4, 0xbc, 0xc8, 0x7c, 0xa6, 0xef, 0x9a, 0xc8, 0xea, 0x0c, 0x74, 0x75,
	0xf2, 0x37, 0xe0, 0x3b, 0x27, 0x55, 0x69, 0x1c, 0x43, 0xe7, 0xf0, 0x5b, 0x6d, 0x7d, 0x5d, 0xa6,
	0x0b, 0x32, 0xc9, 0xf1, 0x7e, 0x52, 0xc4, 0x22, 0x6c, 0xee, 0x78, 0x0f, 0xee, 0x98, 0xa8, 0x3b,
	0x39, 0x97, 0x55, 0x69, 0x5c, 0x40, 0x98, 0x56, 0xd6, 0xab, 0x89, 0xda, 0x6c, 0xd2, 0x0e, 0x18,
	0xf5, 0x75, 0x1e, 0xfc, 0x40, 0xf8, 0x91, 0x12, 0x19, 0x72, 0x89, 0x5b, 0x11, 0x9b, 0xc1, 0x2b,
	0x6a, 0x3b, 0x47, 0x55, 0x69, 0x3c, 0x95, 0xbb, 0x45, 0x6c, 0xa6, 0x6c, 0x72, 0x5f, 0xa3, 0xb2,
	0x4c, 0x3e, 0xa1, 0xa9, 0x1b, 0x17, 0x4d, 0xf4, 0x97, 0x4f, 0x08, 0x44, 0x85, 0xb0, 0xae, 0xd2,
	0xba, 0x78, 0xb8, 0x75, 0x81, 0xba, 0xdf, 0xb6, 0x70, 0x67, 0x65, 0xb8, 0x32, 0x7c, 0x13, 0x1e,
	0xa6, 0x82, 0x65, 0x8b, 0xde, 0x96, 0xe1, 0x5b, 0xc8, 0x0a, 0x5d, 0xd5, 0x69, 0x83, 0x20, 0x1f,
	0x71, 0x7b, 0xe2, 0x66, 0x6e, 0xc2, 0x24, 0xaf, 0xee, 0xf5, 0xbc, 0x2a, 0x8d, 0x01, 0xf0, 0x1a,
	0x43, 0x25, 0x6e, 0x70, 0xe8, 0x1d, 0x88, 0x5c, 0xe3, 0xbd, 0x98, 0xfb, 0xae, 0x08, 0x79, 0xaa,
	0xb7, 0x00, 0x7a, 0x56, 0x95, 0xc6, 0x89, 0x84, 0x36, 0xba, 0xc2, 0xdc, 0x60, 0xd0, 0x25, 0xc5,
	0xf9, 0xfa, 0x73, 0x7e, 0x88, 0x7e, 0xcd, 0x0f, 0xd1, 0xdf, 0xf9, 0x21, 0xc2, 0xcf, 0x42, 0x6e,
	0xfb, 0x3c, 0x1d, 0xc5, 0x05, 0x4b, 0x85, 0x5d, 0x67, 0x03, 0xfe, 0x06, 0xbd, 0x62, 0x64, 0x47,
	0xee, 0x28, 0x72, 0xeb, 0xa8, 0x4c, 0x4f, 0xaf, 0xd1, 0xe7, 0xf3, 0x71, 0x28, 0xbe, 0x14, 0x9e,
	0xed, 0xf3, 0xa4, 0xbf, 0x2c, 0x0a, 0x53, 0xbf, 0x5f, 0x87, 0xaa, 0x97, 0x07, 0x51, 0x6f, 0xcc,
	0x7b, 0xd3, 0xd3, 0xde, 0xa4, 0xf0, 0xe2, 0xd0, 0x7f, 0x59, 0xeb, 0xde, 0x0e, 0x10, 0x9f, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x65, 0x87, 0x80, 0x72, 0x8f, 0x05, 0x00, 0x00,
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintError(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.NestedErrors) > 0 {
		for k, _ := range m.NestedErrors {
			dAtA[i] = 0x1a
			i++
			v := m.NestedErrors[k]
			mapSize := 1 + len(k) + sovError(uint64(len(k))) + 1 + len(v) + sovError(uint64(len(v)))
			i = encodeVarintError(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintError(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintError(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Details) > 0 {
		for _, s := range m.Details {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Stack != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintError(dAtA, i, uint64(m.Stack.Size()))
		n1, err := m.Stack.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.ErrorCode) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.ErrorCode)))
		i += copy(dAtA[i:], m.ErrorCode)
	}
	if m.Source != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintError(dAtA, i, uint64(m.Source.Size()))
		n2, err := m.Source.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.MultiErrors) > 0 {
		for _, msg := range m.MultiErrors {
			dAtA[i] = 0x42
			i++
			i = encodeVarintError(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ErrorSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Pointer) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Pointer)))
		i += copy(dAtA[i:], m.Pointer)
	}
	if len(m.Parameter) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Parameter)))
		i += copy(dAtA[i:], m.Parameter)
	}
	if len(m.Location) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Location)))
		i += copy(dAtA[i:], m.Location)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovError(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if len(m.NestedErrors) > 0 {
		for k, v := range m.NestedErrors {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovError(uint64(len(k))) + 1 + len(v) + sovError(uint64(len(v)))
			n += mapEntrySize + 1 + sovError(uint64(mapEntrySize))
		}
	}
	if len(m.Details) > 0 {
		for _, s := range m.Details {
			l = len(s)
			n += 1 + l + sovError(uint64(l))
		}
	}
	if m.Stack != nil {
		l = m.Stack.Size()
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.ErrorCode)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.Source != nil {
		l = m.Source.Size()
		n += 1 + l + sovError(uint64(l))
	}
	if len(m.MultiErrors) > 0 {
		for _, e := range m.MultiErrors {
			l = e.Size()
			n += 1 + l + sovError(uint64(l))
		}
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ErrorSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pointer)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovError(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NestedErrors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NestedErrors == nil {
				m.NestedErrors = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowError
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowError
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthError
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthError
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowError
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthError
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthError
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipError(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthError
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.NestedErrors[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = append(m.Details, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stack", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v Stack
			m.Stack = &v
			if err := m.Stack.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Source == nil {
				m.Source = &ErrorSource{}
			}
			if err := m.Source.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiErrors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultiErrors = append(m.MultiErrors, &Error{})
			if err := m.MultiErrors[len(m.MultiErrors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: ErrorSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pointer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pointer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowError
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
				return 0, ErrInvalidLengthError
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthError
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowError
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipError(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthError
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthError = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError   = fmt.Errorf("proto: integer overflow")
)
