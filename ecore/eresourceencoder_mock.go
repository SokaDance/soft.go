// Code generated by mockery v2.16.0. DO NOT EDIT.

package ecore

import mock "github.com/stretchr/testify/mock"

// MockEResourceEncoder is an autogenerated mock type for the EResourceEncoder type
type MockEResourceEncoder struct {
	mock.Mock
}

type MockEResourceEncoder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEResourceEncoder) EXPECT() *MockEResourceEncoder_Expecter {
	return &MockEResourceEncoder_Expecter{mock: &_m.Mock}
}

// Encode provides a mock function with given fields:
func (_m *MockEResourceEncoder) Encode() {
	_m.Called()
}

// MockEResourceEncoder_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type MockEResourceEncoder_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
func (_e *MockEResourceEncoder_Expecter) Encode() *MockEResourceEncoder_Encode_Call {
	return &MockEResourceEncoder_Encode_Call{Call: _e.mock.On("Encode")}
}

func (_c *MockEResourceEncoder_Encode_Call) Run(run func()) *MockEResourceEncoder_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEResourceEncoder_Encode_Call) Return() *MockEResourceEncoder_Encode_Call {
	_c.Call.Return()
	return _c
}

// EncodeObject provides a mock function with given fields: object
func (_m *MockEResourceEncoder) EncodeObject(object EObject) error {
	ret := _m.Called(object)

	var r0 error
	if rf, ok := ret.Get(0).(func(EObject) error); ok {
		r0 = rf(object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEResourceEncoder_EncodeObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EncodeObject'
type MockEResourceEncoder_EncodeObject_Call struct {
	*mock.Call
}

// EncodeObject is a helper method to define mock.On call
//   - object EObject
func (_e *MockEResourceEncoder_Expecter) EncodeObject(object interface{}) *MockEResourceEncoder_EncodeObject_Call {
	return &MockEResourceEncoder_EncodeObject_Call{Call: _e.mock.On("EncodeObject", object)}
}

func (_c *MockEResourceEncoder_EncodeObject_Call) Run(run func(object EObject)) *MockEResourceEncoder_EncodeObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject))
	})
	return _c
}

func (_c *MockEResourceEncoder_EncodeObject_Call) Return(_a0 error) *MockEResourceEncoder_EncodeObject_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockEResourceEncoder interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEResourceEncoder creates a new instance of MockEResourceEncoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEResourceEncoder(t mockConstructorTestingTNewMockEResourceEncoder) *MockEResourceEncoder {
	mock := &MockEResourceEncoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
