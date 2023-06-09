// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "server/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockArticleRepository is an autogenerated mock type for the ArticleRepository type
type MockArticleRepository struct {
	mock.Mock
}

type MockArticleRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockArticleRepository) EXPECT() *MockArticleRepository_Expecter {
	return &MockArticleRepository_Expecter{mock: &_m.Mock}
}

// AddTag provides a mock function with given fields: ids, newTag
func (_m *MockArticleRepository) AddTag(ids []uint, newTag *entities.Tag) error {
	ret := _m.Called(ids, newTag)

	var r0 error
	if rf, ok := ret.Get(0).(func([]uint, *entities.Tag) error); ok {
		r0 = rf(ids, newTag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticleRepository_AddTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTag'
type MockArticleRepository_AddTag_Call struct {
	*mock.Call
}

// AddTag is a helper method to define mock.On call
//   - ids []uint
//   - newTag *entities.Tag
func (_e *MockArticleRepository_Expecter) AddTag(ids interface{}, newTag interface{}) *MockArticleRepository_AddTag_Call {
	return &MockArticleRepository_AddTag_Call{Call: _e.mock.On("AddTag", ids, newTag)}
}

func (_c *MockArticleRepository_AddTag_Call) Run(run func(ids []uint, newTag *entities.Tag)) *MockArticleRepository_AddTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]uint), args[1].(*entities.Tag))
	})
	return _c
}

func (_c *MockArticleRepository_AddTag_Call) Return(_a0 error) *MockArticleRepository_AddTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticleRepository_AddTag_Call) RunAndReturn(run func([]uint, *entities.Tag) error) *MockArticleRepository_AddTag_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MockArticleRepository) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticleRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockArticleRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *MockArticleRepository_Expecter) Delete(id interface{}) *MockArticleRepository_Delete_Call {
	return &MockArticleRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockArticleRepository_Delete_Call) Run(run func(id uint)) *MockArticleRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockArticleRepository_Delete_Call) Return(_a0 error) *MockArticleRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticleRepository_Delete_Call) RunAndReturn(run func(uint) error) *MockArticleRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FirstOrCreate provides a mock function with given fields: article
func (_m *MockArticleRepository) FirstOrCreate(article *entities.Article) error {
	ret := _m.Called(article)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Article) error); ok {
		r0 = rf(article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticleRepository_FirstOrCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstOrCreate'
type MockArticleRepository_FirstOrCreate_Call struct {
	*mock.Call
}

// FirstOrCreate is a helper method to define mock.On call
//   - article *entities.Article
func (_e *MockArticleRepository_Expecter) FirstOrCreate(article interface{}) *MockArticleRepository_FirstOrCreate_Call {
	return &MockArticleRepository_FirstOrCreate_Call{Call: _e.mock.On("FirstOrCreate", article)}
}

func (_c *MockArticleRepository_FirstOrCreate_Call) Run(run func(article *entities.Article)) *MockArticleRepository_FirstOrCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Article))
	})
	return _c
}

func (_c *MockArticleRepository_FirstOrCreate_Call) Return(_a0 error) *MockArticleRepository_FirstOrCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticleRepository_FirstOrCreate_Call) RunAndReturn(run func(*entities.Article) error) *MockArticleRepository_FirstOrCreate_Call {
	_c.Call.Return(run)
	return _c
}

// GetArticlesByIDs provides a mock function with given fields: ids
func (_m *MockArticleRepository) GetArticlesByIDs(ids []uint) ([]entities.Article, error) {
	ret := _m.Called(ids)

	var r0 []entities.Article
	var r1 error
	if rf, ok := ret.Get(0).(func([]uint) ([]entities.Article, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]uint) []entities.Article); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Article)
		}
	}

	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticleRepository_GetArticlesByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArticlesByIDs'
type MockArticleRepository_GetArticlesByIDs_Call struct {
	*mock.Call
}

// GetArticlesByIDs is a helper method to define mock.On call
//   - ids []uint
func (_e *MockArticleRepository_Expecter) GetArticlesByIDs(ids interface{}) *MockArticleRepository_GetArticlesByIDs_Call {
	return &MockArticleRepository_GetArticlesByIDs_Call{Call: _e.mock.On("GetArticlesByIDs", ids)}
}

func (_c *MockArticleRepository_GetArticlesByIDs_Call) Run(run func(ids []uint)) *MockArticleRepository_GetArticlesByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]uint))
	})
	return _c
}

func (_c *MockArticleRepository_GetArticlesByIDs_Call) Return(_a0 []entities.Article, _a1 error) *MockArticleRepository_GetArticlesByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticleRepository_GetArticlesByIDs_Call) RunAndReturn(run func([]uint) ([]entities.Article, error)) *MockArticleRepository_GetArticlesByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetCrawledArticleToday provides a mock function with given fields:
func (_m *MockArticleRepository) GetCrawledArticleToday() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticleRepository_GetCrawledArticleToday_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCrawledArticleToday'
type MockArticleRepository_GetCrawledArticleToday_Call struct {
	*mock.Call
}

// GetCrawledArticleToday is a helper method to define mock.On call
func (_e *MockArticleRepository_Expecter) GetCrawledArticleToday() *MockArticleRepository_GetCrawledArticleToday_Call {
	return &MockArticleRepository_GetCrawledArticleToday_Call{Call: _e.mock.On("GetCrawledArticleToday")}
}

func (_c *MockArticleRepository_GetCrawledArticleToday_Call) Run(run func()) *MockArticleRepository_GetCrawledArticleToday_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArticleRepository_GetCrawledArticleToday_Call) Return(_a0 int64, _a1 error) *MockArticleRepository_GetCrawledArticleToday_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticleRepository_GetCrawledArticleToday_Call) RunAndReturn(run func() (int64, error)) *MockArticleRepository_GetCrawledArticleToday_Call {
	_c.Call.Return(run)
	return _c
}

// GetTotalCrawledArticle provides a mock function with given fields:
func (_m *MockArticleRepository) GetTotalCrawledArticle() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticleRepository_GetTotalCrawledArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTotalCrawledArticle'
type MockArticleRepository_GetTotalCrawledArticle_Call struct {
	*mock.Call
}

// GetTotalCrawledArticle is a helper method to define mock.On call
func (_e *MockArticleRepository_Expecter) GetTotalCrawledArticle() *MockArticleRepository_GetTotalCrawledArticle_Call {
	return &MockArticleRepository_GetTotalCrawledArticle_Call{Call: _e.mock.On("GetTotalCrawledArticle")}
}

func (_c *MockArticleRepository_GetTotalCrawledArticle_Call) Run(run func()) *MockArticleRepository_GetTotalCrawledArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArticleRepository_GetTotalCrawledArticle_Call) Return(_a0 int64, _a1 error) *MockArticleRepository_GetTotalCrawledArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticleRepository_GetTotalCrawledArticle_Call) RunAndReturn(run func() (int64, error)) *MockArticleRepository_GetTotalCrawledArticle_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTag provides a mock function with given fields: ids, newTag
func (_m *MockArticleRepository) RemoveTag(ids []uint, newTag *entities.Tag) error {
	ret := _m.Called(ids, newTag)

	var r0 error
	if rf, ok := ret.Get(0).(func([]uint, *entities.Tag) error); ok {
		r0 = rf(ids, newTag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticleRepository_RemoveTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTag'
type MockArticleRepository_RemoveTag_Call struct {
	*mock.Call
}

// RemoveTag is a helper method to define mock.On call
//   - ids []uint
//   - newTag *entities.Tag
func (_e *MockArticleRepository_Expecter) RemoveTag(ids interface{}, newTag interface{}) *MockArticleRepository_RemoveTag_Call {
	return &MockArticleRepository_RemoveTag_Call{Call: _e.mock.On("RemoveTag", ids, newTag)}
}

func (_c *MockArticleRepository_RemoveTag_Call) Run(run func(ids []uint, newTag *entities.Tag)) *MockArticleRepository_RemoveTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]uint), args[1].(*entities.Tag))
	})
	return _c
}

func (_c *MockArticleRepository_RemoveTag_Call) Return(_a0 error) *MockArticleRepository_RemoveTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticleRepository_RemoveTag_Call) RunAndReturn(run func([]uint, *entities.Tag) error) *MockArticleRepository_RemoveTag_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockArticleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockArticleRepository creates a new instance of MockArticleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockArticleRepository(t mockConstructorTestingTNewMockArticleRepository) *MockArticleRepository {
	mock := &MockArticleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
