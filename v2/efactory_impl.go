// Code generated by soft.generator.go. DO NOT EDIT.



package ecore



// eFactoryImpl is the implementation of the model object 'EFactory'
type eFactoryImpl struct {
    eModelElementExt
    
}

// newEFactoryImpl is the constructor of a eFactoryImpl
func newEFactoryImpl() *eFactoryImpl {
    eFactory := new(eFactoryImpl)
	eFactory.SetInterfaces(eFactory)
	eFactory.Initialize()
    return eFactory
}


func (eFactory *eFactoryImpl) asEFactory() EFactory {
	return eFactory.GetInterfaces().(EFactory)
}


func (eFactory *eFactoryImpl) EStaticClass() EClass {
    return GetPackage().GetEFactory()
}

func (eFactory *eFactoryImpl) EStaticFeatureCount() int {
    return EFACTORY_FEATURE_COUNT
}

// ConvertToString default implementation
func (eFactory *eFactoryImpl) ConvertToString(EDataType, any) string {
    panic("ConvertToString not implemented")
}
// Create default implementation
func (eFactory *eFactoryImpl) Create(EClass) EObject {
    panic("Create not implemented")
}
// CreateFromString default implementation
func (eFactory *eFactoryImpl) CreateFromString(EDataType, string) any {
    panic("CreateFromString not implemented")
}

// GetEPackage get the value of ePackage
func (eFactory *eFactoryImpl) GetEPackage() EPackage {
    if eFactory.EContainerFeatureID() == EFACTORY__EPACKAGE {
        return eFactory.EContainer().(EPackage);
    } 
    return nil
}

// SetEPackage set the value of ePackage
func (eFactory *eFactoryImpl) SetEPackage( newEPackage EPackage ) {
    if ( newEPackage != eFactory.EInternalContainer() || (newEPackage != nil && eFactory.EContainerFeatureID() !=  EFACTORY__EPACKAGE)) {
        var notifications ENotificationChain
        if ( eFactory.EInternalContainer() != nil) {
            notifications = eFactory.EBasicRemoveFromContainer(notifications)
        }
        if newEPackageInternal , _ := newEPackage.(EObjectInternal); newEPackageInternal != nil {
            notifications = newEPackageInternal.EInverseAdd( eFactory.AsEObject() , EPACKAGE__EFACTORY_INSTANCE, notifications )
        }
        notifications = eFactory.basicSetEPackage( newEPackage, notifications )
        if ( notifications != nil ) {
            notifications.Dispatch()
        }
    } else if ( eFactory.ENotificationRequired() ) {
        eFactory.ENotify( NewNotificationByFeatureID(eFactory, SET, EFACTORY__EPACKAGE, newEPackage, newEPackage, NO_INDEX) )
    }
}


func (eFactory *eFactoryImpl) basicSetEPackage( newEPackage EPackage , msgs ENotificationChain )  ENotificationChain {
    return eFactory.EBasicSetContainer(newEPackage,EFACTORY__EPACKAGE,msgs) 
}



func (eFactory *eFactoryImpl) EGetFromID(featureID int, resolve bool) any {
    switch featureID {
    case EFACTORY__EPACKAGE:
        return ToAny(eFactory.asEFactory().GetEPackage())
    default:
        return eFactory.eModelElementExt.EGetFromID(featureID, resolve)
    }
}


func (eFactory *eFactoryImpl) ESetFromID(featureID int, value any) {
    switch featureID {
    case EFACTORY__EPACKAGE:
		newValue := FromAny[EPackage](value)
        eFactory.asEFactory().SetEPackage(newValue)
    default:
        eFactory.eModelElementExt.ESetFromID(featureID, value)
    }
}


func (eFactory *eFactoryImpl) EUnsetFromID(featureID int) {
    switch featureID {
    case EFACTORY__EPACKAGE:
        eFactory.asEFactory().SetEPackage(nil)
    default:
        eFactory.eModelElementExt.EUnsetFromID(featureID)
    }
}

func (eFactory *eFactoryImpl) EIsSetFromID(featureID int) bool {
    switch featureID {
    case EFACTORY__EPACKAGE:
        return eFactory.GetEPackage() != nil
    default:
        return eFactory.eModelElementExt.EIsSetFromID(featureID)
    }
}

func (eFactory *eFactoryImpl) EInvokeFromID(operationID int, arguments EList[any]) any {
    switch operationID {
    case EFACTORY__CONVERT_TO_STRING_EDATATYPE_EJAVAOBJECT:
        return eFactory.asEFactory().ConvertToString(arguments.Get(0).(EDataType),arguments.Get(1)) 
    case EFACTORY__CREATE_ECLASS:
        return eFactory.asEFactory().Create(arguments.Get(0).(EClass)) 
    case EFACTORY__CREATE_FROM_STRING_EDATATYPE_ESTRING:
        return eFactory.asEFactory().CreateFromString(arguments.Get(0).(EDataType),arguments.Get(1).(string)) 
    default:
        return eFactory.eModelElementExt.EInvokeFromID(operationID,arguments)
    }
}


func (eFactory *eFactoryImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case EFACTORY__EPACKAGE:
        msgs := notifications
        if eFactory.EInternalContainer() != nil {
            msgs = eFactory.EBasicRemoveFromContainer(msgs)
        }
        return eFactory.basicSetEPackage(otherEnd.(EPackage), msgs)
    default:
        return eFactory.eModelElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
    }
}


func (eFactory *eFactoryImpl) EBasicInverseRemove(otherEnd  EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case EFACTORY__EPACKAGE:
        return eFactory.basicSetEPackage(nil, notifications)
    default:
        return eFactory.eModelElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
    }
}


