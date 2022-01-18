// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

// eParameterImpl is the implementation of the model object 'EParameter'
type eParameterImpl struct {
	eTypedElementExt
}

// newEParameterImpl is the constructor of a eParameterImpl
func newEParameterImpl() *eParameterImpl {
	eParameter := new(eParameterImpl)
	eParameter.SetInterfaces(eParameter)
	eParameter.Initialize()
	return eParameter
}

func (eParameter *eParameterImpl) asEParameter() EParameter {
	return eParameter.GetInterfaces().(EParameter)
}

func (eParameter *eParameterImpl) EStaticClass() EClass {
	return GetPackage().GetEParameter()
}

func (eParameter *eParameterImpl) EStaticFeatureCount() int {
	return EPARAMETER_FEATURE_COUNT
}

// GetEOperation get the value of eOperation
func (eParameter *eParameterImpl) GetEOperation() EOperation {
	if eParameter.EContainerFeatureID() == EPARAMETER__EOPERATION {
		return eParameter.EContainer().(EOperation)
	}
	return nil
}

func (eParameter *eParameterImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case EPARAMETER__EOPERATION:
		return ToAny(eParameter.asEParameter().GetEOperation())
	default:
		return eParameter.eTypedElementExt.EGetFromID(featureID, resolve)
	}
}

func (eParameter *eParameterImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EPARAMETER__EOPERATION:
		return eParameter.GetEOperation() != nil
	default:
		return eParameter.eTypedElementExt.EIsSetFromID(featureID)
	}
}

func (eParameter *eParameterImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EPARAMETER__EOPERATION:
		msgs := notifications
		if eParameter.EInternalContainer() != nil {
			msgs = eParameter.EBasicRemoveFromContainer(msgs)
		}
		return eParameter.EBasicSetContainer(otherEnd, EPARAMETER__EOPERATION, msgs)
	default:
		return eParameter.eTypedElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eParameter *eParameterImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EPARAMETER__EOPERATION:
		return eParameter.EBasicSetContainer(nil, EPARAMETER__EOPERATION, notifications)
	default:
		return eParameter.eTypedElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
