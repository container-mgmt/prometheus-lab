// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package prompb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import _ "google.golang.org/genproto/googleapis/api/annotations"

import time "time"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type TSDBSnapshotRequest struct {
}

func (m *TSDBSnapshotRequest) Reset()                    { *m = TSDBSnapshotRequest{} }
func (m *TSDBSnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*TSDBSnapshotRequest) ProtoMessage()               {}
func (*TSDBSnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

type TSDBSnapshotResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *TSDBSnapshotResponse) Reset()                    { *m = TSDBSnapshotResponse{} }
func (m *TSDBSnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*TSDBSnapshotResponse) ProtoMessage()               {}
func (*TSDBSnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{1} }

type SeriesDeleteRequest struct {
	MinTime  *time.Time     `protobuf:"bytes,1,opt,name=min_time,json=minTime,stdtime" json:"min_time,omitempty"`
	MaxTime  *time.Time     `protobuf:"bytes,2,opt,name=max_time,json=maxTime,stdtime" json:"max_time,omitempty"`
	Matchers []LabelMatcher `protobuf:"bytes,3,rep,name=matchers" json:"matchers"`
}

func (m *SeriesDeleteRequest) Reset()                    { *m = SeriesDeleteRequest{} }
func (m *SeriesDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*SeriesDeleteRequest) ProtoMessage()               {}
func (*SeriesDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{2} }

type SeriesDeleteResponse struct {
}

func (m *SeriesDeleteResponse) Reset()                    { *m = SeriesDeleteResponse{} }
func (m *SeriesDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*SeriesDeleteResponse) ProtoMessage()               {}
func (*SeriesDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{3} }

func init() {
	proto.RegisterType((*TSDBSnapshotRequest)(nil), "prometheus.TSDBSnapshotRequest")
	proto.RegisterType((*TSDBSnapshotResponse)(nil), "prometheus.TSDBSnapshotResponse")
	proto.RegisterType((*SeriesDeleteRequest)(nil), "prometheus.SeriesDeleteRequest")
	proto.RegisterType((*SeriesDeleteResponse)(nil), "prometheus.SeriesDeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Admin service

type AdminClient interface {
	// Snapshot creates a snapshot of all current data into 'snapshots/<datetime>-<rand>' under
	// the TSDB's data directory.
	TSDBSnapshot(ctx context.Context, in *TSDBSnapshotRequest, opts ...grpc.CallOption) (*TSDBSnapshotResponse, error)
	// DeleteSeries deletes data for a selection of series in a time range.
	DeleteSeries(ctx context.Context, in *SeriesDeleteRequest, opts ...grpc.CallOption) (*SeriesDeleteResponse, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) TSDBSnapshot(ctx context.Context, in *TSDBSnapshotRequest, opts ...grpc.CallOption) (*TSDBSnapshotResponse, error) {
	out := new(TSDBSnapshotResponse)
	err := grpc.Invoke(ctx, "/prometheus.Admin/TSDBSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteSeries(ctx context.Context, in *SeriesDeleteRequest, opts ...grpc.CallOption) (*SeriesDeleteResponse, error) {
	out := new(SeriesDeleteResponse)
	err := grpc.Invoke(ctx, "/prometheus.Admin/DeleteSeries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	// Snapshot creates a snapshot of all current data into 'snapshots/<datetime>-<rand>' under
	// the TSDB's data directory.
	TSDBSnapshot(context.Context, *TSDBSnapshotRequest) (*TSDBSnapshotResponse, error)
	// DeleteSeries deletes data for a selection of series in a time range.
	DeleteSeries(context.Context, *SeriesDeleteRequest) (*SeriesDeleteResponse, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_TSDBSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TSDBSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).TSDBSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prometheus.Admin/TSDBSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).TSDBSnapshot(ctx, req.(*TSDBSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeriesDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prometheus.Admin/DeleteSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteSeries(ctx, req.(*SeriesDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prometheus.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TSDBSnapshot",
			Handler:    _Admin_TSDBSnapshot_Handler,
		},
		{
			MethodName: "DeleteSeries",
			Handler:    _Admin_DeleteSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func (m *TSDBSnapshotRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TSDBSnapshotRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *TSDBSnapshotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TSDBSnapshotResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *SeriesDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeriesDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MinTime != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.MinTime)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.MinTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.MaxTime != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.MaxTime)))
		n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.MaxTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Matchers) > 0 {
		for _, msg := range m.Matchers {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SeriesDeleteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeriesDeleteResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TSDBSnapshotRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *TSDBSnapshotResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *SeriesDeleteRequest) Size() (n int) {
	var l int
	_ = l
	if m.MinTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.MinTime)
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.MaxTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.MaxTime)
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Matchers) > 0 {
		for _, e := range m.Matchers {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}

func (m *SeriesDeleteResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovRpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TSDBSnapshotRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TSDBSnapshotRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TSDBSnapshotRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *TSDBSnapshotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TSDBSnapshotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TSDBSnapshotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *SeriesDeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SeriesDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeriesDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinTime == nil {
				m.MinTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.MinTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxTime == nil {
				m.MaxTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.MaxTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matchers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matchers = append(m.Matchers, LabelMatcher{})
			if err := m.Matchers[len(m.Matchers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *SeriesDeleteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SeriesDeleteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeriesDeleteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
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
				next, err := skipRpc(dAtA[start:])
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
	ErrInvalidLengthRpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0xbe, 0x73, 0x2f, 0xa2, 0x0c, 0xac, 0x0a, 0x6a, 0x6d, 0x48, 0x8b, 0xb3, 0x91, 0xb0, 0xe8,
	0x24, 0x75, 0x87, 0x2b, 0x1b, 0x96, 0xba, 0x29, 0xac, 0xdc, 0x90, 0x29, 0x1c, 0x4b, 0x13, 0x3a,
	0x33, 0x76, 0x06, 0x83, 0x2e, 0x7d, 0x02, 0x13, 0x5f, 0x8a, 0xc4, 0x8d, 0x89, 0x7b, 0x7f, 0x88,
	0xef, 0xe0, 0xd6, 0x74, 0x86, 0x2a, 0x35, 0x98, 0xb8, 0x3b, 0x3d, 0xfd, 0xbe, 0xf3, 0xfd, 0x0c,
	0xee, 0x94, 0x72, 0x15, 0xca, 0x52, 0x68, 0xe1, 0x60, 0x59, 0x8a, 0x02, 0xf4, 0x06, 0x76, 0xca,
	0xeb, 0xea, 0x37, 0x12, 0x94, 0xfd, 0xe1, 0x05, 0x99, 0x10, 0xd9, 0x16, 0xa8, 0xf9, 0x4a, 0x77,
	0x2f, 0xa9, 0xce, 0x0b, 0x50, 0x9a, 0x15, 0xf2, 0x04, 0x18, 0x9e, 0x00, 0x4c, 0xe6, 0x94, 0x71,
	0x2e, 0x34, 0xd3, 0xb9, 0xe0, 0x35, 0x7d, 0x90, 0x89, 0x4c, 0x98, 0x91, 0x56, 0x93, 0xdd, 0x92,
	0xbb, 0xb8, 0xbf, 0x98, 0xcf, 0xe2, 0x39, 0x67, 0x52, 0x6d, 0x84, 0x4e, 0xe0, 0xd5, 0x0e, 0x94,
	0x26, 0x13, 0x3c, 0x68, 0xae, 0x95, 0x14, 0x5c, 0x81, 0xe3, 0xe0, 0x16, 0x67, 0x05, 0xb8, 0x68,
	0x84, 0xc6, 0x9d, 0xc4, 0xcc, 0xe4, 0x23, 0xc2, 0xfd, 0x39, 0x94, 0x39, 0xa8, 0x19, 0x6c, 0x41,
	0xc3, 0xe9, 0x86, 0xf3, 0x04, 0xdf, 0x29, 0x72, 0xbe, 0xac, 0x5c, 0x1a, 0x7c, 0x37, 0xf2, 0x42,
	0xeb, 0x30, 0xac, 0x23, 0x84, 0x8b, 0x3a, 0x42, 0xdc, 0x7a, 0xff, 0x35, 0x40, 0xc9, 0xed, 0x22,
	0xe7, 0xd5, 0xce, 0x90, 0xd9, 0xde, 0x92, 0xaf, 0xff, 0x9b, 0xcc, 0xf6, 0x86, 0x3c, 0xad, 0xc8,
	0x7a, 0xb5, 0x81, 0x52, 0xb9, 0x37, 0xa3, 0x9b, 0x71, 0x37, 0x72, 0xc3, 0x3f, 0xad, 0x86, 0xcf,
	0x58, 0x0a, 0xdb, 0xe7, 0x16, 0x10, 0xb7, 0x0e, 0x5f, 0x82, 0xab, 0xe4, 0x37, 0x9e, 0xdc, 0xc3,
	0x83, 0x66, 0x18, 0x9b, 0x3c, 0xfa, 0x89, 0xf0, 0xad, 0xa7, 0xeb, 0x22, 0xe7, 0x4e, 0x89, 0x7b,
	0xe7, 0xdd, 0x38, 0xc1, 0xf9, 0xed, 0x0b, 0x65, 0x7a, 0xa3, 0x7f, 0x03, 0xec, 0x71, 0x12, 0xbc,
	0xfb, 0xfc, 0xe3, 0xc3, 0xf5, 0x03, 0x72, 0x9f, 0xbe, 0x8e, 0x28, 0xab, 0x54, 0xa8, 0x56, 0xeb,
	0x94, 0xaa, 0x5a, 0xe3, 0x2d, 0xee, 0x59, 0x3f, 0xd6, 0x5b, 0x53, 0xf3, 0x42, 0xf9, 0x4d, 0xcd,
	0x4b, 0x81, 0xc8, 0x23, 0xa3, 0xf9, 0x90, 0x0c, 0xff, 0xd2, 0x5c, 0x1b, 0xd8, 0x52, 0x19, 0xce,
	0x14, 0x4d, 0x62, 0xf7, 0xf0, 0xdd, 0xbf, 0x3a, 0x1c, 0x7d, 0xf4, 0xe9, 0xe8, 0xa3, 0x6f, 0x47,
	0x1f, 0xbd, 0x68, 0x57, 0xb7, 0x65, 0x9a, 0xb6, 0xcd, 0x53, 0x3c, 0xfe, 0x15, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0x59, 0x61, 0x63, 0xbe, 0x02, 0x00, 0x00,
}
