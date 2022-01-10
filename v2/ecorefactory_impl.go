// Code generated by soft.generator.go. DO NOT EDIT.



package ecore

import "strconv"
import "time"

type ecoreFactoryInternal interface {
    createEBigDecimalFromString(eDataType EDataType, literalValue string) any
    createEBigIntegerFromString(eDataType EDataType, literalValue string) any
    createEBooleanFromString(eDataType EDataType, literalValue string) any
    createEBooleanObjectFromString(eDataType EDataType, literalValue string) any
    createEByteFromString(eDataType EDataType, literalValue string) any
    createEByteArrayFromString(eDataType EDataType, literalValue string) any
    createEByteObjectFromString(eDataType EDataType, literalValue string) any
    createECharFromString(eDataType EDataType, literalValue string) any
    createECharacterObjectFromString(eDataType EDataType, literalValue string) any
    createEDateFromString(eDataType EDataType, literalValue string) any
    createEDoubleFromString(eDataType EDataType, literalValue string) any
    createEDoubleObjectFromString(eDataType EDataType, literalValue string) any
    createEFloatFromString(eDataType EDataType, literalValue string) any
    createEFloatObjectFromString(eDataType EDataType, literalValue string) any
    createEIntFromString(eDataType EDataType, literalValue string) any
    createEIntegerObjectFromString(eDataType EDataType, literalValue string) any
    createEJavaClassFromString(eDataType EDataType, literalValue string) any
    createEJavaObjectFromString(eDataType EDataType, literalValue string) any
    createELongFromString(eDataType EDataType, literalValue string) any
    createELongObjectFromString(eDataType EDataType, literalValue string) any
    createEShortFromString(eDataType EDataType, literalValue string) any
    createEShortObjectFromString(eDataType EDataType, literalValue string) any
    createEStringFromString(eDataType EDataType, literalValue string) any
    convertEBigDecimalToString(eDataType EDataType, literalValue any) string
    convertEBigIntegerToString(eDataType EDataType, literalValue any) string
    convertEBooleanToString(eDataType EDataType, literalValue any) string
    convertEBooleanObjectToString(eDataType EDataType, literalValue any) string
    convertEByteToString(eDataType EDataType, literalValue any) string
    convertEByteArrayToString(eDataType EDataType, literalValue any) string
    convertEByteObjectToString(eDataType EDataType, literalValue any) string
    convertECharToString(eDataType EDataType, literalValue any) string
    convertECharacterObjectToString(eDataType EDataType, literalValue any) string
    convertEDateToString(eDataType EDataType, literalValue any) string
    convertEDoubleToString(eDataType EDataType, literalValue any) string
    convertEDoubleObjectToString(eDataType EDataType, literalValue any) string
    convertEFloatToString(eDataType EDataType, literalValue any) string
    convertEFloatObjectToString(eDataType EDataType, literalValue any) string
    convertEIntToString(eDataType EDataType, literalValue any) string
    convertEIntegerObjectToString(eDataType EDataType, literalValue any) string
    convertEJavaClassToString(eDataType EDataType, literalValue any) string
    convertEJavaObjectToString(eDataType EDataType, literalValue any) string
    convertELongToString(eDataType EDataType, literalValue any) string
    convertELongObjectToString(eDataType EDataType, literalValue any) string
    convertEShortToString(eDataType EDataType, literalValue any) string
    convertEShortObjectToString(eDataType EDataType, literalValue any) string
    convertEStringToString(eDataType EDataType, literalValue any) string
}

type ecoreFactoryImpl struct {
    EFactoryExt
}

func newEcoreFactoryImpl() *ecoreFactoryImpl {
    factory := new(ecoreFactoryImpl)
    factory.SetInterfaces(factory)
	factory.Initialize()
    return factory 
}

func (ecoreFactoryImpl *ecoreFactoryImpl) Create(eClass EClass) EObject {
	classID := eClass.GetClassifierID()
	switch classID {
    case EANNOTATION:
        return ecoreFactoryImpl.CreateEAnnotation()
    case EATTRIBUTE:
        return ecoreFactoryImpl.CreateEAttribute()
    case ECLASS:
        return ecoreFactoryImpl.CreateEClass()
    case EDATA_TYPE:
        return ecoreFactoryImpl.CreateEDataType()
    case EENUM:
        return ecoreFactoryImpl.CreateEEnum()
    case EENUM_LITERAL:
        return ecoreFactoryImpl.CreateEEnumLiteral()
    case EFACTORY:
        return ecoreFactoryImpl.CreateEFactory()
    case EGENERIC_TYPE:
        return ecoreFactoryImpl.CreateEGenericType()
    case EOBJECT:
        return ecoreFactoryImpl.CreateEObject()
    case EOPERATION:
        return ecoreFactoryImpl.CreateEOperation()
    case EPACKAGE:
        return ecoreFactoryImpl.CreateEPackage()
    case EPARAMETER:
        return ecoreFactoryImpl.CreateEParameter()
    case EREFERENCE:
        return ecoreFactoryImpl.CreateEReference()
    case ESTRING_TO_STRING_MAP_ENTRY:
        return ecoreFactoryImpl.CreateEStringToStringMapEntry()
    case ETYPE_PARAMETER:
        return ecoreFactoryImpl.CreateETypeParameter()
    default:
        panic("Create: " + strconv.Itoa( classID ) + " not found")
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAnnotation() EAnnotation {
	return newEAnnotationImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAnnotationFromContainer(eContainer EModelElement) EAnnotation {
    element := newEAnnotationImpl()
    if eContainer != nil {
        eContainer.GetEAnnotations().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAttribute() EAttribute {
	return newEAttributeExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAttributeFromContainer(eContainer EClass) EAttribute {
    element := newEAttributeExt()
    if eContainer != nil {
        eContainer.GetEStructuralFeatures().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEAttributeFromContainerAndClassID(eContainer EClass, classID int) EAttribute {
    element := newEAttributeExt()
    element.SetFeatureID(classID)
    if eContainer != nil {
        eContainer.GetEStructuralFeatures().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEClass() EClass {
	return newEClassExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEClassFromContainer(eContainer EPackage) EClass {
    element := newEClassExt()
    if eContainer != nil {
        eContainer.GetEClassifiers().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEClassFromContainerAndClassID(eContainer EPackage, classID int) EClass {
    element := newEClassExt()
    element.SetClassifierID(classID)
    if eContainer != nil {
        eContainer.GetEClassifiers().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEDataType() EDataType {
	return newEDataTypeExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEDataTypeFromContainer(eContainer EPackage) EDataType {
    element := newEDataTypeExt()
    if eContainer != nil {
        eContainer.GetEClassifiers().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEDataTypeFromContainerAndClassID(eContainer EPackage, classID int) EDataType {
    element := newEDataTypeExt()
    element.SetClassifierID(classID)
    if eContainer != nil {
        eContainer.GetEClassifiers().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnum() EEnum {
	return newEEnumExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumFromContainer(eContainer EPackage) EEnum {
    element := newEEnumExt()
    if eContainer != nil {
        eContainer.GetEClassifiers().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumFromContainerAndClassID(eContainer EPackage, classID int) EEnum {
    element := newEEnumExt()
    element.SetClassifierID(classID)
    if eContainer != nil {
        eContainer.GetEClassifiers().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumLiteral() EEnumLiteral {
	return newEEnumLiteralExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEEnumLiteralFromContainer(eContainer EEnum) EEnumLiteral {
    element := newEEnumLiteralExt()
    if eContainer != nil {
        eContainer.GetELiterals().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEFactory() EFactory {
	return NewEFactoryExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEFactoryFromContainer(eContainer EPackage) EFactory {
    element := NewEFactoryExt()
    if eContainer != nil {
        eContainer.SetEFactoryInstance(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEGenericType() EGenericType {
	return newEGenericTypeImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEObject() EObject {
	return NewEObjectImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEOperation() EOperation {
	return newEOperationExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEOperationFromContainer(eContainer EClass) EOperation {
    element := newEOperationExt()
    if eContainer != nil {
        eContainer.GetEOperations().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEOperationFromContainerAndClassID(eContainer EClass, classID int) EOperation {
    element := newEOperationExt()
    element.SetOperationID(classID)
    if eContainer != nil {
        eContainer.GetEOperations().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEPackage() EPackage {
	return NewEPackageExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEPackageFromContainer(eContainer EPackage) EPackage {
    element := NewEPackageExt()
    if eContainer != nil {
        eContainer.GetESubPackages().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEParameter() EParameter {
	return newEParameterImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEParameterFromContainer(eContainer EOperation) EParameter {
    element := newEParameterImpl()
    if eContainer != nil {
        eContainer.GetEParameters().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEReference() EReference {
	return newEReferenceExt()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEReferenceFromContainer(eContainer EClass) EReference {
    element := newEReferenceExt()
    if eContainer != nil {
        eContainer.GetEStructuralFeatures().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEReferenceFromContainerAndClassID(eContainer EClass, classID int) EReference {
    element := newEReferenceExt()
    element.SetFeatureID(classID)
    if eContainer != nil {
        eContainer.GetEStructuralFeatures().Add(element)
    }
    return element
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateEStringToStringMapEntry() EStringToStringMapEntry {
	return newEStringToStringMapEntryImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateETypeParameter() ETypeParameter {
	return newETypeParameterImpl()
}
func (ecoreFactoryImpl *ecoreFactoryImpl) CreateFromString(eDataType EDataType, literalValue string) any {
	classID  := eDataType.GetClassifierID()
    internal := ecoreFactoryImpl.GetInterfaces().(ecoreFactoryInternal)
	switch classID {
    case EBIG_DECIMAL:
        return internal.createEBigDecimalFromString(eDataType, literalValue)
    case EBIG_INTEGER:
        return internal.createEBigIntegerFromString(eDataType, literalValue)
    case EBOOLEAN:
        return internal.createEBooleanFromString(eDataType, literalValue)
    case EBOOLEAN_OBJECT:
        return internal.createEBooleanObjectFromString(eDataType, literalValue)
    case EBYTE:
        return internal.createEByteFromString(eDataType, literalValue)
    case EBYTE_ARRAY:
        return internal.createEByteArrayFromString(eDataType, literalValue)
    case EBYTE_OBJECT:
        return internal.createEByteObjectFromString(eDataType, literalValue)
    case ECHAR:
        return internal.createECharFromString(eDataType, literalValue)
    case ECHARACTER_OBJECT:
        return internal.createECharacterObjectFromString(eDataType, literalValue)
    case EDATE:
        return internal.createEDateFromString(eDataType, literalValue)
    case EDOUBLE:
        return internal.createEDoubleFromString(eDataType, literalValue)
    case EDOUBLE_OBJECT:
        return internal.createEDoubleObjectFromString(eDataType, literalValue)
    case EFLOAT:
        return internal.createEFloatFromString(eDataType, literalValue)
    case EFLOAT_OBJECT:
        return internal.createEFloatObjectFromString(eDataType, literalValue)
    case EINT:
        return internal.createEIntFromString(eDataType, literalValue)
    case EINTEGER_OBJECT:
        return internal.createEIntegerObjectFromString(eDataType, literalValue)
    case EJAVA_CLASS:
        return internal.createEJavaClassFromString(eDataType, literalValue)
    case EJAVA_OBJECT:
        return internal.createEJavaObjectFromString(eDataType, literalValue)
    case ELONG:
        return internal.createELongFromString(eDataType, literalValue)
    case ELONG_OBJECT:
        return internal.createELongObjectFromString(eDataType, literalValue)
    case ESHORT:
        return internal.createEShortFromString(eDataType, literalValue)
    case ESHORT_OBJECT:
        return internal.createEShortObjectFromString(eDataType, literalValue)
    case ESTRING:
        return internal.createEStringFromString(eDataType, literalValue)
    default:
        panic("The datatype '" + eDataType.GetName() + "' is not a valid classifier")
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) ConvertToString(eDataType EDataType, instanceValue any) string {
    classID := eDataType.GetClassifierID()
    internal := ecoreFactoryImpl.GetInterfaces().(ecoreFactoryInternal)
    switch classID {
    case EBIG_DECIMAL:
        return internal.convertEBigDecimalToString(eDataType, instanceValue)
    case EBIG_INTEGER:
        return internal.convertEBigIntegerToString(eDataType, instanceValue)
    case EBOOLEAN:
        return internal.convertEBooleanToString(eDataType, instanceValue)
    case EBOOLEAN_OBJECT:
        return internal.convertEBooleanObjectToString(eDataType, instanceValue)
    case EBYTE:
        return internal.convertEByteToString(eDataType, instanceValue)
    case EBYTE_ARRAY:
        return internal.convertEByteArrayToString(eDataType, instanceValue)
    case EBYTE_OBJECT:
        return internal.convertEByteObjectToString(eDataType, instanceValue)
    case ECHAR:
        return internal.convertECharToString(eDataType, instanceValue)
    case ECHARACTER_OBJECT:
        return internal.convertECharacterObjectToString(eDataType, instanceValue)
    case EDATE:
        return internal.convertEDateToString(eDataType, instanceValue)
    case EDOUBLE:
        return internal.convertEDoubleToString(eDataType, instanceValue)
    case EDOUBLE_OBJECT:
        return internal.convertEDoubleObjectToString(eDataType, instanceValue)
    case EFLOAT:
        return internal.convertEFloatToString(eDataType, instanceValue)
    case EFLOAT_OBJECT:
        return internal.convertEFloatObjectToString(eDataType, instanceValue)
    case EINT:
        return internal.convertEIntToString(eDataType, instanceValue)
    case EINTEGER_OBJECT:
        return internal.convertEIntegerObjectToString(eDataType, instanceValue)
    case EJAVA_CLASS:
        return internal.convertEJavaClassToString(eDataType, instanceValue)
    case EJAVA_OBJECT:
        return internal.convertEJavaObjectToString(eDataType, instanceValue)
    case ELONG:
        return internal.convertELongToString(eDataType, instanceValue)
    case ELONG_OBJECT:
        return internal.convertELongObjectToString(eDataType, instanceValue)
    case ESHORT:
        return internal.convertEShortToString(eDataType, instanceValue)
    case ESHORT_OBJECT:
        return internal.convertEShortObjectToString(eDataType, instanceValue)
    case ESTRING:
        return internal.convertEStringToString(eDataType, instanceValue)
    default:
        panic("The datatype '" + eDataType.GetName() + "' is not a valid classifier")
    }
}
func (ecoreFactoryImpl *ecoreFactoryImpl) createEBigDecimalFromString(eDataType EDataType, literalValue string) any {
   	value, _ := strconv.ParseFloat(literalValue, 64)
	return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBigDecimalToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(float64)
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEBigIntegerFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.ParseInt(literalValue, 10, 64)
   return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBigIntegerToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int64)
	return strconv.FormatInt(v, 10)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEBooleanFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.ParseBool(literalValue)
	return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBooleanToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(bool)
	return strconv.FormatBool(v)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEBooleanObjectFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.ParseBool(literalValue)
	return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEBooleanObjectToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(bool)
	return strconv.FormatBool(v)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEByteFromString(eDataType EDataType, literalValue string) any {
   	if len(literalValue) == 0 {
		return "golang\u0000"
   	} else {
	  	return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEByteToString(eDataType EDataType, instanceValue any) string {
   	b := instanceValue.(byte)
	return string([]byte{b})
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEByteArrayFromString(eDataType EDataType, literalValue string) any {
   	return []byte(literalValue)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEByteArrayToString(eDataType EDataType, instanceValue any) string {
   	b := instanceValue.([]byte)
	return string(b)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEByteObjectFromString(eDataType EDataType, literalValue string) any {
   	if len(literalValue) == 0 {
		return "golang\u0000"
   	} else {
	  	return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEByteObjectToString(eDataType EDataType, instanceValue any) string {
   	b := instanceValue.(byte)
	return string([]byte{b})
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createECharFromString(eDataType EDataType, literalValue string) any {
   	if len(literalValue) == 0 {
		return "golang\u0000"
   	} else {
	  	return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertECharToString(eDataType EDataType, instanceValue any) string {
   	b := instanceValue.(byte)
	return string([]byte{b})
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createECharacterObjectFromString(eDataType EDataType, literalValue string) any {
   	if len(literalValue) == 0 {
		return "golang\u0000"
   	} else {
	  	return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertECharacterObjectToString(eDataType EDataType, instanceValue any) string {
   	b := instanceValue.(byte)
	return string([]byte{b})
}

const (
	dateFormat string = "2006-01-02T15:04:05.999Z"
)

func (ecoreFactoryImpl *ecoreFactoryImpl) createEDateFromString(eDataType EDataType, literalValue string) any {
   	t, _ := time.Parse(dateFormat, literalValue)
	return &t
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEDateToString(eDataType EDataType, instanceValue any) string { 
	t , _ := instanceValue.(*time.Time)
	return t.Format(dateFormat)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEDoubleFromString(eDataType EDataType, literalValue string) any {
   	value, _ := strconv.ParseFloat(literalValue, 64)
	return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEDoubleToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(float64)
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEDoubleObjectFromString(eDataType EDataType, literalValue string) any {
   	value, _ := strconv.ParseFloat(literalValue, 64)
	return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEDoubleObjectToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(float64)
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEFloatFromString(eDataType EDataType, literalValue string) any {
   	value, _ := strconv.ParseFloat(literalValue, 32)
	return float32(value)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEFloatToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(float32)
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEFloatObjectFromString(eDataType EDataType, literalValue string) any {
   	value, _ := strconv.ParseFloat(literalValue, 32)
	return float32(value)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEFloatObjectToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(float32)
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEIntFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.Atoi(literalValue)
   return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEIntToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int)
	return strconv.Itoa(v)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEIntegerObjectFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.Atoi(literalValue)
   return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEIntegerObjectToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int)
	return strconv.Itoa(v)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEJavaClassFromString(eDataType EDataType, literalValue string) any {
   panic("NotImplementedException")
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEJavaClassToString(eDataType EDataType, instanceValue any) string { 
	panic("NotImplementedException")
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEJavaObjectFromString(eDataType EDataType, literalValue string) any {
   panic("NotImplementedException")
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEJavaObjectToString(eDataType EDataType, instanceValue any) string { 
	panic("NotImplementedException")
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createELongFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.ParseInt(literalValue, 10, 64)
   return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertELongToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int64)
	return strconv.FormatInt(v, 10)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createELongObjectFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.Atoi(literalValue)
   return value
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertELongObjectToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int)
	return strconv.Itoa(v)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEShortFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.ParseInt(literalValue, 10, 16)
   return int16(value)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEShortToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int16)
	return strconv.FormatInt(int64(v), 10)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEShortObjectFromString(eDataType EDataType, literalValue string) any {
   value, _ := strconv.ParseInt(literalValue, 10, 16)
   return int16(value)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEShortObjectToString(eDataType EDataType, instanceValue any) string {
   	v, _ := instanceValue.(int16)
	return strconv.FormatInt(int64(v), 10)
}

func (ecoreFactoryImpl *ecoreFactoryImpl) createEStringFromString(eDataType EDataType, literalValue string) any {
   return literalValue
}

func (ecoreFactoryImpl *ecoreFactoryImpl) convertEStringToString(eDataType EDataType, instanceValue any) string {
   	v , _ := instanceValue.(string) 
	return v
}

