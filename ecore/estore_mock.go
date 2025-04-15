// Code generated by mockery v2.53.2. DO NOT EDIT.

package ecore

import (
	iter "iter"

	mock "github.com/stretchr/testify/mock"
)

// MockEStore is an autogenerated mock type for the EStore type
type MockEStore struct {
	mock.Mock
}

type MockEStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEStore) EXPECT() *MockEStore_Expecter {
	return &MockEStore_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: object, feature, index, value
func (_m *MockEStore) Add(object EObject, feature EStructuralFeature, index int, value interface{}) {
	_m.Called(object, feature, index, value)
}

// MockEStore_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockEStore_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - index int
//   - value interface{}
func (_e *MockEStore_Expecter) Add(object interface{}, feature interface{}, index interface{}, value interface{}) *MockEStore_Add_Call {
	return &MockEStore_Add_Call{Call: _e.mock.On("Add", object, feature, index, value)}
}

func (_c *MockEStore_Add_Call) Run(run func(object EObject, feature EStructuralFeature, index int, value interface{})) *MockEStore_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(int), args[3].(interface{}))
	})
	return _c
}

func (_c *MockEStore_Add_Call) Return() *MockEStore_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_Add_Call) RunAndReturn(run func(EObject, EStructuralFeature, int, interface{})) *MockEStore_Add_Call {
	_c.Run(run)
	return _c
}

// AddAll provides a mock function with given fields: object, feature, index, collection
func (_m *MockEStore) AddAll(object EObject, feature EStructuralFeature, index int, collection Collection) {
	_m.Called(object, feature, index, collection)
}

// MockEStore_AddAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAll'
type MockEStore_AddAll_Call struct {
	*mock.Call
}

// AddAll is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - index int
//   - collection Collection
func (_e *MockEStore_Expecter) AddAll(object interface{}, feature interface{}, index interface{}, collection interface{}) *MockEStore_AddAll_Call {
	return &MockEStore_AddAll_Call{Call: _e.mock.On("AddAll", object, feature, index, collection)}
}

func (_c *MockEStore_AddAll_Call) Run(run func(object EObject, feature EStructuralFeature, index int, collection Collection)) *MockEStore_AddAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(int), args[3].(Collection))
	})
	return _c
}

func (_c *MockEStore_AddAll_Call) Return() *MockEStore_AddAll_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_AddAll_Call) RunAndReturn(run func(EObject, EStructuralFeature, int, Collection)) *MockEStore_AddAll_Call {
	_c.Run(run)
	return _c
}

// AddRoot provides a mock function with given fields: object
func (_m *MockEStore) AddRoot(object EObject) {
	_m.Called(object)
}

// MockEStore_AddRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRoot'
type MockEStore_AddRoot_Call struct {
	*mock.Call
}

// AddRoot is a helper method to define mock.On call
//   - object EObject
func (_e *MockEStore_Expecter) AddRoot(object interface{}) *MockEStore_AddRoot_Call {
	return &MockEStore_AddRoot_Call{Call: _e.mock.On("AddRoot", object)}
}

func (_c *MockEStore_AddRoot_Call) Run(run func(object EObject)) *MockEStore_AddRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject))
	})
	return _c
}

func (_c *MockEStore_AddRoot_Call) Return() *MockEStore_AddRoot_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_AddRoot_Call) RunAndReturn(run func(EObject)) *MockEStore_AddRoot_Call {
	_c.Run(run)
	return _c
}

// All provides a mock function with given fields: object, feature
func (_m *MockEStore) All(object EObject, feature EStructuralFeature) iter.Seq[interface{}] {
	ret := _m.Called(object, feature)

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 iter.Seq[interface{}]
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature) iter.Seq[interface{}]); ok {
		r0 = rf(object, feature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iter.Seq[interface{}])
		}
	}

	return r0
}

// MockEStore_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type MockEStore_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) All(object interface{}, feature interface{}) *MockEStore_All_Call {
	return &MockEStore_All_Call{Call: _e.mock.On("All", object, feature)}
}

func (_c *MockEStore_All_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_All_Call) Return(_a0 iter.Seq[interface{}]) *MockEStore_All_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_All_Call) RunAndReturn(run func(EObject, EStructuralFeature) iter.Seq[interface{}]) *MockEStore_All_Call {
	_c.Call.Return(run)
	return _c
}

// Clear provides a mock function with given fields: object, feature
func (_m *MockEStore) Clear(object EObject, feature EStructuralFeature) {
	_m.Called(object, feature)
}

// MockEStore_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type MockEStore_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) Clear(object interface{}, feature interface{}) *MockEStore_Clear_Call {
	return &MockEStore_Clear_Call{Call: _e.mock.On("Clear", object, feature)}
}

func (_c *MockEStore_Clear_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_Clear_Call) Return() *MockEStore_Clear_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_Clear_Call) RunAndReturn(run func(EObject, EStructuralFeature)) *MockEStore_Clear_Call {
	_c.Run(run)
	return _c
}

// Contains provides a mock function with given fields: object, feature, value
func (_m *MockEStore) Contains(object EObject, feature EStructuralFeature, value interface{}) bool {
	ret := _m.Called(object, feature, value)

	if len(ret) == 0 {
		panic("no return value specified for Contains")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, interface{}) bool); ok {
		r0 = rf(object, feature, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockEStore_Contains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contains'
type MockEStore_Contains_Call struct {
	*mock.Call
}

// Contains is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - value interface{}
func (_e *MockEStore_Expecter) Contains(object interface{}, feature interface{}, value interface{}) *MockEStore_Contains_Call {
	return &MockEStore_Contains_Call{Call: _e.mock.On("Contains", object, feature, value)}
}

func (_c *MockEStore_Contains_Call) Run(run func(object EObject, feature EStructuralFeature, value interface{})) *MockEStore_Contains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(interface{}))
	})
	return _c
}

func (_c *MockEStore_Contains_Call) Return(_a0 bool) *MockEStore_Contains_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_Contains_Call) RunAndReturn(run func(EObject, EStructuralFeature, interface{}) bool) *MockEStore_Contains_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: object, feature, index
func (_m *MockEStore) Get(object EObject, feature EStructuralFeature, index int) interface{} {
	ret := _m.Called(object, feature, index)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, int) interface{}); ok {
		r0 = rf(object, feature, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockEStore_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEStore_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - index int
func (_e *MockEStore_Expecter) Get(object interface{}, feature interface{}, index interface{}) *MockEStore_Get_Call {
	return &MockEStore_Get_Call{Call: _e.mock.On("Get", object, feature, index)}
}

func (_c *MockEStore_Get_Call) Run(run func(object EObject, feature EStructuralFeature, index int)) *MockEStore_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(int))
	})
	return _c
}

func (_c *MockEStore_Get_Call) Return(_a0 interface{}) *MockEStore_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_Get_Call) RunAndReturn(run func(EObject, EStructuralFeature, int) interface{}) *MockEStore_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetContainer provides a mock function with given fields: object
func (_m *MockEStore) GetContainer(object EObject) (EObject, EStructuralFeature) {
	ret := _m.Called(object)

	if len(ret) == 0 {
		panic("no return value specified for GetContainer")
	}

	var r0 EObject
	var r1 EStructuralFeature
	if rf, ok := ret.Get(0).(func(EObject) (EObject, EStructuralFeature)); ok {
		return rf(object)
	}
	if rf, ok := ret.Get(0).(func(EObject) EObject); ok {
		r0 = rf(object)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	if rf, ok := ret.Get(1).(func(EObject) EStructuralFeature); ok {
		r1 = rf(object)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(EStructuralFeature)
		}
	}

	return r0, r1
}

// MockEStore_GetContainer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContainer'
type MockEStore_GetContainer_Call struct {
	*mock.Call
}

// GetContainer is a helper method to define mock.On call
//   - object EObject
func (_e *MockEStore_Expecter) GetContainer(object interface{}) *MockEStore_GetContainer_Call {
	return &MockEStore_GetContainer_Call{Call: _e.mock.On("GetContainer", object)}
}

func (_c *MockEStore_GetContainer_Call) Run(run func(object EObject)) *MockEStore_GetContainer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject))
	})
	return _c
}

func (_c *MockEStore_GetContainer_Call) Return(_a0 EObject, _a1 EStructuralFeature) *MockEStore_GetContainer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEStore_GetContainer_Call) RunAndReturn(run func(EObject) (EObject, EStructuralFeature)) *MockEStore_GetContainer_Call {
	_c.Call.Return(run)
	return _c
}

// IndexOf provides a mock function with given fields: object, feature, value
func (_m *MockEStore) IndexOf(object EObject, feature EStructuralFeature, value interface{}) int {
	ret := _m.Called(object, feature, value)

	if len(ret) == 0 {
		panic("no return value specified for IndexOf")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, interface{}) int); ok {
		r0 = rf(object, feature, value)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockEStore_IndexOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IndexOf'
type MockEStore_IndexOf_Call struct {
	*mock.Call
}

// IndexOf is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - value interface{}
func (_e *MockEStore_Expecter) IndexOf(object interface{}, feature interface{}, value interface{}) *MockEStore_IndexOf_Call {
	return &MockEStore_IndexOf_Call{Call: _e.mock.On("IndexOf", object, feature, value)}
}

func (_c *MockEStore_IndexOf_Call) Run(run func(object EObject, feature EStructuralFeature, value interface{})) *MockEStore_IndexOf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(interface{}))
	})
	return _c
}

func (_c *MockEStore_IndexOf_Call) Return(_a0 int) *MockEStore_IndexOf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_IndexOf_Call) RunAndReturn(run func(EObject, EStructuralFeature, interface{}) int) *MockEStore_IndexOf_Call {
	_c.Call.Return(run)
	return _c
}

// IsEmpty provides a mock function with given fields: object, feature
func (_m *MockEStore) IsEmpty(object EObject, feature EStructuralFeature) bool {
	ret := _m.Called(object, feature)

	if len(ret) == 0 {
		panic("no return value specified for IsEmpty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature) bool); ok {
		r0 = rf(object, feature)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockEStore_IsEmpty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsEmpty'
type MockEStore_IsEmpty_Call struct {
	*mock.Call
}

// IsEmpty is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) IsEmpty(object interface{}, feature interface{}) *MockEStore_IsEmpty_Call {
	return &MockEStore_IsEmpty_Call{Call: _e.mock.On("IsEmpty", object, feature)}
}

func (_c *MockEStore_IsEmpty_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_IsEmpty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_IsEmpty_Call) Return(_a0 bool) *MockEStore_IsEmpty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_IsEmpty_Call) RunAndReturn(run func(EObject, EStructuralFeature) bool) *MockEStore_IsEmpty_Call {
	_c.Call.Return(run)
	return _c
}

// IsSet provides a mock function with given fields: object, feature
func (_m *MockEStore) IsSet(object EObject, feature EStructuralFeature) bool {
	ret := _m.Called(object, feature)

	if len(ret) == 0 {
		panic("no return value specified for IsSet")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature) bool); ok {
		r0 = rf(object, feature)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockEStore_IsSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSet'
type MockEStore_IsSet_Call struct {
	*mock.Call
}

// IsSet is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) IsSet(object interface{}, feature interface{}) *MockEStore_IsSet_Call {
	return &MockEStore_IsSet_Call{Call: _e.mock.On("IsSet", object, feature)}
}

func (_c *MockEStore_IsSet_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_IsSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_IsSet_Call) Return(_a0 bool) *MockEStore_IsSet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_IsSet_Call) RunAndReturn(run func(EObject, EStructuralFeature) bool) *MockEStore_IsSet_Call {
	_c.Call.Return(run)
	return _c
}

// LastIndexOf provides a mock function with given fields: object, feature, value
func (_m *MockEStore) LastIndexOf(object EObject, feature EStructuralFeature, value interface{}) int {
	ret := _m.Called(object, feature, value)

	if len(ret) == 0 {
		panic("no return value specified for LastIndexOf")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, interface{}) int); ok {
		r0 = rf(object, feature, value)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockEStore_LastIndexOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LastIndexOf'
type MockEStore_LastIndexOf_Call struct {
	*mock.Call
}

// LastIndexOf is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - value interface{}
func (_e *MockEStore_Expecter) LastIndexOf(object interface{}, feature interface{}, value interface{}) *MockEStore_LastIndexOf_Call {
	return &MockEStore_LastIndexOf_Call{Call: _e.mock.On("LastIndexOf", object, feature, value)}
}

func (_c *MockEStore_LastIndexOf_Call) Run(run func(object EObject, feature EStructuralFeature, value interface{})) *MockEStore_LastIndexOf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(interface{}))
	})
	return _c
}

func (_c *MockEStore_LastIndexOf_Call) Return(_a0 int) *MockEStore_LastIndexOf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_LastIndexOf_Call) RunAndReturn(run func(EObject, EStructuralFeature, interface{}) int) *MockEStore_LastIndexOf_Call {
	_c.Call.Return(run)
	return _c
}

// Move provides a mock function with given fields: object, feature, sourceIndex, targetIndex, needResult
func (_m *MockEStore) Move(object EObject, feature EStructuralFeature, sourceIndex int, targetIndex int, needResult bool) interface{} {
	ret := _m.Called(object, feature, sourceIndex, targetIndex, needResult)

	if len(ret) == 0 {
		panic("no return value specified for Move")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, int, int, bool) interface{}); ok {
		r0 = rf(object, feature, sourceIndex, targetIndex, needResult)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockEStore_Move_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Move'
type MockEStore_Move_Call struct {
	*mock.Call
}

// Move is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - sourceIndex int
//   - targetIndex int
//   - needResult bool
func (_e *MockEStore_Expecter) Move(object interface{}, feature interface{}, sourceIndex interface{}, targetIndex interface{}, needResult interface{}) *MockEStore_Move_Call {
	return &MockEStore_Move_Call{Call: _e.mock.On("Move", object, feature, sourceIndex, targetIndex, needResult)}
}

func (_c *MockEStore_Move_Call) Run(run func(object EObject, feature EStructuralFeature, sourceIndex int, targetIndex int, needResult bool)) *MockEStore_Move_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(int), args[3].(int), args[4].(bool))
	})
	return _c
}

func (_c *MockEStore_Move_Call) Return(_a0 interface{}) *MockEStore_Move_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_Move_Call) RunAndReturn(run func(EObject, EStructuralFeature, int, int, bool) interface{}) *MockEStore_Move_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: object, feature, index, needResult
func (_m *MockEStore) Remove(object EObject, feature EStructuralFeature, index int, needResult bool) interface{} {
	ret := _m.Called(object, feature, index, needResult)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, int, bool) interface{}); ok {
		r0 = rf(object, feature, index, needResult)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockEStore_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MockEStore_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - index int
//   - needResult bool
func (_e *MockEStore_Expecter) Remove(object interface{}, feature interface{}, index interface{}, needResult interface{}) *MockEStore_Remove_Call {
	return &MockEStore_Remove_Call{Call: _e.mock.On("Remove", object, feature, index, needResult)}
}

func (_c *MockEStore_Remove_Call) Run(run func(object EObject, feature EStructuralFeature, index int, needResult bool)) *MockEStore_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(int), args[3].(bool))
	})
	return _c
}

func (_c *MockEStore_Remove_Call) Return(_a0 interface{}) *MockEStore_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_Remove_Call) RunAndReturn(run func(EObject, EStructuralFeature, int, bool) interface{}) *MockEStore_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveRoot provides a mock function with given fields: object
func (_m *MockEStore) RemoveRoot(object EObject) {
	_m.Called(object)
}

// MockEStore_RemoveRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRoot'
type MockEStore_RemoveRoot_Call struct {
	*mock.Call
}

// RemoveRoot is a helper method to define mock.On call
//   - object EObject
func (_e *MockEStore_Expecter) RemoveRoot(object interface{}) *MockEStore_RemoveRoot_Call {
	return &MockEStore_RemoveRoot_Call{Call: _e.mock.On("RemoveRoot", object)}
}

func (_c *MockEStore_RemoveRoot_Call) Run(run func(object EObject)) *MockEStore_RemoveRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject))
	})
	return _c
}

func (_c *MockEStore_RemoveRoot_Call) Return() *MockEStore_RemoveRoot_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_RemoveRoot_Call) RunAndReturn(run func(EObject)) *MockEStore_RemoveRoot_Call {
	_c.Run(run)
	return _c
}

// Set provides a mock function with given fields: object, feature, index, value, needResult
func (_m *MockEStore) Set(object EObject, feature EStructuralFeature, index int, value interface{}, needResult bool) interface{} {
	ret := _m.Called(object, feature, index, value, needResult)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature, int, interface{}, bool) interface{}); ok {
		r0 = rf(object, feature, index, value, needResult)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockEStore_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockEStore_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
//   - index int
//   - value interface{}
//   - needResult bool
func (_e *MockEStore_Expecter) Set(object interface{}, feature interface{}, index interface{}, value interface{}, needResult interface{}) *MockEStore_Set_Call {
	return &MockEStore_Set_Call{Call: _e.mock.On("Set", object, feature, index, value, needResult)}
}

func (_c *MockEStore_Set_Call) Run(run func(object EObject, feature EStructuralFeature, index int, value interface{}, needResult bool)) *MockEStore_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature), args[2].(int), args[3].(interface{}), args[4].(bool))
	})
	return _c
}

func (_c *MockEStore_Set_Call) Return(_a0 interface{}) *MockEStore_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_Set_Call) RunAndReturn(run func(EObject, EStructuralFeature, int, interface{}, bool) interface{}) *MockEStore_Set_Call {
	_c.Call.Return(run)
	return _c
}

// SetContainer provides a mock function with given fields: object, container, feature
func (_m *MockEStore) SetContainer(object EObject, container EObject, feature EStructuralFeature) {
	_m.Called(object, container, feature)
}

// MockEStore_SetContainer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetContainer'
type MockEStore_SetContainer_Call struct {
	*mock.Call
}

// SetContainer is a helper method to define mock.On call
//   - object EObject
//   - container EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) SetContainer(object interface{}, container interface{}, feature interface{}) *MockEStore_SetContainer_Call {
	return &MockEStore_SetContainer_Call{Call: _e.mock.On("SetContainer", object, container, feature)}
}

func (_c *MockEStore_SetContainer_Call) Run(run func(object EObject, container EObject, feature EStructuralFeature)) *MockEStore_SetContainer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EObject), args[2].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_SetContainer_Call) Return() *MockEStore_SetContainer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_SetContainer_Call) RunAndReturn(run func(EObject, EObject, EStructuralFeature)) *MockEStore_SetContainer_Call {
	_c.Run(run)
	return _c
}

// Size provides a mock function with given fields: object, feature
func (_m *MockEStore) Size(object EObject, feature EStructuralFeature) int {
	ret := _m.Called(object, feature)

	if len(ret) == 0 {
		panic("no return value specified for Size")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature) int); ok {
		r0 = rf(object, feature)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockEStore_Size_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Size'
type MockEStore_Size_Call struct {
	*mock.Call
}

// Size is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) Size(object interface{}, feature interface{}) *MockEStore_Size_Call {
	return &MockEStore_Size_Call{Call: _e.mock.On("Size", object, feature)}
}

func (_c *MockEStore_Size_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_Size_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_Size_Call) Return(_a0 int) *MockEStore_Size_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_Size_Call) RunAndReturn(run func(EObject, EStructuralFeature) int) *MockEStore_Size_Call {
	_c.Call.Return(run)
	return _c
}

// ToArray provides a mock function with given fields: object, feature
func (_m *MockEStore) ToArray(object EObject, feature EStructuralFeature) []interface{} {
	ret := _m.Called(object, feature)

	if len(ret) == 0 {
		panic("no return value specified for ToArray")
	}

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(EObject, EStructuralFeature) []interface{}); ok {
		r0 = rf(object, feature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// MockEStore_ToArray_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToArray'
type MockEStore_ToArray_Call struct {
	*mock.Call
}

// ToArray is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) ToArray(object interface{}, feature interface{}) *MockEStore_ToArray_Call {
	return &MockEStore_ToArray_Call{Call: _e.mock.On("ToArray", object, feature)}
}

func (_c *MockEStore_ToArray_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_ToArray_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_ToArray_Call) Return(_a0 []interface{}) *MockEStore_ToArray_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEStore_ToArray_Call) RunAndReturn(run func(EObject, EStructuralFeature) []interface{}) *MockEStore_ToArray_Call {
	_c.Call.Return(run)
	return _c
}

// UnSet provides a mock function with given fields: object, feature
func (_m *MockEStore) UnSet(object EObject, feature EStructuralFeature) {
	_m.Called(object, feature)
}

// MockEStore_UnSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnSet'
type MockEStore_UnSet_Call struct {
	*mock.Call
}

// UnSet is a helper method to define mock.On call
//   - object EObject
//   - feature EStructuralFeature
func (_e *MockEStore_Expecter) UnSet(object interface{}, feature interface{}) *MockEStore_UnSet_Call {
	return &MockEStore_UnSet_Call{Call: _e.mock.On("UnSet", object, feature)}
}

func (_c *MockEStore_UnSet_Call) Run(run func(object EObject, feature EStructuralFeature)) *MockEStore_UnSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EObject), args[1].(EStructuralFeature))
	})
	return _c
}

func (_c *MockEStore_UnSet_Call) Return() *MockEStore_UnSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockEStore_UnSet_Call) RunAndReturn(run func(EObject, EStructuralFeature)) *MockEStore_UnSet_Call {
	_c.Run(run)
	return _c
}

// NewMockEStore creates a new instance of MockEStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEStore {
	mock := &MockEStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
