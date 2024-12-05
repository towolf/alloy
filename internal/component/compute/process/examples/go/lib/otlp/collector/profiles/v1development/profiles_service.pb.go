// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.8.0
// source: opentelemetry/proto/collector/profiles/v1development/profiles_service.proto

package opentelemetry_proto_collector_profiles_v1development

import (
	fmt "fmt"
	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	v1development "github.com/grafana/alloy/internal/component/compute/process/examples/go/lib/otlp/profiles/v1development"
	io "io"
)

// Copyright 2023, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

type ExportProfilesServiceRequest struct {
	unknownFields []byte
	// An array of ResourceProfiles.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceProfiles []*v1development.ResourceProfiles `protobuf:"bytes,1,rep,name=resource_profiles,json=resourceProfiles,proto3" json:"resourceProfiles,omitempty"`
}

func (x *ExportProfilesServiceRequest) Reset() {
	*x = ExportProfilesServiceRequest{}
}

func (*ExportProfilesServiceRequest) ProtoMessage() {}

func (x *ExportProfilesServiceRequest) GetResourceProfiles() []*v1development.ResourceProfiles {
	if x != nil {
		return x.ResourceProfiles
	}
	return nil
}

type ExportProfilesServiceResponse struct {
	unknownFields []byte
	// The details of a partially successful export request.
	//
	// If the request is only partially accepted
	// (i.e. when the server accepts only parts of the data and rejects the rest)
	// the server MUST initialize the `partial_success` field and MUST
	// set the `rejected_<signal>` with the number of items it rejected.
	//
	// Servers MAY also make use of the `partial_success` field to convey
	// warnings/suggestions to senders even when the request was fully accepted.
	// In such cases, the `rejected_<signal>` MUST have a value of `0` and
	// the `error_message` MUST be non-empty.
	//
	// A `partial_success` message with an empty value (rejected_<signal> = 0 and
	// `error_message` = "") is equivalent to it not being set/present. Senders
	// SHOULD interpret it the same way as in the full success case.
	PartialSuccess *ExportProfilesPartialSuccess `protobuf:"bytes,1,opt,name=partial_success,json=partialSuccess,proto3" json:"partialSuccess,omitempty"`
}

func (x *ExportProfilesServiceResponse) Reset() {
	*x = ExportProfilesServiceResponse{}
}

func (*ExportProfilesServiceResponse) ProtoMessage() {}

func (x *ExportProfilesServiceResponse) GetPartialSuccess() *ExportProfilesPartialSuccess {
	if x != nil {
		return x.PartialSuccess
	}
	return nil
}

type ExportProfilesPartialSuccess struct {
	unknownFields []byte
	// The number of rejected profiles.
	//
	// A `rejected_<signal>` field holding a `0` value indicates that the
	// request was fully accepted.
	RejectedProfiles int64 `protobuf:"varint,1,opt,name=rejected_profiles,json=rejectedProfiles,proto3" json:"rejectedProfiles,omitempty"`
	// A developer-facing human-readable message in English. It should be used
	// either to explain why the server rejected parts of the data during a partial
	// success or to convey warnings/suggestions during a full success. The message
	// should offer guidance on how users can address such issues.
	//
	// error_message is an optional field. An error_message with an empty value
	// is equivalent to it not being set.
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *ExportProfilesPartialSuccess) Reset() {
	*x = ExportProfilesPartialSuccess{}
}

func (*ExportProfilesPartialSuccess) ProtoMessage() {}

func (x *ExportProfilesPartialSuccess) GetRejectedProfiles() int64 {
	if x != nil {
		return x.RejectedProfiles
	}
	return 0
}

func (x *ExportProfilesPartialSuccess) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (m *ExportProfilesServiceRequest) CloneVT() *ExportProfilesServiceRequest {
	if m == nil {
		return (*ExportProfilesServiceRequest)(nil)
	}
	r := new(ExportProfilesServiceRequest)
	if rhs := m.ResourceProfiles; rhs != nil {
		tmpContainer := make([]*v1development.ResourceProfiles, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ResourceProfiles = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ExportProfilesServiceRequest) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *ExportProfilesServiceResponse) CloneVT() *ExportProfilesServiceResponse {
	if m == nil {
		return (*ExportProfilesServiceResponse)(nil)
	}
	r := new(ExportProfilesServiceResponse)
	r.PartialSuccess = m.PartialSuccess.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ExportProfilesServiceResponse) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *ExportProfilesPartialSuccess) CloneVT() *ExportProfilesPartialSuccess {
	if m == nil {
		return (*ExportProfilesPartialSuccess)(nil)
	}
	r := new(ExportProfilesPartialSuccess)
	r.RejectedProfiles = m.RejectedProfiles
	r.ErrorMessage = m.ErrorMessage
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ExportProfilesPartialSuccess) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *ExportProfilesServiceRequest) EqualVT(that *ExportProfilesServiceRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.ResourceProfiles) != len(that.ResourceProfiles) {
		return false
	}
	for i, vx := range this.ResourceProfiles {
		vy := that.ResourceProfiles[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &v1development.ResourceProfiles{}
			}
			if q == nil {
				q = &v1development.ResourceProfiles{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ExportProfilesServiceRequest) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*ExportProfilesServiceRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ExportProfilesServiceResponse) EqualVT(that *ExportProfilesServiceResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.PartialSuccess.EqualVT(that.PartialSuccess) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ExportProfilesServiceResponse) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*ExportProfilesServiceResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ExportProfilesPartialSuccess) EqualVT(that *ExportProfilesPartialSuccess) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RejectedProfiles != that.RejectedProfiles {
		return false
	}
	if this.ErrorMessage != that.ErrorMessage {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ExportProfilesPartialSuccess) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*ExportProfilesPartialSuccess)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *ExportProfilesServiceRequest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportProfilesServiceRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ExportProfilesServiceRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.ResourceProfiles) > 0 {
		for iNdEx := len(m.ResourceProfiles) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ResourceProfiles[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExportProfilesServiceResponse) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportProfilesServiceResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ExportProfilesServiceResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.PartialSuccess != nil {
		size, err := m.PartialSuccess.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExportProfilesPartialSuccess) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportProfilesPartialSuccess) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ExportProfilesPartialSuccess) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x12
	}
	if m.RejectedProfiles != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.RejectedProfiles))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExportProfilesServiceRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ResourceProfiles) > 0 {
		for _, e := range m.ResourceProfiles {
			l = e.SizeVT()
			n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ExportProfilesServiceResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PartialSuccess != nil {
		l = m.PartialSuccess.SizeVT()
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ExportProfilesPartialSuccess) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RejectedProfiles != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.RejectedProfiles))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ExportProfilesServiceRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: ExportProfilesServiceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportProfilesServiceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceProfiles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceProfiles = append(m.ResourceProfiles, &v1development.ResourceProfiles{})
			if err := m.ResourceProfiles[len(m.ResourceProfiles)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportProfilesServiceResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: ExportProfilesServiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportProfilesServiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartialSuccess", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PartialSuccess == nil {
				m.PartialSuccess = &ExportProfilesPartialSuccess{}
			}
			if err := m.PartialSuccess.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportProfilesPartialSuccess) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: ExportProfilesPartialSuccess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportProfilesPartialSuccess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectedProfiles", wireType)
			}
			m.RejectedProfiles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RejectedProfiles |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}