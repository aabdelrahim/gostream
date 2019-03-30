// Code generated by protoc-gen-go. DO NOT EDIT.
// source: songservice.proto

package songservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddSongRequest struct {
	NewSong              *Song    `protobuf:"bytes,1,opt,name=newSong,proto3" json:"newSong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSongRequest) Reset()         { *m = AddSongRequest{} }
func (m *AddSongRequest) String() string { return proto.CompactTextString(m) }
func (*AddSongRequest) ProtoMessage()    {}
func (*AddSongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{0}
}

func (m *AddSongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSongRequest.Unmarshal(m, b)
}
func (m *AddSongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSongRequest.Marshal(b, m, deterministic)
}
func (m *AddSongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSongRequest.Merge(m, src)
}
func (m *AddSongRequest) XXX_Size() int {
	return xxx_messageInfo_AddSongRequest.Size(m)
}
func (m *AddSongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSongRequest proto.InternalMessageInfo

func (m *AddSongRequest) GetNewSong() *Song {
	if m != nil {
		return m.NewSong
	}
	return nil
}

type UpdateSongRequest struct {
	SongID               string   `protobuf:"bytes,1,opt,name=songID,proto3" json:"songID,omitempty"`
	UpdatedSong          *Song    `protobuf:"bytes,2,opt,name=updatedSong,proto3" json:"updatedSong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSongRequest) Reset()         { *m = UpdateSongRequest{} }
func (m *UpdateSongRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSongRequest) ProtoMessage()    {}
func (*UpdateSongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{1}
}

func (m *UpdateSongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSongRequest.Unmarshal(m, b)
}
func (m *UpdateSongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSongRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSongRequest.Merge(m, src)
}
func (m *UpdateSongRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSongRequest.Size(m)
}
func (m *UpdateSongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSongRequest proto.InternalMessageInfo

func (m *UpdateSongRequest) GetSongID() string {
	if m != nil {
		return m.SongID
	}
	return ""
}

func (m *UpdateSongRequest) GetUpdatedSong() *Song {
	if m != nil {
		return m.UpdatedSong
	}
	return nil
}

type GetSongRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Artists              []string `protobuf:"bytes,2,rep,name=artists,proto3" json:"artists,omitempty"`
	SongID               string   `protobuf:"bytes,3,opt,name=songID,proto3" json:"songID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSongRequest) Reset()         { *m = GetSongRequest{} }
func (m *GetSongRequest) String() string { return proto.CompactTextString(m) }
func (*GetSongRequest) ProtoMessage()    {}
func (*GetSongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{2}
}

func (m *GetSongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSongRequest.Unmarshal(m, b)
}
func (m *GetSongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSongRequest.Marshal(b, m, deterministic)
}
func (m *GetSongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSongRequest.Merge(m, src)
}
func (m *GetSongRequest) XXX_Size() int {
	return xxx_messageInfo_GetSongRequest.Size(m)
}
func (m *GetSongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSongRequest proto.InternalMessageInfo

func (m *GetSongRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetSongRequest) GetArtists() []string {
	if m != nil {
		return m.Artists
	}
	return nil
}

func (m *GetSongRequest) GetSongID() string {
	if m != nil {
		return m.SongID
	}
	return ""
}

type GetSongResponse struct {
	Result               []*GetSongResponse_GetSongResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetSongResponse) Reset()         { *m = GetSongResponse{} }
func (m *GetSongResponse) String() string { return proto.CompactTextString(m) }
func (*GetSongResponse) ProtoMessage()    {}
func (*GetSongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{3}
}

func (m *GetSongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSongResponse.Unmarshal(m, b)
}
func (m *GetSongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSongResponse.Marshal(b, m, deterministic)
}
func (m *GetSongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSongResponse.Merge(m, src)
}
func (m *GetSongResponse) XXX_Size() int {
	return xxx_messageInfo_GetSongResponse.Size(m)
}
func (m *GetSongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSongResponse proto.InternalMessageInfo

func (m *GetSongResponse) GetResult() []*GetSongResponse_GetSongResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetSongResponse_GetSongResult struct {
	Song                 *Song    `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	SongID               string   `protobuf:"bytes,2,opt,name=songID,proto3" json:"songID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSongResponse_GetSongResult) Reset()         { *m = GetSongResponse_GetSongResult{} }
func (m *GetSongResponse_GetSongResult) String() string { return proto.CompactTextString(m) }
func (*GetSongResponse_GetSongResult) ProtoMessage()    {}
func (*GetSongResponse_GetSongResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{3, 0}
}

func (m *GetSongResponse_GetSongResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSongResponse_GetSongResult.Unmarshal(m, b)
}
func (m *GetSongResponse_GetSongResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSongResponse_GetSongResult.Marshal(b, m, deterministic)
}
func (m *GetSongResponse_GetSongResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSongResponse_GetSongResult.Merge(m, src)
}
func (m *GetSongResponse_GetSongResult) XXX_Size() int {
	return xxx_messageInfo_GetSongResponse_GetSongResult.Size(m)
}
func (m *GetSongResponse_GetSongResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSongResponse_GetSongResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetSongResponse_GetSongResult proto.InternalMessageInfo

func (m *GetSongResponse_GetSongResult) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *GetSongResponse_GetSongResult) GetSongID() string {
	if m != nil {
		return m.SongID
	}
	return ""
}

type DeleteSongRequest struct {
	SongID               string   `protobuf:"bytes,1,opt,name=SongID,proto3" json:"SongID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSongRequest) Reset()         { *m = DeleteSongRequest{} }
func (m *DeleteSongRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSongRequest) ProtoMessage()    {}
func (*DeleteSongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{4}
}

func (m *DeleteSongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSongRequest.Unmarshal(m, b)
}
func (m *DeleteSongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSongRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSongRequest.Merge(m, src)
}
func (m *DeleteSongRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSongRequest.Size(m)
}
func (m *DeleteSongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSongRequest proto.InternalMessageInfo

func (m *DeleteSongRequest) GetSongID() string {
	if m != nil {
		return m.SongID
	}
	return ""
}

type Song struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Artists              []string `protobuf:"bytes,2,rep,name=artists,proto3" json:"artists,omitempty"`
	Audio                []byte   `protobuf:"bytes,3,opt,name=audio,proto3" json:"audio,omitempty"`
	Audioformate         string   `protobuf:"bytes,4,opt,name=audioformate,proto3" json:"audioformate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{5}
}

func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (m *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(m, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Song) GetArtists() []string {
	if m != nil {
		return m.Artists
	}
	return nil
}

func (m *Song) GetAudio() []byte {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *Song) GetAudioformate() string {
	if m != nil {
		return m.Audioformate
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e82c45bb8447e0, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddSongRequest)(nil), "songservice.AddSongRequest")
	proto.RegisterType((*UpdateSongRequest)(nil), "songservice.UpdateSongRequest")
	proto.RegisterType((*GetSongRequest)(nil), "songservice.GetSongRequest")
	proto.RegisterType((*GetSongResponse)(nil), "songservice.GetSongResponse")
	proto.RegisterType((*GetSongResponse_GetSongResult)(nil), "songservice.GetSongResponse.GetSongResult")
	proto.RegisterType((*DeleteSongRequest)(nil), "songservice.DeleteSongRequest")
	proto.RegisterType((*Song)(nil), "songservice.Song")
	proto.RegisterType((*Empty)(nil), "songservice.Empty")
}

func init() { proto.RegisterFile("songservice.proto", fileDescriptor_c8e82c45bb8447e0) }

var fileDescriptor_c8e82c45bb8447e0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x7d, 0xf9, 0x68, 0x4a, 0x6f, 0xfa, 0xfa, 0xc8, 0xe5, 0x21, 0xa1, 0x8a, 0x94, 0x01, 0xa1,
	0x58, 0xe8, 0xa2, 0xdd, 0xb8, 0x51, 0x68, 0xa9, 0x14, 0x37, 0x2e, 0x52, 0x74, 0x6d, 0x34, 0xd7,
	0x52, 0x68, 0x33, 0x31, 0x33, 0x51, 0xfc, 0x0d, 0xfe, 0x0e, 0xff, 0xa7, 0x64, 0xd2, 0xca, 0x4c,
	0x6b, 0x0b, 0xee, 0xee, 0xc7, 0xb9, 0xe7, 0x9e, 0x9c, 0xb9, 0x81, 0x40, 0xf0, 0x74, 0x2e, 0x28,
	0x7f, 0x5d, 0x3c, 0x51, 0x3f, 0xcb, 0xb9, 0xe4, 0xe8, 0x6b, 0x25, 0x76, 0x09, 0xad, 0x51, 0x92,
	0xcc, 0x78, 0x3a, 0x8f, 0xe8, 0xa5, 0x20, 0x21, 0xb1, 0x07, 0xf5, 0x94, 0xde, 0xca, 0x4a, 0x68,
	0x75, 0xac, 0xae, 0x3f, 0x08, 0xfa, 0x3a, 0x87, 0x82, 0x6e, 0x10, 0xec, 0x01, 0x82, 0xbb, 0x2c,
	0x89, 0x25, 0xe9, 0x0c, 0x47, 0xe0, 0x95, 0x13, 0x37, 0x13, 0x45, 0xd0, 0x88, 0xd6, 0x19, 0x0e,
	0xc1, 0x2f, 0x14, 0x58, 0xed, 0x0b, 0xed, 0x7d, 0xec, 0x3a, 0x8a, 0xdd, 0x43, 0x6b, 0x4a, 0x52,
	0xa7, 0x47, 0x70, 0xd3, 0x78, 0x45, 0x6b, 0x72, 0x15, 0x63, 0x08, 0xf5, 0x38, 0x97, 0x0b, 0x21,
	0x45, 0x68, 0x77, 0x9c, 0x6e, 0x23, 0xda, 0xa4, 0x9a, 0x18, 0x47, 0x17, 0xc3, 0x3e, 0x2d, 0xf8,
	0xf7, 0x4d, 0x2c, 0x32, 0x9e, 0x0a, 0xc2, 0x31, 0x78, 0x39, 0x89, 0x62, 0x29, 0x43, 0xab, 0xe3,
	0x74, 0xfd, 0xc1, 0xb9, 0xa1, 0x6d, 0x0b, 0xad, 0xe5, 0xc5, 0x52, 0x46, 0xeb, 0xc9, 0xf6, 0x2d,
	0xfc, 0x35, 0x1a, 0x78, 0x06, 0xae, 0x38, 0x68, 0xa6, 0x6a, 0x6b, 0x3a, 0x6d, 0x43, 0x67, 0x0f,
	0x82, 0x09, 0x2d, 0x69, 0xc7, 0xe1, 0x99, 0xe1, 0x70, 0x95, 0xb1, 0x14, 0xdc, 0x32, 0xfa, 0xa5,
	0x45, 0xff, 0xa1, 0x16, 0x17, 0xc9, 0x82, 0x2b, 0x87, 0x9a, 0x51, 0x95, 0x20, 0x83, 0xa6, 0x0a,
	0x9e, 0x79, 0xbe, 0x8a, 0x25, 0x85, 0xae, 0xe2, 0x32, 0x6a, 0xac, 0x0e, 0xb5, 0xeb, 0x55, 0x26,
	0xdf, 0x07, 0x1f, 0x36, 0xf8, 0xe5, 0xe6, 0x59, 0xf5, 0x61, 0x78, 0x01, 0xce, 0x28, 0x49, 0xf0,
	0xd8, 0xf8, 0x5a, 0xf3, 0xd0, 0xda, 0x68, 0x34, 0x15, 0x0f, 0xfb, 0x83, 0x57, 0xe0, 0x55, 0x17,
	0x85, 0xa7, 0x46, 0x7f, 0xe7, 0xcc, 0xf6, 0xcc, 0x8f, 0xc1, 0x99, 0x92, 0xdc, 0xda, 0x6c, 0x5e,
	0x50, 0xfb, 0xe4, 0xd0, 0xbb, 0x56, 0x1a, 0x2a, 0xcf, 0xb7, 0x34, 0xec, 0x3c, 0xc4, 0xcf, 0x1a,
	0x1e, 0x3d, 0xf5, 0xa3, 0x0d, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0x92, 0xf9, 0x2a, 0x7d,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SongServiceClient is the client API for SongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SongServiceClient interface {
	Add(ctx context.Context, in *AddSongRequest, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *GetSongRequest, opts ...grpc.CallOption) (*GetSongResponse, error)
	Delete(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*Empty, error)
}

type songServiceClient struct {
	cc *grpc.ClientConn
}

func NewSongServiceClient(cc *grpc.ClientConn) SongServiceClient {
	return &songServiceClient{cc}
}

func (c *songServiceClient) Add(ctx context.Context, in *AddSongRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/songservice.SongService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Update(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/songservice.SongService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Get(ctx context.Context, in *GetSongRequest, opts ...grpc.CallOption) (*GetSongResponse, error) {
	out := new(GetSongResponse)
	err := c.cc.Invoke(ctx, "/songservice.SongService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Delete(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/songservice.SongService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongServiceServer is the server API for SongService service.
type SongServiceServer interface {
	Add(context.Context, *AddSongRequest) (*Empty, error)
	Update(context.Context, *UpdateSongRequest) (*Empty, error)
	Get(context.Context, *GetSongRequest) (*GetSongResponse, error)
	Delete(context.Context, *DeleteSongRequest) (*Empty, error)
}

func RegisterSongServiceServer(s *grpc.Server, srv SongServiceServer) {
	s.RegisterService(&_SongService_serviceDesc, srv)
}

func _SongService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songservice.SongService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Add(ctx, req.(*AddSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songservice.SongService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Update(ctx, req.(*UpdateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songservice.SongService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Get(ctx, req.(*GetSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songservice.SongService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Delete(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SongService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "songservice.SongService",
	HandlerType: (*SongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _SongService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SongService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SongService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SongService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "songservice.proto",
}
