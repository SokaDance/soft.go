package sql

import "github.com/masagroup/soft.go/ecore"

type SQLFeatureKind int

const (
	sfkTransient SQLFeatureKind = iota
	sfkFloat64
	sfkFloat32
	sfkInt
	sfkInt64
	sfkInt32
	sfkInt16
	sfkByte
	sfkBool
	sfkString
	sfkByteArray
	sfkEnum
	sfkDate
	sfkData
	sfkDataList
	sfkObject
	sfkObjectList
	sfkObjectReference
	sfkObjectReferenceList
)

func GetSQLCodecFeatureKind(eFeature ecore.EStructuralFeature) SQLFeatureKind {
	if eFeature.IsTransient() {
		return sfkTransient
	} else if eReference, _ := eFeature.(ecore.EReference); eReference != nil {
		if eReference.IsContainment() {
			if eReference.IsMany() {
				return sfkObjectList
			} else {
				return sfkObject
			}
		}
		opposite := eReference.GetEOpposite()
		if opposite != nil && opposite.IsContainment() {
			return sfkTransient
		}
		if eReference.IsResolveProxies() {
			if eReference.IsMany() {
				return sfkObjectReferenceList
			} else {
				return sfkObjectReference
			}
		}
		if eReference.IsContainer() {
			return sfkTransient
		}
		if eReference.IsMany() {
			return sfkObjectList
		} else {
			return sfkObject
		}
	} else if eAttribute, _ := eFeature.(ecore.EAttribute); eAttribute != nil {
		if eAttribute.IsMany() {
			return sfkDataList
		} else {
			eDataType := eAttribute.GetEAttributeType()
			if eEnum, _ := eDataType.(ecore.EEnum); eEnum != nil {
				return sfkEnum
			}

			switch eDataType.GetInstanceTypeName() {
			case "float64", "java.lang.Double", "double":
				return sfkFloat64
			case "float32", "java.lang.Float", "float":
				return sfkFloat32
			case "int", "java.lang.Integer":
				return sfkInt
			case "int64", "java.lang.Long", "java.math.BigInteger", "long":
				return sfkInt64
			case "int32":
				return sfkInt32
			case "int16", "java.lang.Short", "short":
				return sfkInt16
			case "byte":
				return sfkByte
			case "bool", "java.lang.Boolean", "boolean":
				return sfkBool
			case "string", "java.lang.String":
				return sfkString
			case "[]byte", "java.util.ByteArray":
				return sfkByteArray
			case "*time/time.Time", "java.util.Date":
				return sfkDate
			}

			return sfkData
		}
	}
	return -1
}
