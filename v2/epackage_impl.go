// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

// ePackageImpl is the implementation of the model object 'EPackage'
type ePackageImpl struct {
	eNamedElementImpl
	eClassifiers     EList[EClassifier]
	eFactoryInstance EFactory
	eSubPackages     EList[EPackage]
	nsPrefix         string
	nsURI            string
}
type ePackageImplInitializers interface {
	initEClassifiers() EList[EClassifier]
	initESubPackages() EList[EPackage]
}

// newEPackageImpl is the constructor of a ePackageImpl
func newEPackageImpl() *ePackageImpl {
	ePackage := new(ePackageImpl)
	ePackage.SetInterfaces(ePackage)
	ePackage.Initialize()
	return ePackage
}

func (ePackage *ePackageImpl) Initialize() {
	ePackage.eNamedElementImpl.Initialize()
	ePackage.nsPrefix = ""
	ePackage.nsURI = ""

}

func (ePackage *ePackageImpl) asEPackage() EPackage {
	return ePackage.GetInterfaces().(EPackage)
}

func (ePackage *ePackageImpl) asInitializers() ePackageImplInitializers {
	return ePackage.AsEObject().(ePackageImplInitializers)
}

func (ePackage *ePackageImpl) EStaticClass() EClass {
	return GetPackage().GetEPackage()
}

func (ePackage *ePackageImpl) EStaticFeatureCount() int {
	return EPACKAGE_FEATURE_COUNT
}

// GetEClassifier default implementation
func (ePackage *ePackageImpl) GetEClassifier(string) EClassifier {
	panic("GetEClassifier not implemented")
}

// GetEClassifiers get the value of eClassifiers
func (ePackage *ePackageImpl) GetEClassifiers() EList[EClassifier] {
	if ePackage.eClassifiers == nil {
		ePackage.eClassifiers = ePackage.asInitializers().initEClassifiers()
	}
	return ePackage.eClassifiers
}

// GetEFactoryInstance get the value of eFactoryInstance
func (ePackage *ePackageImpl) GetEFactoryInstance() EFactory {
	return ePackage.eFactoryInstance
}

// SetEFactoryInstance set the value of eFactoryInstance
func (ePackage *ePackageImpl) SetEFactoryInstance(newEFactoryInstance EFactory) {
	if newEFactoryInstance != ePackage.eFactoryInstance {
		var notifications ENotificationChain
		if oldEFactoryInstanceInternal, _ := ePackage.eFactoryInstance.(EObjectInternal); oldEFactoryInstanceInternal != nil {
			notifications = oldEFactoryInstanceInternal.EInverseRemove(ePackage, EFACTORY__EPACKAGE, notifications)
		}
		if newEFactoryInstanceInternal, _ := newEFactoryInstance.(EObjectInternal); newEFactoryInstanceInternal != nil {
			notifications = newEFactoryInstanceInternal.EInverseAdd(ePackage.AsEObject(), EFACTORY__EPACKAGE, notifications)
		}
		notifications = ePackage.basicSetEFactoryInstance(newEFactoryInstance, notifications)
		if notifications != nil {
			notifications.Dispatch()
		}
	}
}

func (ePackage *ePackageImpl) basicSetEFactoryInstance(newEFactoryInstance EFactory, msgs ENotificationChain) ENotificationChain {
	oldEFactoryInstance := ePackage.eFactoryInstance
	ePackage.eFactoryInstance = newEFactoryInstance
	notifications := msgs
	if ePackage.ENotificationRequired() {
		notification := NewNotificationByFeatureID(ePackage.AsEObject(), SET, EPACKAGE__EFACTORY_INSTANCE, oldEFactoryInstance, newEFactoryInstance, NO_INDEX)
		if notifications != nil {
			notifications.Add(notification)
		} else {
			notifications = notification
		}
	}
	return notifications
}

// GetESubPackages get the value of eSubPackages
func (ePackage *ePackageImpl) GetESubPackages() EList[EPackage] {
	if ePackage.eSubPackages == nil {
		ePackage.eSubPackages = ePackage.asInitializers().initESubPackages()
	}
	return ePackage.eSubPackages
}

// GetESuperPackage get the value of eSuperPackage
func (ePackage *ePackageImpl) GetESuperPackage() EPackage {
	if ePackage.EContainerFeatureID() == EPACKAGE__ESUPER_PACKAGE {
		return ePackage.EContainer().(EPackage)
	}
	return nil
}

// GetNsPrefix get the value of nsPrefix
func (ePackage *ePackageImpl) GetNsPrefix() string {
	return ePackage.nsPrefix
}

// SetNsPrefix set the value of nsPrefix
func (ePackage *ePackageImpl) SetNsPrefix(newNsPrefix string) {
	oldNsPrefix := ePackage.nsPrefix
	ePackage.nsPrefix = newNsPrefix
	if ePackage.ENotificationRequired() {
		ePackage.ENotify(NewNotificationByFeatureID(ePackage.AsEObject(), SET, EPACKAGE__NS_PREFIX, oldNsPrefix, newNsPrefix, NO_INDEX))
	}
}

// GetNsURI get the value of nsURI
func (ePackage *ePackageImpl) GetNsURI() string {
	return ePackage.nsURI
}

// SetNsURI set the value of nsURI
func (ePackage *ePackageImpl) SetNsURI(newNsURI string) {
	oldNsURI := ePackage.nsURI
	ePackage.nsURI = newNsURI
	if ePackage.ENotificationRequired() {
		ePackage.ENotify(NewNotificationByFeatureID(ePackage.AsEObject(), SET, EPACKAGE__NS_URI, oldNsURI, newNsURI, NO_INDEX))
	}
}

func (ePackage *ePackageImpl) initEClassifiers() EList[EClassifier] {
	return NewBasicEObjectList[EClassifier](ePackage.AsEObjectInternal(), EPACKAGE__ECLASSIFIERS, ECLASSIFIER__EPACKAGE, true, true, true, false, false)
}

func (ePackage *ePackageImpl) initESubPackages() EList[EPackage] {
	return NewBasicEObjectList[EPackage](ePackage.AsEObjectInternal(), EPACKAGE__ESUB_PACKAGES, EPACKAGE__ESUPER_PACKAGE, true, true, true, false, false)
}

func (ePackage *ePackageImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		return ToAnyList(ePackage.asEPackage().GetEClassifiers())
	case EPACKAGE__EFACTORY_INSTANCE:
		return ToAny(ePackage.asEPackage().GetEFactoryInstance())
	case EPACKAGE__ESUB_PACKAGES:
		return ToAnyList(ePackage.asEPackage().GetESubPackages())
	case EPACKAGE__ESUPER_PACKAGE:
		return ToAny(ePackage.asEPackage().GetESuperPackage())
	case EPACKAGE__NS_PREFIX:
		return ToAny(ePackage.asEPackage().GetNsPrefix())
	case EPACKAGE__NS_URI:
		return ToAny(ePackage.asEPackage().GetNsURI())
	default:
		return ePackage.eNamedElementImpl.EGetFromID(featureID, resolve)
	}
}

func (ePackage *ePackageImpl) ESetFromID(featureID int, value any) {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		newList := FromAnyList[EClassifier](value.(EList[any]))
		l := ePackage.asEPackage().GetEClassifiers()
		l.Clear()
		l.AddAll(newList)
	case EPACKAGE__EFACTORY_INSTANCE:
		newValue := FromAny[EFactory](value)
		ePackage.asEPackage().SetEFactoryInstance(newValue)
	case EPACKAGE__ESUB_PACKAGES:
		newList := FromAnyList[EPackage](value.(EList[any]))
		l := ePackage.asEPackage().GetESubPackages()
		l.Clear()
		l.AddAll(newList)
	case EPACKAGE__NS_PREFIX:
		newValue := FromAny[string](value)
		ePackage.asEPackage().SetNsPrefix(newValue)
	case EPACKAGE__NS_URI:
		newValue := FromAny[string](value)
		ePackage.asEPackage().SetNsURI(newValue)
	default:
		ePackage.eNamedElementImpl.ESetFromID(featureID, value)
	}
}

func (ePackage *ePackageImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		ePackage.asEPackage().GetEClassifiers().Clear()
	case EPACKAGE__EFACTORY_INSTANCE:
		ePackage.asEPackage().SetEFactoryInstance(nil)
	case EPACKAGE__ESUB_PACKAGES:
		ePackage.asEPackage().GetESubPackages().Clear()
	case EPACKAGE__NS_PREFIX:
		ePackage.asEPackage().SetNsPrefix("")
	case EPACKAGE__NS_URI:
		ePackage.asEPackage().SetNsURI("")
	default:
		ePackage.eNamedElementImpl.EUnsetFromID(featureID)
	}
}

func (ePackage *ePackageImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		return ePackage.eClassifiers != nil && ePackage.eClassifiers.Size() != 0
	case EPACKAGE__EFACTORY_INSTANCE:
		return ePackage.eFactoryInstance != nil
	case EPACKAGE__ESUB_PACKAGES:
		return ePackage.eSubPackages != nil && ePackage.eSubPackages.Size() != 0
	case EPACKAGE__ESUPER_PACKAGE:
		return ePackage.GetESuperPackage() != nil
	case EPACKAGE__NS_PREFIX:
		return ePackage.nsPrefix != ""
	case EPACKAGE__NS_URI:
		return ePackage.nsURI != ""
	default:
		return ePackage.eNamedElementImpl.EIsSetFromID(featureID)
	}
}

func (ePackage *ePackageImpl) EInvokeFromID(operationID int, arguments EList[any]) any {
	switch operationID {
	case EPACKAGE__GET_ECLASSIFIER_ESTRING:
		return ePackage.asEPackage().GetEClassifier(arguments.Get(0).(string))
	default:
		return ePackage.eNamedElementImpl.EInvokeFromID(operationID, arguments)
	}
}

func (ePackage *ePackageImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		list := ePackage.GetEClassifiers().(ENotifyingList[EClassifier])
		end := otherEnd.(EClassifier)
		return list.AddWithNotification(end, notifications)
	case EPACKAGE__EFACTORY_INSTANCE:
		msgs := notifications
		eFactoryInstance := ePackage.eFactoryInstance
		if eFactoryInstance != nil {
			msgs = eFactoryInstance.(EObjectInternal).EInverseRemove(ePackage.AsEObject(), EOPPOSITE_FEATURE_BASE-EPACKAGE__EFACTORY_INSTANCE, msgs)
		}
		return ePackage.basicSetEFactoryInstance(otherEnd.(EFactory), msgs)
	case EPACKAGE__ESUB_PACKAGES:
		list := ePackage.GetESubPackages().(ENotifyingList[EPackage])
		end := otherEnd.(EPackage)
		return list.AddWithNotification(end, notifications)
	case EPACKAGE__ESUPER_PACKAGE:
		msgs := notifications
		if ePackage.EInternalContainer() != nil {
			msgs = ePackage.EBasicRemoveFromContainer(msgs)
		}
		return ePackage.EBasicSetContainer(otherEnd, EPACKAGE__ESUPER_PACKAGE, msgs)
	default:
		return ePackage.eNamedElementImpl.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (ePackage *ePackageImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		list := ePackage.GetEClassifiers().(ENotifyingList[EClassifier])
		end := otherEnd.(EClassifier)
		return list.RemoveWithNotification(end, notifications)
	case EPACKAGE__EFACTORY_INSTANCE:
		return ePackage.basicSetEFactoryInstance(nil, notifications)
	case EPACKAGE__ESUB_PACKAGES:
		list := ePackage.GetESubPackages().(ENotifyingList[EPackage])
		end := otherEnd.(EPackage)
		return list.RemoveWithNotification(end, notifications)
	case EPACKAGE__ESUPER_PACKAGE:
		return ePackage.EBasicSetContainer(nil, EPACKAGE__ESUPER_PACKAGE, notifications)
	default:
		return ePackage.eNamedElementImpl.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
