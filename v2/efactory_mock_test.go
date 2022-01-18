// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func discardMockEFactory() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEFactoryGetEPackage tests method GetEPackage
func TestMockEFactoryGetEPackage(t *testing.T) {
	o := &MockEFactory{}
	r := &MockEPackage{}
	o.On("GetEPackage").Once().Return(r)
	o.On("GetEPackage").Once().Return(func() EPackage {
		return r
	})
	assert.Equal(t, r, o.GetEPackage())
	assert.Equal(t, r, o.GetEPackage())
	o.AssertExpectations(t)
}

// TestMockEFactorySetEPackage tests method SetEPackage
func TestMockEFactorySetEPackage(t *testing.T) {
	o := &MockEFactory{}
	v := &MockEPackage{}
	o.On("SetEPackage", v).Once()
	o.SetEPackage(v)
	o.AssertExpectations(t)
}

// TestMockEFactoryConvertToString tests method ConvertToString
func TestMockEFactoryConvertToString(t *testing.T) {
	o := &MockEFactory{}
	eDataType := &MockEDataType{}
	instanceValue := any(nil)
	r := "Test String"
	o.On("ConvertToString", eDataType, instanceValue).Return(r).Once()
	o.On("ConvertToString", eDataType, instanceValue).Return(func() string {
		return r
	}).Once()
	assert.Equal(t, r, o.ConvertToString(eDataType, instanceValue))
	assert.Equal(t, r, o.ConvertToString(eDataType, instanceValue))
	o.AssertExpectations(t)
}

// TestMockEFactoryCreate tests method Create
func TestMockEFactoryCreate(t *testing.T) {
	o := &MockEFactory{}
	eClass := &MockEClass{}
	r := &MockEObjectInternal{}
	o.On("Create", eClass).Return(r).Once()
	o.On("Create", eClass).Return(func() EObject {
		return r
	}).Once()
	assert.Equal(t, r, o.Create(eClass))
	assert.Equal(t, r, o.Create(eClass))
	o.AssertExpectations(t)
}

// TestMockEFactoryCreateFromString tests method CreateFromString
func TestMockEFactoryCreateFromString(t *testing.T) {
	o := &MockEFactory{}
	eDataType := &MockEDataType{}
	literalValue := "Test String"
	r := any(nil)
	o.On("CreateFromString", eDataType, literalValue).Return(r).Once()
	o.On("CreateFromString", eDataType, literalValue).Return(func() any {
		return r
	}).Once()
	assert.Equal(t, r, o.CreateFromString(eDataType, literalValue))
	assert.Equal(t, r, o.CreateFromString(eDataType, literalValue))
	o.AssertExpectations(t)
}
