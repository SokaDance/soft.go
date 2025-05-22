// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package bin

import (
	"io"

	"github.com/masagroup/soft.go/ecore"
)

const (
	BINARY_OPTION_ID_ATTRIBUTE        = "ID_ATTRIBUTE"       // if true, save id attribute of the object
	BINARY_OPTION_NAMESPACE_ATTRIBUTE = "NAMESPACE_ATTIBUTE" // if true, namespaces informations are encoded
)

type BinaryCodec struct {
}

func (bc *BinaryCodec) NewEncoder(resource ecore.EResource, w io.Writer, options map[string]any) ecore.EEncoder {
	return NewBinaryEncoder(resource, w, options)
}
func (bc *BinaryCodec) NewDecoder(resource ecore.EResource, r io.Reader, options map[string]any) ecore.EDecoder {
	return NewBinaryDecoder(resource, r, options)
}

func init() {
	registryCodec := ecore.GetCodecRegistry()
	codecs := registryCodec.GetExtensionToCodecMap()
	codecs["bin"] = &BinaryCodec{}
}

type binaryFeatureKind int

const (
	bfkObjectContainer binaryFeatureKind = iota
	bfkObjectContainerProxy
	bfkObject
	bfkObjectProxy
	bfkObjectList
	bfkObjectListProxy
	bfkObjectContainment
	bfkObjectContainmentProxy
	bfkObjectContainmentList
	bfkObjectContainmentListProxy
	bfkFloat64
	bfkFloat32
	bfkInt
	bfkInt64
	bfkInt32
	bfkInt16
	bfkByte
	bfkBool
	bfkString
	bfkByteArray
	bfkData
	bfkDataList
	bfkEnum
	bfkDate
)

func getBinaryCodecFeatureKind(eFeature ecore.EStructuralFeature) binaryFeatureKind {
	if eReference, _ := eFeature.(ecore.EReference); eReference != nil {
		if eReference.IsContainment() {
			if eReference.IsResolveProxies() {
				if eReference.IsMany() {
					return bfkObjectContainmentListProxy
				} else {
					return bfkObjectContainmentProxy
				}
			} else {
				if eReference.IsMany() {
					return bfkObjectContainmentList
				} else {
					return bfkObjectContainment
				}
			}
		} else if eReference.IsContainer() {
			if eReference.IsResolveProxies() {
				return bfkObjectContainerProxy
			} else {
				return bfkObjectContainer
			}
		} else if eReference.IsResolveProxies() {
			if eReference.IsMany() {
				return bfkObjectListProxy
			} else {
				return bfkObjectProxy
			}
		} else {
			if eReference.IsMany() {
				return bfkObjectList
			} else {
				return bfkObject
			}
		}
	} else if eAttribute, _ := eFeature.(ecore.EAttribute); eAttribute != nil {
		if eAttribute.IsMany() {
			return bfkDataList
		} else {
			eDataType := eAttribute.GetEAttributeType()
			if eEnum, _ := eDataType.(ecore.EEnum); eEnum != nil {
				return bfkEnum
			}

			switch eDataType.GetInstanceTypeName() {
			case "float64", "java.lang.Double", "double":
				return bfkFloat64
			case "float32", "java.lang.Float", "float":
				return bfkFloat32
			case "int", "java.lang.Integer":
				return bfkInt
			case "int64", "java.lang.Long", "java.math.BigInteger", "long":
				return bfkInt64
			case "int32":
				return bfkInt32
			case "int16", "java.lang.Short", "short":
				return bfkInt16
			case "byte":
				return bfkByte
			case "bool", "java.lang.Boolean", "boolean":
				return bfkBool
			case "string", "java.lang.String":
				return bfkString
			case "[]byte", "java.util.ByteArray":
				return bfkByteArray
			case "*time/time.Time", "java.util.Date":
				return bfkDate
			}

			return bfkData
		}
	}
	return -1
}
