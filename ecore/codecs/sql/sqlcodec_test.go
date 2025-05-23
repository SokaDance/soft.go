package sql

import (
	"testing"

	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/require"
)

func TestGetSQLCodecFeatureKind_Transient(t *testing.T) {
	mockFeature := ecore.NewMockEStructuralFeature(t)
	mockFeature.EXPECT().IsTransient().Return(true).Once()
	require.Equal(t, sfkTransient, GetSQLCodecFeatureKind(mockFeature))
	mockFeature.EXPECT().IsTransient().Return(false).Once()
	require.Equal(t, SQLFeatureKind(-1), GetSQLCodecFeatureKind(mockFeature))
}

func TestGetSQLCodecFeatureKind_Attribute(t *testing.T) {
	mockAttribute := ecore.NewMockEAttribute(t)
	mockDataType := ecore.NewMockEDataType(t)
	mockEnumType := ecore.NewMockEEnum(t)

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, sfkDataList, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("").Once()
	require.Equal(t, sfkData, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("float64").Once()
	require.Equal(t, sfkFloat64, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("float32").Once()
	require.Equal(t, sfkFloat32, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("int").Once()
	require.Equal(t, sfkInt, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("int64").Once()
	require.Equal(t, sfkInt64, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("int32").Once()
	require.Equal(t, sfkInt32, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("int16").Once()
	require.Equal(t, sfkInt16, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("byte").Once()
	require.Equal(t, sfkByte, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("bool").Once()
	require.Equal(t, sfkBool, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("string").Once()
	require.Equal(t, sfkString, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("[]byte").Once()
	require.Equal(t, sfkByteArray, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockDataType).Once()
	mockDataType.EXPECT().GetInstanceTypeName().Return("java.util.Date").Once()
	require.Equal(t, sfkDate, GetSQLCodecFeatureKind(mockAttribute))

	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	mockAttribute.EXPECT().GetEAttributeType().Return(mockEnumType).Once()
	require.Equal(t, sfkEnum, GetSQLCodecFeatureKind(mockAttribute))
}

func TestGetSQLCodecFeatureKind_Reference(t *testing.T) {
	mockReference := ecore.NewMockEReference(t)

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, sfkObject, GetSQLCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, sfkObjectList, GetSQLCodecFeatureKind(mockReference))

	mockOpposite := ecore.NewMockEReference(t)
	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(mockOpposite).Once()
	mockOpposite.EXPECT().IsContainment().Return(true).Once()
	require.Equal(t, sfkTransient, GetSQLCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, sfkObjectReferenceList, GetSQLCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, sfkObjectReference, GetSQLCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(false).Once()
	mockReference.EXPECT().IsContainer().Return(true).Once()
	require.Equal(t, sfkTransient, GetSQLCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(false).Once()
	mockReference.EXPECT().IsContainer().Return(false).Once()
	mockReference.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, sfkObject, GetSQLCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(false).Once()
	mockReference.EXPECT().IsContainer().Return(false).Once()
	mockReference.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, sfkObjectList, GetSQLCodecFeatureKind(mockReference))
}
