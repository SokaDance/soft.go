// Code generated by soft.generator.go. DO NOT EDIT.



package ecore



// eAnnotationImpl is the implementation of the model object 'EAnnotation'
type eAnnotationImpl struct {
    eModelElementExt
    contents EList[EObject]
    details EMap[string,string]
    references EList[EObject]
    source string
    
}
type eAnnotationImplInitializers interface {
    initContents() EList[EObject]
    initDetails() EMap[string,string]
    initReferences() EList[EObject]
    
}

// newEAnnotationImpl is the constructor of a eAnnotationImpl
func newEAnnotationImpl() *eAnnotationImpl {
    eAnnotation := new(eAnnotationImpl)
	eAnnotation.SetInterfaces(eAnnotation)
	eAnnotation.Initialize()
    return eAnnotation
}

func (eAnnotation *eAnnotationImpl) Initialize() {
	eAnnotation.eModelElementExt.Initialize()
	eAnnotation.source = ""
	
}

func (eAnnotation *eAnnotationImpl) asEAnnotation() EAnnotation {
	return eAnnotation.GetInterfaces().(EAnnotation)
}

func (eAnnotation *eAnnotationImpl) asInitializers() eAnnotationImplInitializers {
    return eAnnotation.AsEObject().(eAnnotationImplInitializers)
}

func (eAnnotation *eAnnotationImpl) EStaticClass() EClass {
    return GetPackage().GetEAnnotationClass()
}

func (eAnnotation *eAnnotationImpl) EStaticFeatureCount() int {
    return EANNOTATION_FEATURE_COUNT
}


// GetContents get the value of contents
func (eAnnotation *eAnnotationImpl) GetContents() EList[EObject] {
    if eAnnotation.contents == nil {
        eAnnotation.contents = eAnnotation.asInitializers().initContents()
    }
    return eAnnotation.contents
}

// GetDetails get the value of details
func (eAnnotation *eAnnotationImpl) GetDetails() EMap[string,string] {
    if eAnnotation.details == nil {
        eAnnotation.details = eAnnotation.asInitializers().initDetails()
    }
    return eAnnotation.details
}

// GetEModelElement get the value of eModelElement
func (eAnnotation *eAnnotationImpl) GetEModelElement() EModelElement {
    if eAnnotation.EContainerFeatureID() == EANNOTATION__EMODEL_ELEMENT {
        return eAnnotation.EContainer().(EModelElement);
    } 
    return nil
}

// SetEModelElement set the value of eModelElement
func (eAnnotation *eAnnotationImpl) SetEModelElement( newEModelElement EModelElement ) {
    if ( newEModelElement != eAnnotation.EInternalContainer() || (newEModelElement != nil && eAnnotation.EContainerFeatureID() !=  EANNOTATION__EMODEL_ELEMENT)) {
        var notifications ENotificationChain
        if ( eAnnotation.EInternalContainer() != nil) {
            notifications = eAnnotation.EBasicRemoveFromContainer(notifications)
        }
        if newEModelElementInternal , _ := newEModelElement.(EObjectInternal); newEModelElementInternal != nil {
            notifications = newEModelElementInternal.EInverseAdd( eAnnotation.AsEObject() , EMODEL_ELEMENT__EANNOTATIONS, notifications )
        }
        notifications = eAnnotation.basicSetEModelElement( newEModelElement, notifications )
        if ( notifications != nil ) {
            notifications.Dispatch()
        }
    } else if ( eAnnotation.ENotificationRequired() ) {
        eAnnotation.ENotify( NewNotificationByFeatureID(eAnnotation, SET, EANNOTATION__EMODEL_ELEMENT, newEModelElement, newEModelElement, NO_INDEX) )
    }
}


func (eAnnotation *eAnnotationImpl) basicSetEModelElement( newEModelElement EModelElement , msgs ENotificationChain )  ENotificationChain {
    return eAnnotation.EBasicSetContainer(newEModelElement,EANNOTATION__EMODEL_ELEMENT,msgs) 
}

// GetReferences get the value of references
func (eAnnotation *eAnnotationImpl) GetReferences() EList[EObject] {
    if eAnnotation.references == nil {
        eAnnotation.references = eAnnotation.asInitializers().initReferences()
    }
    return eAnnotation.references
}

// GetSource get the value of source
func (eAnnotation *eAnnotationImpl) GetSource() string {
    return eAnnotation.source
}

// SetSource set the value of source
func (eAnnotation *eAnnotationImpl) SetSource( newSource string ) {
    oldSource := eAnnotation.source
    eAnnotation.source = newSource
    if eAnnotation.ENotificationRequired() {
        eAnnotation.ENotify(NewNotificationByFeatureID(eAnnotation.AsEObject(), SET, EANNOTATION__SOURCE, oldSource, newSource, NO_INDEX))
    }
}



func (eAnnotation *eAnnotationImpl) initContents() EList[EObject] {
    return NewBasicEObjectList[EObject](eAnnotation.AsEObjectInternal() , EANNOTATION__CONTENTS , -1, true , true , false , false , false );
}

func (eAnnotation *eAnnotationImpl) initDetails() EMap[string,string] {
	return NewBasicEObjectMap[string,string](GetPackage().GetEStringToStringMapEntry());
}

func (eAnnotation *eAnnotationImpl) initReferences() EList[EObject] {
    return NewBasicEObjectList[EObject](eAnnotation.AsEObjectInternal() , EANNOTATION__REFERENCES , -1 , false , false , false , true , false );
}


func (eAnnotation *eAnnotationImpl) EGetFromID(featureID int, resolve bool) any {
    switch featureID {
    case EANNOTATION__CONTENTS:
        return eAnnotation.asEAnnotation().GetContents()
    case EANNOTATION__DETAILS:
        return eAnnotation.asEAnnotation().GetDetails()
    case EANNOTATION__EMODEL_ELEMENT:
        return eAnnotation.asEAnnotation().GetEModelElement()
    case EANNOTATION__REFERENCES:
		list := eAnnotation.asEAnnotation().GetReferences();
		if !resolve {
			if objects , _ := list.(EObjectList[EObject]); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
    case EANNOTATION__SOURCE:
        return eAnnotation.asEAnnotation().GetSource()
    default:
        return eAnnotation.eModelElementExt.EGetFromID(featureID, resolve)
    }
}

func (eAnnotation *eAnnotationImpl) ESetFromID(featureID int, newValue any) {
    switch featureID {
    case EANNOTATION__CONTENTS:
		l := eAnnotation.asEAnnotation().GetContents()
        l.Clear()
        l.AddAll(newValue.(EList[EObject]))
    case EANNOTATION__DETAILS:
		m := eAnnotation.asEAnnotation().GetDetails()
        m.Clear()
        m.AddAll(newValue.(EMap[string,string]))
    case EANNOTATION__EMODEL_ELEMENT:
        eAnnotation.asEAnnotation().SetEModelElement(newValue.(EModelElement))
    case EANNOTATION__REFERENCES:
		l := eAnnotation.asEAnnotation().GetReferences()
        l.Clear()
        l.AddAll(newValue.(EList[EObject]))
    case EANNOTATION__SOURCE:
        eAnnotation.asEAnnotation().SetSource(newValue.(string))
    default:
        eAnnotation.eModelElementExt.ESetFromID(featureID, newValue)
    }
}

func (eAnnotation *eAnnotationImpl) EUnsetFromID(featureID int) {
    switch featureID {
    case EANNOTATION__CONTENTS:
        eAnnotation.asEAnnotation().GetContents().Clear()
    case EANNOTATION__DETAILS:
        eAnnotation.asEAnnotation().GetDetails().Clear()
    case EANNOTATION__EMODEL_ELEMENT:
        eAnnotation.asEAnnotation().SetEModelElement(nil)
    case EANNOTATION__REFERENCES:
        eAnnotation.asEAnnotation().GetReferences().Clear()
    case EANNOTATION__SOURCE:
        eAnnotation.asEAnnotation().SetSource("")
    default:
        eAnnotation.eModelElementExt.EUnsetFromID(featureID)
    }
}

func (eAnnotation *eAnnotationImpl) EIsSetFromID(featureID int) bool {
    switch featureID {
    case EANNOTATION__CONTENTS:
        return eAnnotation.contents != nil && eAnnotation.contents.Size() != 0
    case EANNOTATION__DETAILS:
        return eAnnotation.details != nil && eAnnotation.details.Size() != 0
    case EANNOTATION__EMODEL_ELEMENT:
        return eAnnotation.GetEModelElement() != nil
    case EANNOTATION__REFERENCES:
        return eAnnotation.references != nil && eAnnotation.references.Size() != 0
    case EANNOTATION__SOURCE:
        return eAnnotation.source != ""
    default:
        return eAnnotation.eModelElementExt.EIsSetFromID(featureID)
    }
}

func (eAnnotation *eAnnotationImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case EANNOTATION__EMODEL_ELEMENT:
        msgs := notifications
        if eAnnotation.EInternalContainer() != nil {
            msgs = eAnnotation.EBasicRemoveFromContainer(msgs)
        }
        return eAnnotation.basicSetEModelElement(otherEnd.(EModelElement), msgs)
    default:
        return eAnnotation.eModelElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
    }
}


func (eAnnotation *eAnnotationImpl) EBasicInverseRemove(otherEnd  EObject, featureID int, notifications ENotificationChain)  ENotificationChain {
    switch featureID {
    case EANNOTATION__CONTENTS:
        list := eAnnotation.GetContents().(ENotifyingList[EObject])
		end := otherEnd.(EObject)
        return list.RemoveWithNotification(end, notifications)
    case EANNOTATION__DETAILS:
		return notifications
    case EANNOTATION__EMODEL_ELEMENT:
        return eAnnotation.basicSetEModelElement(nil, notifications)
    default:
        return eAnnotation.eModelElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
    }
}


