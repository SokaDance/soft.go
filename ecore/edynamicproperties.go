package ecore

type EDynamicProperties interface {
	EDynamicGet(dynamicFeatureID int) any
	EDynamicSet(dynamicFeatureID int, newValue any)
	EDynamicUnset(dynamicFeatureID int)
}
