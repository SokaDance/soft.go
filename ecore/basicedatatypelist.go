// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type basicEDataTypeList[T any] struct {
	BasicENotifyingList[T]
	owner     EObjectInternal
	featureID int
}

func NewBasicEDataTypeList[T any](owner EObjectInternal, featureID int, isUnique bool) *basicEDataTypeList[T] {
	l := new(basicEDataTypeList[T])
	l.interfaces = l
	l.data = []T{}
	l.owner = owner
	l.featureID = featureID
	l.isUnique = isUnique
	return l
}

// GetNotifier ...
func (list *basicEDataTypeList[T]) GetNotifier() ENotifier {
	return list.owner
}

// GetFeature ...
func (list *basicEDataTypeList[T]) GetFeature() EStructuralFeature {
	if list.owner != nil {
		return list.owner.EClass().GetEStructuralFeature(list.featureID)
	}
	return nil
}

// GetFeatureID ...
func (list *basicEDataTypeList[T]) GetFeatureID() int {
	return list.featureID
}
