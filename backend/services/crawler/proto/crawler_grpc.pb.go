// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: crawler.proto

package crawlerproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CrawlerService_GetArticles_FullMethodName       = "/crawlerproto.CrawlerService/GetArticles"
	CrawlerService_GetSchedulesOnDay_FullMethodName = "/crawlerproto.CrawlerService/GetSchedulesOnDay"
	CrawlerService_GetMatchDetail_FullMethodName    = "/crawlerproto.CrawlerService/GetMatchDetail"
)

// CrawlerServiceClient is the client API for CrawlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlerServiceClient interface {
	GetArticles(ctx context.Context, in *KeywordToSearch, opts ...grpc.CallOption) (CrawlerService_GetArticlesClient, error)
	GetSchedulesOnDay(ctx context.Context, in *Date, opts ...grpc.CallOption) (*SchedulesReponse, error)
	GetMatchDetail(ctx context.Context, in *MatchURLs, opts ...grpc.CallOption) (CrawlerService_GetMatchDetailClient, error)
}

type crawlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlerServiceClient(cc grpc.ClientConnInterface) CrawlerServiceClient {
	return &crawlerServiceClient{cc}
}

func (c *crawlerServiceClient) GetArticles(ctx context.Context, in *KeywordToSearch, opts ...grpc.CallOption) (CrawlerService_GetArticlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrawlerService_ServiceDesc.Streams[0], CrawlerService_GetArticles_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &crawlerServiceGetArticlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrawlerService_GetArticlesClient interface {
	Recv() (*ArticlesReponse, error)
	grpc.ClientStream
}

type crawlerServiceGetArticlesClient struct {
	grpc.ClientStream
}

func (x *crawlerServiceGetArticlesClient) Recv() (*ArticlesReponse, error) {
	m := new(ArticlesReponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crawlerServiceClient) GetSchedulesOnDay(ctx context.Context, in *Date, opts ...grpc.CallOption) (*SchedulesReponse, error) {
	out := new(SchedulesReponse)
	err := c.cc.Invoke(ctx, CrawlerService_GetSchedulesOnDay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerServiceClient) GetMatchDetail(ctx context.Context, in *MatchURLs, opts ...grpc.CallOption) (CrawlerService_GetMatchDetailClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrawlerService_ServiceDesc.Streams[1], CrawlerService_GetMatchDetail_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &crawlerServiceGetMatchDetailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrawlerService_GetMatchDetailClient interface {
	Recv() (*MatchDetail, error)
	grpc.ClientStream
}

type crawlerServiceGetMatchDetailClient struct {
	grpc.ClientStream
}

func (x *crawlerServiceGetMatchDetailClient) Recv() (*MatchDetail, error) {
	m := new(MatchDetail)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CrawlerServiceServer is the server API for CrawlerService service.
// All implementations must embed UnimplementedCrawlerServiceServer
// for forward compatibility
type CrawlerServiceServer interface {
	GetArticles(*KeywordToSearch, CrawlerService_GetArticlesServer) error
	GetSchedulesOnDay(context.Context, *Date) (*SchedulesReponse, error)
	GetMatchDetail(*MatchURLs, CrawlerService_GetMatchDetailServer) error
	mustEmbedUnimplementedCrawlerServiceServer()
}

// UnimplementedCrawlerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrawlerServiceServer struct {
}

func (UnimplementedCrawlerServiceServer) GetArticles(*KeywordToSearch, CrawlerService_GetArticlesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedCrawlerServiceServer) GetSchedulesOnDay(context.Context, *Date) (*SchedulesReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedulesOnDay not implemented")
}
func (UnimplementedCrawlerServiceServer) GetMatchDetail(*MatchURLs, CrawlerService_GetMatchDetailServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMatchDetail not implemented")
}
func (UnimplementedCrawlerServiceServer) mustEmbedUnimplementedCrawlerServiceServer() {}

// UnsafeCrawlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlerServiceServer will
// result in compilation errors.
type UnsafeCrawlerServiceServer interface {
	mustEmbedUnimplementedCrawlerServiceServer()
}

func RegisterCrawlerServiceServer(s grpc.ServiceRegistrar, srv CrawlerServiceServer) {
	s.RegisterService(&CrawlerService_ServiceDesc, srv)
}

func _CrawlerService_GetArticles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeywordToSearch)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlerServiceServer).GetArticles(m, &crawlerServiceGetArticlesServer{stream})
}

type CrawlerService_GetArticlesServer interface {
	Send(*ArticlesReponse) error
	grpc.ServerStream
}

type crawlerServiceGetArticlesServer struct {
	grpc.ServerStream
}

func (x *crawlerServiceGetArticlesServer) Send(m *ArticlesReponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrawlerService_GetSchedulesOnDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).GetSchedulesOnDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlerService_GetSchedulesOnDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).GetSchedulesOnDay(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlerService_GetMatchDetail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchURLs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlerServiceServer).GetMatchDetail(m, &crawlerServiceGetMatchDetailServer{stream})
}

type CrawlerService_GetMatchDetailServer interface {
	Send(*MatchDetail) error
	grpc.ServerStream
}

type crawlerServiceGetMatchDetailServer struct {
	grpc.ServerStream
}

func (x *crawlerServiceGetMatchDetailServer) Send(m *MatchDetail) error {
	return x.ServerStream.SendMsg(m)
}

// CrawlerService_ServiceDesc is the grpc.ServiceDesc for CrawlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crawlerproto.CrawlerService",
	HandlerType: (*CrawlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchedulesOnDay",
			Handler:    _CrawlerService_GetSchedulesOnDay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArticles",
			Handler:       _CrawlerService_GetArticles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMatchDetail",
			Handler:       _CrawlerService_GetMatchDetail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crawler.proto",
}
