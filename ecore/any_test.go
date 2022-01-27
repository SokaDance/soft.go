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
