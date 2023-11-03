// Code generated by mockery v2.35.2. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"
)

// MockFeedbagHandler is an autogenerated mock type for the FeedbagHandler type
type MockFeedbagHandler struct {
	mock.Mock
}

type MockFeedbagHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFeedbagHandler) EXPECT() *MockFeedbagHandler_Expecter {
	return &MockFeedbagHandler_Expecter{mock: &_m.Mock}
}

// DeleteItemHandler provides a mock function with given fields: ctx, sm, sess, fm, snacPayloadIn
func (_m *MockFeedbagHandler) DeleteItemHandler(ctx context.Context, sm SessionManager, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x0A_FeedbagDeleteItem) (XMessage, error) {
	ret := _m.Called(ctx, sm, sess, fm, snacPayloadIn)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x0A_FeedbagDeleteItem) (XMessage, error)); ok {
		return rf(ctx, sm, sess, fm, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x0A_FeedbagDeleteItem) XMessage); ok {
		r0 = rf(ctx, sm, sess, fm, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x0A_FeedbagDeleteItem) error); ok {
		r1 = rf(ctx, sm, sess, fm, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFeedbagHandler_DeleteItemHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItemHandler'
type MockFeedbagHandler_DeleteItemHandler_Call struct {
	*mock.Call
}

// DeleteItemHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sm SessionManager
//   - sess *Session
//   - fm FeedbagManager
//   - snacPayloadIn oscar.SNAC_0x13_0x0A_FeedbagDeleteItem
func (_e *MockFeedbagHandler_Expecter) DeleteItemHandler(ctx interface{}, sm interface{}, sess interface{}, fm interface{}, snacPayloadIn interface{}) *MockFeedbagHandler_DeleteItemHandler_Call {
	return &MockFeedbagHandler_DeleteItemHandler_Call{Call: _e.mock.On("DeleteItemHandler", ctx, sm, sess, fm, snacPayloadIn)}
}

func (_c *MockFeedbagHandler_DeleteItemHandler_Call) Run(run func(ctx context.Context, sm SessionManager, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x0A_FeedbagDeleteItem)) *MockFeedbagHandler_DeleteItemHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(SessionManager), args[2].(*Session), args[3].(FeedbagManager), args[4].(oscar.SNAC_0x13_0x0A_FeedbagDeleteItem))
	})
	return _c
}

func (_c *MockFeedbagHandler_DeleteItemHandler_Call) Return(_a0 XMessage, _a1 error) *MockFeedbagHandler_DeleteItemHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFeedbagHandler_DeleteItemHandler_Call) RunAndReturn(run func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x0A_FeedbagDeleteItem) (XMessage, error)) *MockFeedbagHandler_DeleteItemHandler_Call {
	_c.Call.Return(run)
	return _c
}

// InsertItemHandler provides a mock function with given fields: ctx, sm, sess, fm, snacPayloadIn
func (_m *MockFeedbagHandler) InsertItemHandler(ctx context.Context, sm SessionManager, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x08_FeedbagInsertItem) (XMessage, error) {
	ret := _m.Called(ctx, sm, sess, fm, snacPayloadIn)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x08_FeedbagInsertItem) (XMessage, error)); ok {
		return rf(ctx, sm, sess, fm, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x08_FeedbagInsertItem) XMessage); ok {
		r0 = rf(ctx, sm, sess, fm, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x08_FeedbagInsertItem) error); ok {
		r1 = rf(ctx, sm, sess, fm, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFeedbagHandler_InsertItemHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertItemHandler'
type MockFeedbagHandler_InsertItemHandler_Call struct {
	*mock.Call
}

// InsertItemHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sm SessionManager
//   - sess *Session
//   - fm FeedbagManager
//   - snacPayloadIn oscar.SNAC_0x13_0x08_FeedbagInsertItem
func (_e *MockFeedbagHandler_Expecter) InsertItemHandler(ctx interface{}, sm interface{}, sess interface{}, fm interface{}, snacPayloadIn interface{}) *MockFeedbagHandler_InsertItemHandler_Call {
	return &MockFeedbagHandler_InsertItemHandler_Call{Call: _e.mock.On("InsertItemHandler", ctx, sm, sess, fm, snacPayloadIn)}
}

func (_c *MockFeedbagHandler_InsertItemHandler_Call) Run(run func(ctx context.Context, sm SessionManager, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x08_FeedbagInsertItem)) *MockFeedbagHandler_InsertItemHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(SessionManager), args[2].(*Session), args[3].(FeedbagManager), args[4].(oscar.SNAC_0x13_0x08_FeedbagInsertItem))
	})
	return _c
}

func (_c *MockFeedbagHandler_InsertItemHandler_Call) Return(_a0 XMessage, _a1 error) *MockFeedbagHandler_InsertItemHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFeedbagHandler_InsertItemHandler_Call) RunAndReturn(run func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x08_FeedbagInsertItem) (XMessage, error)) *MockFeedbagHandler_InsertItemHandler_Call {
	_c.Call.Return(run)
	return _c
}

// QueryHandler provides a mock function with given fields: ctx, sess, fm
func (_m *MockFeedbagHandler) QueryHandler(ctx context.Context, sess *Session, fm FeedbagManager) (XMessage, error) {
	ret := _m.Called(ctx, sess, fm)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, FeedbagManager) (XMessage, error)); ok {
		return rf(ctx, sess, fm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, FeedbagManager) XMessage); ok {
		r0 = rf(ctx, sess, fm)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, FeedbagManager) error); ok {
		r1 = rf(ctx, sess, fm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFeedbagHandler_QueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryHandler'
type MockFeedbagHandler_QueryHandler_Call struct {
	*mock.Call
}

// QueryHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - fm FeedbagManager
func (_e *MockFeedbagHandler_Expecter) QueryHandler(ctx interface{}, sess interface{}, fm interface{}) *MockFeedbagHandler_QueryHandler_Call {
	return &MockFeedbagHandler_QueryHandler_Call{Call: _e.mock.On("QueryHandler", ctx, sess, fm)}
}

func (_c *MockFeedbagHandler_QueryHandler_Call) Run(run func(ctx context.Context, sess *Session, fm FeedbagManager)) *MockFeedbagHandler_QueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(FeedbagManager))
	})
	return _c
}

func (_c *MockFeedbagHandler_QueryHandler_Call) Return(_a0 XMessage, _a1 error) *MockFeedbagHandler_QueryHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFeedbagHandler_QueryHandler_Call) RunAndReturn(run func(context.Context, *Session, FeedbagManager) (XMessage, error)) *MockFeedbagHandler_QueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// QueryIfModifiedHandler provides a mock function with given fields: ctx, sess, fm, snacPayloadIn
func (_m *MockFeedbagHandler) QueryIfModifiedHandler(ctx context.Context, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x05_FeedbagQueryIfModified) (XMessage, error) {
	ret := _m.Called(ctx, sess, fm, snacPayloadIn)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Session, FeedbagManager, oscar.SNAC_0x13_0x05_FeedbagQueryIfModified) (XMessage, error)); ok {
		return rf(ctx, sess, fm, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Session, FeedbagManager, oscar.SNAC_0x13_0x05_FeedbagQueryIfModified) XMessage); ok {
		r0 = rf(ctx, sess, fm, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Session, FeedbagManager, oscar.SNAC_0x13_0x05_FeedbagQueryIfModified) error); ok {
		r1 = rf(ctx, sess, fm, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFeedbagHandler_QueryIfModifiedHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryIfModifiedHandler'
type MockFeedbagHandler_QueryIfModifiedHandler_Call struct {
	*mock.Call
}

// QueryIfModifiedHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *Session
//   - fm FeedbagManager
//   - snacPayloadIn oscar.SNAC_0x13_0x05_FeedbagQueryIfModified
func (_e *MockFeedbagHandler_Expecter) QueryIfModifiedHandler(ctx interface{}, sess interface{}, fm interface{}, snacPayloadIn interface{}) *MockFeedbagHandler_QueryIfModifiedHandler_Call {
	return &MockFeedbagHandler_QueryIfModifiedHandler_Call{Call: _e.mock.On("QueryIfModifiedHandler", ctx, sess, fm, snacPayloadIn)}
}

func (_c *MockFeedbagHandler_QueryIfModifiedHandler_Call) Run(run func(ctx context.Context, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x05_FeedbagQueryIfModified)) *MockFeedbagHandler_QueryIfModifiedHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Session), args[2].(FeedbagManager), args[3].(oscar.SNAC_0x13_0x05_FeedbagQueryIfModified))
	})
	return _c
}

func (_c *MockFeedbagHandler_QueryIfModifiedHandler_Call) Return(_a0 XMessage, _a1 error) *MockFeedbagHandler_QueryIfModifiedHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFeedbagHandler_QueryIfModifiedHandler_Call) RunAndReturn(run func(context.Context, *Session, FeedbagManager, oscar.SNAC_0x13_0x05_FeedbagQueryIfModified) (XMessage, error)) *MockFeedbagHandler_QueryIfModifiedHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RightsQueryHandler provides a mock function with given fields: _a0
func (_m *MockFeedbagHandler) RightsQueryHandler(_a0 context.Context) XMessage {
	ret := _m.Called(_a0)

	var r0 XMessage
	if rf, ok := ret.Get(0).(func(context.Context) XMessage); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	return r0
}

// MockFeedbagHandler_RightsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RightsQueryHandler'
type MockFeedbagHandler_RightsQueryHandler_Call struct {
	*mock.Call
}

// RightsQueryHandler is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockFeedbagHandler_Expecter) RightsQueryHandler(_a0 interface{}) *MockFeedbagHandler_RightsQueryHandler_Call {
	return &MockFeedbagHandler_RightsQueryHandler_Call{Call: _e.mock.On("RightsQueryHandler", _a0)}
}

func (_c *MockFeedbagHandler_RightsQueryHandler_Call) Run(run func(_a0 context.Context)) *MockFeedbagHandler_RightsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockFeedbagHandler_RightsQueryHandler_Call) Return(_a0 XMessage) *MockFeedbagHandler_RightsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFeedbagHandler_RightsQueryHandler_Call) RunAndReturn(run func(context.Context) XMessage) *MockFeedbagHandler_RightsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// StartClusterHandler provides a mock function with given fields: _a0, _a1
func (_m *MockFeedbagHandler) StartClusterHandler(_a0 context.Context, _a1 oscar.SNAC_0x13_0x11_FeedbagStartCluster) {
	_m.Called(_a0, _a1)
}

// MockFeedbagHandler_StartClusterHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartClusterHandler'
type MockFeedbagHandler_StartClusterHandler_Call struct {
	*mock.Call
}

// StartClusterHandler is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 oscar.SNAC_0x13_0x11_FeedbagStartCluster
func (_e *MockFeedbagHandler_Expecter) StartClusterHandler(_a0 interface{}, _a1 interface{}) *MockFeedbagHandler_StartClusterHandler_Call {
	return &MockFeedbagHandler_StartClusterHandler_Call{Call: _e.mock.On("StartClusterHandler", _a0, _a1)}
}

func (_c *MockFeedbagHandler_StartClusterHandler_Call) Run(run func(_a0 context.Context, _a1 oscar.SNAC_0x13_0x11_FeedbagStartCluster)) *MockFeedbagHandler_StartClusterHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNAC_0x13_0x11_FeedbagStartCluster))
	})
	return _c
}

func (_c *MockFeedbagHandler_StartClusterHandler_Call) Return() *MockFeedbagHandler_StartClusterHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFeedbagHandler_StartClusterHandler_Call) RunAndReturn(run func(context.Context, oscar.SNAC_0x13_0x11_FeedbagStartCluster)) *MockFeedbagHandler_StartClusterHandler_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItemHandler provides a mock function with given fields: ctx, sm, sess, fm, snacPayloadIn
func (_m *MockFeedbagHandler) UpdateItemHandler(ctx context.Context, sm SessionManager, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x09_FeedbagUpdateItem) (XMessage, error) {
	ret := _m.Called(ctx, sm, sess, fm, snacPayloadIn)

	var r0 XMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x09_FeedbagUpdateItem) (XMessage, error)); ok {
		return rf(ctx, sm, sess, fm, snacPayloadIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x09_FeedbagUpdateItem) XMessage); ok {
		r0 = rf(ctx, sm, sess, fm, snacPayloadIn)
	} else {
		r0 = ret.Get(0).(XMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x09_FeedbagUpdateItem) error); ok {
		r1 = rf(ctx, sm, sess, fm, snacPayloadIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFeedbagHandler_UpdateItemHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItemHandler'
type MockFeedbagHandler_UpdateItemHandler_Call struct {
	*mock.Call
}

// UpdateItemHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sm SessionManager
//   - sess *Session
//   - fm FeedbagManager
//   - snacPayloadIn oscar.SNAC_0x13_0x09_FeedbagUpdateItem
func (_e *MockFeedbagHandler_Expecter) UpdateItemHandler(ctx interface{}, sm interface{}, sess interface{}, fm interface{}, snacPayloadIn interface{}) *MockFeedbagHandler_UpdateItemHandler_Call {
	return &MockFeedbagHandler_UpdateItemHandler_Call{Call: _e.mock.On("UpdateItemHandler", ctx, sm, sess, fm, snacPayloadIn)}
}

func (_c *MockFeedbagHandler_UpdateItemHandler_Call) Run(run func(ctx context.Context, sm SessionManager, sess *Session, fm FeedbagManager, snacPayloadIn oscar.SNAC_0x13_0x09_FeedbagUpdateItem)) *MockFeedbagHandler_UpdateItemHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(SessionManager), args[2].(*Session), args[3].(FeedbagManager), args[4].(oscar.SNAC_0x13_0x09_FeedbagUpdateItem))
	})
	return _c
}

func (_c *MockFeedbagHandler_UpdateItemHandler_Call) Return(_a0 XMessage, _a1 error) *MockFeedbagHandler_UpdateItemHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFeedbagHandler_UpdateItemHandler_Call) RunAndReturn(run func(context.Context, SessionManager, *Session, FeedbagManager, oscar.SNAC_0x13_0x09_FeedbagUpdateItem) (XMessage, error)) *MockFeedbagHandler_UpdateItemHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFeedbagHandler creates a new instance of MockFeedbagHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFeedbagHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFeedbagHandler {
	mock := &MockFeedbagHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
