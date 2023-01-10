// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dataocean/dataocean/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the dataocean module's genesis state.
type GenesisState struct {
	Params        Params      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	VideoList     []Video     `protobuf:"bytes,2,rep,name=videoList,proto3" json:"videoList"`
	VideoCount    uint64      `protobuf:"varint,3,opt,name=videoCount,proto3" json:"videoCount,omitempty"`
	VideoLinkList []VideoLink `protobuf:"bytes,4,rep,name=videoLinkList,proto3" json:"videoLinkList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_853f1a06e0a3a47a, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetVideoList() []Video {
	if m != nil {
		return m.VideoList
	}
	return nil
}

func (m *GenesisState) GetVideoCount() uint64 {
	if m != nil {
		return m.VideoCount
	}
	return 0
}

func (m *GenesisState) GetVideoLinkList() []VideoLink {
	if m != nil {
		return m.VideoLinkList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dataocean.dataocean.GenesisState")
}

func init() { proto.RegisterFile("dataocean/dataocean/genesis.proto", fileDescriptor_853f1a06e0a3a47a) }

var fileDescriptor_853f1a06e0a3a47a = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x49, 0x2c, 0x49,
	0xcc, 0x4f, 0x4e, 0x4d, 0xcc, 0xd3, 0x47, 0xb0, 0xd2, 0x53, 0xf3, 0x52, 0x8b, 0x33, 0x8b, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0x12, 0x7a, 0x70, 0x96, 0x94, 0x48, 0x7a, 0x7e,
	0x7a, 0x3e, 0x58, 0x5e, 0x1f, 0xc4, 0x82, 0x28, 0x95, 0x52, 0xc0, 0x66, 0x5a, 0x41, 0x62, 0x51,
	0x62, 0x2e, 0xd4, 0x30, 0x29, 0x79, 0x6c, 0x2a, 0xca, 0x32, 0x53, 0x52, 0x61, 0x46, 0xa8, 0xe0,
	0x54, 0x10, 0x9f, 0x93, 0x99, 0x97, 0x0d, 0x51, 0xa5, 0xf4, 0x95, 0x91, 0x8b, 0xc7, 0x1d, 0xe2,
	0xca, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x4b, 0x2e, 0x36, 0x88, 0x3d, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0xd2, 0x7a, 0x58, 0x5c, 0xad, 0x17, 0x00, 0x56, 0xe2, 0xc4, 0x72, 0xe2, 0x9e,
	0x3c, 0x43, 0x10, 0x54, 0x83, 0x90, 0x1d, 0x17, 0x27, 0xd8, 0x7c, 0x9f, 0xcc, 0xe2, 0x12, 0x09,
	0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x29, 0xac, 0xba, 0xc3, 0x40, 0xaa, 0xa0, 0x9a, 0x11, 0x5a,
	0x84, 0xe4, 0xb8, 0xb8, 0xc0, 0x1c, 0xe7, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0x96, 0x20, 0x24, 0x11, 0x21, 0x2f, 0x2e, 0x5e, 0xa8, 0xe2, 0xbc, 0x6c, 0xb0, 0x1d, 0x2c, 0x60,
	0x3b, 0xe4, 0x70, 0xdb, 0x01, 0x52, 0x09, 0xb5, 0x07, 0x55, 0xab, 0x93, 0xe9, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x49, 0x23, 0x42, 0xab, 0x02, 0x29, 0xe4, 0x4a, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xa1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x28,
	0xba, 0xf4, 0xee, 0x01, 0x00, 0x00,
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
	if len(m.VideoLinkList) > 0 {
		for iNdEx := len(m.VideoLinkList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VideoLinkList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.VideoCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VideoCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.VideoList) > 0 {
		for iNdEx := len(m.VideoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VideoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.VideoList) > 0 {
		for _, e := range m.VideoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.VideoCount != 0 {
		n += 1 + sovGenesis(uint64(m.VideoCount))
	}
	if len(m.VideoLinkList) > 0 {
		for _, e := range m.VideoLinkList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoList", wireType)
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
			m.VideoList = append(m.VideoList, Video{})
			if err := m.VideoList[len(m.VideoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoCount", wireType)
			}
			m.VideoCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VideoCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoLinkList", wireType)
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
			m.VideoLinkList = append(m.VideoLinkList, VideoLink{})
			if err := m.VideoLinkList[len(m.VideoLinkList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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