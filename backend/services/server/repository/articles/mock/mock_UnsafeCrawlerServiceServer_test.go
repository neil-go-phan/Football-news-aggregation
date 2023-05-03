// Code generated by mockery v2.20.0. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockUnsafeCrawlerServiceServer is an autogenerated mock type for the UnsafeCrawlerServiceServer type
type MockUnsafeCrawlerServiceServer struct {
	mock.Mock
}

type MockUnsafeCrawlerServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeCrawlerServiceServer) EXPECT() *MockUnsafeCrawlerServiceServer_Expecter {
	return &MockUnsafeCrawlerServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedCrawlerServiceServer provides a mock function with given fields:
func (_m *MockUnsafeCrawlerServiceServer) mustEmbedUnimplementedCrawlerServiceServer() {
	_m.Called()
}

// MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedCrawlerServiceServer'
type MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedCrawlerServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeCrawlerServiceServer_Expecter) mustEmbedUnimplementedCrawlerServiceServer() *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call {
	return &MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedCrawlerServiceServer")}
}

func (_c *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call) Run(run func()) *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call) Return() *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call) RunAndReturn(run func()) *MockUnsafeCrawlerServiceServer_mustEmbedUnimplementedCrawlerServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockUnsafeCrawlerServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUnsafeCrawlerServiceServer creates a new instance of MockUnsafeCrawlerServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeCrawlerServiceServer(t mockConstructorTestingTNewMockUnsafeCrawlerServiceServer) *MockUnsafeCrawlerServiceServer {
	mock := &MockUnsafeCrawlerServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
