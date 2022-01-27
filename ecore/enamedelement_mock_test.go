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

func discardMockENamedElement() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockENamedElementGetName tests method GetName
func TestMockENamedElementGetName(t *testing.T) {
	o := &MockENamedElement{}
	r := "Test String"
	o.On("GetName").Once().Return(r)
	o.On("GetName").Once().Return(func() string {
		return r
	})
	assert.Equal(t, r, o.GetName())
	assert.Equal(t, r, o.GetName())
	o.AssertExpectations(t)
}

// TestMockENamedElementSetName tests method SetName
func TestMockENamedElementSetName(t *testing.T) {
	o := &MockENamedElement{}
	v := "Test String"
	o.On("SetName", v).Once()
	o.SetName(v)
	o.AssertExpectations(t)
}
