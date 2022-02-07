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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockEObjectIDManagerClear(t *testing.T) {
	rm := &MockEObjectIDManager{}
	rm.On("Clear").Once()
	rm.Clear()
	mock.AssertExpectationsForObjects(t, rm)
}

func TestMockEObjectIDManagerRegister(t *testing.T) {
	rm := &MockEObjectIDManager{}
	o := &MockEObject{}
	rm.On("Register", o).Once()
	rm.Register(o)
	mock.AssertExpectationsForObjects(t, rm, o)
}

func TestMockEObjectIDManagerUnRegister(t *testing.T) {
	rm := &MockEObjectIDManager{}
	o := &MockEObject{}
	rm.On("UnRegister", o).Once()
	rm.UnRegister(o)
	mock.AssertExpectationsForObjects(t, rm, o)
}

func TestMockEObjectIDManagerGetID(t *testing.T) {
	rm := &MockEObjectIDManager{}
	o := &MockEObject{}
	rm.On("GetID", o).Return("id1").Once()
	rm.On("GetID", o).Return(func(EObject) any {
		return "id2"
	}).Once()
	assert.Equal(t, "id1", rm.GetID(o))
	assert.Equal(t, "id2", rm.GetID(o))
	mock.AssertExpectationsForObjects(t, rm, o)
}

func TestMockEObjectIDManagerSetID(t *testing.T) {
	rm := &MockEObjectIDManager{}
	o := &MockEObject{}
	rm.On("SetID", o, 1).Return(nil).Once()
	rm.On("SetID", o, 1).Return(func(EObject, any) error {
		return nil
	}).Once()
	assert.Nil(t, rm.SetID(o, 1))
	assert.Nil(t, rm.SetID(o, 1))
	mock.AssertExpectationsForObjects(t, rm, o)
}

func TestMockEObjectIDManagerGetDetachedID(t *testing.T) {
	rm := &MockEObjectIDManager{}
	o := &MockEObject{}
	rm.On("GetDetachedID", o).Return("id1").Once()
	rm.On("GetDetachedID", o).Return(func(EObject) any {
		return "id2"
	}).Once()
	assert.Equal(t, "id1", rm.GetDetachedID(o))
	assert.Equal(t, "id2", rm.GetDetachedID(o))
	mock.AssertExpectationsForObjects(t, rm, o)
}

func TestMockEObjectIDManagerGetEObject(t *testing.T) {
	rm := &MockEObjectIDManager{}
	o := &MockEObject{}
	rm.On("GetEObject", "id1").Return(o).Once()
	rm.On("GetEObject", "id2").Return(func(any) EObject {
		return o
	}).Once()
	assert.Equal(t, o, rm.GetEObject("id1"))
	assert.Equal(t, o, rm.GetEObject("id2"))
	mock.AssertExpectationsForObjects(t, rm, o)
}
