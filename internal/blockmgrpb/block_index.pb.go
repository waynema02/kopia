// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/blockmgrpb/block_index.proto

/*
	Package blockmgrpb is a generated protocol buffer package.

	It is generated from these files:
		internal/blockmgrpb/block_index.proto

	It has these top-level messages:
		Index
		Indexes
*/
package blockmgrpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Index struct {
	PackBlockId     string            `protobuf:"bytes,1,opt,name=pack_block_id,json=packBlockId,proto3" json:"pack_block_id,omitempty"`
	PackLength      uint64            `protobuf:"varint,2,opt,name=pack_length,json=packLength,proto3" json:"pack_length,omitempty"`
	CreateTimeNanos uint64            `protobuf:"varint,3,opt,name=create_time_nanos,json=createTimeNanos,proto3" json:"create_time_nanos,omitempty"`
	Items           map[string]uint64 `protobuf:"bytes,4,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DeletedItems    []string          `protobuf:"bytes,5,rep,name=deleted_items,json=deletedItems" json:"deleted_items,omitempty"`
}

func (m *Index) Reset()                    { *m = Index{} }
func (m *Index) String() string            { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()               {}
func (*Index) Descriptor() ([]byte, []int) { return fileDescriptorBlockIndex, []int{0} }

func (m *Index) GetPackBlockId() string {
	if m != nil {
		return m.PackBlockId
	}
	return ""
}

func (m *Index) GetPackLength() uint64 {
	if m != nil {
		return m.PackLength
	}
	return 0
}

func (m *Index) GetCreateTimeNanos() uint64 {
	if m != nil {
		return m.CreateTimeNanos
	}
	return 0
}

func (m *Index) GetItems() map[string]uint64 {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Index) GetDeletedItems() []string {
	if m != nil {
		return m.DeletedItems
	}
	return nil
}

type Indexes struct {
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
}

func (m *Indexes) Reset()                    { *m = Indexes{} }
func (m *Indexes) String() string            { return proto.CompactTextString(m) }
func (*Indexes) ProtoMessage()               {}
func (*Indexes) Descriptor() ([]byte, []int) { return fileDescriptorBlockIndex, []int{1} }

func (m *Indexes) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func init() {
	proto.RegisterType((*Index)(nil), "Index")
	proto.RegisterType((*Indexes)(nil), "Indexes")
}
func (m *Index) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Index) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PackBlockId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlockIndex(dAtA, i, uint64(len(m.PackBlockId)))
		i += copy(dAtA[i:], m.PackBlockId)
	}
	if m.PackLength != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBlockIndex(dAtA, i, uint64(m.PackLength))
	}
	if m.CreateTimeNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBlockIndex(dAtA, i, uint64(m.CreateTimeNanos))
	}
	if len(m.Items) > 0 {
		for k, _ := range m.Items {
			dAtA[i] = 0x22
			i++
			v := m.Items[k]
			mapSize := 1 + len(k) + sovBlockIndex(uint64(len(k))) + 1 + sovBlockIndex(uint64(v))
			i = encodeVarintBlockIndex(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(v))
		}
	}
	if len(m.DeletedItems) > 0 {
		for _, s := range m.DeletedItems {
			dAtA[i] = 0x2a
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
	return i, nil
}

func (m *Indexes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Indexes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		for _, msg := range m.Indexes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintBlockIndex(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Index) Size() (n int) {
	var l int
	_ = l
	l = len(m.PackBlockId)
	if l > 0 {
		n += 1 + l + sovBlockIndex(uint64(l))
	}
	if m.PackLength != 0 {
		n += 1 + sovBlockIndex(uint64(m.PackLength))
	}
	if m.CreateTimeNanos != 0 {
		n += 1 + sovBlockIndex(uint64(m.CreateTimeNanos))
	}
	if len(m.Items) > 0 {
		for k, v := range m.Items {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovBlockIndex(uint64(len(k))) + 1 + sovBlockIndex(uint64(v))
			n += mapEntrySize + 1 + sovBlockIndex(uint64(mapEntrySize))
		}
	}
	if len(m.DeletedItems) > 0 {
		for _, s := range m.DeletedItems {
			l = len(s)
			n += 1 + l + sovBlockIndex(uint64(l))
		}
	}
	return n
}

func (m *Indexes) Size() (n int) {
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		for _, e := range m.Indexes {
			l = e.Size()
			n += 1 + l + sovBlockIndex(uint64(l))
		}
	}
	return n
}

func sovBlockIndex(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlockIndex(x uint64) (n int) {
	return sovBlockIndex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Index) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockIndex
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Index: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Index: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackBlockId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackBlockId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackLength", wireType)
			}
			m.PackLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PackLength |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTimeNanos", wireType)
			}
			m.CreateTimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateTimeNanos |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Items == nil {
				m.Items = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBlockIndex
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockIndex
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBlockIndex
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockIndex
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBlockIndex(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBlockIndex
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Items[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedItems", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedItems = append(m.DeletedItems, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockIndex
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
func (m *Indexes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockIndex
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Indexes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Indexes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexes = append(m.Indexes, &Index{})
			if err := m.Indexes[len(m.Indexes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockIndex
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
func skipBlockIndex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockIndex
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
					return 0, ErrIntOverflowBlockIndex
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
					return 0, ErrIntOverflowBlockIndex
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBlockIndex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlockIndex
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
				next, err := skipBlockIndex(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthBlockIndex = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockIndex   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("internal/blockmgrpb/block_index.proto", fileDescriptorBlockIndex) }

var fileDescriptorBlockIndex = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x9d, 0xa6, 0x69, 0xe9, 0xad, 0x45, 0x3b, 0xb8, 0x08, 0x22, 0x31, 0x44, 0xc4, 0xa0,
	0x10, 0x41, 0x37, 0xc5, 0x65, 0xc1, 0x45, 0x40, 0x5c, 0x04, 0x57, 0x6e, 0x86, 0xfc, 0x5c, 0xea,
	0x90, 0x64, 0x12, 0x92, 0x51, 0xec, 0xce, 0xc7, 0xf0, 0x91, 0x5c, 0xfa, 0x08, 0x12, 0x5f, 0x44,
	0x66, 0x26, 0xd2, 0xdd, 0x39, 0xdf, 0xbd, 0x87, 0x39, 0x73, 0xe1, 0x9c, 0x0b, 0x89, 0xad, 0x48,
	0xca, 0xeb, 0xb4, 0xac, 0xb3, 0xa2, 0xda, 0xb4, 0x4d, 0x6a, 0x24, 0xe3, 0x22, 0xc7, 0xf7, 0xb0,
	0x69, 0x6b, 0x59, 0xfb, 0x1f, 0x23, 0xb0, 0x23, 0xe5, 0xa9, 0x0f, 0x8b, 0x26, 0xc9, 0x0a, 0x36,
	0xec, 0xe4, 0x0e, 0xf1, 0x48, 0x30, 0x8b, 0xe7, 0x0a, 0xae, 0x15, 0x8b, 0x72, 0x7a, 0x0a, 0xda,
	0xb2, 0x12, 0xc5, 0x46, 0xbe, 0x38, 0x23, 0x8f, 0x04, 0xe3, 0x18, 0x14, 0x7a, 0xd0, 0x84, 0x5e,
	0xc2, 0x32, 0x6b, 0x31, 0x91, 0xc8, 0x24, 0xaf, 0x90, 0x89, 0x44, 0xd4, 0x9d, 0x63, 0xe9, 0xb5,
	0x03, 0x33, 0x78, 0xe2, 0x15, 0x3e, 0x2a, 0x4c, 0x2f, 0xc0, 0xe6, 0x12, 0xab, 0xce, 0x19, 0x7b,
	0x56, 0x30, 0xbf, 0x59, 0x86, 0xba, 0x47, 0x18, 0x29, 0x76, 0x2f, 0x64, 0xbb, 0x8d, 0xcd, 0x9c,
	0x9e, 0xc1, 0x22, 0xc7, 0x12, 0x25, 0xe6, 0xcc, 0x04, 0x6c, 0xcf, 0x0a, 0x66, 0xf1, 0xfe, 0x00,
	0x75, 0xe0, 0x78, 0x05, 0xb0, 0x4b, 0xd2, 0x43, 0xb0, 0x0a, 0xdc, 0x0e, 0x5f, 0x50, 0x92, 0x1e,
	0x81, 0xfd, 0x96, 0x94, 0xaf, 0x38, 0x94, 0x36, 0xe6, 0x6e, 0xb4, 0x22, 0xfe, 0x15, 0x4c, 0xf5,
	0xcb, 0xd8, 0x51, 0x0f, 0xa6, 0xdc, 0x48, 0x87, 0xe8, 0x52, 0x13, 0x53, 0x2a, 0xfe, 0xc7, 0xeb,
	0x93, 0xaf, 0xde, 0x25, 0xdf, 0xbd, 0x4b, 0x7e, 0x7a, 0x97, 0x7c, 0xfe, 0xba, 0x7b, 0xcf, 0xb0,
	0xbb, 0x6f, 0x3a, 0xd1, 0x47, 0xbd, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x26, 0xd1, 0xd9, 0xf1,
	0x7d, 0x01, 0x00, 0x00,
}