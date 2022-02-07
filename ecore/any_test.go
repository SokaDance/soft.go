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
)

func TestFromAny(t *testing.T) {
	v := 1
	a := any(v)
	assert.Equal(t, 1, FromAny[int](a))
	assert.Equal(t, a, FromAny[any](a))
}

func TestToAnyCollection(t *testing.T) {
	c1 := &MockECollection[any]{}
	assert.Nil(t, ToAnyCollection[any](nil))
	assert.Equal(t, c1, ToAnyCollection[any](c1))
	c2 := &MockECollection[EObject]{}
	assert.Nil(t, ToAnyCollection[EObject](nil))
	assert.NotNil(t, ToAnyCollection[EObject](c2))
}

func TestFromAnyCollection(t *testing.T) {
	c1 := &MockECollection[EObject]{}
	assert.Nil(t, FromAnyCollection[any](nil))
	assert.Equal(t, c1, FromAnyCollection[EObject](c1))
	c2 := &MockECollection[any]{}
	assert.Nil(t, FromAnyCollection[EObject](nil))
	assert.NotNil(t, FromAnyCollection[EObject](c2))
}

func TestToAnyList(t *testing.T) {
	assert.Nil(t, ToAnyList[any](nil))
	assert.Nil(t, ToAnyList[EObject](nil))

	c1 := &MockEList[any]{}
	assert.Equal(t, c1, ToAnyList[any](c1))

	c2 := &MockEList[EObject]{}
	assert.NotNil(t, ToAnyList[EObject](c2))

	c3 := &MockENotifyingList[EObject]{}
	anyC3 := ToAnyList[EObject](c3)
	assert.NotNil(t, anyC3)
	_, isNotifyingList := anyC3.(ENotifyingList[any])
	assert.True(t, isNotifyingList)

	var c4 EList[EObject]
	anyC4 := ToAnyList(c4)
	assert.Nil(t, anyC4)
}

func TestFromAnyList(t *testing.T) {
	assert.Nil(t, FromAnyList[EObject](nil))
	assert.Nil(t, FromAnyList[any](nil))

	c1 := &MockEList[any]{}
	assert.Equal(t, c1, FromAnyList[any](c1))

	c2 := &MockEList[EObject]{}
	assert.NotNil(t, FromAnyList[EObject](c2))

	c3 := &MockENotifyingList[any]{}
	objectC3 := FromAnyList[EObject](c3)
	assert.NotNil(t, objectC3)
	_, isNotifyingList := objectC3.(ENotifyingList[EObject])
	assert.True(t, isNotifyingList)

	var c4 EList[any]
	objectC4 := FromAnyList[EObject](c4)
	assert.Nil(t, objectC4)
}

func TestToAnyMap(t *testing.T) {
	assert.Nil(t, ToAnyMap[any, any](nil))
	assert.Nil(t, ToAnyMap[EObject, EObject](nil))

	m1 := &MockEMap[any, any]{}
	assert.Equal(t, m1, ToAnyMap[any, any](m1))

	m2 := &MockEMap[EObject, EObject]{}
	r2 := ToAnyMap[EObject, EObject](m2)
	assert.NotNil(t, r2)
	_, isMap := r2.(EMap[any, any])
	assert.True(t, isMap)
}

func TestFromAnyMap(t *testing.T) {
	assert.Nil(t, FromAnyMap[any, any](nil))
	assert.Nil(t, FromAnyMap[EObject, EObject](nil))

	m1 := &MockEMap[EObject, EObject]{}
	assert.Equal(t, m1, FromAnyMap[EObject, EObject](m1))

	m2 := &MockEMap[any, any]{}
	r2 := FromAnyMap[EObject, EObject](m2)
	assert.NotNil(t, r2)
	_, isMap := r2.(EMap[EObject, EObject])
	assert.True(t, isMap)
}
