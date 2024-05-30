// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: player_graph_service.proto

package playergraphpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The response message containing the data.
type PlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiotIdGameName string `protobuf:"bytes,1,opt,name=RiotIdGameName,proto3" json:"RiotIdGameName,omitempty"`
	RiotIdTagline  string `protobuf:"bytes,2,opt,name=RiotIdTagline,proto3" json:"RiotIdTagline,omitempty"`
}

func (x *PlayerResponse) Reset() {
	*x = PlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_graph_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerResponse) ProtoMessage() {}

func (x *PlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_player_graph_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerResponse.ProtoReflect.Descriptor instead.
func (*PlayerResponse) Descriptor() ([]byte, []int) {
	return file_player_graph_service_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerResponse) GetRiotIdGameName() string {
	if x != nil {
		return x.RiotIdGameName
	}
	return ""
}

func (x *PlayerResponse) GetRiotIdTagline() string {
	if x != nil {
		return x.RiotIdTagline
	}
	return ""
}

type GraphRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *GraphRequest) Reset() {
	*x = GraphRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_graph_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphRequest) ProtoMessage() {}

func (x *GraphRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_graph_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphRequest.ProtoReflect.Descriptor instead.
func (*GraphRequest) Descriptor() ([]byte, []int) {
	return file_player_graph_service_proto_rawDescGZIP(), []int{1}
}

func (x *GraphRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type GraphResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players     []*Player     `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Connections []*Connection `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *GraphResponse) Reset() {
	*x = GraphResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_graph_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphResponse) ProtoMessage() {}

func (x *GraphResponse) ProtoReflect() protoreflect.Message {
	mi := &file_player_graph_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphResponse.ProtoReflect.Descriptor instead.
func (*GraphResponse) Descriptor() ([]byte, []int) {
	return file_player_graph_service_proto_rawDescGZIP(), []int{2}
}

func (x *GraphResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GraphResponse) GetConnections() []*Connection {
	if x != nil {
		return x.Connections
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid          string `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	RiotIdName     string `protobuf:"bytes,2,opt,name=riotIdName,proto3" json:"riotIdName,omitempty"`
	RiotIdGameName string `protobuf:"bytes,3,opt,name=riotIdGameName,proto3" json:"riotIdGameName,omitempty"`
	LastSeen       string `protobuf:"bytes,4,opt,name=lastSeen,proto3" json:"lastSeen,omitempty"`
	Role           string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_graph_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_graph_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_graph_service_proto_rawDescGZIP(), []int{3}
}

func (x *Player) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *Player) GetRiotIdName() string {
	if x != nil {
		return x.RiotIdName
	}
	return ""
}

func (x *Player) GetRiotIdGameName() string {
	if x != nil {
		return x.RiotIdGameName
	}
	return ""
}

func (x *Player) GetLastSeen() string {
	if x != nil {
		return x.LastSeen
	}
	return ""
}

func (x *Player) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerOneUuid string `protobuf:"bytes,1,opt,name=playerOneUuid,proto3" json:"playerOneUuid,omitempty"`
	PlayerTwoUuid string `protobuf:"bytes,2,opt,name=playerTwoUuid,proto3" json:"playerTwoUuid,omitempty"`
	GameId        string `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
	TimesPlayed   int32  `protobuf:"varint,4,opt,name=timesPlayed,proto3" json:"timesPlayed,omitempty"`
	LastPlayed    string `protobuf:"bytes,5,opt,name=lastPlayed,proto3" json:"lastPlayed,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_graph_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_player_graph_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_player_graph_service_proto_rawDescGZIP(), []int{4}
}

func (x *Connection) GetPlayerOneUuid() string {
	if x != nil {
		return x.PlayerOneUuid
	}
	return ""
}

func (x *Connection) GetPlayerTwoUuid() string {
	if x != nil {
		return x.PlayerTwoUuid
	}
	return ""
}

func (x *Connection) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *Connection) GetTimesPlayed() int32 {
	if x != nil {
		return x.TimesPlayed
	}
	return 0
}

func (x *Connection) GetLastPlayed() string {
	if x != nil {
		return x.LastPlayed
	}
	return ""
}

var File_player_graph_service_proto protoreflect.FileDescriptor

var file_player_graph_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x69, 0x6f, 0x74,
	0x49, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x52, 0x69, 0x6f, 0x74, 0x49, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x52, 0x69, 0x6f, 0x74, 0x49, 0x64, 0x54, 0x61, 0x67, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x69, 0x6f, 0x74, 0x49, 0x64, 0x54,
	0x61, 0x67, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x79,
	0x0a, 0x0d, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x39,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x06, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x69,
	0x6f, 0x74, 0x49, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x69, 0x6f, 0x74, 0x49, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x69,
	0x6f, 0x74, 0x49, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x69, 0x6f, 0x74, 0x49, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x55, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4f, 0x6e, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x54, 0x77, 0x6f, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x32, 0xa5, 0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x79,
	0x6c, 0x61, 0x6e, 0x44, 0x75, 0x6e, 0x68, 0x61, 0x6d, 0x30, 0x33, 0x2f, 0x4c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2d, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_graph_service_proto_rawDescOnce sync.Once
	file_player_graph_service_proto_rawDescData = file_player_graph_service_proto_rawDesc
)

func file_player_graph_service_proto_rawDescGZIP() []byte {
	file_player_graph_service_proto_rawDescOnce.Do(func() {
		file_player_graph_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_graph_service_proto_rawDescData)
	})
	return file_player_graph_service_proto_rawDescData
}

var file_player_graph_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_player_graph_service_proto_goTypes = []interface{}{
	(*PlayerResponse)(nil), // 0: playergraph.PlayerResponse
	(*GraphRequest)(nil),   // 1: playergraph.GraphRequest
	(*GraphResponse)(nil),  // 2: playergraph.GraphResponse
	(*Player)(nil),         // 3: playergraph.Player
	(*Connection)(nil),     // 4: playergraph.Connection
	(*emptypb.Empty)(nil),  // 5: google.protobuf.Empty
}
var file_player_graph_service_proto_depIdxs = []int32{
	3, // 0: playergraph.GraphResponse.players:type_name -> playergraph.Player
	4, // 1: playergraph.GraphResponse.connections:type_name -> playergraph.Connection
	5, // 2: playergraph.PlayerGraphService.GetPlayerData:input_type -> google.protobuf.Empty
	1, // 3: playergraph.PlayerGraphService.GetPlayerGraph:input_type -> playergraph.GraphRequest
	0, // 4: playergraph.PlayerGraphService.GetPlayerData:output_type -> playergraph.PlayerResponse
	2, // 5: playergraph.PlayerGraphService.GetPlayerGraph:output_type -> playergraph.GraphResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_player_graph_service_proto_init() }
func file_player_graph_service_proto_init() {
	if File_player_graph_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_graph_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_graph_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_graph_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_graph_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_graph_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_player_graph_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_graph_service_proto_goTypes,
		DependencyIndexes: file_player_graph_service_proto_depIdxs,
		MessageInfos:      file_player_graph_service_proto_msgTypes,
	}.Build()
	File_player_graph_service_proto = out.File
	file_player_graph_service_proto_rawDesc = nil
	file_player_graph_service_proto_goTypes = nil
	file_player_graph_service_proto_depIdxs = nil
}