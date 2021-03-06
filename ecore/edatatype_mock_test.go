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

func discardMockEDataType() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEDataTypeIsSerializable tests method IsSerializable
func TestMockEDataTypeIsSerializable(t *testing.T) {
	o := &MockEDataType{}
	r := bool(true)
	o.On("IsSerializable").Once().Return(r)
	o.On("IsSerializable").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsSerializable())
	assert.Equal(t, r, o.IsSerializable())
	o.AssertExpectations(t)
}

// TestMockEDataTypeSetSerializable tests method SetSerializable
func TestMockEDataTypeSetSerializable(t *testing.T) {
	o := &MockEDataType{}
	v := bool(true)
	o.On("SetSerializable", v).Once()
	o.SetSerializable(v)
	o.AssertExpectations(t)
}
