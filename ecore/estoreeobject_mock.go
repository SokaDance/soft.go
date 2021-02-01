package ecore

// MockEStoreEObject is an autogenerated mock type for the EStoreEObject type
type MockEStoreEObject struct {
	MockEObject
}

// EStore provides a mock function with given fields:
func (_m *MockEStoreEObject) EStore() EStore {
	ret := _m.Called()

	var r0 EStore
	if rf, ok := ret.Get(0).(func() EStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EStore)
		}
	}

	return r0
}