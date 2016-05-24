// Code generated by protoc-gen-go.
// source: pdpb.proto
// DO NOT EDIT!

/*
Package pdpb is a generated protocol buffer package.

It is generated from these files:
	pdpb.proto

It has these top-level messages:
	Leader
	TsoRequest
	Timestamp
	TsoResponse
	BootstrapRequest
	BootstrapResponse
	IsBootstrappedRequest
	IsBootstrappedResponse
	AllocIdRequest
	AllocIdResponse
	GetStoreRequest
	GetStoreResponse
	GetRegionRequest
	GetRegionResponse
	GetClusterConfigRequest
	GetClusterConfigResponse
	PutStoreRequest
	PutStoreResponse
	RegionHeartbeatRequest
	ChangePeer
	RegionHeartbeatResponse
	PutClusterConfigRequest
	PutClusterConfigResponse
	AskSplitRequest
	AskSplitResponse
	RequestHeader
	ResponseHeader
	Request
	Response
	BootstrappedError
	Error
*/
package pdpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import metapb "github.com/pingcap/kvproto/pkg/metapb"
import raftpb "github.com/pingcap/kvproto/pkg/raftpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CommandType int32

const (
	CommandType_Invalid          CommandType = 0
	CommandType_Tso              CommandType = 1
	CommandType_Bootstrap        CommandType = 2
	CommandType_IsBootstrapped   CommandType = 3
	CommandType_AllocId          CommandType = 4
	CommandType_GetStore         CommandType = 5
	CommandType_PutStore         CommandType = 6
	CommandType_AskSplit         CommandType = 7
	CommandType_GetRegion        CommandType = 8
	CommandType_RegionHeartbeat  CommandType = 9
	CommandType_GetClusterConfig CommandType = 10
	CommandType_PutClusterConfig CommandType = 11
)

var CommandType_name = map[int32]string{
	0:  "Invalid",
	1:  "Tso",
	2:  "Bootstrap",
	3:  "IsBootstrapped",
	4:  "AllocId",
	5:  "GetStore",
	6:  "PutStore",
	7:  "AskSplit",
	8:  "GetRegion",
	9:  "RegionHeartbeat",
	10: "GetClusterConfig",
	11: "PutClusterConfig",
}
var CommandType_value = map[string]int32{
	"Invalid":          0,
	"Tso":              1,
	"Bootstrap":        2,
	"IsBootstrapped":   3,
	"AllocId":          4,
	"GetStore":         5,
	"PutStore":         6,
	"AskSplit":         7,
	"GetRegion":        8,
	"RegionHeartbeat":  9,
	"GetClusterConfig": 10,
	"PutClusterConfig": 11,
}

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}
func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (x *CommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandType_value, data, "CommandType")
	if err != nil {
		return err
	}
	*x = CommandType(value)
	return nil
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Leader struct {
	Addr             *string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Pid              *int64  `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Leader) Reset()                    { *m = Leader{} }
func (m *Leader) String() string            { return proto.CompactTextString(m) }
func (*Leader) ProtoMessage()               {}
func (*Leader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Leader) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *Leader) GetPid() int64 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

type TsoRequest struct {
	Count            *uint32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsoRequest) Reset()                    { *m = TsoRequest{} }
func (m *TsoRequest) String() string            { return proto.CompactTextString(m) }
func (*TsoRequest) ProtoMessage()               {}
func (*TsoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TsoRequest) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type Timestamp struct {
	Physical         *int64 `protobuf:"varint,1,opt,name=physical" json:"physical,omitempty"`
	Logical          *int64 `protobuf:"varint,2,opt,name=logical" json:"logical,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Timestamp) GetPhysical() int64 {
	if m != nil && m.Physical != nil {
		return *m.Physical
	}
	return 0
}

func (m *Timestamp) GetLogical() int64 {
	if m != nil && m.Logical != nil {
		return *m.Logical
	}
	return 0
}

type TsoResponse struct {
	Timestamps       []*Timestamp `protobuf:"bytes,1,rep,name=timestamps" json:"timestamps,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TsoResponse) Reset()                    { *m = TsoResponse{} }
func (m *TsoResponse) String() string            { return proto.CompactTextString(m) }
func (*TsoResponse) ProtoMessage()               {}
func (*TsoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TsoResponse) GetTimestamps() []*Timestamp {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type BootstrapRequest struct {
	Store            *metapb.Store  `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	Region           *metapb.Region `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *BootstrapRequest) Reset()                    { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string            { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()               {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BootstrapRequest) GetStore() *metapb.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *BootstrapRequest) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type BootstrapResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BootstrapResponse) Reset()                    { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string            { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()               {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type IsBootstrappedRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsBootstrappedRequest) Reset()                    { *m = IsBootstrappedRequest{} }
func (m *IsBootstrappedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsBootstrappedRequest) ProtoMessage()               {}
func (*IsBootstrappedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type IsBootstrappedResponse struct {
	Bootstrapped     *bool  `protobuf:"varint,1,opt,name=bootstrapped" json:"bootstrapped,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsBootstrappedResponse) Reset()                    { *m = IsBootstrappedResponse{} }
func (m *IsBootstrappedResponse) String() string            { return proto.CompactTextString(m) }
func (*IsBootstrappedResponse) ProtoMessage()               {}
func (*IsBootstrappedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsBootstrappedResponse) GetBootstrapped() bool {
	if m != nil && m.Bootstrapped != nil {
		return *m.Bootstrapped
	}
	return false
}

type AllocIdRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AllocIdRequest) Reset()                    { *m = AllocIdRequest{} }
func (m *AllocIdRequest) String() string            { return proto.CompactTextString(m) }
func (*AllocIdRequest) ProtoMessage()               {}
func (*AllocIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AllocIdResponse struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AllocIdResponse) Reset()                    { *m = AllocIdResponse{} }
func (m *AllocIdResponse) String() string            { return proto.CompactTextString(m) }
func (*AllocIdResponse) ProtoMessage()               {}
func (*AllocIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AllocIdResponse) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type GetStoreRequest struct {
	StoreId          *uint64 `protobuf:"varint,1,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetStoreRequest) Reset()                    { *m = GetStoreRequest{} }
func (m *GetStoreRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStoreRequest) ProtoMessage()               {}
func (*GetStoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetStoreRequest) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

type GetStoreResponse struct {
	Store            *metapb.Store `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GetStoreResponse) Reset()                    { *m = GetStoreResponse{} }
func (m *GetStoreResponse) String() string            { return proto.CompactTextString(m) }
func (*GetStoreResponse) ProtoMessage()               {}
func (*GetStoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetStoreResponse) GetStore() *metapb.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type GetRegionRequest struct {
	RegionKey        []byte `protobuf:"bytes,1,opt,name=region_key" json:"region_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetRegionRequest) Reset()                    { *m = GetRegionRequest{} }
func (m *GetRegionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRegionRequest) ProtoMessage()               {}
func (*GetRegionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetRegionRequest) GetRegionKey() []byte {
	if m != nil {
		return m.RegionKey
	}
	return nil
}

type GetRegionResponse struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GetRegionResponse) Reset()                    { *m = GetRegionResponse{} }
func (m *GetRegionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRegionResponse) ProtoMessage()               {}
func (*GetRegionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetRegionResponse) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type GetClusterConfigRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetClusterConfigRequest) Reset()                    { *m = GetClusterConfigRequest{} }
func (m *GetClusterConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetClusterConfigRequest) ProtoMessage()               {}
func (*GetClusterConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type GetClusterConfigResponse struct {
	Cluster          *metapb.Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetClusterConfigResponse) Reset()                    { *m = GetClusterConfigResponse{} }
func (m *GetClusterConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetClusterConfigResponse) ProtoMessage()               {}
func (*GetClusterConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetClusterConfigResponse) GetCluster() *metapb.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type PutStoreRequest struct {
	Store            *metapb.Store `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PutStoreRequest) Reset()                    { *m = PutStoreRequest{} }
func (m *PutStoreRequest) String() string            { return proto.CompactTextString(m) }
func (*PutStoreRequest) ProtoMessage()               {}
func (*PutStoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *PutStoreRequest) GetStore() *metapb.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type PutStoreResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PutStoreResponse) Reset()                    { *m = PutStoreResponse{} }
func (m *PutStoreResponse) String() string            { return proto.CompactTextString(m) }
func (*PutStoreResponse) ProtoMessage()               {}
func (*PutStoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type RegionHeartbeatRequest struct {
	Region *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	// Leader Peer sending the heartbeat.
	Leader           *metapb.Peer `protobuf:"bytes,2,opt,name=leader" json:"leader,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RegionHeartbeatRequest) Reset()                    { *m = RegionHeartbeatRequest{} }
func (m *RegionHeartbeatRequest) String() string            { return proto.CompactTextString(m) }
func (*RegionHeartbeatRequest) ProtoMessage()               {}
func (*RegionHeartbeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *RegionHeartbeatRequest) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *RegionHeartbeatRequest) GetLeader() *metapb.Peer {
	if m != nil {
		return m.Leader
	}
	return nil
}

type ChangePeer struct {
	ChangeType       *raftpb.ConfChangeType `protobuf:"varint,1,opt,name=change_type,enum=raftpb.ConfChangeType" json:"change_type,omitempty"`
	Peer             *metapb.Peer           `protobuf:"bytes,2,opt,name=peer" json:"peer,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ChangePeer) Reset()                    { *m = ChangePeer{} }
func (m *ChangePeer) String() string            { return proto.CompactTextString(m) }
func (*ChangePeer) ProtoMessage()               {}
func (*ChangePeer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *ChangePeer) GetChangeType() raftpb.ConfChangeType {
	if m != nil && m.ChangeType != nil {
		return *m.ChangeType
	}
	return raftpb.ConfChangeType_AddNode
}

func (m *ChangePeer) GetPeer() *metapb.Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type RegionHeartbeatResponse struct {
	// Notice, Pd only allows handling reported epoch >= current pd's.
	// Leader peer reports region status with RegionHeartbeatRequest
	// to pd regularly, pd will determine whether this region
	// should do ChangePeer or not.
	// E,g, max peer number is 3, region A, first only peer 1 in A.
	// 1. Pd region state -> Peers (1), ConfVer (1).
	// 2. Leader peer 1 reports region state to pd, pd finds the
	// peer number is < 3, so first changes its current region
	// state -> Peers (1, 2), ConfVer (1), and returns ChangePeer Adding 2.
	// 3. Leader does ChangePeer, then reports Peers (1, 2), ConfVer (2),
	// pd updates its state -> Peers (1, 2), ConfVer (2).
	// 4. Leader may report old Peers (1), ConfVer (1) to pd before ConfChange
	// finished, pd stills responses ChangePeer Adding 2, of course, we must
	// guarantee the second ChangePeer can't be applied in TiKV.
	ChangePeer       *ChangePeer `protobuf:"bytes,1,opt,name=change_peer" json:"change_peer,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RegionHeartbeatResponse) Reset()                    { *m = RegionHeartbeatResponse{} }
func (m *RegionHeartbeatResponse) String() string            { return proto.CompactTextString(m) }
func (*RegionHeartbeatResponse) ProtoMessage()               {}
func (*RegionHeartbeatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *RegionHeartbeatResponse) GetChangePeer() *ChangePeer {
	if m != nil {
		return m.ChangePeer
	}
	return nil
}

type PutClusterConfigRequest struct {
	Cluster          *metapb.Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PutClusterConfigRequest) Reset()                    { *m = PutClusterConfigRequest{} }
func (m *PutClusterConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*PutClusterConfigRequest) ProtoMessage()               {}
func (*PutClusterConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *PutClusterConfigRequest) GetCluster() *metapb.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type PutClusterConfigResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PutClusterConfigResponse) Reset()                    { *m = PutClusterConfigResponse{} }
func (m *PutClusterConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*PutClusterConfigResponse) ProtoMessage()               {}
func (*PutClusterConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

type AskSplitRequest struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *AskSplitRequest) Reset()                    { *m = AskSplitRequest{} }
func (m *AskSplitRequest) String() string            { return proto.CompactTextString(m) }
func (*AskSplitRequest) ProtoMessage()               {}
func (*AskSplitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *AskSplitRequest) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type AskSplitResponse struct {
	// We split the region into two, first uses the origin
	// parent region id, and the second uses the new_region_id.
	// We must guarantee that the new_region_id is global unique.
	NewRegionId *uint64 `protobuf:"varint,1,opt,name=new_region_id" json:"new_region_id,omitempty"`
	// The peer ids for the new split region.
	NewPeerIds       []uint64 `protobuf:"varint,2,rep,name=new_peer_ids" json:"new_peer_ids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AskSplitResponse) Reset()                    { *m = AskSplitResponse{} }
func (m *AskSplitResponse) String() string            { return proto.CompactTextString(m) }
func (*AskSplitResponse) ProtoMessage()               {}
func (*AskSplitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *AskSplitResponse) GetNewRegionId() uint64 {
	if m != nil && m.NewRegionId != nil {
		return *m.NewRegionId
	}
	return 0
}

func (m *AskSplitResponse) GetNewPeerIds() []uint64 {
	if m != nil {
		return m.NewPeerIds
	}
	return nil
}

type RequestHeader struct {
	// 16 bytes, to distinguish request.
	Uuid             []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ClusterId        *uint64 `protobuf:"varint,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *RequestHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *RequestHeader) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

type ResponseHeader struct {
	// 16 bytes, to distinguish request.
	Uuid             []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ClusterId        *uint64 `protobuf:"varint,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	Error            *Error  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

func (m *ResponseHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *ResponseHeader) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *ResponseHeader) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Request struct {
	Header           *RequestHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CmdType          *CommandType             `protobuf:"varint,2,opt,name=cmd_type,enum=pdpb.CommandType" json:"cmd_type,omitempty"`
	Tso              *TsoRequest              `protobuf:"bytes,3,opt,name=tso" json:"tso,omitempty"`
	Bootstrap        *BootstrapRequest        `protobuf:"bytes,4,opt,name=bootstrap" json:"bootstrap,omitempty"`
	IsBootstrapped   *IsBootstrappedRequest   `protobuf:"bytes,5,opt,name=is_bootstrapped" json:"is_bootstrapped,omitempty"`
	AllocId          *AllocIdRequest          `protobuf:"bytes,6,opt,name=alloc_id" json:"alloc_id,omitempty"`
	GetStore         *GetStoreRequest         `protobuf:"bytes,7,opt,name=get_store" json:"get_store,omitempty"`
	PutStore         *PutStoreRequest         `protobuf:"bytes,8,opt,name=put_store" json:"put_store,omitempty"`
	AskSplit         *AskSplitRequest         `protobuf:"bytes,9,opt,name=ask_split" json:"ask_split,omitempty"`
	GetRegion        *GetRegionRequest        `protobuf:"bytes,10,opt,name=get_region" json:"get_region,omitempty"`
	RegionHeartbeat  *RegionHeartbeatRequest  `protobuf:"bytes,11,opt,name=region_heartbeat" json:"region_heartbeat,omitempty"`
	GetClusterConfig *GetClusterConfigRequest `protobuf:"bytes,12,opt,name=get_cluster_config" json:"get_cluster_config,omitempty"`
	PutClusterConfig *PutClusterConfigRequest `protobuf:"bytes,13,opt,name=put_cluster_config" json:"put_cluster_config,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

func (m *Request) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Request) GetTso() *TsoRequest {
	if m != nil {
		return m.Tso
	}
	return nil
}

func (m *Request) GetBootstrap() *BootstrapRequest {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Request) GetIsBootstrapped() *IsBootstrappedRequest {
	if m != nil {
		return m.IsBootstrapped
	}
	return nil
}

func (m *Request) GetAllocId() *AllocIdRequest {
	if m != nil {
		return m.AllocId
	}
	return nil
}

func (m *Request) GetGetStore() *GetStoreRequest {
	if m != nil {
		return m.GetStore
	}
	return nil
}

func (m *Request) GetPutStore() *PutStoreRequest {
	if m != nil {
		return m.PutStore
	}
	return nil
}

func (m *Request) GetAskSplit() *AskSplitRequest {
	if m != nil {
		return m.AskSplit
	}
	return nil
}

func (m *Request) GetGetRegion() *GetRegionRequest {
	if m != nil {
		return m.GetRegion
	}
	return nil
}

func (m *Request) GetRegionHeartbeat() *RegionHeartbeatRequest {
	if m != nil {
		return m.RegionHeartbeat
	}
	return nil
}

func (m *Request) GetGetClusterConfig() *GetClusterConfigRequest {
	if m != nil {
		return m.GetClusterConfig
	}
	return nil
}

func (m *Request) GetPutClusterConfig() *PutClusterConfigRequest {
	if m != nil {
		return m.PutClusterConfig
	}
	return nil
}

type Response struct {
	Header           *ResponseHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CmdType          *CommandType              `protobuf:"varint,2,opt,name=cmd_type,enum=pdpb.CommandType" json:"cmd_type,omitempty"`
	Tso              *TsoResponse              `protobuf:"bytes,3,opt,name=tso" json:"tso,omitempty"`
	Bootstrap        *BootstrapResponse        `protobuf:"bytes,4,opt,name=bootstrap" json:"bootstrap,omitempty"`
	IsBootstrapped   *IsBootstrappedResponse   `protobuf:"bytes,5,opt,name=is_bootstrapped" json:"is_bootstrapped,omitempty"`
	AllocId          *AllocIdResponse          `protobuf:"bytes,6,opt,name=alloc_id" json:"alloc_id,omitempty"`
	GetStore         *GetStoreResponse         `protobuf:"bytes,7,opt,name=get_store" json:"get_store,omitempty"`
	PutStore         *PutStoreResponse         `protobuf:"bytes,8,opt,name=put_store" json:"put_store,omitempty"`
	AskSplit         *AskSplitResponse         `protobuf:"bytes,9,opt,name=ask_split" json:"ask_split,omitempty"`
	GetRegion        *GetRegionResponse        `protobuf:"bytes,10,opt,name=get_region" json:"get_region,omitempty"`
	RegionHeartbeat  *RegionHeartbeatResponse  `protobuf:"bytes,11,opt,name=region_heartbeat" json:"region_heartbeat,omitempty"`
	GetClusterConfig *GetClusterConfigResponse `protobuf:"bytes,12,opt,name=get_cluster_config" json:"get_cluster_config,omitempty"`
	PutClusterConfig *PutClusterConfigResponse `protobuf:"bytes,13,opt,name=put_cluster_config" json:"put_cluster_config,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{28} }

func (m *Response) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Response) GetTso() *TsoResponse {
	if m != nil {
		return m.Tso
	}
	return nil
}

func (m *Response) GetBootstrap() *BootstrapResponse {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Response) GetIsBootstrapped() *IsBootstrappedResponse {
	if m != nil {
		return m.IsBootstrapped
	}
	return nil
}

func (m *Response) GetAllocId() *AllocIdResponse {
	if m != nil {
		return m.AllocId
	}
	return nil
}

func (m *Response) GetGetStore() *GetStoreResponse {
	if m != nil {
		return m.GetStore
	}
	return nil
}

func (m *Response) GetPutStore() *PutStoreResponse {
	if m != nil {
		return m.PutStore
	}
	return nil
}

func (m *Response) GetAskSplit() *AskSplitResponse {
	if m != nil {
		return m.AskSplit
	}
	return nil
}

func (m *Response) GetGetRegion() *GetRegionResponse {
	if m != nil {
		return m.GetRegion
	}
	return nil
}

func (m *Response) GetRegionHeartbeat() *RegionHeartbeatResponse {
	if m != nil {
		return m.RegionHeartbeat
	}
	return nil
}

func (m *Response) GetGetClusterConfig() *GetClusterConfigResponse {
	if m != nil {
		return m.GetClusterConfig
	}
	return nil
}

func (m *Response) GetPutClusterConfig() *PutClusterConfigResponse {
	if m != nil {
		return m.PutClusterConfig
	}
	return nil
}

type BootstrappedError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BootstrappedError) Reset()                    { *m = BootstrappedError{} }
func (m *BootstrappedError) String() string            { return proto.CompactTextString(m) }
func (*BootstrappedError) ProtoMessage()               {}
func (*BootstrappedError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{29} }

type Error struct {
	Message          *string            `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Bootstrapped     *BootstrappedError `protobuf:"bytes,2,opt,name=bootstrapped" json:"bootstrapped,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{30} }

func (m *Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Error) GetBootstrapped() *BootstrappedError {
	if m != nil {
		return m.Bootstrapped
	}
	return nil
}

func init() {
	proto.RegisterType((*Leader)(nil), "pdpb.Leader")
	proto.RegisterType((*TsoRequest)(nil), "pdpb.TsoRequest")
	proto.RegisterType((*Timestamp)(nil), "pdpb.Timestamp")
	proto.RegisterType((*TsoResponse)(nil), "pdpb.TsoResponse")
	proto.RegisterType((*BootstrapRequest)(nil), "pdpb.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "pdpb.BootstrapResponse")
	proto.RegisterType((*IsBootstrappedRequest)(nil), "pdpb.IsBootstrappedRequest")
	proto.RegisterType((*IsBootstrappedResponse)(nil), "pdpb.IsBootstrappedResponse")
	proto.RegisterType((*AllocIdRequest)(nil), "pdpb.AllocIdRequest")
	proto.RegisterType((*AllocIdResponse)(nil), "pdpb.AllocIdResponse")
	proto.RegisterType((*GetStoreRequest)(nil), "pdpb.GetStoreRequest")
	proto.RegisterType((*GetStoreResponse)(nil), "pdpb.GetStoreResponse")
	proto.RegisterType((*GetRegionRequest)(nil), "pdpb.GetRegionRequest")
	proto.RegisterType((*GetRegionResponse)(nil), "pdpb.GetRegionResponse")
	proto.RegisterType((*GetClusterConfigRequest)(nil), "pdpb.GetClusterConfigRequest")
	proto.RegisterType((*GetClusterConfigResponse)(nil), "pdpb.GetClusterConfigResponse")
	proto.RegisterType((*PutStoreRequest)(nil), "pdpb.PutStoreRequest")
	proto.RegisterType((*PutStoreResponse)(nil), "pdpb.PutStoreResponse")
	proto.RegisterType((*RegionHeartbeatRequest)(nil), "pdpb.RegionHeartbeatRequest")
	proto.RegisterType((*ChangePeer)(nil), "pdpb.ChangePeer")
	proto.RegisterType((*RegionHeartbeatResponse)(nil), "pdpb.RegionHeartbeatResponse")
	proto.RegisterType((*PutClusterConfigRequest)(nil), "pdpb.PutClusterConfigRequest")
	proto.RegisterType((*PutClusterConfigResponse)(nil), "pdpb.PutClusterConfigResponse")
	proto.RegisterType((*AskSplitRequest)(nil), "pdpb.AskSplitRequest")
	proto.RegisterType((*AskSplitResponse)(nil), "pdpb.AskSplitResponse")
	proto.RegisterType((*RequestHeader)(nil), "pdpb.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "pdpb.ResponseHeader")
	proto.RegisterType((*Request)(nil), "pdpb.Request")
	proto.RegisterType((*Response)(nil), "pdpb.Response")
	proto.RegisterType((*BootstrappedError)(nil), "pdpb.BootstrappedError")
	proto.RegisterType((*Error)(nil), "pdpb.Error")
	proto.RegisterEnum("pdpb.CommandType", CommandType_name, CommandType_value)
}

var fileDescriptor0 = []byte{
	// 1008 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xeb, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0xb1, 0x73, 0x3b, 0xce, 0x65, 0xea, 0xb6, 0x89, 0xe9, 0xb6, 0xd5, 0xca, 0x85, 0xa5,
	0xdb, 0x15, 0x81, 0x0d, 0x37, 0x71, 0x91, 0xb8, 0x54, 0x68, 0xb7, 0x12, 0x42, 0xd5, 0x6e, 0xe1,
	0x6f, 0xe4, 0xc6, 0xb3, 0xa9, 0xd5, 0x24, 0x36, 0x1e, 0x07, 0xd4, 0x87, 0xe4, 0x17, 0x2f, 0xc0,
	0xa3, 0x30, 0x97, 0x33, 0x8e, 0x33, 0x4e, 0x4a, 0xfe, 0xc5, 0x67, 0xbe, 0xef, 0x5c, 0xbf, 0x39,
	0x19, 0x80, 0x24, 0x4c, 0x6e, 0x87, 0x49, 0x1a, 0x67, 0xb1, 0x6b, 0x8b, 0xdf, 0x47, 0xed, 0x39,
	0xcd, 0x02, 0x6d, 0x3b, 0x6a, 0xa7, 0xc1, 0xbb, 0x4c, 0x7f, 0xf9, 0x67, 0x50, 0xff, 0x85, 0x06,
	0x21, 0x4d, 0xdd, 0x36, 0xd8, 0x41, 0x18, 0xa6, 0x5e, 0xe5, 0x69, 0xe5, 0xbc, 0xe5, 0x3a, 0x60,
	0x25, 0x51, 0xe8, 0x55, 0xf9, 0x87, 0xe5, 0x3f, 0x01, 0xb8, 0x61, 0xf1, 0x1b, 0xfa, 0xc7, 0x92,
	0xb2, 0xcc, 0xed, 0x40, 0x6d, 0x12, 0x2f, 0x17, 0x99, 0x44, 0x76, 0xfc, 0x21, 0xb4, 0x6e, 0xa2,
	0x39, 0x3f, 0x08, 0xe6, 0x89, 0x4b, 0xa0, 0x99, 0xdc, 0x3d, 0xb0, 0x68, 0x12, 0xcc, 0xe4, 0xb1,
	0xe5, 0xf6, 0xa0, 0x31, 0x8b, 0xa7, 0xd2, 0xa0, 0x9c, 0x8d, 0xc0, 0x91, 0xce, 0x58, 0x12, 0x2f,
	0x18, 0x75, 0xcf, 0x00, 0x32, 0x4d, 0x67, 0x9c, 0x63, 0x9d, 0x3b, 0xa3, 0xde, 0x50, 0xd6, 0x90,
	0xbb, 0xf5, 0xaf, 0x81, 0xfc, 0x14, 0xc7, 0x19, 0xcb, 0xd2, 0x20, 0xd1, 0x69, 0x1c, 0x43, 0x8d,
	0x65, 0x71, 0x4a, 0x65, 0x1c, 0x67, 0xd4, 0x19, 0x62, 0x95, 0x6f, 0x85, 0xd1, 0x3d, 0x85, 0x7a,
	0x4a, 0xa7, 0x51, 0xbc, 0x90, 0x51, 0x9d, 0x51, 0x57, 0x1f, 0xbf, 0x91, 0x56, 0x7f, 0x1f, 0xf6,
	0x0a, 0x1e, 0x55, 0x2e, 0xfe, 0x00, 0x0e, 0xaf, 0x58, 0x6e, 0x4e, 0x68, 0x88, 0xb1, 0x78, 0x8d,
	0x7d, 0xf3, 0x00, 0xd3, 0x3f, 0x80, 0xf6, 0x6d, 0xc1, 0x2e, 0x93, 0x69, 0xfa, 0x04, 0xba, 0x3f,
	0xce, 0x66, 0xf1, 0xe4, 0x2a, 0xf7, 0x70, 0x02, 0xbd, 0xdc, 0x82, 0x54, 0x80, 0x6a, 0xa4, 0x08,
	0x36, 0x1f, 0x43, 0xef, 0x15, 0xcd, 0x64, 0xea, 0xba, 0x3e, 0xde, 0x4a, 0x59, 0xdf, 0x38, 0x07,
	0x7d, 0x0a, 0x64, 0x05, 0x42, 0x27, 0x8f, 0x76, 0xc1, 0x7f, 0x26, 0x19, 0xaa, 0x64, 0xed, 0x97,
	0xc7, 0x55, 0x9d, 0x19, 0xdf, 0xd3, 0x07, 0x49, 0x6b, 0xfb, 0x9f, 0xc1, 0x5e, 0x01, 0x87, 0xae,
	0x57, 0x2d, 0xac, 0x6c, 0x6c, 0xe1, 0xfb, 0x30, 0xe0, 0xa4, 0xcb, 0xd9, 0x92, 0x65, 0x34, 0xbd,
	0x8c, 0x17, 0xef, 0xa2, 0xa9, 0xae, 0xf6, 0x3b, 0xf0, 0xca, 0x47, 0xe8, 0xf6, 0x29, 0x34, 0x26,
	0xea, 0x00, 0xfd, 0xf6, 0xb4, 0x5f, 0xc4, 0xfb, 0x9f, 0x40, 0xef, 0x7a, 0xb9, 0xde, 0x8c, 0xc7,
	0xcb, 0x74, 0x81, 0xac, 0x08, 0x38, 0xcb, 0xdf, 0xa1, 0xaf, 0xf2, 0x7c, 0x4d, 0x83, 0x34, 0xbb,
	0xa5, 0x41, 0xa6, 0x7d, 0xfd, 0x4f, 0x5d, 0x3c, 0x56, 0x7d, 0x26, 0xaf, 0x04, 0x4a, 0xa7, 0xad,
	0xcf, 0xaf, 0x29, 0x4f, 0xee, 0x37, 0x80, 0xcb, 0xbb, 0x60, 0x31, 0xa5, 0xe2, 0xcb, 0x7d, 0x01,
	0xce, 0x44, 0x7e, 0x8d, 0xb3, 0x87, 0x44, 0x65, 0xd7, 0x1d, 0xf5, 0x87, 0x78, 0xc5, 0x44, 0xe5,
	0x0a, 0x7c, 0xc3, 0x4f, 0xdd, 0x23, 0xb0, 0x13, 0xba, 0xc5, 0xed, 0x0f, 0x30, 0x28, 0xa5, 0x8b,
	0x0d, 0xfb, 0x30, 0x8f, 0x21, 0xd9, 0x2a, 0x69, 0xa2, 0xae, 0xc8, 0x2a, 0x15, 0xff, 0x5b, 0x18,
	0xf0, 0x26, 0x6c, 0x1a, 0xc7, 0x0e, 0x2d, 0x3f, 0x02, 0xaf, 0x4c, 0xc6, 0x4e, 0xbe, 0xe4, 0xd2,
	0x65, 0xf7, 0x6f, 0x93, 0x59, 0xb4, 0x6b, 0x0b, 0xfd, 0xef, 0x81, 0xac, 0x28, 0x58, 0xc6, 0x21,
	0x74, 0x16, 0xf4, 0xaf, 0x31, 0x6a, 0x4f, 0x8b, 0x5a, 0x5c, 0x20, 0x61, 0x16, 0xa5, 0x71, 0x23,
	0xe3, 0xcd, 0xb1, 0xb8, 0xd4, 0x5f, 0x42, 0x07, 0x63, 0xbd, 0xce, 0xb7, 0xd3, 0x72, 0x89, 0xa4,
	0xb6, 0xd0, 0x30, 0x16, 0x34, 0xc6, 0x25, 0x65, 0xfb, 0xbf, 0x42, 0x57, 0xc7, 0xda, 0x95, 0xc3,
	0x27, 0x52, 0xa3, 0x69, 0x1a, 0xa7, 0x9e, 0x25, 0xcb, 0x70, 0x54, 0x53, 0x7f, 0x16, 0x26, 0xff,
	0x1f, 0x1b, 0x1a, 0xba, 0x5e, 0xbe, 0x25, 0xef, 0x94, 0x24, 0x54, 0xbd, 0xfb, 0x0a, 0xb8, 0x9e,
	0xe2, 0x19, 0x34, 0x27, 0xf3, 0x50, 0x09, 0xa1, 0x2a, 0x85, 0xb0, 0x87, 0x43, 0x8a, 0xe7, 0xf3,
	0x60, 0x11, 0x4a, 0x0d, 0x9c, 0x80, 0x95, 0xb1, 0x18, 0xe3, 0xe1, 0x10, 0x0b, 0xbb, 0xf5, 0x39,
	0xb4, 0xf2, 0x75, 0xe2, 0xd9, 0x12, 0xd4, 0x57, 0xa0, 0xd2, 0xfe, 0xfb, 0x1c, 0x7a, 0x11, 0x1b,
	0xaf, 0x2d, 0x9f, 0x9a, 0x24, 0x3c, 0x51, 0x84, 0x8d, 0x9b, 0xcc, 0x7d, 0x06, 0xcd, 0x40, 0xec,
	0x21, 0xd1, 0x83, 0xba, 0x84, 0x1f, 0x28, 0xf8, 0xfa, 0xbe, 0x72, 0xcf, 0xa1, 0x35, 0xa5, 0xd9,
	0x58, 0x5d, 0xba, 0x86, 0x04, 0x1e, 0x2a, 0xa0, 0xb9, 0xa7, 0x38, 0x32, 0x59, 0x6a, 0x64, 0xb3,
	0x88, 0x34, 0x2f, 0x31, 0x47, 0x06, 0xec, 0x7e, 0xcc, 0x84, 0x2c, 0xbc, 0x56, 0x11, 0x69, 0xea,
	0xeb, 0x02, 0x40, 0x44, 0x47, 0x8d, 0x41, 0xb1, 0x0f, 0xa5, 0x7d, 0xf6, 0x25, 0x10, 0xd4, 0xd4,
	0x9d, 0xbe, 0x3a, 0x9e, 0x23, 0x19, 0xc7, 0x7a, 0x4a, 0x1b, 0xd7, 0xc0, 0xd7, 0xe0, 0x8a, 0x18,
	0x5a, 0x13, 0x13, 0x29, 0x7a, 0xaf, 0x2d, 0x99, 0x27, 0x79, 0xac, 0x8d, 0xf7, 0x89, 0x53, 0x45,
	0xc9, 0x06, 0xb5, 0x53, 0xa4, 0x6e, 0xb9, 0x8a, 0xfe, 0xbf, 0x36, 0x34, 0xf3, 0x2b, 0xf1, 0x81,
	0x21, 0xab, 0x03, 0x9d, 0xf0, 0x9a, 0x8c, 0x77, 0xd2, 0xd5, 0x69, 0x51, 0x57, 0x7b, 0x05, 0x5d,
	0x61, 0xa8, 0x8b, 0xb2, 0xb0, 0x06, 0x25, 0x61, 0x21, 0xf6, 0x8b, 0x6d, 0xca, 0x3a, 0xde, 0xac,
	0x2c, 0xa4, 0x7d, 0x54, 0x92, 0xd6, 0xa1, 0x21, 0x2d, 0x04, 0x3e, 0x2f, 0x6b, 0xab, 0x6f, 0x6a,
	0x6b, 0x05, 0x35, 0xc5, 0xd5, 0x37, 0xc5, 0xb5, 0x82, 0x9a, 0xea, 0xea, 0x9b, 0xea, 0x42, 0xe8,
	0x8b, 0x0d, 0xf2, 0x1a, 0x94, 0xe4, 0x85, 0xe0, 0xaf, 0xb6, 0xea, 0xeb, 0x64, 0x8b, 0xbe, 0x90,
	0xf8, 0xcd, 0x23, 0x02, 0x3b, 0xdd, 0x26, 0xb0, 0x15, 0x77, 0xab, 0xc2, 0x4e, 0xb7, 0x29, 0x0c,
	0xf7, 0x75, 0xf1, 0x69, 0xc3, 0xe7, 0xa3, 0xb6, 0xd9, 0x2b, 0xa8, 0xc9, 0x1f, 0xe2, 0x3d, 0xc6,
	0x5f, 0x55, 0x2c, 0x98, 0x52, 0x7c, 0xe9, 0x7d, 0x6c, 0xbc, 0x60, 0xaa, 0x1b, 0xc5, 0xa1, 0x1d,
	0x5d, 0xfc, 0x5d, 0x01, 0xa7, 0x28, 0x3c, 0x07, 0x1a, 0x57, 0x8b, 0x3f, 0x83, 0x59, 0x14, 0x92,
	0xf7, 0xdc, 0x06, 0x58, 0x5c, 0x74, 0xa4, 0xc2, 0xdf, 0x88, 0xad, 0x9c, 0x4a, 0xaa, 0x7c, 0xf7,
	0x76, 0xd7, 0x45, 0x43, 0x2c, 0x41, 0x44, 0x61, 0x10, 0x9b, 0xaf, 0xea, 0xa6, 0x9e, 0x3d, 0xa9,
	0x89, 0x2f, 0x3d, 0x5e, 0x52, 0x17, 0x5f, 0x7a, 0x82, 0xa4, 0x21, 0x3c, 0xe7, 0x33, 0x22, 0x4d,
	0x77, 0x1f, 0x7a, 0x46, 0xff, 0x49, 0x8b, 0xff, 0xa7, 0x10, 0xb3, 0xb3, 0x04, 0x84, 0xd5, 0xec,
	0x19, 0x71, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x2c, 0x36, 0x36, 0x2f, 0x0b, 0x00, 0x00,
}
