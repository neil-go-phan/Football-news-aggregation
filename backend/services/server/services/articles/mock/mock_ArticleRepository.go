// Code generated by mockery v2.20.0. DO NOT EDIT.

package articlesservices

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

// AddTagForAllArticle provides a mock function with given fields: tag
func (_m *MockArticleRepository) AddTagForAllArticle(tag string) error {
	ret := _m.Called(tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticleRepository_AddTagForAllArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTagForAllArticle'
type MockArticleRepository_AddTagForAllArticle_Call struct {
	*mock.Call
}

// AddTagForAllArticle is a helper method to define mock.On call
//   - tag string
func (_e *MockArticleRepository_Expecter) AddTagForAllArticle(tag interface{}) *MockArticleRepository_AddTagForAllArticle_Call {
	return &MockArticleRepository_AddTagForAllArticle_Call{Call: _e.mock.On("AddTagForAllArticle", tag)}
}

func (_c *MockArticleRepository_AddTagForAllArticle_Call) Run(run func(tag string)) *MockArticleRepository_AddTagForAllArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockArticleRepository_AddTagForAllArticle_Call) Return(_a0 error) *MockArticleRepository_AddTagForAllArticle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticleRepository_AddTagForAllArticle_Call) RunAndReturn(run func(string) error) *MockArticleRepository_AddTagForAllArticle_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteArticle provides a mock function with given fields: title
func (_m *MockArticleRepository) DeleteArticle(title string) error {
	ret := _m.Called(title)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticleRepository_DeleteArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteArticle'
type MockArticleRepository_DeleteArticle_Call struct {
	*mock.Call
}

// DeleteArticle is a helper method to define mock.On call
//   - title string
func (_e *MockArticleRepository_Expecter) DeleteArticle(title interface{}) *MockArticleRepository_DeleteArticle_Call {
	return &MockArticleRepository_DeleteArticle_Call{Call: _e.mock.On("DeleteArticle", title)}
}

func (_c *MockArticleRepository_DeleteArticle_Call) Run(run func(title string)) *MockArticleRepository_DeleteArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockArticleRepository_DeleteArticle_Call) Return(_a0 error) *MockArticleRepository_DeleteArticle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticleRepository_DeleteArticle_Call) RunAndReturn(run func(string) error) *MockArticleRepository_DeleteArticle_Call {
	_c.Call.Return(run)
	return _c
}

// GetArticleCount provides a mock function with given fields:
func (_m *MockArticleRepository) GetArticleCount() (float64, float64, error) {
	ret := _m.Called()

	var r0 float64
	var r1 float64
	var r2 error
	if rf, ok := ret.Get(0).(func() (float64, float64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockArticleRepository_GetArticleCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArticleCount'
type MockArticleRepository_GetArticleCount_Call struct {
	*mock.Call
}

// GetArticleCount is a helper method to define mock.On call
func (_e *MockArticleRepository_Expecter) GetArticleCount() *MockArticleRepository_GetArticleCount_Call {
	return &MockArticleRepository_GetArticleCount_Call{Call: _e.mock.On("GetArticleCount")}
}

func (_c *MockArticleRepository_GetArticleCount_Call) Run(run func()) *MockArticleRepository_GetArticleCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArticleRepository_GetArticleCount_Call) Return(total float64, today float64, err error) *MockArticleRepository_GetArticleCount_Call {
	_c.Call.Return(total, today, err)
	return _c
}

func (_c *MockArticleRepository_GetArticleCount_Call) RunAndReturn(run func() (float64, float64, error)) *MockArticleRepository_GetArticleCount_Call {
	_c.Call.Return(run)
	return _c
}

// GetArticles provides a mock function with given fields: keywords
func (_m *MockArticleRepository) GetArticles(keywords []string) {
	_m.Called(keywords)
}

// MockArticleRepository_GetArticles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArticles'
type MockArticleRepository_GetArticles_Call struct {
	*mock.Call
}

// GetArticles is a helper method to define mock.On call
//   - keywords []string
func (_e *MockArticleRepository_Expecter) GetArticles(keywords interface{}) *MockArticleRepository_GetArticles_Call {
	return &MockArticleRepository_GetArticles_Call{Call: _e.mock.On("GetArticles", keywords)}
}

func (_c *MockArticleRepository_GetArticles_Call) Run(run func(keywords []string)) *MockArticleRepository_GetArticles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockArticleRepository_GetArticles_Call) Return() *MockArticleRepository_GetArticles_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockArticleRepository_GetArticles_Call) RunAndReturn(run func([]string)) *MockArticleRepository_GetArticles_Call {
	_c.Call.Return(run)
	return _c
}

// GetFirstPageOfLeagueRelatedArticle provides a mock function with given fields: leagueName
func (_m *MockArticleRepository) GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]entities.Article, error) {
	ret := _m.Called(leagueName)

	var r0 []entities.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]entities.Article, error)); ok {
		return rf(leagueName)
	}
	if rf, ok := ret.Get(0).(func(string) []entities.Article); ok {
		r0 = rf(leagueName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(leagueName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirstPageOfLeagueRelatedArticle'
type MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call struct {
	*mock.Call
}

// GetFirstPageOfLeagueRelatedArticle is a helper method to define mock.On call
//   - leagueName string
func (_e *MockArticleRepository_Expecter) GetFirstPageOfLeagueRelatedArticle(leagueName interface{}) *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call {
	return &MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call{Call: _e.mock.On("GetFirstPageOfLeagueRelatedArticle", leagueName)}
}

func (_c *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call) Run(run func(leagueName string)) *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call) Return(_a0 []entities.Article, _a1 error) *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call) RunAndReturn(run func(string) ([]entities.Article, error)) *MockArticleRepository_GetFirstPageOfLeagueRelatedArticle_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshCache provides a mock function with given fields:
func (_m *MockArticleRepository) RefreshCache() {
	_m.Called()
}

// MockArticleRepository_RefreshCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshCache'
type MockArticleRepository_RefreshCache_Call struct {
	*mock.Call
}

// RefreshCache is a helper method to define mock.On call
func (_e *MockArticleRepository_Expecter) RefreshCache() *MockArticleRepository_RefreshCache_Call {
	return &MockArticleRepository_RefreshCache_Call{Call: _e.mock.On("RefreshCache")}
}

func (_c *MockArticleRepository_RefreshCache_Call) Run(run func()) *MockArticleRepository_RefreshCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArticleRepository_RefreshCache_Call) Return() *MockArticleRepository_RefreshCache_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockArticleRepository_RefreshCache_Call) RunAndReturn(run func()) *MockArticleRepository_RefreshCache_Call {
	_c.Call.Return(run)
	return _c
}

// SearchArticlesTagsAndKeyword provides a mock function with given fields: keyword, formatedTags, from
func (_m *MockArticleRepository) SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, float64, error) {
	ret := _m.Called(keyword, formatedTags, from)

	var r0 []entities.Article
	var r1 float64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, []string, int) ([]entities.Article, float64, error)); ok {
		return rf(keyword, formatedTags, from)
	}
	if rf, ok := ret.Get(0).(func(string, []string, int) []entities.Article); ok {
		r0 = rf(keyword, formatedTags, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, int) float64); ok {
		r1 = rf(keyword, formatedTags, from)
	} else {
		r1 = ret.Get(1).(float64)
	}

	if rf, ok := ret.Get(2).(func(string, []string, int) error); ok {
		r2 = rf(keyword, formatedTags, from)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockArticleRepository_SearchArticlesTagsAndKeyword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchArticlesTagsAndKeyword'
type MockArticleRepository_SearchArticlesTagsAndKeyword_Call struct {
	*mock.Call
}

// SearchArticlesTagsAndKeyword is a helper method to define mock.On call
//   - keyword string
//   - formatedTags []string
//   - from int
func (_e *MockArticleRepository_Expecter) SearchArticlesTagsAndKeyword(keyword interface{}, formatedTags interface{}, from interface{}) *MockArticleRepository_SearchArticlesTagsAndKeyword_Call {
	return &MockArticleRepository_SearchArticlesTagsAndKeyword_Call{Call: _e.mock.On("SearchArticlesTagsAndKeyword", keyword, formatedTags, from)}
}

func (_c *MockArticleRepository_SearchArticlesTagsAndKeyword_Call) Run(run func(keyword string, formatedTags []string, from int)) *MockArticleRepository_SearchArticlesTagsAndKeyword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(int))
	})
	return _c
}

func (_c *MockArticleRepository_SearchArticlesTagsAndKeyword_Call) Return(_a0 []entities.Article, _a1 float64, _a2 error) *MockArticleRepository_SearchArticlesTagsAndKeyword_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockArticleRepository_SearchArticlesTagsAndKeyword_Call) RunAndReturn(run func(string, []string, int) ([]entities.Article, float64, error)) *MockArticleRepository_SearchArticlesTagsAndKeyword_Call {
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
