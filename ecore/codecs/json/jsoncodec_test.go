package json

import (
	"testing"

	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func diagnosticError(errors ecore.EList) string {
	if errors.Empty() {
		return ""
	} else {
		return errors.Get(0).(ecore.EDiagnostic).GetMessage()
	}
}

func loadPackage(packageFileName string) ecore.EPackage {
	xmiProcessor := ecore.NewXMIProcessor()
	eResource := xmiProcessor.Load(ecore.NewURI("testdata/" + packageFileName))
	if eResource.IsLoaded() && eResource.GetContents().Size() > 0 {
		ePackage, _ := eResource.GetContents().Get(0).(ecore.EPackage)
		return ePackage
	} else {
		return nil
	}
}

func loadTestPackage(t *testing.T, resourceSet ecore.EResourceSet, packageURI *ecore.URI) (ecore.EResource, ecore.EPackage) {
	// load package
	r := resourceSet.CreateResource(packageURI)
	r.SetObjectIDManager(ecore.NewIncrementalIDManager())
	r.Load()
	assert.True(t, r.IsLoaded())
	assert.True(t, r.GetErrors().Empty(), diagnosticError(r.GetErrors()))
	assert.True(t, r.GetWarnings().Empty(), diagnosticError(r.GetWarnings()))

	// retrieve package
	ePackage, _ := r.GetContents().Get(0).(ecore.EPackage)
	require.NotNil(t, ePackage)
	resourceSet.GetPackageRegistry().RegisterPackage(ePackage)
	return r, ePackage
}

func loadTestModel(t *testing.T, resourceSet ecore.EResourceSet, modelURI *ecore.URI) (ecore.EResource, ecore.EObject) {
	// load package
	r := resourceSet.CreateResource(modelURI)
	r.SetObjectIDManager(ecore.NewIncrementalIDManager())
	r.Load()
	require.True(t, r.IsLoaded())
	require.True(t, r.GetErrors().Empty(), diagnosticError(r.GetErrors()))
	require.True(t, r.GetWarnings().Empty(), diagnosticError(r.GetWarnings()))
	require.Equal(t, 1, r.GetContents().Size())

	// retrieve root object
	return r, r.GetContents().Get(0).(ecore.EObject)
}

func TestJSONCodec_NewEncoder(t *testing.T) {
	mockResource := ecore.NewMockEResource(t)
	codec := &JSONCodec{}
	require.NotNil(t, codec.NewEncoder(mockResource, nil, nil))
}

func TestJSONCodec_NewDecoder(t *testing.T) {
	require.Panics(t, func() {
		mockResource := ecore.NewMockEResource(t)
		codec := &JSONCodec{}
		codec.NewDecoder(mockResource, nil, nil)
	})
}

func TestGetJSONCodecFeatureKind_Transient(t *testing.T) {
	mockFeature := ecore.NewMockEStructuralFeature(t)
	mockFeature.EXPECT().IsTransient().Return(true).Once()
	require.Equal(t, jfkTransient, getJSONCodecFeatureKind(mockFeature))
	mockFeature.EXPECT().IsTransient().Return(false).Once()
	require.Equal(t, jsonFeatureKind(-1), getJSONCodecFeatureKind(mockFeature))
}

func TestGetJSONCodecFeatureKind_Attribute(t *testing.T) {
	mockAttribute := ecore.NewMockEAttribute(t)
	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, jfkData, getJSONCodecFeatureKind(mockAttribute))
	mockAttribute.EXPECT().IsTransient().Return(false).Once()
	mockAttribute.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, jfkDataList, getJSONCodecFeatureKind(mockAttribute))
}

func TestGetJSONCodecFeatureKind_Reference(t *testing.T) {
	mockReference := ecore.NewMockEReference(t)

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, jfkObject, getJSONCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, jfkObjectList, getJSONCodecFeatureKind(mockReference))

	mockOpposite := ecore.NewMockEReference(t)
	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(mockOpposite).Once()
	mockOpposite.EXPECT().IsContainment().Return(true).Once()
	require.Equal(t, jfkTransient, getJSONCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, jfkObjectReferenceList, getJSONCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(true).Once()
	mockReference.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, jfkObjectReference, getJSONCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(false).Once()
	mockReference.EXPECT().IsMany().Return(false).Once()
	require.Equal(t, jfkObject, getJSONCodecFeatureKind(mockReference))

	mockReference.EXPECT().IsTransient().Return(false).Once()
	mockReference.EXPECT().IsContainment().Return(false).Once()
	mockReference.EXPECT().GetEOpposite().Return(nil).Once()
	mockReference.EXPECT().IsResolveProxies().Return(false).Once()
	mockReference.EXPECT().IsMany().Return(true).Once()
	require.Equal(t, jfkObjectList, getJSONCodecFeatureKind(mockReference))
}
