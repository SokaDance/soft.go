// Code generated by soft.generator.go. DO NOT EDIT.



package ecore



// eClassImpl is the implementation of the model object 'EClass'
type eClassImpl struct {
    eClassifierExt
    eAllAttributes EList[EAttribute]
    eAllContainments EList[EReference]
    eAllOperations EList[EOperation]
    eAllReferences EList[EReference]
    eAllStructuralFeatures EList[EStructuralFeature]
    eAllSuperTypes EList[EClass]
    eAttributes EList[EAttribute]
    eContainmentFeatures EList[EStructuralFeature]
    eCrossReferenceFeatures EList[EStructuralFeature]
    eIDAttribute EAttribute
    eOperations EList[EOperation]
    eReferences EList[EReference]
    eStructuralFeatures EList[EStructuralFeature]
    eSuperTypes EList[EClass]
    isAbstract bool
    isInterface bool
    
}
type eClassImplInitializers interface {
    initEAllAttributes()
    initEAllContainments()
    initEAllOperations()
    initEAllReferences()
    initEAllStructuralFeatures()
    initEAllSuperTypes()
    initEAttributes()
    initEContainmentFeatures()
    initECrossReferenceFeatures()
    initEIDAttribute()
    initEOperations() EList[EOperation]
    initEReferences()
    initEStructuralFeatures() EList[EStructuralFeature]
    initESuperTypes() EList[EClass]
    
}

// newEClassImpl is the constructor of a eClassImpl
func newEClassImpl() *eClassImpl {
    eClass := new(eClassImpl)
	eClass.SetInterfaces(eClass)
	eClass.Initialize()
    return eClass
}

func (eClass *eClassImpl) Initialize() {
	eClass.eClassifierExt.Initialize()
	eClass.isAbstract = false
	eClass.isInterface = false
	
}

func (eClass *eClassImpl) asEClass() EClass {
	return eClass.GetInterfaces().(EClass)
}

func (eClass *eClassImpl) asInitializers() eClassImplInitializers {
    return eClass.AsEObject().(eClassImplInitializers)
}

func (eClass *eClassImpl) EStaticClass() EClass {
    return GetPackage().GetEClass()
}

func (eClass *eClassImpl) EStaticFeatureCount() int {
    return ECLASS_FEATURE_COUNT
}

// GetEOperation default implementation
func (eClass *eClassImpl) GetEOperation(int) EOperation {
    panic("GetEOperation not implemented")
}
// GetEStructuralFeature default implementation
func (eClass *eClassImpl) GetEStructuralFeature(int) EStructuralFeature {
    panic("GetEStructuralFeature not implemented")
}
// GetEStructuralFeatureFromName default implementation
func (eClass *eClassImpl) GetEStructuralFeatureFromName(string) EStructuralFeature {
    panic("GetEStructuralFeatureFromName not implemented")
}
// GetFeatureCount default implementation
func (eClass *eClassImpl) GetFeatureCount() int {
    panic("GetFeatureCount not implemented")
}
// GetFeatureID default implementation
func (eClass *eClassImpl) GetFeatureID(EStructuralFeature) int {
    panic("GetFeatureID not implemented")
}
// GetFeatureType default implementation
func (eClass *eClassImpl) GetFeatureType(EStructuralFeature) EClassifier {
    panic("GetFeatureType not implemented")
}
// GetOperationCount default implementation
func (eClass *eClassImpl) GetOperationCount() int {
    panic("GetOperationCount not implemented")
}
// GetOperationID default implementation
func (eClass *eClassImpl) GetOperationID(EOperation) int {
    panic("GetOperationID not implemented")
}
// GetOverride default implementation
func (eClass *eClassImpl) GetOverride(EOperation) EOperation {
    panic("GetOverride not implemented")
}
// IsSuperTypeOf default implementation
func (eClass *eClassImpl) IsSuperTypeOf(EClass) bool {
    panic("IsSuperTypeOf not implemented")
}

// GetEAllAttributes get the value of eAllAttributes
func (eClass *eClassImpl) GetEAllAttributes() EList[EAttribute] {
    eClass.asInitializers().initEAllAttributes()
    return eClass.eAllAttributes
}

// GetEAllContainments get the value of eAllContainments
func (eClass *eClassImpl) GetEAllContainments() EList[EReference] {
    eClass.asInitializers().initEAllContainments()
    return eClass.eAllContainments
}

// GetEAllOperations get the value of eAllOperations
func (eClass *eClassImpl) GetEAllOperations() EList[EOperation] {
    eClass.asInitializers().initEAllOperations()
    return eClass.eAllOperations
}

// GetEAllReferences get the value of eAllReferences
func (eClass *eClassImpl) GetEAllReferences() EList[EReference] {
    eClass.asInitializers().initEAllReferences()
    return eClass.eAllReferences
}

// GetEAllStructuralFeatures get the value of eAllStructuralFeatures
func (eClass *eClassImpl) GetEAllStructuralFeatures() EList[EStructuralFeature] {
    eClass.asInitializers().initEAllStructuralFeatures()
    return eClass.eAllStructuralFeatures
}

// GetEAllSuperTypes get the value of eAllSuperTypes
func (eClass *eClassImpl) GetEAllSuperTypes() EList[EClass] {
    eClass.asInitializers().initEAllSuperTypes()
    return eClass.eAllSuperTypes
}

// GetEAttributes get the value of eAttributes
func (eClass *eClassImpl) GetEAttributes() EList[EAttribute] {
    eClass.asInitializers().initEAttributes()
    return eClass.eAttributes
}

// GetEContainmentFeatures get the value of eContainmentFeatures
func (eClass *eClassImpl) GetEContainmentFeatures() EList[EStructuralFeature] {
    eClass.asInitializers().initEContainmentFeatures()
    return eClass.eContainmentFeatures
}

// GetECrossReferenceFeatures get the value of eCrossReferenceFeatures
func (eClass *eClassImpl) GetECrossReferenceFeatures() EList[EStructuralFeature] {
    eClass.asInitializers().initECrossReferenceFeatures()
    return eClass.eCrossReferenceFeatures
}

// GetEIDAttribute get the value of eIDAttribute
func (eClass *eClassImpl) GetEIDAttribute() EAttribute {
    eClass.asInitializers().initEIDAttribute()
    return eClass.eIDAttribute
}

// GetEOperations get the value of eOperations
func (eClass *eClassImpl) GetEOperations() EList[EOperation] {
    if eClass.eOperations == nil {
        eClass.eOperations = eClass.asInitializers().initEOperations()
    }
    return eClass.eOperations
}

// GetEReferences get the value of eReferences
func (eClass *eClassImpl) GetEReferences() EList[EReference] {
    eClass.asInitializers().initEReferences()
    return eClass.eReferences
}

// GetEStructuralFeatures get the value of eStructuralFeatures
func (eClass *eClassImpl) GetEStructuralFeatures() EList[EStructuralFeature] {
    if eClass.eStructuralFeatures == nil {
        eClass.eStructuralFeatures = eClass.asInitializers().initEStructuralFeatures()
    }
    return eClass.eStructuralFeatures
}

// GetESuperTypes get the value of eSuperTypes
func (eClass *eClassImpl) GetESuperTypes() EList[EClass] {
    if eClass.eSuperTypes == nil {
        eClass.eSuperTypes = eClass.asInitializers().initESuperTypes()
    }
    return eClass.eSuperTypes
}

// IsAbstract get the value of isAbstract
func (eClass *eClassImpl) IsAbstract() bool {
    return eClass.isAbstract
}

// SetAbstract set the value of isAbstract
func (eClass *eClassImpl) SetAbstract( newIsAbstract bool ) {
    oldIsAbstract := eClass.isAbstract
    eClass.isAbstract = newIsAbstract
    if eClass.ENotificationRequired() {
        eClass.ENotify(NewNotificationByFeatureID(eClass.AsEObject(), SET, ECLASS__ABSTRACT, oldIsAbstract, newIsAbstract, NO_INDEX))
    }
}


// IsInterface get the value of isInterface
func (eClass *eClassImpl) IsInterface() bool {
    return eClass.isInterface
}

// SetInterface set the value of isInterface
func (eClass *eClassImpl) SetInterface( newIsInterface bool ) {
    oldIsInterface := eClass.isInterface
    eClass.isInterface = newIsInterface
    if eClass.ENotificationRequired() {
        eClass.ENotify(NewNotificationByFeatureID(eClass.AsEObject(), SET, ECLASS__INTERFACE, oldIsInterface, newIsInterface, NO_INDEX))
    }
}



func (eClass *eClassImpl) initEAllAttributes() {
    eClass.eAllAttributes = NewBasicEObjectList[EAttribute](eClass.AsEObjectInternal() , ECLASS__EALL_ATTRIBUTES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEAllContainments() {
    eClass.eAllContainments = NewBasicEObjectList[EReference](eClass.AsEObjectInternal() , ECLASS__EALL_CONTAINMENTS , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEAllOperations() {
    eClass.eAllOperations = NewBasicEObjectList[EOperation](eClass.AsEObjectInternal() , ECLASS__EALL_OPERATIONS , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEAllReferences() {
    eClass.eAllReferences = NewBasicEObjectList[EReference](eClass.AsEObjectInternal() , ECLASS__EALL_REFERENCES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEAllStructuralFeatures() {
    eClass.eAllStructuralFeatures = NewBasicEObjectList[EStructuralFeature](eClass.AsEObjectInternal() , ECLASS__EALL_STRUCTURAL_FEATURES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEAllSuperTypes() {
    eClass.eAllSuperTypes = NewBasicEObjectList[EClass](eClass.AsEObjectInternal() , ECLASS__EALL_SUPER_TYPES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEAttributes() {
    eClass.eAttributes = NewBasicEObjectList[EAttribute](eClass.AsEObjectInternal() , ECLASS__EATTRIBUTES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEContainmentFeatures() {
    eClass.eContainmentFeatures = NewBasicEObjectList[EStructuralFeature](eClass.AsEObjectInternal() , ECLASS__ECONTAINMENT_FEATURES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initECrossReferenceFeatures() {
    eClass.eCrossReferenceFeatures = NewBasicEObjectList[EStructuralFeature](eClass.AsEObjectInternal() , ECLASS__ECROSS_REFERENCE_FEATURES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEIDAttribute() {
    panic("initEIDAttribute not implemented")
}

func (eClass *eClassImpl) initEOperations() EList[EOperation] {
    return NewBasicEObjectList[EOperation](eClass.AsEObjectInternal() , ECLASS__EOPERATIONS , EOPERATION__ECONTAINING_CLASS, true , true , true , false , false );
}

func (eClass *eClassImpl) initEReferences() {
    eClass.eReferences = NewBasicEObjectList[EReference](eClass.AsEObjectInternal() , ECLASS__EREFERENCES , -1 , false , false , false , true , false );
}

func (eClass *eClassImpl) initEStructuralFeatures() EList[EStructuralFeature] {
    return NewBasicEObjectList[EStructuralFeature](eClass.AsEObjectInternal() , ECLASS__ESTRUCTURAL_FEATURES , ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, true , true , true , false , false );
}

func (eClass *eClassImpl) initESuperTypes() EList[EClass] {
    return NewBasicEObjectList[EClass](eClass.AsEObjectInternal() , ECLASS__ESUPER_TYPES , -1 , false , false , false , true , false );
}


func (eClass *eClassImpl) EGetFromID(featureID int, resolve bool) any {
    switch featureID {
    case ECLASS__ABSTRACT:
        return eClass.asEClass().IsAbstract()
    case ECLASS__EALL_ATTRIBUTES:
		list := eClass.asEClass().GetEAllAttributes();
		if !resolve {
			if objects , _ := list.(EObjectList[EAttribute]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EALL_CONTAINMENTS:
		list := eClass.asEClass().GetEAllContainments();
		if !resolve {
			if objects , _ := list.(EObjectList[EReference]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EALL_OPERATIONS:
		list := eClass.asEClass().GetEAllOperations();
		if !resolve {
			if objects , _ := list.(EObjectList[EOperation]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EALL_REFERENCES:
		list := eClass.asEClass().GetEAllReferences();
		if !resolve {
			if objects , _ := list.(EObjectList[EReference]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EALL_STRUCTURAL_FEATURES:
		list := eClass.asEClass().GetEAllStructuralFeatures();
		if !resolve {
			if objects , _ := list.(EObjectList[EStructuralFeature]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EALL_SUPER_TYPES:
		list := eClass.asEClass().GetEAllSuperTypes();
		if !resolve {
			if objects , _ := list.(EObjectList[EClass]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EATTRIBUTES:
		list := eClass.asEClass().GetEAttributes();
		if !resolve {
			if objects , _ := list.(EObjectList[EAttribute]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__ECONTAINMENT_FEATURES:
		list := eClass.asEClass().GetEContainmentFeatures();
		if !resolve {
			if objects , _ := list.(EObjectList[EStructuralFeature]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__ECROSS_REFERENCE_FEATURES:
		list := eClass.asEClass().GetECrossReferenceFeatures();
		if !resolve {
			if objects , _ := list.(EObjectList[EStructuralFeature]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__EID_ATTRIBUTE:
        return eClass.asEClass().GetEIDAttribute()
    case ECLASS__EOPERATIONS:
        return eClass.asEClass().GetEOperations()
    case ECLASS__EREFERENCES:
		list := eClass.asEClass().GetEReferences();
		if !resolve {
			if objects , _ := list.(EObjectList[EReference]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__ESTRUCTURAL_FEATURES:
        return eClass.asEClass().GetEStructuralFeatures()
    case ECLASS__ESUPER_TYPES:
		list := eClass.asEClass().GetESuperTypes();
		if !resolve {
			if objects , _ := list.(EObjectList[EClass]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case ECLASS__INTERFACE:
        return eClass.asEClass().IsInterface()
    default:
        return eClass.eClassifierExt.EGetFromID(featureID, resolve)
    }
}

func (eClass *eClassImpl) ESetFromID(featureID int, newValue any) {
    switch featureID {
    case ECLASS__ABSTRACT:
        eClass.asEClass().SetAbstract(newValue.(bool))
    case ECLASS__EOPERATIONS:
		list := eClass.asEClass().GetEOperations()
        list.Clear()
        list.AddAll(newValue.(EList[EOperation]))
    case ECLASS__ESTRUCTURAL_FEATURES:
		list := eClass.asEClass().GetEStructuralFeatures()
        list.Clear()
        list.AddAll(newValue.(EList[EStructuralFeature]))
    case ECLASS__ESUPER_TYPES:
		list := eClass.asEClass().GetESuperTypes()
        list.Clear()
        list.AddAll(newValue.(EList[EClass]))
    case ECLASS__INTERFACE:
        eClass.asEClass().SetInterface(newValue.(bool))
    default:
        eClass.eClassifierExt.ESetFromID(featureID, newValue)
    }
}

func (eClass *eClassImpl) EUnsetFromID(featureID int) {
    switch featureID {
    case ECLASS__ABSTRACT:
        eClass.asEClass().SetAbstract(false)
    case ECLASS__EOPERATIONS:
        eClass.asEClass().GetEOperations().Clear()
    case ECLASS__ESTRUCTURAL_FEATURES:
        eClass.asEClass().GetEStructuralFeatures().Clear()
    case ECLASS__ESUPER_TYPES:
        eClass.asEClass().GetESuperTypes().Clear()
    case ECLASS__INTERFACE:
        eClass.asEClass().SetInterface(false)
    default:
        eClass.eClassifierExt.EUnsetFromID(featureID)
    }
}

func (eClass *eClassImpl) EIsSetFromID(featureID int) bool {
    switch featureID {
    case ECLASS__ABSTRACT:
        return eClass.isAbstract != false
    case ECLASS__EALL_ATTRIBUTES:
        return eClass.eAllAttributes != nil && eClass.eAllAttributes.Size() != 0
    case ECLASS__EALL_CONTAINMENTS:
        return eClass.eAllContainments != nil && eClass.eAllContainments.Size() != 0
    case ECLASS__EALL_OPERATIONS:
        return eClass.eAllOperations != nil && eClass.eAllOperations.Size() != 0
    case ECLASS__EALL_REFERENCES:
        return eClass.eAllReferences != nil && eClass.eAllReferences.Size() != 0
    case ECLASS__EALL_STRUCTURAL_FEATURES:
        return eClass.eAllStructuralFeatures != nil && eClass.eAllStructuralFeatures.Size() != 0
    case ECLASS__EALL_SUPER_TYPES:
        return eClass.eAllSuperTypes != nil && eClass.eAllSuperTypes.Size() != 0
    case ECLASS__EATTRIBUTES:
        return eClass.eAttributes != nil && eClass.eAttributes.Size() != 0
    case ECLASS__ECONTAINMENT_FEATURES:
        return eClass.eContainmentFeatures != nil && eClass.eContainmentFeatures.Size() != 0
    case ECLASS__ECROSS_REFERENCE_FEATURES:
        return eClass.eCrossReferenceFeatures != nil && eClass.eCrossReferenceFeatures.Size() != 0
    case ECLASS__EID_ATTRIBUTE:
        return eClass.eIDAttribute != nil
    case ECLASS__EOPERATIONS:
        return eClass.eOperations != nil && eClass.eOperations.Size() != 0
    case ECLASS__EREFERENCES:
        return eClass.eReferences != nil && eClass.eReferences.Size() != 0
    case ECLASS__ESTRUCTURAL_FEATURES:
        return eClass.eStructuralFeatures != nil && eClass.eStructuralFeatures.Size() != 0
    case ECLASS__ESUPER_TYPES:
        return eClass.eSuperTypes != nil && eClass.eSuperTypes.Size() != 0
    case ECLASS__INTERFACE:
        return eClass.isInterface != false
    default:
        return eClass.eClassifierExt.EIsSetFromID(featureID)
    }
}

func (eClass *eClassImpl) EInvokeFromID(operationID int, arguments EList[any]) any {
    switch operationID {
    case ECLASS__GET_EOPERATION_EINT:
        return eClass.asEClass().GetEOperation(arguments.Get(0).(int)) 
    case ECLASS__GET_ESTRUCTURAL_FEATURE_EINT:
        return eClass.asEClass().GetEStructuralFeature(arguments.Get(0).(int)) 
    case ECLASS__GET_ESTRUCTURAL_FEATURE_ESTRING:
        return eClass.asEClass().GetEStructuralFeatureFromName(arguments.Get(0).(string)) 
    case ECLASS__GET_FEATURE_COUNT:
        return eClass.asEClass().GetFeatureCount() 
    case ECLASS__GET_FEATURE_ID_ESTRUCTURALFEATURE:
        return eClass.asEClass().GetFeatureID(arguments.Get(0).(EStructuralFeature)) 
    case ECLASS__GET_FEATURE_TYPE_ESTRUCTURALFEATURE:
        return eClass.asEClass().GetFeatureType(arguments.Get(0).(EStructuralFeature)) 
    case ECLASS__GET_OPERATION_COUNT:
        return eClass.asEClass().GetOperationCount() 
    case ECLASS__GET_OPERATION_ID_EOPERATION:
        return eClass.asEClass().GetOperationID(arguments.Get(0).(EOperation)) 
    case ECLASS__GET_OVERRIDE_EOPERATION:
        return eClass.asEClass().GetOverride(arguments.Get(0).(EOperation)) 
    case ECLASS__IS_SUPER_TYPE_OF_ECLASS:
        return eClass.asEClass().IsSuperTypeOf(arguments.Get(0).(EClass)) 
    default:
        return eClass.eClassifierExt.EInvokeFromID(operationID,arguments)
    }
}


func (eClass *eClassImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case ECLASS__EOPERATIONS:
        list := eClass.GetEOperations().(ENotifyingList[EOperation])
		end := otherEnd.(EOperation)
		return list.AddWithNotification(end, notifications)
    case ECLASS__ESTRUCTURAL_FEATURES:
        list := eClass.GetEStructuralFeatures().(ENotifyingList[EStructuralFeature])
		end := otherEnd.(EStructuralFeature)
		return list.AddWithNotification(end, notifications)
    default:
        return eClass.eClassifierExt.EBasicInverseAdd(otherEnd, featureID, notifications)
    }
}


func (eClass *eClassImpl) EBasicInverseRemove(otherEnd  EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case ECLASS__EOPERATIONS:
        list := eClass.GetEOperations().(ENotifyingList[EOperation])
		end := otherEnd.(EOperation)
        return list.RemoveWithNotification(end, notifications)
    case ECLASS__ESTRUCTURAL_FEATURES:
        list := eClass.GetEStructuralFeatures().(ENotifyingList[EStructuralFeature])
		end := otherEnd.(EStructuralFeature)
        return list.RemoveWithNotification(end, notifications)
    default:
        return eClass.eClassifierExt.EBasicInverseRemove(otherEnd, featureID, notifications)
    }
}


