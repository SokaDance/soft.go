// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

// eStringToStringMapEntryImpl is the implementation of the model object 'EStringToStringMapEntry'
type eStringToStringMapEntryImpl struct {
	CompactEObjectContainer
	key   string
	value string
}

// newEStringToStringMapEntryImpl is the constructor of a eStringToStringMapEntryImpl
func newEStringToStringMapEntryImpl() *eStringToStringMapEntryImpl {
	eStringToStringMapEntry := new(eStringToStringMapEntryImpl)
	eStringToStringMapEntry.SetInterfaces(eStringToStringMapEntry)
	eStringToStringMapEntry.Initialize()
	return eStringToStringMapEntry
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) Initialize() {
	eStringToStringMapEntry.CompactEObjectContainer.Initialize()
	eStringToStringMapEntry.key = ""
	eStringToStringMapEntry.value = ""

}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) asEStringToStringMapEntry() EStringToStringMapEntry {
	return eStringToStringMapEntry.GetInterfaces().(EStringToStringMapEntry)
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EStaticClass() EClass {
	return GetPackage().GetEStringToStringMapEntry()
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EStaticFeatureCount() int {
	return ESTRING_TO_STRING_MAP_ENTRY_FEATURE_COUNT
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) SetAnyKey(key any) {
	eStringToStringMapEntry.SetKey(key.(string))
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) SetAnyValue(value any) {
	eStringToStringMapEntry.SetValue(value.(string))
}

// GetKey get the value of key
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) GetKey() string {
	return eStringToStringMapEntry.key
}

// SetKey set the value of key
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) SetKey(newKey string) {
	oldKey := eStringToStringMapEntry.key
	eStringToStringMapEntry.key = newKey
	if eStringToStringMapEntry.ENotificationRequired() {
		eStringToStringMapEntry.ENotify(NewNotificationByFeatureID(eStringToStringMapEntry.AsEObject(), SET, ESTRING_TO_STRING_MAP_ENTRY__KEY, oldKey, newKey, NO_INDEX))
	}
}

// GetValue get the value of value
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) GetValue() string {
	return eStringToStringMapEntry.value
}

// SetValue set the value of value
func (eStringToStringMapEntry *eStringToStringMapEntryImpl) SetValue(newValue string) {
	oldValue := eStringToStringMapEntry.value
	eStringToStringMapEntry.value = newValue
	if eStringToStringMapEntry.ENotificationRequired() {
		eStringToStringMapEntry.ENotify(NewNotificationByFeatureID(eStringToStringMapEntry.AsEObject(), SET, ESTRING_TO_STRING_MAP_ENTRY__VALUE, oldValue, newValue, NO_INDEX))
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		return ToAny(eStringToStringMapEntry.asEStringToStringMapEntry().GetKey())
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		return ToAny(eStringToStringMapEntry.asEStringToStringMapEntry().GetValue())
	default:
		return eStringToStringMapEntry.CompactEObjectContainer.EGetFromID(featureID, resolve)
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) ESetFromID(featureID int, value any) {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		newValue := FromAny[string](value)
		eStringToStringMapEntry.asEStringToStringMapEntry().SetKey(newValue)
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		newValue := FromAny[string](value)
		eStringToStringMapEntry.asEStringToStringMapEntry().SetValue(newValue)
	default:
		eStringToStringMapEntry.CompactEObjectContainer.ESetFromID(featureID, value)
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		eStringToStringMapEntry.asEStringToStringMapEntry().SetKey("")
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		eStringToStringMapEntry.asEStringToStringMapEntry().SetValue("")
	default:
		eStringToStringMapEntry.CompactEObjectContainer.EUnsetFromID(featureID)
	}
}

func (eStringToStringMapEntry *eStringToStringMapEntryImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		return eStringToStringMapEntry.key != ""
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		return eStringToStringMapEntry.value != ""
	default:
		return eStringToStringMapEntry.CompactEObjectContainer.EIsSetFromID(featureID)
	}
}
