// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto

	It has these top-level messages:
		TcpProxy
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_config_filter_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_api_v2_core1 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import google_protobuf2 "github.com/gogo/protobuf/types"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TcpProxy struct {
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_tcp_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The upstream cluster to connect to.
	//
	// .. note::
	//
	//  Once full filter chain matching is implemented in listeners, this field will become the only
	//  way to configure the target cluster. All other matching will be done via :ref:`filter chain
	//  matching rules <envoy_api_msg_listener.FilterChainMatch>`. For very simple configurations,
	//  this field can still be used to select the cluster when no other matching rules are required.
	//  Otherwise, a :ref:`deprecated_v1
	//  <envoy_api_field_config.filter.network.tcp_proxy.v2.TcpProxy.deprecated_v1>` configuration is
	//  required to use more complex routing in the interim.
	//
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Optional endpoint metadata match criteria. Only endpoints in the upstream
	// cluster with metadata matching that set in metadata_match will be
	// considered. The filter name should be specified as *envoy.lb*.
	MetadataMatch *envoy_api_v2_core1.Metadata `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch" json:"metadata_match,omitempty"`
	// The idle timeout for connections managed by the TCP proxy filter. The idle timeout
	// is defined as the period in which there are no bytes sent or received on either
	// the upstream or downstream connection. If not set, connections will never be closed
	// by the TCP proxy due to being idle.
	IdleTimeout *time.Duration `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,stdduration" json:"idle_timeout,omitempty"`
	// [#not-implemented-hide:] The idle timeout for connections managed by the TCP proxy
	// filter. The idle timeout is defined as the period in which there is no
	// active traffic. If not set, there is no idle timeout. When the idle timeout
	// is reached the connection will be closed. The distinction between
	// downstream_idle_timeout/upstream_idle_timeout provides a means to set
	// timeout based on the last byte sent on the downstream/upstream connection.
	DownstreamIdleTimeout *google_protobuf2.Duration `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout" json:"downstream_idle_timeout,omitempty"`
	// [#not-implemented-hide:]
	UpstreamIdleTimeout *google_protobuf2.Duration `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout" json:"upstream_idle_timeout,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>`
	// emitted by the this tcp_proxy.
	AccessLog []*envoy_config_filter_accesslog_v2.AccessLog `protobuf:"bytes,5,rep,name=access_log,json=accessLog" json:"access_log,omitempty"`
	// TCP Proxy filter configuration using deprecated V1 format. This is required for complex
	// routing until filter chain matching in the listener is implemented.
	//
	// .. attention::
	//
	//   Using this field will lead to `problems loading the configuration
	//   <https://github.com/envoyproxy/envoy/issues/2441>`_. If you want to configure the filter
	//   using v1 config structure, please make this field a boolean with value ``true`` and configure
	//   via the opaque ``value`` field like is suggested in :api:`envoy/config/filter/README.md`.
	DeprecatedV1 *TcpProxy_DeprecatedV1 `protobuf:"bytes,6,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
	// The maximum number of unsuccessful connection attempts that will be made before
	// giving up. If the parameter is not specified, 1 connection attempt will be made.
	MaxConnectAttempts *google_protobuf.UInt32Value `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts" json:"max_connect_attempts,omitempty"`
}

func (m *TcpProxy) Reset()                    { *m = TcpProxy{} }
func (m *TcpProxy) String() string            { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()               {}
func (*TcpProxy) Descriptor() ([]byte, []int) { return fileDescriptorTcpProxy, []int{0} }

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *TcpProxy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy) GetMetadataMatch() *envoy_api_v2_core1.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *time.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *google_protobuf2.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *google_protobuf2.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*envoy_config_filter_accesslog_v2.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

func (m *TcpProxy) GetDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *google_protobuf.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

// TCP Proxy filter configuration using V1 format, until Envoy gets the
// ability to match source/destination at the listener level (called
// :ref:`filter chain match <envoy_api_msg_listener.FilterChainMatch>`).
type TcpProxy_DeprecatedV1 struct {
	// The route table for the filter. All filter instances must have a route
	// table, even if it is empty.
	Routes []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *TcpProxy_DeprecatedV1) Reset()                    { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()               {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptorTcpProxy, []int{0, 0} }

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

// A TCP proxy route consists of a set of optional L4 criteria and the
// name of a cluster. If a downstream connection matches all the
// specified criteria, the cluster in the route is used for the
// corresponding upstream connection. Routes are tried in the order
// specified until a match is found. If no match is found, the connection
// is closed. A route with no criteria is valid and always produces a
// match.
type TcpProxy_DeprecatedV1_TCPRoute struct {
	// The cluster to connect to when a the downstream network connection
	// matches the specified criteria.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// An optional list of IP address subnets in the form
	// “ip_address/xx”. The criteria is satisfied if the destination IP
	// address of the downstream connection is contained in at least one of
	// the specified subnets. If the parameter is not specified or the list
	// is empty, the destination IP address is ignored. The destination IP
	// address of the downstream connection might be different from the
	// addresses on which the proxy is listening if the connection has been
	// redirected.
	DestinationIpList []*envoy_api_v2_core.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList" json:"destination_ip_list,omitempty"`
	// An optional string containing a comma-separated list of port numbers
	// or ranges. The criteria is satisfied if the destination port of the
	// downstream connection is contained in at least one of the specified
	// ranges. If the parameter is not specified, the destination port is
	// ignored. The destination port address of the downstream connection
	// might be different from the port on which the proxy is listening if
	// the connection has been redirected.
	DestinationPorts string `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	// An optional list of IP address subnets in the form
	// “ip_address/xx”. The criteria is satisfied if the source IP address
	// of the downstream connection is contained in at least one of the
	// specified subnets. If the parameter is not specified or the list is
	// empty, the source IP address is ignored.
	SourceIpList []*envoy_api_v2_core.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList" json:"source_ip_list,omitempty"`
	// An optional string containing a comma-separated list of port numbers
	// or ranges. The criteria is satisfied if the source port of the
	// downstream connection is contained in at least one of the specified
	// ranges. If the parameter is not specified, the source port is
	// ignored.
	SourcePorts string `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptorTcpProxy, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*envoy_api_v2_core.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*envoy_api_v2_core.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1.TCPRoute")
}
func (m *TcpProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpProxy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if len(m.Cluster) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(len(m.Cluster)))
		i += copy(dAtA[i:], m.Cluster)
	}
	if m.DownstreamIdleTimeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(m.DownstreamIdleTimeout.Size()))
		n1, err := m.DownstreamIdleTimeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.UpstreamIdleTimeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(m.UpstreamIdleTimeout.Size()))
		n2, err := m.UpstreamIdleTimeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.AccessLog) > 0 {
		for _, msg := range m.AccessLog {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintTcpProxy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.DeprecatedV1 != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(m.DeprecatedV1.Size()))
		n3, err := m.DeprecatedV1.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.MaxConnectAttempts != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(m.MaxConnectAttempts.Size()))
		n4, err := m.MaxConnectAttempts.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.IdleTimeout != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.IdleTimeout)))
		n5, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.IdleTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.MetadataMatch != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(m.MetadataMatch.Size()))
		n6, err := m.MetadataMatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *TcpProxy_DeprecatedV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpProxy_DeprecatedV1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Routes) > 0 {
		for _, msg := range m.Routes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTcpProxy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(len(m.Cluster)))
		i += copy(dAtA[i:], m.Cluster)
	}
	if len(m.DestinationIpList) > 0 {
		for _, msg := range m.DestinationIpList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTcpProxy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DestinationPorts) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(len(m.DestinationPorts)))
		i += copy(dAtA[i:], m.DestinationPorts)
	}
	if len(m.SourceIpList) > 0 {
		for _, msg := range m.SourceIpList {
			dAtA[i] = 0x22
			i++
			i = encodeVarintTcpProxy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.SourcePorts) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTcpProxy(dAtA, i, uint64(len(m.SourcePorts)))
		i += copy(dAtA[i:], m.SourcePorts)
	}
	return i, nil
}

func encodeVarintTcpProxy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TcpProxy) Size() (n int) {
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if m.DownstreamIdleTimeout != nil {
		l = m.DownstreamIdleTimeout.Size()
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if m.UpstreamIdleTimeout != nil {
		l = m.UpstreamIdleTimeout.Size()
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if len(m.AccessLog) > 0 {
		for _, e := range m.AccessLog {
			l = e.Size()
			n += 1 + l + sovTcpProxy(uint64(l))
		}
	}
	if m.DeprecatedV1 != nil {
		l = m.DeprecatedV1.Size()
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if m.MaxConnectAttempts != nil {
		l = m.MaxConnectAttempts.Size()
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if m.IdleTimeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.IdleTimeout)
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if m.MetadataMatch != nil {
		l = m.MetadataMatch.Size()
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	return n
}

func (m *TcpProxy_DeprecatedV1) Size() (n int) {
	var l int
	_ = l
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovTcpProxy(uint64(l))
		}
	}
	return n
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Size() (n int) {
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if len(m.DestinationIpList) > 0 {
		for _, e := range m.DestinationIpList {
			l = e.Size()
			n += 1 + l + sovTcpProxy(uint64(l))
		}
	}
	l = len(m.DestinationPorts)
	if l > 0 {
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	if len(m.SourceIpList) > 0 {
		for _, e := range m.SourceIpList {
			l = e.Size()
			n += 1 + l + sovTcpProxy(uint64(l))
		}
	}
	l = len(m.SourcePorts)
	if l > 0 {
		n += 1 + l + sovTcpProxy(uint64(l))
	}
	return n
}

func sovTcpProxy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTcpProxy(x uint64) (n int) {
	return sovTcpProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TcpProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTcpProxy
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
			return fmt.Errorf("proto: TcpProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TcpProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownstreamIdleTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DownstreamIdleTimeout == nil {
				m.DownstreamIdleTimeout = &google_protobuf2.Duration{}
			}
			if err := m.DownstreamIdleTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpstreamIdleTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpstreamIdleTimeout == nil {
				m.UpstreamIdleTimeout = &google_protobuf2.Duration{}
			}
			if err := m.UpstreamIdleTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessLog", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLog = append(m.AccessLog, &envoy_config_filter_accesslog_v2.AccessLog{})
			if err := m.AccessLog[len(m.AccessLog)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedV1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeprecatedV1 == nil {
				m.DeprecatedV1 = &TcpProxy_DeprecatedV1{}
			}
			if err := m.DeprecatedV1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnectAttempts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnectAttempts == nil {
				m.MaxConnectAttempts = &google_protobuf.UInt32Value{}
			}
			if err := m.MaxConnectAttempts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdleTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IdleTimeout == nil {
				m.IdleTimeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.IdleTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetadataMatch == nil {
				m.MetadataMatch = &envoy_api_v2_core1.Metadata{}
			}
			if err := m.MetadataMatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTcpProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTcpProxy
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
func (m *TcpProxy_DeprecatedV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTcpProxy
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
			return fmt.Errorf("proto: DeprecatedV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeprecatedV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, &TcpProxy_DeprecatedV1_TCPRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTcpProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTcpProxy
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
func (m *TcpProxy_DeprecatedV1_TCPRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTcpProxy
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
			return fmt.Errorf("proto: TCPRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TCPRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationIpList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationIpList = append(m.DestinationIpList, &envoy_api_v2_core.CidrRange{})
			if err := m.DestinationIpList[len(m.DestinationIpList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPorts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationPorts = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceIpList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceIpList = append(m.SourceIpList, &envoy_api_v2_core.CidrRange{})
			if err := m.SourceIpList[len(m.SourceIpList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePorts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpProxy
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
				return ErrInvalidLengthTcpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePorts = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTcpProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTcpProxy
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
func skipTcpProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTcpProxy
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
					return 0, ErrIntOverflowTcpProxy
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
					return 0, ErrIntOverflowTcpProxy
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
				return 0, ErrInvalidLengthTcpProxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTcpProxy
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
				next, err := skipTcpProxy(dAtA[start:])
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
	ErrInvalidLengthTcpProxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTcpProxy   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto", fileDescriptorTcpProxy)
}

var fileDescriptorTcpProxy = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0xe3, 0xf4, 0x27, 0x99, 0xa4, 0x55, 0xaf, 0xdb, 0xaa, 0xbe, 0xb9, 0x55, 0x1a, 0x60,
	0x13, 0xb5, 0xd2, 0x98, 0xa6, 0x1b, 0x76, 0xa8, 0x69, 0x17, 0x14, 0xa5, 0x52, 0xb0, 0x4a, 0x25,
	0xd8, 0x58, 0x53, 0x7b, 0xe2, 0x8e, 0xb0, 0x3d, 0xa3, 0x99, 0xb1, 0x9b, 0xbe, 0x05, 0xb0, 0xe2,
	0x11, 0x10, 0x8f, 0xc0, 0x8a, 0x25, 0x4b, 0xde, 0x00, 0xd4, 0x1d, 0x6f, 0x81, 0x3c, 0xb6, 0x13,
	0xa3, 0x06, 0xb5, 0x12, 0xbb, 0x39, 0x3f, 0xdf, 0xf7, 0xf9, 0x7c, 0xe7, 0x18, 0x3e, 0x21, 0x71,
	0xca, 0xae, 0x6d, 0x8f, 0xc5, 0x63, 0x1a, 0xd8, 0x63, 0x1a, 0x2a, 0x22, 0xec, 0x98, 0xa8, 0x2b,
	0x26, 0xde, 0xd8, 0xca, 0xe3, 0x2e, 0x17, 0x6c, 0x72, 0x6d, 0xa7, 0xfd, 0x59, 0x80, 0xb8, 0x60,
	0x8a, 0x99, 0x3d, 0x8d, 0x44, 0x39, 0x12, 0xe5, 0x48, 0x54, 0x20, 0xd1, 0xac, 0x39, 0xed, 0xb7,
	0x1f, 0xcf, 0xd3, 0xc0, 0x9e, 0x47, 0xa4, 0x0c, 0x59, 0x90, 0x71, 0x4f, 0x83, 0x9c, 0xbb, 0xbd,
	0x93, 0x23, 0x30, 0xa7, 0x59, 0xd5, 0x63, 0x82, 0xd8, 0xd8, 0xf7, 0x05, 0x91, 0xb2, 0x68, 0xd8,
	0xbe, 0xdd, 0x70, 0x81, 0x25, 0x29, 0xaa, 0x9d, 0x80, 0xb1, 0x20, 0x24, 0xb6, 0x8e, 0x2e, 0x92,
	0xb1, 0xed, 0x27, 0x02, 0x2b, 0xca, 0xe2, 0x3f, 0xd5, 0xaf, 0x04, 0xe6, 0x9c, 0x88, 0x92, 0x7d,
	0x2b, 0xc5, 0x21, 0xf5, 0xb1, 0x22, 0x76, 0xf9, 0x28, 0x0a, 0x1b, 0x01, 0x0b, 0x98, 0x7e, 0xda,
	0xd9, 0x2b, 0xcf, 0x3e, 0xfc, 0x58, 0x87, 0xf5, 0x33, 0x8f, 0x8f, 0xb2, 0x79, 0xcd, 0x5d, 0xd8,
	0x94, 0x0a, 0x2b, 0x97, 0x0b, 0x32, 0xa6, 0x13, 0x0b, 0x74, 0x41, 0xaf, 0x31, 0x68, 0x7c, 0xfe,
	0xf9, 0xa5, 0xb6, 0x20, 0x8c, 0x2e, 0x70, 0x60, 0x56, 0x1d, 0xe9, 0xa2, 0x69, 0xc1, 0x65, 0x2f,
	0x4c, 0xa4, 0x22, 0xc2, 0x32, 0xb2, 0x3e, 0xa7, 0x0c, 0xcd, 0x17, 0x70, 0xcb, 0x67, 0x57, 0xb1,
	0x54, 0x82, 0xe0, 0xc8, 0xa5, 0x7e, 0x48, 0x5c, 0x45, 0x23, 0xc2, 0x12, 0x65, 0xd5, 0xba, 0xa0,
	0xd7, 0xec, 0xff, 0x87, 0xf2, 0x19, 0x50, 0x39, 0x03, 0x3a, 0x2e, 0x66, 0x74, 0x36, 0x67, 0xc8,
	0x13, 0x3f, 0x24, 0x67, 0x39, 0xce, 0x3c, 0x85, 0x9b, 0x09, 0x9f, 0x47, 0xb8, 0x70, 0x17, 0xe1,
	0x7a, 0x89, 0xab, 0xd2, 0x3d, 0x87, 0x30, 0xdf, 0x9a, 0x1b, 0xb2, 0xc0, 0x5a, 0xec, 0xd6, 0x7a,
	0xcd, 0xfe, 0x1e, 0x9a, 0x77, 0x13, 0xb3, 0xe5, 0xa6, 0x7d, 0x74, 0xa8, 0x83, 0x21, 0x0b, 0x9c,
	0x06, 0x2e, 0x9f, 0xe6, 0x25, 0x5c, 0xf1, 0x09, 0x17, 0xc4, 0xc3, 0x8a, 0xf8, 0x6e, 0xba, 0x6f,
	0x2d, 0xe9, 0x4f, 0x7a, 0x8a, 0xee, 0x7b, 0x62, 0xa8, 0xb4, 0x1f, 0x1d, 0x4f, 0x79, 0xce, 0xf7,
	0x07, 0x86, 0x05, 0x9c, 0x96, 0x5f, 0xc9, 0x98, 0xaf, 0xe0, 0x46, 0x84, 0x27, 0xae, 0xc7, 0xe2,
	0x98, 0x78, 0xca, 0xc5, 0x4a, 0x91, 0x88, 0x2b, 0x69, 0x2d, 0x6b, 0xc1, 0xed, 0x5b, 0x1e, 0xbc,
	0x3c, 0x89, 0xd5, 0x41, 0xff, 0x1c, 0x87, 0x09, 0x29, 0x96, 0xb8, 0x6b, 0xf4, 0x80, 0x63, 0x46,
	0x78, 0x72, 0x94, 0x73, 0x1c, 0x16, 0x14, 0xe6, 0x10, 0xb6, 0x7e, 0xb3, 0xb5, 0x7e, 0x87, 0xad,
	0x83, 0xd5, 0x0f, 0xdf, 0x77, 0x40, 0xc6, 0xb9, 0xf8, 0x09, 0x18, 0xbb, 0xff, 0x38, 0x4d, 0x5a,
	0xb1, 0x77, 0x00, 0x57, 0x23, 0xa2, 0xb0, 0x8f, 0x15, 0x76, 0x23, 0xac, 0xbc, 0x4b, 0xab, 0xa1,
	0xf9, 0xfe, 0x2f, 0x3c, 0xc1, 0x9c, 0x66, 0x73, 0x67, 0x97, 0x8f, 0x4e, 0x8b, 0x46, 0x67, 0xa5,
	0x84, 0x9c, 0x66, 0x88, 0xf6, 0xdb, 0x1a, 0x6c, 0x55, 0xfd, 0x30, 0x43, 0xb8, 0x24, 0x58, 0xa2,
	0x88, 0xb4, 0x80, 0xde, 0xd7, 0xb3, 0xbf, 0x34, 0x18, 0x9d, 0x1d, 0x8d, 0x9c, 0x8c, 0x70, 0x00,
	0xf5, 0x1c, 0xef, 0x81, 0x51, 0x07, 0x4e, 0xa1, 0xd1, 0x7e, 0x67, 0xc0, 0x7a, 0xd9, 0x60, 0x3e,
	0x9a, 0x9d, 0xfa, 0xad, 0x5f, 0x62, 0x7a, 0xf5, 0x43, 0xb8, 0xee, 0x13, 0xa9, 0x68, 0xac, 0x0d,
	0x72, 0x29, 0x77, 0x43, 0x2a, 0x95, 0x65, 0xe8, 0x8f, 0xdd, 0x9e, 0x33, 0xf9, 0x11, 0xf5, 0x85,
	0x83, 0xe3, 0x80, 0x38, 0xff, 0x56, 0x80, 0x27, 0x7c, 0x48, 0xa5, 0x32, 0xf7, 0x60, 0x35, 0xe9,
	0x72, 0x26, 0x94, 0xd4, 0x7f, 0x4f, 0xc3, 0x59, 0xab, 0x14, 0x46, 0x59, 0x3e, 0xf3, 0x5b, 0xb2,
	0x44, 0x78, 0x64, 0xaa, 0xba, 0x70, 0x0f, 0xd5, 0x56, 0x8e, 0x29, 0x04, 0x1f, 0xc0, 0x22, 0x2e,
	0xb4, 0x16, 0xb5, 0x56, 0x33, 0xcf, 0x69, 0x99, 0xc1, 0xda, 0xd7, 0x9b, 0x0e, 0xf8, 0x76, 0xd3,
	0x01, 0x3f, 0x6e, 0x3a, 0xe0, 0xb5, 0x91, 0xf6, 0x2f, 0x96, 0xf4, 0x61, 0x1c, 0xfc, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x36, 0xb9, 0xaf, 0x89, 0x05, 0x00, 0x00,
}