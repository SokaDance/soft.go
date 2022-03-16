// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type deepCopy struct {
	objects            map[EObject]EObject
	resolve            bool
	originalReferences bool
}

func newDeepCopy(resolve bool, originalReferences bool) *deepCopy {
	return &deepCopy{
		objects:            make(map[EObject]EObject),
		resolve:            resolve,
		originalReferences: originalReferences,
	}
}

func (dC *deepCopy) copy(eObject EObject) EObject {
	if eObject != nil {
		copyEObject := dC.createCopy(eObject)
		if copyEObject != nil {
			dC.objects[eObject] = copyEObject
			eClass := eObject.EClass()
			for it := eClass.GetEAllStructuralFeatures().Iterator(); it.HasNext(); {
				eFeature := it.Next().(EStructuralFeature)
				if eFeature.IsChangeable() && !eFeature.IsDerived() {
					if eAttribute, _ := eFeature.(EAttribute); eAttribute != nil {
						dC.copyAttribute(eAttribute, eObject, copyEObject)
					} else if eReference, _ := eFeature.(EReference); eReference != nil && eReference.IsContainment() {
						dC.copyContainment(eReference, eObject, copyEObject)
					}
				}
			}
			dC.copyProxyURI(eObject, copyEObject)
		}
		return copyEObject
	}
	return nil
}

func (dC *deepCopy) copyAll(eObjects EList[EObject]) EList[EObject] {
	copies := []EObject{}
	for it := eObjects.Iterator(); it.HasNext(); {
		eObject := it.Next()
		copies = append(copies, dC.copy(eObject))
	}
	return NewImmutableEList(copies)
}

func (dC *deepCopy) createCopy(eObject EObject) EObject {
	eClass := eObject.EClass()
	eFactory := eClass.GetEPackage().GetEFactoryInstance()
	return eFactory.Create(eClass)
}

func (dC *deepCopy) copyProxyURI(eObject EObject, copyEObject EObject) {
	if eObject.EIsProxy() {
		eObjectInternal := eObject.(EObjectInternal)
		eCopyInternal := copyEObject.(EObjectInternal)
		eCopyInternal.ESetProxyURI(eObjectInternal.EProxyURI())
	}
}

func (dC *deepCopy) copyAttribute(eAttribute EAttribute, eObject EObject, copyEObject EObject) {
	if eObject.EIsSet(eAttribute) {
		copyEObject.ESet(eAttribute, eObject.EGet(eAttribute))
	}
}

func (dC *deepCopy) copyContainment(eReference EReference, eObject EObject, copyEObject EObject) {
	if eObject.EIsSet(eReference) {
		value := eObject.EGetResolve(eReference, dC.resolve)
		if eReference.IsMany() {
			list := FromAnyList[EObject](value)
			copyEObject.ESet(eReference, dC.copyAll(list))
		} else {
			object := FromAny[EObject](value)
			copyEObject.ESet(eReference, dC.copy(object))
		}
	}
}

func (dC *deepCopy) copyReferences() {
	for eObject, copyEObject := range dC.objects {
		for it := eObject.EClass().GetEReferences().Iterator(); it.HasNext(); {
			eReference := it.Next().(EReference)
			if eReference.IsChangeable() && !eReference.IsDerived() && !eReference.IsContainment() && !eReference.IsContainer() {
				dC.copyReference(eReference, eObject, copyEObject)
			}
		}
	}
}

func (dC *deepCopy) copyReference(eReference EReference, eObject EObject, copyEObject EObject) {
	if eObject.EIsSet(eReference) {
		value := eObject.EGetResolve(eReference, dC.resolve)
		if eReference.IsMany() {
			listSource := FromAnyList[EObject](value).(EObjectList[EObject])
			listTarget := FromAnyList[EObject](copyEObject.EGetResolve(eReference, false)).(EObjectList[EObject])
			if !dC.resolve {
				listSource = listSource.GetUnResolvedList()
			}
			listTarget = listTarget.GetUnResolvedList()
			if listSource.Empty() {
				listTarget.Clear()
			} else {
				isBidirectional := eReference.GetEOpposite() != nil
				index := 0
				for it := listSource.Iterator(); it.HasNext(); {
					referencedObject := it.Next()
					copyReferencedEObject := dC.objects[referencedObject]
					if copyReferencedEObject != nil {
						if isBidirectional {
							position := listTarget.IndexOf(copyReferencedEObject)
							if position == -1 {
								listTarget.Insert(index, copyReferencedEObject)
							} else if index != position {
								listTarget.MoveObject(index, copyReferencedEObject)
							}

						} else {
							listTarget.Insert(index, copyReferencedEObject)
						}
						index++
					} else {
						if dC.originalReferences && !isBidirectional {
							listTarget.Insert(index, referencedObject)
							index++
						}
					}
				}
			}
		} else {
			object, _ := value.(EObject)
			if object != nil {
				copyReferencedEObject := dC.objects[object]
				if copyReferencedEObject != nil {
					copyEObject.ESet(eReference, copyReferencedEObject)
				} else {
					if dC.originalReferences && eReference.GetEOpposite() == nil {
						copyEObject.ESet(eReference, object)
					}
				}
			} else {
				copyEObject.ESet(eReference, object)
			}
		}
	}
}
