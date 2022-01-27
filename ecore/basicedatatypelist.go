package ecore

type basicEDataTypeList[T comparable] struct {
	BasicENotifyingList[T]
	owner     EObjectInternal
	featureID int
}

func NewBasicEDataTypeList[T comparable](owner EObjectInternal, featureID int, isUnique bool) *basicEDataTypeList[T] {
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
