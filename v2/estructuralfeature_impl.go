// Code generated by soft.generator.go. DO NOT EDIT.



package ecore

import "reflect"


// eStructuralFeatureImpl is the implementation of the model object 'EStructuralFeature'
type eStructuralFeatureImpl struct {
    eTypedElementExt
    defaultValueLiteral string
    featureID int
    isChangeable bool
    isDerived bool
    isTransient bool
    isUnsettable bool
    isVolatile bool
    
}

// newEStructuralFeatureImpl is the constructor of a eStructuralFeatureImpl
func newEStructuralFeatureImpl() *eStructuralFeatureImpl {
    eStructuralFeature := new(eStructuralFeatureImpl)
	eStructuralFeature.SetInterfaces(eStructuralFeature)
	eStructuralFeature.Initialize()
    return eStructuralFeature
}

func (eStructuralFeature *eStructuralFeatureImpl) Initialize() {
	eStructuralFeature.eTypedElementExt.Initialize()
	eStructuralFeature.defaultValueLiteral = ""
	eStructuralFeature.featureID = -1
	eStructuralFeature.isChangeable = true
	eStructuralFeature.isDerived = false
	eStructuralFeature.isTransient = false
	eStructuralFeature.isUnsettable = false
	eStructuralFeature.isVolatile = false
	
}

func (eStructuralFeature *eStructuralFeatureImpl) asEStructuralFeature() EStructuralFeature {
	return eStructuralFeature.GetInterfaces().(EStructuralFeature)
}


func (eStructuralFeature *eStructuralFeatureImpl) EStaticClass() EClass {
    return GetPackage().GetEStructuralFeature()
}

func (eStructuralFeature *eStructuralFeatureImpl) EStaticFeatureCount() int {
    return ESTRUCTURAL_FEATURE_FEATURE_COUNT
}

// GetContainerClass default implementation
func (eStructuralFeature *eStructuralFeatureImpl) GetContainerClass() reflect.Type {
    panic("GetContainerClass not implemented")
}

// GetDefaultValue get the value of defaultValue
func (eStructuralFeature *eStructuralFeatureImpl) GetDefaultValue() any {
    panic("GetDefaultValue not implemented")
}

// SetDefaultValue set the value of defaultValue
func (eStructuralFeature *eStructuralFeatureImpl) SetDefaultValue( newDefaultValue any ) {
    panic("SetDefaultValue not implemented")
}


// GetDefaultValueLiteral get the value of defaultValueLiteral
func (eStructuralFeature *eStructuralFeatureImpl) GetDefaultValueLiteral() string {
    return eStructuralFeature.defaultValueLiteral
}

// SetDefaultValueLiteral set the value of defaultValueLiteral
func (eStructuralFeature *eStructuralFeatureImpl) SetDefaultValueLiteral( newDefaultValueLiteral string ) {
    oldDefaultValueLiteral := eStructuralFeature.defaultValueLiteral
    eStructuralFeature.defaultValueLiteral = newDefaultValueLiteral
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL, oldDefaultValueLiteral, newDefaultValueLiteral, NO_INDEX))
    }
}


// GetEContainingClass get the value of eContainingClass
func (eStructuralFeature *eStructuralFeatureImpl) GetEContainingClass() EClass {
    if eStructuralFeature.EContainerFeatureID() == ESTRUCTURAL_FEATURE__ECONTAINING_CLASS {
        return eStructuralFeature.EContainer().(EClass);
    } 
    return nil
}

// GetFeatureID get the value of featureID
func (eStructuralFeature *eStructuralFeatureImpl) GetFeatureID() int {
    return eStructuralFeature.featureID
}

// SetFeatureID set the value of featureID
func (eStructuralFeature *eStructuralFeatureImpl) SetFeatureID( newFeatureID int ) {
    oldFeatureID := eStructuralFeature.featureID
    eStructuralFeature.featureID = newFeatureID
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__FEATURE_ID, oldFeatureID, newFeatureID, NO_INDEX))
    }
}


// IsChangeable get the value of isChangeable
func (eStructuralFeature *eStructuralFeatureImpl) IsChangeable() bool {
    return eStructuralFeature.isChangeable
}

// SetChangeable set the value of isChangeable
func (eStructuralFeature *eStructuralFeatureImpl) SetChangeable( newIsChangeable bool ) {
    oldIsChangeable := eStructuralFeature.isChangeable
    eStructuralFeature.isChangeable = newIsChangeable
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__CHANGEABLE, oldIsChangeable, newIsChangeable, NO_INDEX))
    }
}


// IsDerived get the value of isDerived
func (eStructuralFeature *eStructuralFeatureImpl) IsDerived() bool {
    return eStructuralFeature.isDerived
}

// SetDerived set the value of isDerived
func (eStructuralFeature *eStructuralFeatureImpl) SetDerived( newIsDerived bool ) {
    oldIsDerived := eStructuralFeature.isDerived
    eStructuralFeature.isDerived = newIsDerived
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__DERIVED, oldIsDerived, newIsDerived, NO_INDEX))
    }
}


// IsTransient get the value of isTransient
func (eStructuralFeature *eStructuralFeatureImpl) IsTransient() bool {
    return eStructuralFeature.isTransient
}

// SetTransient set the value of isTransient
func (eStructuralFeature *eStructuralFeatureImpl) SetTransient( newIsTransient bool ) {
    oldIsTransient := eStructuralFeature.isTransient
    eStructuralFeature.isTransient = newIsTransient
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__TRANSIENT, oldIsTransient, newIsTransient, NO_INDEX))
    }
}


// IsUnsettable get the value of isUnsettable
func (eStructuralFeature *eStructuralFeatureImpl) IsUnsettable() bool {
    return eStructuralFeature.isUnsettable
}

// SetUnsettable set the value of isUnsettable
func (eStructuralFeature *eStructuralFeatureImpl) SetUnsettable( newIsUnsettable bool ) {
    oldIsUnsettable := eStructuralFeature.isUnsettable
    eStructuralFeature.isUnsettable = newIsUnsettable
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__UNSETTABLE, oldIsUnsettable, newIsUnsettable, NO_INDEX))
    }
}


// IsVolatile get the value of isVolatile
func (eStructuralFeature *eStructuralFeatureImpl) IsVolatile() bool {
    return eStructuralFeature.isVolatile
}

// SetVolatile set the value of isVolatile
func (eStructuralFeature *eStructuralFeatureImpl) SetVolatile( newIsVolatile bool ) {
    oldIsVolatile := eStructuralFeature.isVolatile
    eStructuralFeature.isVolatile = newIsVolatile
    if eStructuralFeature.ENotificationRequired() {
        eStructuralFeature.ENotify(NewNotificationByFeatureID(eStructuralFeature.AsEObject(), SET, ESTRUCTURAL_FEATURE__VOLATILE, oldIsVolatile, newIsVolatile, NO_INDEX))
    }
}




func (eStructuralFeature *eStructuralFeatureImpl) EGetFromID(featureID int, resolve bool) any {
    switch featureID {
    case ESTRUCTURAL_FEATURE__CHANGEABLE:
        return eStructuralFeature.asEStructuralFeature().IsChangeable()
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
        return eStructuralFeature.asEStructuralFeature().GetDefaultValue()
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
        return eStructuralFeature.asEStructuralFeature().GetDefaultValueLiteral()
    case ESTRUCTURAL_FEATURE__DERIVED:
        return eStructuralFeature.asEStructuralFeature().IsDerived()
    case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
        return eStructuralFeature.asEStructuralFeature().GetEContainingClass()
    case ESTRUCTURAL_FEATURE__FEATURE_ID:
        return eStructuralFeature.asEStructuralFeature().GetFeatureID()
    case ESTRUCTURAL_FEATURE__TRANSIENT:
        return eStructuralFeature.asEStructuralFeature().IsTransient()
    case ESTRUCTURAL_FEATURE__UNSETTABLE:
        return eStructuralFeature.asEStructuralFeature().IsUnsettable()
    case ESTRUCTURAL_FEATURE__VOLATILE:
        return eStructuralFeature.asEStructuralFeature().IsVolatile()
    default:
        return eStructuralFeature.eTypedElementExt.EGetFromID(featureID, resolve)
    }
}

func (eStructuralFeature *eStructuralFeatureImpl) ESetFromID(featureID int, newValue any) {
    switch featureID {
    case ESTRUCTURAL_FEATURE__CHANGEABLE:
        eStructuralFeature.asEStructuralFeature().SetChangeable(newValue.(bool))
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
        eStructuralFeature.asEStructuralFeature().SetDefaultValue(newValue)
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
        eStructuralFeature.asEStructuralFeature().SetDefaultValueLiteral(newValue.(string))
    case ESTRUCTURAL_FEATURE__DERIVED:
        eStructuralFeature.asEStructuralFeature().SetDerived(newValue.(bool))
    case ESTRUCTURAL_FEATURE__FEATURE_ID:
        eStructuralFeature.asEStructuralFeature().SetFeatureID(newValue.(int))
    case ESTRUCTURAL_FEATURE__TRANSIENT:
        eStructuralFeature.asEStructuralFeature().SetTransient(newValue.(bool))
    case ESTRUCTURAL_FEATURE__UNSETTABLE:
        eStructuralFeature.asEStructuralFeature().SetUnsettable(newValue.(bool))
    case ESTRUCTURAL_FEATURE__VOLATILE:
        eStructuralFeature.asEStructuralFeature().SetVolatile(newValue.(bool))
    default:
        eStructuralFeature.eTypedElementExt.ESetFromID(featureID, newValue)
    }
}

func (eStructuralFeature *eStructuralFeatureImpl) EUnsetFromID(featureID int) {
    switch featureID {
    case ESTRUCTURAL_FEATURE__CHANGEABLE:
        eStructuralFeature.asEStructuralFeature().SetChangeable(true)
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
        eStructuralFeature.asEStructuralFeature().SetDefaultValue(nil)
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
        eStructuralFeature.asEStructuralFeature().SetDefaultValueLiteral("")
    case ESTRUCTURAL_FEATURE__DERIVED:
        eStructuralFeature.asEStructuralFeature().SetDerived(false)
    case ESTRUCTURAL_FEATURE__FEATURE_ID:
        eStructuralFeature.asEStructuralFeature().SetFeatureID(-1)
    case ESTRUCTURAL_FEATURE__TRANSIENT:
        eStructuralFeature.asEStructuralFeature().SetTransient(false)
    case ESTRUCTURAL_FEATURE__UNSETTABLE:
        eStructuralFeature.asEStructuralFeature().SetUnsettable(false)
    case ESTRUCTURAL_FEATURE__VOLATILE:
        eStructuralFeature.asEStructuralFeature().SetVolatile(false)
    default:
        eStructuralFeature.eTypedElementExt.EUnsetFromID(featureID)
    }
}

func (eStructuralFeature *eStructuralFeatureImpl) EIsSetFromID(featureID int) bool {
    switch featureID {
    case ESTRUCTURAL_FEATURE__CHANGEABLE:
        return eStructuralFeature.isChangeable != true
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
        return eStructuralFeature.GetDefaultValue() != nil
    case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
        return eStructuralFeature.defaultValueLiteral != ""
    case ESTRUCTURAL_FEATURE__DERIVED:
        return eStructuralFeature.isDerived != false
    case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
        return eStructuralFeature.GetEContainingClass() != nil
    case ESTRUCTURAL_FEATURE__FEATURE_ID:
        return eStructuralFeature.featureID != -1
    case ESTRUCTURAL_FEATURE__TRANSIENT:
        return eStructuralFeature.isTransient != false
    case ESTRUCTURAL_FEATURE__UNSETTABLE:
        return eStructuralFeature.isUnsettable != false
    case ESTRUCTURAL_FEATURE__VOLATILE:
        return eStructuralFeature.isVolatile != false
    default:
        return eStructuralFeature.eTypedElementExt.EIsSetFromID(featureID)
    }
}

func (eStructuralFeature *eStructuralFeatureImpl) EInvokeFromID(operationID int, arguments EList[any]) any {
    switch operationID {
    case ESTRUCTURAL_FEATURE__GET_CONTAINER_CLASS:
        return eStructuralFeature.asEStructuralFeature().GetContainerClass() 
    default:
        return eStructuralFeature.eTypedElementExt.EInvokeFromID(operationID,arguments)
    }
}


func (eStructuralFeature *eStructuralFeatureImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
        msgs := notifications
        if eStructuralFeature.EInternalContainer() != nil {
            msgs = eStructuralFeature.EBasicRemoveFromContainer(msgs)
        }
        return eStructuralFeature.EBasicSetContainer(otherEnd, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, msgs)
    default:
        return eStructuralFeature.eTypedElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
    }
}


func (eStructuralFeature *eStructuralFeatureImpl) EBasicInverseRemove(otherEnd  EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
        return eStructuralFeature.EBasicSetContainer(nil, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, notifications)
    default:
        return eStructuralFeature.eTypedElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
    }
}

