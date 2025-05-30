// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	dto "ingestor/internal/core/dto"

	mock "github.com/stretchr/testify/mock"

	model "ingestor/internal/model"
)

// Aggregator is an autogenerated mock type for the Aggregator type
type Aggregator struct {
	mock.Mock
}

type Aggregator_Expecter struct {
	mock *mock.Mock
}

func (_m *Aggregator) EXPECT() *Aggregator_Expecter {
	return &Aggregator_Expecter{mock: &_m.Mock}
}

// AddPulse provides a mock function with given fields: pulse
func (_m *Aggregator) AddPulse(pulse model.Pulse) {
	_m.Called(pulse)
}

// Aggregator_AddPulse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPulse'
type Aggregator_AddPulse_Call struct {
	*mock.Call
}

// AddPulse is a helper method to define mock.On call
//   - pulse model.Pulse
func (_e *Aggregator_Expecter) AddPulse(pulse interface{}) *Aggregator_AddPulse_Call {
	return &Aggregator_AddPulse_Call{Call: _e.mock.On("AddPulse", pulse)}
}

func (_c *Aggregator_AddPulse_Call) Run(run func(pulse model.Pulse)) *Aggregator_AddPulse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Pulse))
	})
	return _c
}

func (_c *Aggregator_AddPulse_Call) Return() *Aggregator_AddPulse_Call {
	_c.Call.Return()
	return _c
}

func (_c *Aggregator_AddPulse_Call) RunAndReturn(run func(model.Pulse)) *Aggregator_AddPulse_Call {
	_c.Run(run)
	return _c
}

// FlushAggregates provides a mock function with no fields
func (_m *Aggregator) FlushAggregates() {
	_m.Called()
}

// Aggregator_FlushAggregates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FlushAggregates'
type Aggregator_FlushAggregates_Call struct {
	*mock.Call
}

// FlushAggregates is a helper method to define mock.On call
func (_e *Aggregator_Expecter) FlushAggregates() *Aggregator_FlushAggregates_Call {
	return &Aggregator_FlushAggregates_Call{Call: _e.mock.On("FlushAggregates")}
}

func (_c *Aggregator_FlushAggregates_Call) Run(run func()) *Aggregator_FlushAggregates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Aggregator_FlushAggregates_Call) Return() *Aggregator_FlushAggregates_Call {
	_c.Call.Return()
	return _c
}

func (_c *Aggregator_FlushAggregates_Call) RunAndReturn(run func()) *Aggregator_FlushAggregates_Call {
	_c.Run(run)
	return _c
}

// GetAggregatedData provides a mock function with no fields
func (_m *Aggregator) GetAggregatedData() []dto.AggregatedPulse {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAggregatedData")
	}

	var r0 []dto.AggregatedPulse
	if rf, ok := ret.Get(0).(func() []dto.AggregatedPulse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.AggregatedPulse)
		}
	}

	return r0
}

// Aggregator_GetAggregatedData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAggregatedData'
type Aggregator_GetAggregatedData_Call struct {
	*mock.Call
}

// GetAggregatedData is a helper method to define mock.On call
func (_e *Aggregator_Expecter) GetAggregatedData() *Aggregator_GetAggregatedData_Call {
	return &Aggregator_GetAggregatedData_Call{Call: _e.mock.On("GetAggregatedData")}
}

func (_c *Aggregator_GetAggregatedData_Call) Run(run func()) *Aggregator_GetAggregatedData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Aggregator_GetAggregatedData_Call) Return(_a0 []dto.AggregatedPulse) *Aggregator_GetAggregatedData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Aggregator_GetAggregatedData_Call) RunAndReturn(run func() []dto.AggregatedPulse) *Aggregator_GetAggregatedData_Call {
	_c.Call.Return(run)
	return _c
}

// NewAggregator creates a new instance of Aggregator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAggregator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Aggregator {
	mock := &Aggregator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
