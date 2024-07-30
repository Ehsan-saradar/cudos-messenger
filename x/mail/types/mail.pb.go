// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mail/mail.proto

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

type Mail struct {
	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From            string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To              string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Subject         string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body            string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAtHeight int64  `protobuf:"varint,6,opt,name=createdAtHeight,proto3" json:"createdAtHeight,omitempty"`
}

func (m *Mail) Reset()         { *m = Mail{} }
func (m *Mail) String() string { return proto.CompactTextString(m) }
func (*Mail) ProtoMessage()    {}
func (*Mail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ac4131798133d9, []int{0}
}
func (m *Mail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mail.Merge(m, src)
}
func (m *Mail) XXX_Size() int {
	return m.Size()
}
func (m *Mail) XXX_DiscardUnknown() {
	xxx_messageInfo_Mail.DiscardUnknown(m)
}

var xxx_messageInfo_Mail proto.InternalMessageInfo

func (m *Mail) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mail) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Mail) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Mail) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Mail) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Mail) GetCreatedAtHeight() int64 {
	if m != nil {
		return m.CreatedAtHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Mail)(nil), "CudoVentures.cudosnode.mail.Mail")
}

func init() { proto.RegisterFile("mail/mail.proto", fileDescriptor_d0ac4131798133d9) }

var fileDescriptor_d0ac4131798133d9 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0xcc, 0xcc,
	0xd1, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xd2, 0xce, 0xa5, 0x29, 0xf9, 0x61,
	0xa9, 0x79, 0x25, 0xa5, 0x45, 0xa9, 0xc5, 0x7a, 0xc9, 0xa5, 0x29, 0xf9, 0xc5, 0x79, 0xf9, 0x29,
	0xa9, 0x7a, 0x20, 0x25, 0x4a, 0x93, 0x18, 0xb9, 0x58, 0x7c, 0x13, 0x33, 0x73, 0x84, 0xf8, 0xb8,
	0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8,
	0x58, 0xd2, 0x8a, 0xf2, 0x73, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x9a,
	0x92, 0x7c, 0x09, 0x66, 0xb0, 0x08, 0x53, 0x49, 0xbe, 0x90, 0x04, 0x17, 0x7b, 0x71, 0x69, 0x52,
	0x56, 0x6a, 0x72, 0x89, 0x04, 0x0b, 0x58, 0x10, 0xc6, 0x05, 0xe9, 0x4e, 0xca, 0x4f, 0xa9, 0x94,
	0x60, 0x85, 0xe8, 0x06, 0xb1, 0x85, 0x34, 0xb8, 0xf8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53,
	0x1c, 0x4b, 0x3c, 0x52, 0x33, 0xd3, 0x33, 0x4a, 0x24, 0xd8, 0x14, 0x18, 0x35, 0x98, 0x83, 0xd0,
	0x85, 0x9d, 0x3c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3f, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xd9, 0x5b, 0xfa, 0x60, 0x6f, 0xe9,
	0x82, 0xfc, 0xa5, 0x5f, 0x01, 0xf6, 0xbc, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38,
	0x0c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0xa5, 0xe0, 0x33, 0x16, 0x01, 0x00, 0x00,
}

func (m *Mail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAtHeight != 0 {
		i = encodeVarintMail(dAtA, i, uint64(m.CreatedAtHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintMail(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintMail(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintMail(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMail(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMail(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMail(dAtA []byte, offset int, v uint64) int {
	offset -= sovMail(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMail(uint64(m.Id))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMail(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovMail(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovMail(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovMail(uint64(l))
	}
	if m.CreatedAtHeight != 0 {
		n += 1 + sovMail(uint64(m.CreatedAtHeight))
	}
	return n
}

func sovMail(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMail(x uint64) (n int) {
	return sovMail(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMail
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
			return fmt.Errorf("proto: Mail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMail
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
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMail
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
				return ErrInvalidLengthMail
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMail
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
				return ErrInvalidLengthMail
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMail
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
				return ErrInvalidLengthMail
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMail
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
				return ErrInvalidLengthMail
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAtHeight", wireType)
			}
			m.CreatedAtHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMail
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAtHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMail(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMail
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMail
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
func skipMail(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMail
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
					return 0, ErrIntOverflowMail
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
					return 0, ErrIntOverflowMail
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
				return 0, ErrInvalidLengthMail
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMail
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMail
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMail        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMail          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMail = fmt.Errorf("proto: unexpected end of group")
)
