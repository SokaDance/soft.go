// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

// ENotifyingList ...
type ENotifyingList[T any] interface {
	EList[T]

	GetNotifier() ENotifier

	GetFeature() EStructuralFeature

	GetFeatureID() int

	AddWithNotification(object T, notifications ENotificationChain) ENotificationChain

	RemoveWithNotification(object T, notifications ENotificationChain) ENotificationChain

	SetWithNotification(index int, object T, notifications ENotificationChain) ENotificationChain
}
