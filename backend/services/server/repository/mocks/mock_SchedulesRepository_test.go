// Code generated by mockery v2.20.0. DO NOT EDIT.

package repository

import (
	entities "server/entities"
	serverproto "server/proto"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockSchedulesRepository is an autogenerated mock type for the SchedulesRepository type
type MockSchedulesRepository struct {
	mock.Mock
}

type MockSchedulesRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSchedulesRepository) EXPECT() *MockSchedulesRepository_Expecter {
	return &MockSchedulesRepository_Expecter{mock: &_m.Mock}
}

// ClearMatchURLsOnDay provides a mock function with given fields:
func (_m *MockSchedulesRepository) ClearMatchURLsOnDay() {
	_m.Called()
}

// MockSchedulesRepository_ClearMatchURLsOnDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearMatchURLsOnDay'
type MockSchedulesRepository_ClearMatchURLsOnDay_Call struct {
	*mock.Call
}

// ClearMatchURLsOnDay is a helper method to define mock.On call
func (_e *MockSchedulesRepository_Expecter) ClearMatchURLsOnDay() *MockSchedulesRepository_ClearMatchURLsOnDay_Call {
	return &MockSchedulesRepository_ClearMatchURLsOnDay_Call{Call: _e.mock.On("ClearMatchURLsOnDay")}
}

func (_c *MockSchedulesRepository_ClearMatchURLsOnDay_Call) Run(run func()) *MockSchedulesRepository_ClearMatchURLsOnDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSchedulesRepository_ClearMatchURLsOnDay_Call) Return() *MockSchedulesRepository_ClearMatchURLsOnDay_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSchedulesRepository_ClearMatchURLsOnDay_Call) RunAndReturn(run func()) *MockSchedulesRepository_ClearMatchURLsOnDay_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllScheduleLeagueOnDay provides a mock function with given fields: date
func (_m *MockSchedulesRepository) GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error) {
	ret := _m.Called(date)

	var r0 entities.ScheduleOnDay
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time) (entities.ScheduleOnDay, error)); ok {
		return rf(date)
	}
	if rf, ok := ret.Get(0).(func(time.Time) entities.ScheduleOnDay); ok {
		r0 = rf(date)
	} else {
		r0 = ret.Get(0).(entities.ScheduleOnDay)
	}

	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllScheduleLeagueOnDay'
type MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call struct {
	*mock.Call
}

// GetAllScheduleLeagueOnDay is a helper method to define mock.On call
//   - date time.Time
func (_e *MockSchedulesRepository_Expecter) GetAllScheduleLeagueOnDay(date interface{}) *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call {
	return &MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call{Call: _e.mock.On("GetAllScheduleLeagueOnDay", date)}
}

func (_c *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call) Run(run func(date time.Time)) *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time))
	})
	return _c
}

func (_c *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call) Return(_a0 entities.ScheduleOnDay, _a1 error) *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call) RunAndReturn(run func(time.Time) (entities.ScheduleOnDay, error)) *MockSchedulesRepository_GetAllScheduleLeagueOnDay_Call {
	_c.Call.Return(run)
	return _c
}

// GetMatchURLsOnDay provides a mock function with given fields:
func (_m *MockSchedulesRepository) GetMatchURLsOnDay() entities.MatchURLsOnDay {
	ret := _m.Called()

	var r0 entities.MatchURLsOnDay
	if rf, ok := ret.Get(0).(func() entities.MatchURLsOnDay); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entities.MatchURLsOnDay)
	}

	return r0
}

// MockSchedulesRepository_GetMatchURLsOnDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMatchURLsOnDay'
type MockSchedulesRepository_GetMatchURLsOnDay_Call struct {
	*mock.Call
}

// GetMatchURLsOnDay is a helper method to define mock.On call
func (_e *MockSchedulesRepository_Expecter) GetMatchURLsOnDay() *MockSchedulesRepository_GetMatchURLsOnDay_Call {
	return &MockSchedulesRepository_GetMatchURLsOnDay_Call{Call: _e.mock.On("GetMatchURLsOnDay")}
}

func (_c *MockSchedulesRepository_GetMatchURLsOnDay_Call) Run(run func()) *MockSchedulesRepository_GetMatchURLsOnDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSchedulesRepository_GetMatchURLsOnDay_Call) Return(_a0 entities.MatchURLsOnDay) *MockSchedulesRepository_GetMatchURLsOnDay_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSchedulesRepository_GetMatchURLsOnDay_Call) RunAndReturn(run func() entities.MatchURLsOnDay) *MockSchedulesRepository_GetMatchURLsOnDay_Call {
	_c.Call.Return(run)
	return _c
}

// GetScheduleLeagueOnDay provides a mock function with given fields: date, league
func (_m *MockSchedulesRepository) GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error) {
	ret := _m.Called(date, league)

	var r0 entities.ScheduleOnDay
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, string) (entities.ScheduleOnDay, error)); ok {
		return rf(date, league)
	}
	if rf, ok := ret.Get(0).(func(time.Time, string) entities.ScheduleOnDay); ok {
		r0 = rf(date, league)
	} else {
		r0 = ret.Get(0).(entities.ScheduleOnDay)
	}

	if rf, ok := ret.Get(1).(func(time.Time, string) error); ok {
		r1 = rf(date, league)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSchedulesRepository_GetScheduleLeagueOnDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetScheduleLeagueOnDay'
type MockSchedulesRepository_GetScheduleLeagueOnDay_Call struct {
	*mock.Call
}

// GetScheduleLeagueOnDay is a helper method to define mock.On call
//   - date time.Time
//   - league string
func (_e *MockSchedulesRepository_Expecter) GetScheduleLeagueOnDay(date interface{}, league interface{}) *MockSchedulesRepository_GetScheduleLeagueOnDay_Call {
	return &MockSchedulesRepository_GetScheduleLeagueOnDay_Call{Call: _e.mock.On("GetScheduleLeagueOnDay", date, league)}
}

func (_c *MockSchedulesRepository_GetScheduleLeagueOnDay_Call) Run(run func(date time.Time, league string)) *MockSchedulesRepository_GetScheduleLeagueOnDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time), args[1].(string))
	})
	return _c
}

func (_c *MockSchedulesRepository_GetScheduleLeagueOnDay_Call) Return(_a0 entities.ScheduleOnDay, _a1 error) *MockSchedulesRepository_GetScheduleLeagueOnDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSchedulesRepository_GetScheduleLeagueOnDay_Call) RunAndReturn(run func(time.Time, string) (entities.ScheduleOnDay, error)) *MockSchedulesRepository_GetScheduleLeagueOnDay_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchedules provides a mock function with given fields: date
func (_m *MockSchedulesRepository) GetSchedules(date *serverproto.Date) {
	_m.Called(date)
}

// MockSchedulesRepository_GetSchedules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchedules'
type MockSchedulesRepository_GetSchedules_Call struct {
	*mock.Call
}

// GetSchedules is a helper method to define mock.On call
//   - date *serverproto.Date
func (_e *MockSchedulesRepository_Expecter) GetSchedules(date interface{}) *MockSchedulesRepository_GetSchedules_Call {
	return &MockSchedulesRepository_GetSchedules_Call{Call: _e.mock.On("GetSchedules", date)}
}

func (_c *MockSchedulesRepository_GetSchedules_Call) Run(run func(date *serverproto.Date)) *MockSchedulesRepository_GetSchedules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*serverproto.Date))
	})
	return _c
}

func (_c *MockSchedulesRepository_GetSchedules_Call) Return() *MockSchedulesRepository_GetSchedules_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSchedulesRepository_GetSchedules_Call) RunAndReturn(run func(*serverproto.Date)) *MockSchedulesRepository_GetSchedules_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockSchedulesRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSchedulesRepository creates a new instance of MockSchedulesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSchedulesRepository(t mockConstructorTestingTNewMockSchedulesRepository) *MockSchedulesRepository {
	mock := &MockSchedulesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
