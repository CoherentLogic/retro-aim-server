// Code generated by mockery v2.35.2. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockBuddyHandler is an autogenerated mock type for the BuddyHandler type
type MockBuddyHandler struct {
	mock.Mock
}

type MockBuddyHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBuddyHandler) EXPECT() *MockBuddyHandler_Expecter {
	return &MockBuddyHandler_Expecter{mock: &_m.Mock}
}

// RightsQueryHandler provides a mock function with given fields: ctx
func (_m *MockBuddyHandler) RightsQueryHandler(ctx context.Context) XMessage {
	ret := _m.Called(ctx)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context) XMessage); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockBuddyHandler_RightsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RightsQueryHandler'
type MockBuddyHandler_RightsQueryHandler_Call struct {
	*mock.Call
}

// RightsQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBuddyHandler_Expecter) RightsQueryHandler(ctx interface{}) *MockBuddyHandler_RightsQueryHandler_Call {
	return &MockBuddyHandler_RightsQueryHandler_Call{Call: _e.mock.On("RightsQueryHandler", ctx)}
}

func (_c *MockBuddyHandler_RightsQueryHandler_Call) Run(run func(ctx context.Context)) *MockBuddyHandler_RightsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBuddyHandler_RightsQueryHandler_Call) Return(_a0 XMessage) *MockBuddyHandler_RightsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBuddyHandler_RightsQueryHandler_Call) RunAndReturn(run func(context.Context) XMessage) *MockBuddyHandler_RightsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBuddyHandler creates a new instance of MockBuddyHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBuddyHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBuddyHandler {
	mock := &MockBuddyHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
