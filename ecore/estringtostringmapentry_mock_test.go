// Code generated by soft.generator.go. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func discardMockEStringToStringMapEntry() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEStringToStringMapEntryGetStringKey tests method GetStringKey
func TestMockEStringToStringMapEntryGetStringKey(t *testing.T) {
	o := &MockEStringToStringMapEntry{}
	r := "Test String"
	o.On("GetStringKey").Once().Return(r)
	o.On("GetStringKey").Once().Return(func() string {
		return r
	})
	assert.Equal(t, r, o.GetStringKey())
	assert.Equal(t, r, o.GetStringKey())
	o.AssertExpectations(t)
}

// TestMockEStringToStringMapEntrySetStringKey tests method SetStringKey
func TestMockEStringToStringMapEntrySetStringKey(t *testing.T) {
	o := &MockEStringToStringMapEntry{}
	v := "Test String"
	o.On("SetStringKey", v).Once()
	o.SetStringKey(v)
	o.AssertExpectations(t)
}

// TestMockEStringToStringMapEntryGetStringValue tests method GetStringValue
func TestMockEStringToStringMapEntryGetStringValue(t *testing.T) {
	o := &MockEStringToStringMapEntry{}
	r := "Test String"
	o.On("GetStringValue").Once().Return(r)
	o.On("GetStringValue").Once().Return(func() string {
		return r
	})
	assert.Equal(t, r, o.GetStringValue())
	assert.Equal(t, r, o.GetStringValue())
	o.AssertExpectations(t)
}

// TestMockEStringToStringMapEntrySetStringValue tests method SetStringValue
func TestMockEStringToStringMapEntrySetStringValue(t *testing.T) {
	o := &MockEStringToStringMapEntry{}
	v := "Test String"
	o.On("SetStringValue", v).Once()
	o.SetStringValue(v)
	o.AssertExpectations(t)
}
