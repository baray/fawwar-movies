// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/movie/pb/movie.proto

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
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[0]
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
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{0}
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

// create movie
type CreateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	UserId      uint64                 `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateMovieRequest) Reset() {
	*x = CreateMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieRequest) ProtoMessage() {}

func (x *CreateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieRequest.ProtoReflect.Descriptor instead.
func (*CreateMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMovieRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMovieRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMovieRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *CreateMovieRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// delete the movie you created
type MovieByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId uint32 `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MovieByUserRequest) Reset() {
	*x = MovieByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieByUserRequest) ProtoMessage() {}

func (x *MovieByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieByUserRequest.ProtoReflect.Descriptor instead.
func (*MovieByUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{2}
}

func (x *MovieByUserRequest) GetMovieId() uint32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *MovieByUserRequest) GetUserId() uint64 {
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
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[3]
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
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{3}
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

// get movie information
type FindMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindMovieRequest) Reset() {
	*x = FindMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMovieRequest) ProtoMessage() {}

func (x *FindMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMovieRequest.ProtoReflect.Descriptor instead.
func (*FindMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{4}
}

func (x *FindMovieRequest) GetId() uint32 {
	if x != nil {
		return x.Id
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
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResponse) ProtoMessage() {}

func (x *MovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[5]
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
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{5}
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

// movies list request
type MoviesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortBy    string `protobuf:"bytes,1,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
	Page      uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size      uint32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *MoviesListRequest) Reset() {
	*x = MoviesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviesListRequest) ProtoMessage() {}

func (x *MoviesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviesListRequest.ProtoReflect.Descriptor instead.
func (*MoviesListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{6}
}

func (x *MoviesListRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *MoviesListRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *MoviesListRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *MoviesListRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// list of movies response
type MoviesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	Status string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MoviesListResponse) Reset() {
	*x = MoviesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviesListResponse) ProtoMessage() {}

func (x *MoviesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviesListResponse.ProtoReflect.Descriptor instead.
func (*MoviesListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{7}
}

func (x *MoviesListResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *MoviesListResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MoviesListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// update movie
type UpdateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId     uint32                 `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	UserId      uint64                 `protobuf:"varint,5,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateMovieRequest) Reset() {
	*x = UpdateMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieRequest) ProtoMessage() {}

func (x *UpdateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieRequest.ProtoReflect.Descriptor instead.
func (*UpdateMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateMovieRequest) GetMovieId() uint32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *UpdateMovieRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMovieRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateMovieRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *UpdateMovieRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// update cover of movie request
type UpdateCoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId   uint32 `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	CoverName string `protobuf:"bytes,2,opt,name=coverName,proto3" json:"coverName,omitempty"`
	UserId    uint64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateCoverRequest) Reset() {
	*x = UpdateCoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_movie_pb_movie_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoverRequest) ProtoMessage() {}

func (x *UpdateCoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_movie_pb_movie_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoverRequest.ProtoReflect.Descriptor instead.
func (*UpdateCoverRequest) Descriptor() ([]byte, []int) {
	return file_pkg_movie_pb_movie_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCoverRequest) GetMovieId() uint32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *UpdateCoverRequest) GetCoverName() string {
	if x != nil {
		return x.CoverName
	}
	return ""
}

func (x *UpdateCoverRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_pkg_movie_pb_movie_proto protoreflect.FileDescriptor

var file_pkg_movie_pb_movie_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x12, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x22, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x71, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x68, 0x0a, 0x12, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x9d, 0x03, 0x0a, 0x0c, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x19, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x17,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_movie_pb_movie_proto_rawDescOnce sync.Once
	file_pkg_movie_pb_movie_proto_rawDescData = file_pkg_movie_pb_movie_proto_rawDesc
)

func file_pkg_movie_pb_movie_proto_rawDescGZIP() []byte {
	file_pkg_movie_pb_movie_proto_rawDescOnce.Do(func() {
		file_pkg_movie_pb_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_movie_pb_movie_proto_rawDescData)
	})
	return file_pkg_movie_pb_movie_proto_rawDescData
}

var file_pkg_movie_pb_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_movie_pb_movie_proto_goTypes = []interface{}{
	(*Movie)(nil),                 // 0: movie.Movie
	(*CreateMovieRequest)(nil),    // 1: movie.CreateMovieRequest
	(*MovieByUserRequest)(nil),    // 2: movie.MovieByUserRequest
	(*StatusResponse)(nil),        // 3: movie.StatusResponse
	(*FindMovieRequest)(nil),      // 4: movie.FindMovieRequest
	(*MovieResponse)(nil),         // 5: movie.MovieResponse
	(*MoviesListRequest)(nil),     // 6: movie.MoviesListRequest
	(*MoviesListResponse)(nil),    // 7: movie.MoviesListResponse
	(*UpdateMovieRequest)(nil),    // 8: movie.UpdateMovieRequest
	(*UpdateCoverRequest)(nil),    // 9: movie.UpdateCoverRequest
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_pkg_movie_pb_movie_proto_depIdxs = []int32{
	10, // 0: movie.Movie.date:type_name -> google.protobuf.Timestamp
	10, // 1: movie.CreateMovieRequest.date:type_name -> google.protobuf.Timestamp
	0,  // 2: movie.MovieResponse.movie:type_name -> movie.Movie
	0,  // 3: movie.MoviesListResponse.movies:type_name -> movie.Movie
	10, // 4: movie.UpdateMovieRequest.date:type_name -> google.protobuf.Timestamp
	1,  // 5: movie.MovieService.CreateMovie:input_type -> movie.CreateMovieRequest
	2,  // 6: movie.MovieService.DeleteMovie:input_type -> movie.MovieByUserRequest
	4,  // 7: movie.MovieService.FindMovie:input_type -> movie.FindMovieRequest
	6,  // 8: movie.MovieService.MoviesList:input_type -> movie.MoviesListRequest
	8,  // 9: movie.MovieService.UpdateMovie:input_type -> movie.UpdateMovieRequest
	9,  // 10: movie.MovieService.UpdateCover:input_type -> movie.UpdateCoverRequest
	3,  // 11: movie.MovieService.CreateMovie:output_type -> movie.StatusResponse
	3,  // 12: movie.MovieService.DeleteMovie:output_type -> movie.StatusResponse
	5,  // 13: movie.MovieService.FindMovie:output_type -> movie.MovieResponse
	7,  // 14: movie.MovieService.MoviesList:output_type -> movie.MoviesListResponse
	3,  // 15: movie.MovieService.UpdateMovie:output_type -> movie.StatusResponse
	3,  // 16: movie.MovieService.UpdateCover:output_type -> movie.StatusResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_movie_pb_movie_proto_init() }
func file_pkg_movie_pb_movie_proto_init() {
	if File_pkg_movie_pb_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_movie_pb_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_movie_pb_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMovieRequest); i {
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
		file_pkg_movie_pb_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieByUserRequest); i {
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
		file_pkg_movie_pb_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_movie_pb_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMovieRequest); i {
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
		file_pkg_movie_pb_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_movie_pb_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoviesListRequest); i {
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
		file_pkg_movie_pb_movie_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoviesListResponse); i {
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
		file_pkg_movie_pb_movie_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMovieRequest); i {
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
		file_pkg_movie_pb_movie_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoverRequest); i {
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
			RawDescriptor: file_pkg_movie_pb_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_movie_pb_movie_proto_goTypes,
		DependencyIndexes: file_pkg_movie_pb_movie_proto_depIdxs,
		MessageInfos:      file_pkg_movie_pb_movie_proto_msgTypes,
	}.Build()
	File_pkg_movie_pb_movie_proto = out.File
	file_pkg_movie_pb_movie_proto_rawDesc = nil
	file_pkg_movie_pb_movie_proto_goTypes = nil
	file_pkg_movie_pb_movie_proto_depIdxs = nil
}
