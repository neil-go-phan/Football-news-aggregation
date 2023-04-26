// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockrepository

import (
	entities "server/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockHtmlClassesRepository is an autogenerated mock type for the HtmlClassesRepository type
type MockHtmlClassesRepository struct {
	mock.Mock
}

type MockHtmlClassesRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHtmlClassesRepository) EXPECT() *MockHtmlClassesRepository_Expecter {
	return &MockHtmlClassesRepository_Expecter{mock: &_m.Mock}
}

// GetHtmlClasses provides a mock function with given fields:
func (_m *MockHtmlClassesRepository) GetHtmlClasses() entities.HtmlClasses {
	ret := _m.Called()

	var r0 entities.HtmlClasses
	if rf, ok := ret.Get(0).(func() entities.HtmlClasses); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entities.HtmlClasses)
	}

	return r0
}

// MockHtmlClassesRepository_GetHtmlClasses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHtmlClasses'
type MockHtmlClassesRepository_GetHtmlClasses_Call struct {
	*mock.Call
}

// GetHtmlClasses is a helper method to define mock.On call
func (_e *MockHtmlClassesRepository_Expecter) GetHtmlClasses() *MockHtmlClassesRepository_GetHtmlClasses_Call {
	return &MockHtmlClassesRepository_GetHtmlClasses_Call{Call: _e.mock.On("GetHtmlClasses")}
}

func (_c *MockHtmlClassesRepository_GetHtmlClasses_Call) Run(run func()) *MockHtmlClassesRepository_GetHtmlClasses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHtmlClassesRepository_GetHtmlClasses_Call) Return(_a0 entities.HtmlClasses) *MockHtmlClassesRepository_GetHtmlClasses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHtmlClassesRepository_GetHtmlClasses_Call) RunAndReturn(run func() entities.HtmlClasses) *MockHtmlClassesRepository_GetHtmlClasses_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockHtmlClassesRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockHtmlClassesRepository creates a new instance of MockHtmlClassesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockHtmlClassesRepository(t mockConstructorTestingTNewMockHtmlClassesRepository) *MockHtmlClassesRepository {
	mock := &MockHtmlClassesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}