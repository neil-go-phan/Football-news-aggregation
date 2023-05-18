// Code generated by mockery v2.20.0. DO NOT EDIT.

package repository

import (
	entities "server/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockClubRepository is an autogenerated mock type for the ClubRepository type
type MockClubRepository struct {
	mock.Mock
}

type MockClubRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClubRepository) EXPECT() *MockClubRepository_Expecter {
	return &MockClubRepository_Expecter{mock: &_m.Mock}
}

// FirstOrCreate provides a mock function with given fields: clubName, logo
func (_m *MockClubRepository) FirstOrCreate(clubName string, logo string) (*entities.Club, error) {
	ret := _m.Called(clubName, logo)

	var r0 *entities.Club
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*entities.Club, error)); ok {
		return rf(clubName, logo)
	}
	if rf, ok := ret.Get(0).(func(string, string) *entities.Club); ok {
		r0 = rf(clubName, logo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Club)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(clubName, logo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClubRepository_FirstOrCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstOrCreate'
type MockClubRepository_FirstOrCreate_Call struct {
	*mock.Call
}

// FirstOrCreate is a helper method to define mock.On call
//   - clubName string
//   - logo string
func (_e *MockClubRepository_Expecter) FirstOrCreate(clubName interface{}, logo interface{}) *MockClubRepository_FirstOrCreate_Call {
	return &MockClubRepository_FirstOrCreate_Call{Call: _e.mock.On("FirstOrCreate", clubName, logo)}
}

func (_c *MockClubRepository_FirstOrCreate_Call) Run(run func(clubName string, logo string)) *MockClubRepository_FirstOrCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockClubRepository_FirstOrCreate_Call) Return(_a0 *entities.Club, _a1 error) *MockClubRepository_FirstOrCreate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClubRepository_FirstOrCreate_Call) RunAndReturn(run func(string, string) (*entities.Club, error)) *MockClubRepository_FirstOrCreate_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: clubName
func (_m *MockClubRepository) GetByName(clubName string) (*entities.Club, error) {
	ret := _m.Called(clubName)

	var r0 *entities.Club
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Club, error)); ok {
		return rf(clubName)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Club); ok {
		r0 = rf(clubName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Club)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clubName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClubRepository_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockClubRepository_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - clubName string
func (_e *MockClubRepository_Expecter) GetByName(clubName interface{}) *MockClubRepository_GetByName_Call {
	return &MockClubRepository_GetByName_Call{Call: _e.mock.On("GetByName", clubName)}
}

func (_c *MockClubRepository_GetByName_Call) Run(run func(clubName string)) *MockClubRepository_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClubRepository_GetByName_Call) Return(_a0 *entities.Club, _a1 error) *MockClubRepository_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClubRepository_GetByName_Call) RunAndReturn(run func(string) (*entities.Club, error)) *MockClubRepository_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockClubRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClubRepository creates a new instance of MockClubRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClubRepository(t mockConstructorTestingTNewMockClubRepository) *MockClubRepository {
	mock := &MockClubRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}