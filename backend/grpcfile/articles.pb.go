// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: articles.proto

package grpcfile

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

type HTMLClasses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleClass     string `protobuf:"bytes,1,opt,name=article_class,json=articleClass,proto3" json:"article_class,omitempty"`
	TitleClass       string `protobuf:"bytes,2,opt,name=title_class,json=titleClass,proto3" json:"title_class,omitempty"`
	DescriptionClass string `protobuf:"bytes,3,opt,name=description_class,json=descriptionClass,proto3" json:"description_class,omitempty"`
	ThumbnailClass   string `protobuf:"bytes,4,opt,name=thumbnail_class,json=thumbnailClass,proto3" json:"thumbnail_class,omitempty"`
	LinkClass        string `protobuf:"bytes,5,opt,name=link_class,json=linkClass,proto3" json:"link_class,omitempty"`
}

func (x *HTMLClasses) Reset() {
	*x = HTMLClasses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTMLClasses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTMLClasses) ProtoMessage() {}

func (x *HTMLClasses) ProtoReflect() protoreflect.Message {
	mi := &file_articles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTMLClasses.ProtoReflect.Descriptor instead.
func (*HTMLClasses) Descriptor() ([]byte, []int) {
	return file_articles_proto_rawDescGZIP(), []int{0}
}

func (x *HTMLClasses) GetArticleClass() string {
	if x != nil {
		return x.ArticleClass
	}
	return ""
}

func (x *HTMLClasses) GetTitleClass() string {
	if x != nil {
		return x.TitleClass
	}
	return ""
}

func (x *HTMLClasses) GetDescriptionClass() string {
	if x != nil {
		return x.DescriptionClass
	}
	return ""
}

func (x *HTMLClasses) GetThumbnailClass() string {
	if x != nil {
		return x.ThumbnailClass
	}
	return ""
}

func (x *HTMLClasses) GetLinkClass() string {
	if x != nil {
		return x.LinkClass
	}
	return ""
}

// The request message containing keywords and html class
type AllConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword     string       `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	HtmlClasses *HTMLClasses `protobuf:"bytes,2,opt,name=html_classes,json=htmlClasses,proto3" json:"html_classes,omitempty"`
}

func (x *AllConfigs) Reset() {
	*x = AllConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllConfigs) ProtoMessage() {}

func (x *AllConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_articles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllConfigs.ProtoReflect.Descriptor instead.
func (*AllConfigs) Descriptor() ([]byte, []int) {
	return file_articles_proto_rawDescGZIP(), []int{1}
}

func (x *AllConfigs) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *AllConfigs) GetHtmlClasses() *HTMLClasses {
	if x != nil {
		return x.HtmlClasses
	}
	return nil
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Thumbnail   string `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Link        string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_articles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_articles_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Article) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Article) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

// The response message containing the articles
type ArticlesReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *ArticlesReponse) Reset() {
	*x = ArticlesReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlesReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlesReponse) ProtoMessage() {}

func (x *ArticlesReponse) ProtoReflect() protoreflect.Message {
	mi := &file_articles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlesReponse.ProtoReflect.Descriptor instead.
func (*ArticlesReponse) Descriptor() ([]byte, []int) {
	return file_articles_proto_rawDescGZIP(), []int{3}
}

func (x *ArticlesReponse) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

var File_articles_proto protoreflect.FileDescriptor

var file_articles_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x72, 0x70, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x0b, 0x48,
	0x54, 0x4d, 0x4c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x6b,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x38, 0x0a,
	0x0c, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x48,
	0x54, 0x4d, 0x4c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x0b, 0x68, 0x74, 0x6d, 0x6c,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x22, 0x73, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x40, 0x0a, 0x0f,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x32, 0x54,
	0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12,
	0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x69, 0x6c, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x68, 0x61, 0x6e, 0x2f,
	0x46, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x2d, 0x6e, 0x65, 0x77, 0x73, 0x2d, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_articles_proto_rawDescOnce sync.Once
	file_articles_proto_rawDescData = file_articles_proto_rawDesc
)

func file_articles_proto_rawDescGZIP() []byte {
	file_articles_proto_rawDescOnce.Do(func() {
		file_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_articles_proto_rawDescData)
	})
	return file_articles_proto_rawDescData
}

var file_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_articles_proto_goTypes = []interface{}{
	(*HTMLClasses)(nil),     // 0: grpcfile.HTMLClasses
	(*AllConfigs)(nil),      // 1: grpcfile.AllConfigs
	(*Article)(nil),         // 2: grpcfile.Article
	(*ArticlesReponse)(nil), // 3: grpcfile.ArticlesReponse
}
var file_articles_proto_depIdxs = []int32{
	0, // 0: grpcfile.AllConfigs.html_classes:type_name -> grpcfile.HTMLClasses
	2, // 1: grpcfile.ArticlesReponse.articles:type_name -> grpcfile.Article
	1, // 2: grpcfile.ArticleService.GetArticles:input_type -> grpcfile.AllConfigs
	3, // 3: grpcfile.ArticleService.GetArticles:output_type -> grpcfile.ArticlesReponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_articles_proto_init() }
func file_articles_proto_init() {
	if File_articles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_articles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTMLClasses); i {
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
		file_articles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllConfigs); i {
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
		file_articles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_articles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlesReponse); i {
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
			RawDescriptor: file_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_articles_proto_goTypes,
		DependencyIndexes: file_articles_proto_depIdxs,
		MessageInfos:      file_articles_proto_msgTypes,
	}.Build()
	File_articles_proto = out.File
	file_articles_proto_rawDesc = nil
	file_articles_proto_goTypes = nil
	file_articles_proto_depIdxs = nil
}
