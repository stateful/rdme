// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: runme/parser/v1/parser.proto

package parserv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CellKind int32

const (
	CellKind_CELL_KIND_UNSPECIFIED CellKind = 0
	CellKind_CELL_KIND_MARKUP      CellKind = 1
	CellKind_CELL_KIND_CODE        CellKind = 2
)

// Enum value maps for CellKind.
var (
	CellKind_name = map[int32]string{
		0: "CELL_KIND_UNSPECIFIED",
		1: "CELL_KIND_MARKUP",
		2: "CELL_KIND_CODE",
	}
	CellKind_value = map[string]int32{
		"CELL_KIND_UNSPECIFIED": 0,
		"CELL_KIND_MARKUP":      1,
		"CELL_KIND_CODE":        2,
	}
)

func (x CellKind) Enum() *CellKind {
	p := new(CellKind)
	*p = x
	return p
}

func (x CellKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CellKind) Descriptor() protoreflect.EnumDescriptor {
	return file_runme_parser_v1_parser_proto_enumTypes[0].Descriptor()
}

func (CellKind) Type() protoreflect.EnumType {
	return &file_runme_parser_v1_parser_proto_enumTypes[0]
}

func (x CellKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CellKind.Descriptor instead.
func (CellKind) EnumDescriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{0}
}

type RunmeIdentity int32

const (
	RunmeIdentity_RUNME_IDENTITY_UNSPECIFIED RunmeIdentity = 0 // aka NONE
	RunmeIdentity_RUNME_IDENTITY_ALL         RunmeIdentity = 1
	RunmeIdentity_RUNME_IDENTITY_DOCUMENT    RunmeIdentity = 2
	RunmeIdentity_RUNME_IDENTITY_CELL        RunmeIdentity = 3
)

// Enum value maps for RunmeIdentity.
var (
	RunmeIdentity_name = map[int32]string{
		0: "RUNME_IDENTITY_UNSPECIFIED",
		1: "RUNME_IDENTITY_ALL",
		2: "RUNME_IDENTITY_DOCUMENT",
		3: "RUNME_IDENTITY_CELL",
	}
	RunmeIdentity_value = map[string]int32{
		"RUNME_IDENTITY_UNSPECIFIED": 0,
		"RUNME_IDENTITY_ALL":         1,
		"RUNME_IDENTITY_DOCUMENT":    2,
		"RUNME_IDENTITY_CELL":        3,
	}
)

func (x RunmeIdentity) Enum() *RunmeIdentity {
	p := new(RunmeIdentity)
	*p = x
	return p
}

func (x RunmeIdentity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunmeIdentity) Descriptor() protoreflect.EnumDescriptor {
	return file_runme_parser_v1_parser_proto_enumTypes[1].Descriptor()
}

func (RunmeIdentity) Type() protoreflect.EnumType {
	return &file_runme_parser_v1_parser_proto_enumTypes[1]
}

func (x RunmeIdentity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunmeIdentity.Descriptor instead.
func (RunmeIdentity) EnumDescriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{1}
}

type Runme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Runme) Reset() {
	*x = Runme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Runme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runme) ProtoMessage() {}

func (x *Runme) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runme.ProtoReflect.Descriptor instead.
func (*Runme) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{0}
}

func (x *Runme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Runme) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Notebook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells       []*Cell           `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Frontmatter *Frontmatter      `protobuf:"bytes,3,opt,name=frontmatter,proto3" json:"frontmatter,omitempty"`
}

func (x *Notebook) Reset() {
	*x = Notebook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notebook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notebook) ProtoMessage() {}

func (x *Notebook) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notebook.ProtoReflect.Descriptor instead.
func (*Notebook) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{1}
}

func (x *Notebook) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

func (x *Notebook) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Notebook) GetFrontmatter() *Frontmatter {
	if x != nil {
		return x.Frontmatter
	}
	return nil
}

type TextRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   uint32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TextRange) Reset() {
	*x = TextRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextRange) ProtoMessage() {}

func (x *TextRange) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextRange.ProtoReflect.Descriptor instead.
func (*TextRange) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{2}
}

func (x *TextRange) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TextRange) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

type Cell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind       CellKind          `protobuf:"varint,1,opt,name=kind,proto3,enum=runme.parser.v1.CellKind" json:"kind,omitempty"`
	Value      string            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	LanguageId string            `protobuf:"bytes,3,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
	Metadata   map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TextRange  *TextRange        `protobuf:"bytes,5,opt,name=text_range,json=textRange,proto3" json:"text_range,omitempty"`
}

func (x *Cell) Reset() {
	*x = Cell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{3}
}

func (x *Cell) GetKind() CellKind {
	if x != nil {
		return x.Kind
	}
	return CellKind_CELL_KIND_UNSPECIFIED
}

func (x *Cell) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Cell) GetLanguageId() string {
	if x != nil {
		return x.LanguageId
	}
	return ""
}

func (x *Cell) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Cell) GetTextRange() *TextRange {
	if x != nil {
		return x.TextRange
	}
	return nil
}

type Frontmatter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shell       string `protobuf:"bytes,1,opt,name=shell,proto3" json:"shell,omitempty"`
	Cwd         string `protobuf:"bytes,2,opt,name=cwd,proto3" json:"cwd,omitempty"`
	SkipPrompts bool   `protobuf:"varint,3,opt,name=skip_prompts,json=skipPrompts,proto3" json:"skip_prompts,omitempty"`
	Runme       *Runme `protobuf:"bytes,4,opt,name=runme,proto3" json:"runme,omitempty"`
}

func (x *Frontmatter) Reset() {
	*x = Frontmatter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frontmatter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frontmatter) ProtoMessage() {}

func (x *Frontmatter) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frontmatter.ProtoReflect.Descriptor instead.
func (*Frontmatter) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{4}
}

func (x *Frontmatter) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

func (x *Frontmatter) GetCwd() string {
	if x != nil {
		return x.Cwd
	}
	return ""
}

func (x *Frontmatter) GetSkipPrompts() bool {
	if x != nil {
		return x.SkipPrompts
	}
	return false
}

func (x *Frontmatter) GetRunme() *Runme {
	if x != nil {
		return x.Runme
	}
	return nil
}

type DeserializeRequestOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity RunmeIdentity `protobuf:"varint,1,opt,name=identity,proto3,enum=runme.parser.v1.RunmeIdentity" json:"identity,omitempty"`
}

func (x *DeserializeRequestOptions) Reset() {
	*x = DeserializeRequestOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeserializeRequestOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeserializeRequestOptions) ProtoMessage() {}

func (x *DeserializeRequestOptions) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeserializeRequestOptions.ProtoReflect.Descriptor instead.
func (*DeserializeRequestOptions) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{5}
}

func (x *DeserializeRequestOptions) GetIdentity() RunmeIdentity {
	if x != nil {
		return x.Identity
	}
	return RunmeIdentity_RUNME_IDENTITY_UNSPECIFIED
}

type DeserializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  []byte                     `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Options *DeserializeRequestOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *DeserializeRequest) Reset() {
	*x = DeserializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeserializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeserializeRequest) ProtoMessage() {}

func (x *DeserializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeserializeRequest.ProtoReflect.Descriptor instead.
func (*DeserializeRequest) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{6}
}

func (x *DeserializeRequest) GetSource() []byte {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *DeserializeRequest) GetOptions() *DeserializeRequestOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type DeserializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notebook *Notebook `protobuf:"bytes,1,opt,name=notebook,proto3" json:"notebook,omitempty"`
}

func (x *DeserializeResponse) Reset() {
	*x = DeserializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeserializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeserializeResponse) ProtoMessage() {}

func (x *DeserializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeserializeResponse.ProtoReflect.Descriptor instead.
func (*DeserializeResponse) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{7}
}

func (x *DeserializeResponse) GetNotebook() *Notebook {
	if x != nil {
		return x.Notebook
	}
	return nil
}

type SerializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notebook *Notebook `protobuf:"bytes,1,opt,name=notebook,proto3" json:"notebook,omitempty"`
}

func (x *SerializeRequest) Reset() {
	*x = SerializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializeRequest) ProtoMessage() {}

func (x *SerializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializeRequest.ProtoReflect.Descriptor instead.
func (*SerializeRequest) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{8}
}

func (x *SerializeRequest) GetNotebook() *Notebook {
	if x != nil {
		return x.Notebook
	}
	return nil
}

type SerializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SerializeResponse) Reset() {
	*x = SerializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_parser_v1_parser_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializeResponse) ProtoMessage() {}

func (x *SerializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runme_parser_v1_parser_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializeResponse.ProtoReflect.Descriptor instead.
func (*SerializeResponse) Descriptor() ([]byte, []int) {
	return file_runme_parser_v1_parser_proto_rawDescGZIP(), []int{9}
}

func (x *SerializeResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_runme_parser_v1_parser_proto protoreflect.FileDescriptor

var file_runme_parser_v1_parser_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22,
	0x31, 0x0a, 0x05, 0x52, 0x75, 0x6e, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xf9, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12,
	0x2b, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x3e, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x6d, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x33,
	0x0a, 0x09, 0x54, 0x65, 0x78, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0xa5, 0x02, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x2d, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x75, 0x6e,
	0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6c,
	0x6c, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x86, 0x01, 0x0a, 0x0b,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x68, 0x65, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x77, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x6d, 0x65, 0x52, 0x05, 0x72,
	0x75, 0x6e, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x6d, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x72, 0x0a,
	0x12, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72,
	0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x4c, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x75, 0x6e,
	0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x22,
	0x49, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x4f, 0x0a, 0x08, 0x43, 0x65, 0x6c, 0x6c, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4d, 0x41, 0x52, 0x4b,
	0x55, 0x50, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x2a, 0x7d, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x6d,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x55, 0x4e,
	0x4d, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e,
	0x4d, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x55, 0x4e, 0x4d, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x54, 0x59, 0x5f, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x52, 0x55, 0x4e, 0x4d, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59,
	0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x10, 0x03, 0x32, 0xc1, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65,
	0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x12, 0x21, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x72,
	0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runme_parser_v1_parser_proto_rawDescOnce sync.Once
	file_runme_parser_v1_parser_proto_rawDescData = file_runme_parser_v1_parser_proto_rawDesc
)

func file_runme_parser_v1_parser_proto_rawDescGZIP() []byte {
	file_runme_parser_v1_parser_proto_rawDescOnce.Do(func() {
		file_runme_parser_v1_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_runme_parser_v1_parser_proto_rawDescData)
	})
	return file_runme_parser_v1_parser_proto_rawDescData
}

var file_runme_parser_v1_parser_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_runme_parser_v1_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_runme_parser_v1_parser_proto_goTypes = []interface{}{
	(CellKind)(0),                     // 0: runme.parser.v1.CellKind
	(RunmeIdentity)(0),                // 1: runme.parser.v1.RunmeIdentity
	(*Runme)(nil),                     // 2: runme.parser.v1.Runme
	(*Notebook)(nil),                  // 3: runme.parser.v1.Notebook
	(*TextRange)(nil),                 // 4: runme.parser.v1.TextRange
	(*Cell)(nil),                      // 5: runme.parser.v1.Cell
	(*Frontmatter)(nil),               // 6: runme.parser.v1.Frontmatter
	(*DeserializeRequestOptions)(nil), // 7: runme.parser.v1.DeserializeRequestOptions
	(*DeserializeRequest)(nil),        // 8: runme.parser.v1.DeserializeRequest
	(*DeserializeResponse)(nil),       // 9: runme.parser.v1.DeserializeResponse
	(*SerializeRequest)(nil),          // 10: runme.parser.v1.SerializeRequest
	(*SerializeResponse)(nil),         // 11: runme.parser.v1.SerializeResponse
	nil,                               // 12: runme.parser.v1.Notebook.MetadataEntry
	nil,                               // 13: runme.parser.v1.Cell.MetadataEntry
}
var file_runme_parser_v1_parser_proto_depIdxs = []int32{
	5,  // 0: runme.parser.v1.Notebook.cells:type_name -> runme.parser.v1.Cell
	12, // 1: runme.parser.v1.Notebook.metadata:type_name -> runme.parser.v1.Notebook.MetadataEntry
	6,  // 2: runme.parser.v1.Notebook.frontmatter:type_name -> runme.parser.v1.Frontmatter
	0,  // 3: runme.parser.v1.Cell.kind:type_name -> runme.parser.v1.CellKind
	13, // 4: runme.parser.v1.Cell.metadata:type_name -> runme.parser.v1.Cell.MetadataEntry
	4,  // 5: runme.parser.v1.Cell.text_range:type_name -> runme.parser.v1.TextRange
	2,  // 6: runme.parser.v1.Frontmatter.runme:type_name -> runme.parser.v1.Runme
	1,  // 7: runme.parser.v1.DeserializeRequestOptions.identity:type_name -> runme.parser.v1.RunmeIdentity
	7,  // 8: runme.parser.v1.DeserializeRequest.options:type_name -> runme.parser.v1.DeserializeRequestOptions
	3,  // 9: runme.parser.v1.DeserializeResponse.notebook:type_name -> runme.parser.v1.Notebook
	3,  // 10: runme.parser.v1.SerializeRequest.notebook:type_name -> runme.parser.v1.Notebook
	8,  // 11: runme.parser.v1.ParserService.Deserialize:input_type -> runme.parser.v1.DeserializeRequest
	10, // 12: runme.parser.v1.ParserService.Serialize:input_type -> runme.parser.v1.SerializeRequest
	9,  // 13: runme.parser.v1.ParserService.Deserialize:output_type -> runme.parser.v1.DeserializeResponse
	11, // 14: runme.parser.v1.ParserService.Serialize:output_type -> runme.parser.v1.SerializeResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_runme_parser_v1_parser_proto_init() }
func file_runme_parser_v1_parser_proto_init() {
	if File_runme_parser_v1_parser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runme_parser_v1_parser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Runme); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notebook); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextRange); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cell); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frontmatter); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeserializeRequestOptions); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeserializeRequest); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeserializeResponse); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializeRequest); i {
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
		file_runme_parser_v1_parser_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializeResponse); i {
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
			RawDescriptor: file_runme_parser_v1_parser_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runme_parser_v1_parser_proto_goTypes,
		DependencyIndexes: file_runme_parser_v1_parser_proto_depIdxs,
		EnumInfos:         file_runme_parser_v1_parser_proto_enumTypes,
		MessageInfos:      file_runme_parser_v1_parser_proto_msgTypes,
	}.Build()
	File_runme_parser_v1_parser_proto = out.File
	file_runme_parser_v1_parser_proto_rawDesc = nil
	file_runme_parser_v1_parser_proto_goTypes = nil
	file_runme_parser_v1_parser_proto_depIdxs = nil
}
