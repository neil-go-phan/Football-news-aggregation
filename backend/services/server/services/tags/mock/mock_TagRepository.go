// Code generated by mockery v2.20.0. DO NOT EDIT.

package repository

import (
	entities "server/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockTagRepository is an autogenerated mock type for the TagRepository type
type MockTagRepository struct {
	mock.Mock
}

type MockTagRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTagRepository) EXPECT() *MockTagRepository_Expecter {
	return &MockTagRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: tag
func (_m *MockTagRepository) Create(tag *entities.Tag) error {
	ret := _m.Called(tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Tag) error); ok {
		r0 = rf(tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTagRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - tag *entities.Tag
func (_e *MockTagRepository_Expecter) Create(tag interface{}) *MockTagRepository_Create_Call {
	return &MockTagRepository_Create_Call{Call: _e.mock.On("Create", tag)}
}

func (_c *MockTagRepository_Create_Call) Run(run func(tag *entities.Tag)) *MockTagRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Tag))
	})
	return _c
}

func (_c *MockTagRepository_Create_Call) Return(_a0 error) *MockTagRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagRepository_Create_Call) RunAndReturn(run func(*entities.Tag) error) *MockTagRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: tagName
func (_m *MockTagRepository) Delete(tagName string) error {
	ret := _m.Called(tagName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tagName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockTagRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - tagName string
func (_e *MockTagRepository_Expecter) Delete(tagName interface{}) *MockTagRepository_Delete_Call {
	return &MockTagRepository_Delete_Call{Call: _e.mock.On("Delete", tagName)}
}

func (_c *MockTagRepository_Delete_Call) Run(run func(tagName string)) *MockTagRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTagRepository_Delete_Call) Return(_a0 error) *MockTagRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagRepository_Delete_Call) RunAndReturn(run func(string) error) *MockTagRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: tagName
func (_m *MockTagRepository) Get(tagName string) (*entities.Tag, error) {
	ret := _m.Called(tagName)

	var r0 *entities.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Tag, error)); ok {
		return rf(tagName)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Tag); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tagName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockTagRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - tagName string
func (_e *MockTagRepository_Expecter) Get(tagName interface{}) *MockTagRepository_Get_Call {
	return &MockTagRepository_Get_Call{Call: _e.mock.On("Get", tagName)}
}

func (_c *MockTagRepository_Get_Call) Run(run func(tagName string)) *MockTagRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTagRepository_Get_Call) Return(_a0 *entities.Tag, _a1 error) *MockTagRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagRepository_Get_Call) RunAndReturn(run func(string) (*entities.Tag, error)) *MockTagRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagsByTagNames provides a mock function with given fields: tagNames
func (_m *MockTagRepository) GetTagsByTagNames(tagNames []string) (*[]entities.Tag, error) {
	ret := _m.Called(tagNames)

	var r0 *[]entities.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) (*[]entities.Tag, error)); ok {
		return rf(tagNames)
	}
	if rf, ok := ret.Get(0).(func([]string) *[]entities.Tag); ok {
		r0 = rf(tagNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(tagNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagRepository_GetTagsByTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagsByTagNames'
type MockTagRepository_GetTagsByTagNames_Call struct {
	*mock.Call
}

// GetTagsByTagNames is a helper method to define mock.On call
//   - tagNames []string
func (_e *MockTagRepository_Expecter) GetTagsByTagNames(tagNames interface{}) *MockTagRepository_GetTagsByTagNames_Call {
	return &MockTagRepository_GetTagsByTagNames_Call{Call: _e.mock.On("GetTagsByTagNames", tagNames)}
}

func (_c *MockTagRepository_GetTagsByTagNames_Call) Run(run func(tagNames []string)) *MockTagRepository_GetTagsByTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockTagRepository_GetTagsByTagNames_Call) Return(_a0 *[]entities.Tag, _a1 error) *MockTagRepository_GetTagsByTagNames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagRepository_GetTagsByTagNames_Call) RunAndReturn(run func([]string) (*[]entities.Tag, error)) *MockTagRepository_GetTagsByTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields:
func (_m *MockTagRepository) List() (*[]entities.Tag, error) {
	ret := _m.Called()

	var r0 *[]entities.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func() (*[]entities.Tag, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *[]entities.Tag); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockTagRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
func (_e *MockTagRepository_Expecter) List() *MockTagRepository_List_Call {
	return &MockTagRepository_List_Call{Call: _e.mock.On("List")}
}

func (_c *MockTagRepository_List_Call) Run(run func()) *MockTagRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTagRepository_List_Call) Return(_a0 *[]entities.Tag, _a1 error) *MockTagRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagRepository_List_Call) RunAndReturn(run func() (*[]entities.Tag, error)) *MockTagRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockTagRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTagRepository creates a new instance of MockTagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTagRepository(t mockConstructorTestingTNewMockTagRepository) *MockTagRepository {
	mock := &MockTagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
