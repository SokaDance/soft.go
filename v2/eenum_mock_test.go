// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func discardMockEEnum() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEEnumGetELiterals tests method GetELiterals
func TestMockEEnumGetELiterals(t *testing.T) {
	o := &MockEEnum{}
	l := &MockEList[EEnumLiteral]{}
	// return a value
	o.On("GetELiterals").Once().Return(l)
	o.On("GetELiterals").Once().Return(func() EList[EEnumLiteral] {
		return l
	})
	assert.Equal(t, l, o.GetELiterals())
	assert.Equal(t, l, o.GetELiterals())
	o.AssertExpectations(t)
}

// TestMockEEnumGetEEnumLiteralByLiteral tests method GetEEnumLiteralByLiteral
func TestMockEEnumGetEEnumLiteralByLiteral(t *testing.T) {
	o := &MockEEnum{}
	literal := "Test String"
	r := &MockEEnumLiteral{}
	o.On("GetEEnumLiteralByLiteral", literal).Return(r).Once()
	o.On("GetEEnumLiteralByLiteral", literal).Return(func() EEnumLiteral {
		return r
	}).Once()
	assert.Equal(t, r, o.GetEEnumLiteralByLiteral(literal))
	assert.Equal(t, r, o.GetEEnumLiteralByLiteral(literal))
	o.AssertExpectations(t)
}

// TestMockEEnumGetEEnumLiteralByName tests method GetEEnumLiteralByName
func TestMockEEnumGetEEnumLiteralByName(t *testing.T) {
	o := &MockEEnum{}
	name := "Test String"
	r := &MockEEnumLiteral{}
	o.On("GetEEnumLiteralByName", name).Return(r).Once()
	o.On("GetEEnumLiteralByName", name).Return(func() EEnumLiteral {
		return r
	}).Once()
	assert.Equal(t, r, o.GetEEnumLiteralByName(name))
	assert.Equal(t, r, o.GetEEnumLiteralByName(name))
	o.AssertExpectations(t)
}

// TestMockEEnumGetEEnumLiteralByValue tests method GetEEnumLiteralByValue
func TestMockEEnumGetEEnumLiteralByValue(t *testing.T) {
	o := &MockEEnum{}
	value := int(45)
	r := &MockEEnumLiteral{}
	o.On("GetEEnumLiteralByValue", value).Return(r).Once()
	o.On("GetEEnumLiteralByValue", value).Return(func() EEnumLiteral {
		return r
	}).Once()
	assert.Equal(t, r, o.GetEEnumLiteralByValue(value))
	assert.Equal(t, r, o.GetEEnumLiteralByValue(value))
	o.AssertExpectations(t)
}
