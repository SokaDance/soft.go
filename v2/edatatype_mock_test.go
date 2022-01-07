// Code generated by soft.generator.go. DO NOT EDIT.



package ecore

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func discardMockEDataType() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEDataTypeIsSerializable tests method IsSerializable
func TestMockEDataTypeIsSerializable(t *testing.T) {
	o := &MockEDataType{}
	r := true
	o.On("IsSerializable").Once().Return(r)
	o.On("IsSerializable").Once().Return( func () bool {
		return r
	})
	assert.Equal(t,r,o.IsSerializable())
	assert.Equal(t,r,o.IsSerializable())
	o.AssertExpectations(t)
}

// TestMockEDataTypeSetSerializable tests method SetSerializable
func TestMockEDataTypeSetSerializable(t *testing.T) {
	o := &MockEDataType{}
	v := true
	o.On("SetSerializable",v).Once()
	o.SetSerializable(v)
	o.AssertExpectations(t)
}




