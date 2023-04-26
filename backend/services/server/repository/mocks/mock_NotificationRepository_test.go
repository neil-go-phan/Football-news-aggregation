// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockrepository

import mock "github.com/stretchr/testify/mock"

// MockNotificationRepository is an autogenerated mock type for the NotificationRepository type
type MockNotificationRepository struct {
	mock.Mock
}

type MockNotificationRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNotificationRepository) EXPECT() *MockNotificationRepository_Expecter {
	return &MockNotificationRepository_Expecter{mock: &_m.Mock}
}

// Send provides a mock function with given fields: title, notiType, message
func (_m *MockNotificationRepository) Send(title string, notiType string, message string) {
	_m.Called(title, notiType, message)
}

// MockNotificationRepository_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockNotificationRepository_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - title string
//   - notiType string
//   - message string
func (_e *MockNotificationRepository_Expecter) Send(title interface{}, notiType interface{}, message interface{}) *MockNotificationRepository_Send_Call {
	return &MockNotificationRepository_Send_Call{Call: _e.mock.On("Send", title, notiType, message)}
}

func (_c *MockNotificationRepository_Send_Call) Run(run func(title string, notiType string, message string)) *MockNotificationRepository_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockNotificationRepository_Send_Call) Return() *MockNotificationRepository_Send_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNotificationRepository_Send_Call) RunAndReturn(run func(string, string, string)) *MockNotificationRepository_Send_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockNotificationRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockNotificationRepository creates a new instance of MockNotificationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockNotificationRepository(t mockConstructorTestingTNewMockNotificationRepository) *MockNotificationRepository {
	mock := &MockNotificationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
