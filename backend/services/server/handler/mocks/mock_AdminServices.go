// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	adminservices "server/services/admin"

	mock "github.com/stretchr/testify/mock"
)

// MockAdminServices is an autogenerated mock type for the AdminServices type
type MockAdminServices struct {
	mock.Mock
}

type MockAdminServices_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAdminServices) EXPECT() *MockAdminServices_Expecter {
	return &MockAdminServices_Expecter{mock: &_m.Mock}
}

// ChangePassword provides a mock function with given fields: admin, usernameToken
func (_m *MockAdminServices) ChangePassword(admin *adminservices.AdminWithConfirmPassword, usernameToken string) error {
	ret := _m.Called(admin, usernameToken)

	var r0 error
	if rf, ok := ret.Get(0).(func(*adminservices.AdminWithConfirmPassword, string) error); ok {
		r0 = rf(admin, usernameToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdminServices_ChangePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangePassword'
type MockAdminServices_ChangePassword_Call struct {
	*mock.Call
}

// ChangePassword is a helper method to define mock.On call
//   - admin *adminservices.AdminWithConfirmPassword
//   - usernameToken string
func (_e *MockAdminServices_Expecter) ChangePassword(admin interface{}, usernameToken interface{}) *MockAdminServices_ChangePassword_Call {
	return &MockAdminServices_ChangePassword_Call{Call: _e.mock.On("ChangePassword", admin, usernameToken)}
}

func (_c *MockAdminServices_ChangePassword_Call) Run(run func(admin *adminservices.AdminWithConfirmPassword, usernameToken string)) *MockAdminServices_ChangePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*adminservices.AdminWithConfirmPassword), args[1].(string))
	})
	return _c
}

func (_c *MockAdminServices_ChangePassword_Call) Return(_a0 error) *MockAdminServices_ChangePassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAdminServices_ChangePassword_Call) RunAndReturn(run func(*adminservices.AdminWithConfirmPassword, string) error) *MockAdminServices_ChangePassword_Call {
	_c.Call.Return(run)
	return _c
}

// CheckAdminUsernameToken provides a mock function with given fields: username
func (_m *MockAdminServices) CheckAdminUsernameToken(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdminServices_CheckAdminUsernameToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckAdminUsernameToken'
type MockAdminServices_CheckAdminUsernameToken_Call struct {
	*mock.Call
}

// CheckAdminUsernameToken is a helper method to define mock.On call
//   - username string
func (_e *MockAdminServices_Expecter) CheckAdminUsernameToken(username interface{}) *MockAdminServices_CheckAdminUsernameToken_Call {
	return &MockAdminServices_CheckAdminUsernameToken_Call{Call: _e.mock.On("CheckAdminUsernameToken", username)}
}

func (_c *MockAdminServices_CheckAdminUsernameToken_Call) Run(run func(username string)) *MockAdminServices_CheckAdminUsernameToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAdminServices_CheckAdminUsernameToken_Call) Return(_a0 error) *MockAdminServices_CheckAdminUsernameToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAdminServices_CheckAdminUsernameToken_Call) RunAndReturn(run func(string) error) *MockAdminServices_CheckAdminUsernameToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetAdminUsername provides a mock function with given fields: username
func (_m *MockAdminServices) GetAdminUsername(username string) (string, error) {
	ret := _m.Called(username)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdminServices_GetAdminUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAdminUsername'
type MockAdminServices_GetAdminUsername_Call struct {
	*mock.Call
}

// GetAdminUsername is a helper method to define mock.On call
//   - username string
func (_e *MockAdminServices_Expecter) GetAdminUsername(username interface{}) *MockAdminServices_GetAdminUsername_Call {
	return &MockAdminServices_GetAdminUsername_Call{Call: _e.mock.On("GetAdminUsername", username)}
}

func (_c *MockAdminServices_GetAdminUsername_Call) Run(run func(username string)) *MockAdminServices_GetAdminUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAdminServices_GetAdminUsername_Call) Return(_a0 string, _a1 error) *MockAdminServices_GetAdminUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdminServices_GetAdminUsername_Call) RunAndReturn(run func(string) (string, error)) *MockAdminServices_GetAdminUsername_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: admin
func (_m *MockAdminServices) Login(admin *adminservices.Admin) (string, error) {
	ret := _m.Called(admin)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*adminservices.Admin) (string, error)); ok {
		return rf(admin)
	}
	if rf, ok := ret.Get(0).(func(*adminservices.Admin) string); ok {
		r0 = rf(admin)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*adminservices.Admin) error); ok {
		r1 = rf(admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdminServices_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockAdminServices_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - admin *adminservices.Admin
func (_e *MockAdminServices_Expecter) Login(admin interface{}) *MockAdminServices_Login_Call {
	return &MockAdminServices_Login_Call{Call: _e.mock.On("Login", admin)}
}

func (_c *MockAdminServices_Login_Call) Run(run func(admin *adminservices.Admin)) *MockAdminServices_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*adminservices.Admin))
	})
	return _c
}

func (_c *MockAdminServices_Login_Call) Return(_a0 string, _a1 error) *MockAdminServices_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdminServices_Login_Call) RunAndReturn(run func(*adminservices.Admin) (string, error)) *MockAdminServices_Login_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAdminServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAdminServices creates a new instance of MockAdminServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAdminServices(t mockConstructorTestingTNewMockAdminServices) *MockAdminServices {
	mock := &MockAdminServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
