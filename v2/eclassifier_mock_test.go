// Code generated by soft.generator.go. DO NOT EDIT.



package ecore

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func discardMockEClassifier() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEClassifierGetClassifierID tests method GetClassifierID
func TestMockEClassifierGetClassifierID(t *testing.T) {
	o := &MockEClassifier{}
	r := int(45)
	o.On("GetClassifierID").Once().Return(r)
	o.On("GetClassifierID").Once().Return( func () int {
		return r
	})
	assert.Equal(t,r,o.GetClassifierID())
	assert.Equal(t,r,o.GetClassifierID())
	o.AssertExpectations(t)
}

// TestMockEClassifierSetClassifierID tests method SetClassifierID
func TestMockEClassifierSetClassifierID(t *testing.T) {
	o := &MockEClassifier{}
	v := int(45)
	o.On("SetClassifierID",v).Once()
	o.SetClassifierID(v)
	o.AssertExpectations(t)
}

// TestMockEClassifierGetDefaultValue tests method GetDefaultValue
func TestMockEClassifierGetDefaultValue(t *testing.T) {
	o := &MockEClassifier{}
	r := any(nil)
	o.On("GetDefaultValue").Once().Return(r)
	o.On("GetDefaultValue").Once().Return( func () any {
		return r
	})
	assert.Equal(t,r,o.GetDefaultValue())
	assert.Equal(t,r,o.GetDefaultValue())
	o.AssertExpectations(t)
}

// TestMockEClassifierGetEPackage tests method GetEPackage
func TestMockEClassifierGetEPackage(t *testing.T) {
	o := &MockEClassifier{}
	r := &MockEPackage{}
	o.On("GetEPackage").Once().Return(r)
	o.On("GetEPackage").Once().Return( func () EPackage {
		return r
	})
	assert.Equal(t,r,o.GetEPackage())
	assert.Equal(t,r,o.GetEPackage())
	o.AssertExpectations(t)
}

// TestMockEClassifierGetInstanceClass tests method GetInstanceClass
func TestMockEClassifierGetInstanceClass(t *testing.T) {
	o := &MockEClassifier{}
	r := reflect.TypeOf("")
	o.On("GetInstanceClass").Once().Return(r)
	o.On("GetInstanceClass").Once().Return( func () reflect.Type {
		return r
	})
	assert.Equal(t,r,o.GetInstanceClass())
	assert.Equal(t,r,o.GetInstanceClass())
	o.AssertExpectations(t)
}

// TestMockEClassifierSetInstanceClass tests method SetInstanceClass
func TestMockEClassifierSetInstanceClass(t *testing.T) {
	o := &MockEClassifier{}
	v := reflect.TypeOf("")
	o.On("SetInstanceClass",v).Once()
	o.SetInstanceClass(v)
	o.AssertExpectations(t)
}

// TestMockEClassifierGetInstanceTypeName tests method GetInstanceTypeName
func TestMockEClassifierGetInstanceTypeName(t *testing.T) {
	o := &MockEClassifier{}
	r := "Test String"
	o.On("GetInstanceTypeName").Once().Return(r)
	o.On("GetInstanceTypeName").Once().Return( func () string {
		return r
	})
	assert.Equal(t,r,o.GetInstanceTypeName())
	assert.Equal(t,r,o.GetInstanceTypeName())
	o.AssertExpectations(t)
}

// TestMockEClassifierSetInstanceTypeName tests method SetInstanceTypeName
func TestMockEClassifierSetInstanceTypeName(t *testing.T) {
	o := &MockEClassifier{}
	v := "Test String"
	o.On("SetInstanceTypeName",v).Once()
	o.SetInstanceTypeName(v)
	o.AssertExpectations(t)
}


// TestMockEClassifierIsInstance tests method IsInstance
func TestMockEClassifierIsInstance(t *testing.T) {
	o := &MockEClassifier{}
	object := any(nil)
	r := true
	o.On("IsInstance",object).Return(r).Once()	
	o.On("IsInstance",object).Return( func () bool {
		return r
	}).Once()	
	assert.Equal(t , r, o.IsInstance(object) )
	assert.Equal(t , r, o.IsInstance(object) )
	o.AssertExpectations(t)
}


