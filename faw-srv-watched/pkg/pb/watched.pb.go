// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/watched/pb/watched.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Rate        float32                `protobuf:"fixed32,5,opt,name=rate,proto3" json:"rate,omitempty"`
	Cover       string                 `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
	UserId      uint64                 `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Movie) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Movie) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Movie) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// movie response
type MovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie  *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MovieResponse) Reset() {
	*x = MovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResponse) ProtoMessage() {}

func (x *MovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResponse.ProtoReflect.Descriptor instead.
func (*MovieResponse) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{1}
}

func (x *MovieResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *MovieResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MovieResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type WatchMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId uint32 `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *WatchMovieRequest) Reset() {
	*x = WatchMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchMovieRequest) ProtoMessage() {}

func (x *WatchMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchMovieRequest.ProtoReflect.Descriptor instead.
func (*WatchMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{2}
}

func (x *WatchMovieRequest) GetMovieId() uint32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *WatchMovieRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// status response for movie
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{3}
}

func (x *StatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StatusResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// list of movies response
type WatchedListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	Status string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *WatchedListResponse) Reset() {
	*x = WatchedListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchedListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchedListResponse) ProtoMessage() {}

func (x *WatchedListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchedListResponse.ProtoReflect.Descriptor instead.
func (*WatchedListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{4}
}

func (x *WatchedListResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *WatchedListResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WatchedListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RatingMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId uint32 `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	Rating  uint32 `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Review  string `protobuf:"bytes,3,opt,name=review,proto3" json:"review,omitempty"`
	UserId  uint64 `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RatingMovieRequest) Reset() {
	*x = RatingMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingMovieRequest) ProtoMessage() {}

func (x *RatingMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingMovieRequest.ProtoReflect.Descriptor instead.
func (*RatingMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{5}
}

func (x *RatingMovieRequest) GetMovieId() uint32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *RatingMovieRequest) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *RatingMovieRequest) GetReview() string {
	if x != nil {
		return x.Review
	}
	return ""
}

func (x *RatingMovieRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// movies list request
type WatchedListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortBy    string `protobuf:"bytes,1,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
	Page      uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size      uint32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	UserId    uint64 `protobuf:"varint,5,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *WatchedListRequest) Reset() {
	*x = WatchedListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_watched_pb_watched_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchedListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchedListRequest) ProtoMessage() {}

func (x *WatchedListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_watched_pb_watched_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchedListRequest.ProtoReflect.Descriptor instead.
func (*WatchedListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_watched_pb_watched_proto_rawDescGZIP(), []int{6}
}

func (x *WatchedListRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *WatchedListRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *WatchedListRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *WatchedListRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WatchedListRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_pkg_watched_pb_watched_proto protoreflect.FileDescriptor

var file_pkg_watched_pb_watched_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x0d, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x45, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6b, 0x0a, 0x13, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x76, 0x0a, 0x12, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x12,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xe7, 0x01, 0x0a, 0x0e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x1a, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x1b, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_watched_pb_watched_proto_rawDescOnce sync.Once
	file_pkg_watched_pb_watched_proto_rawDescData = file_pkg_watched_pb_watched_proto_rawDesc
)

func file_pkg_watched_pb_watched_proto_rawDescGZIP() []byte {
	file_pkg_watched_pb_watched_proto_rawDescOnce.Do(func() {
		file_pkg_watched_pb_watched_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_watched_pb_watched_proto_rawDescData)
	})
	return file_pkg_watched_pb_watched_proto_rawDescData
}

var file_pkg_watched_pb_watched_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_watched_pb_watched_proto_goTypes = []interface{}{
	(*Movie)(nil),                 // 0: watched.Movie
	(*MovieResponse)(nil),         // 1: watched.MovieResponse
	(*WatchMovieRequest)(nil),     // 2: watched.WatchMovieRequest
	(*StatusResponse)(nil),        // 3: watched.StatusResponse
	(*WatchedListResponse)(nil),   // 4: watched.WatchedListResponse
	(*RatingMovieRequest)(nil),    // 5: watched.RatingMovieRequest
	(*WatchedListRequest)(nil),    // 6: watched.WatchedListRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_pkg_watched_pb_watched_proto_depIdxs = []int32{
	7, // 0: watched.Movie.date:type_name -> google.protobuf.Timestamp
	0, // 1: watched.MovieResponse.movie:type_name -> watched.Movie
	0, // 2: watched.WatchedListResponse.movies:type_name -> watched.Movie
	2, // 3: watched.WatchedService.WatchMovie:input_type -> watched.WatchMovieRequest
	5, // 4: watched.WatchedService.RatingMovie:input_type -> watched.RatingMovieRequest
	6, // 5: watched.WatchedService.WatchedList:input_type -> watched.WatchedListRequest
	3, // 6: watched.WatchedService.WatchMovie:output_type -> watched.StatusResponse
	1, // 7: watched.WatchedService.RatingMovie:output_type -> watched.MovieResponse
	4, // 8: watched.WatchedService.WatchedList:output_type -> watched.WatchedListResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_watched_pb_watched_proto_init() }
func file_pkg_watched_pb_watched_proto_init() {
	if File_pkg_watched_pb_watched_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_watched_pb_watched_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_pkg_watched_pb_watched_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResponse); i {
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
		file_pkg_watched_pb_watched_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchMovieRequest); i {
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
		file_pkg_watched_pb_watched_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_pkg_watched_pb_watched_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchedListResponse); i {
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
		file_pkg_watched_pb_watched_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingMovieRequest); i {
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
		file_pkg_watched_pb_watched_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchedListRequest); i {
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
			RawDescriptor: file_pkg_watched_pb_watched_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_watched_pb_watched_proto_goTypes,
		DependencyIndexes: file_pkg_watched_pb_watched_proto_depIdxs,
		MessageInfos:      file_pkg_watched_pb_watched_proto_msgTypes,
	}.Build()
	File_pkg_watched_pb_watched_proto = out.File
	file_pkg_watched_pb_watched_proto_rawDesc = nil
	file_pkg_watched_pb_watched_proto_goTypes = nil
	file_pkg_watched_pb_watched_proto_depIdxs = nil
}
