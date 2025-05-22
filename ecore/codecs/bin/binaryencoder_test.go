package bin

import (
	"bytes"
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinaryEncoder_Complex(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	// load resource
	xmlProcessor := ecore.NewXMLProcessor(ecore.XMLProcessorPackages([]ecore.EPackage{ePackage}))
	eResource := xmlProcessor.LoadWithOptions(ecore.NewURI(path.Join(testdataPath, "library.complex.xml")), nil)
	require.NotNil(t, eResource)
	require.True(t, eResource.IsLoaded())
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	require.True(t, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	w := &bytes.Buffer{}
	binaryEncoder := NewBinaryEncoder(eResource, w, nil)
	binaryEncoder.EncodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	//os.WriteFile(path.Join(testdataPath,"library.complex.bin", w.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "library.complex.bin"))
	assert.Nil(t, err)
	assert.Equal(t, bytes, w.Bytes())
}

func TestBinaryEncoder_ComplexWithID(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	// load resource
	uri := ecore.NewURI(path.Join(testdataPath, "library.complex.id.xml"))
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResource.SetObjectIDManager(ecore.NewUUIDManager())
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)
	eResource.LoadWithOptions(map[string]any{ecore.XML_OPTION_ID_ATTRIBUTE_NAME: "id"})
	require.NotNil(t, eResource)
	require.True(t, eResource.IsLoaded())
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	require.True(t, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	// set DocumentRoot uuid, because it is not defined in xml and is always regenerated
	eDocumentRoot, _ := eResource.GetContents().Get(0).(ecore.EObject)
	require.NotNil(t, eDocumentRoot)
	require.Nil(t, eResource.GetObjectIDManager().SetID(eDocumentRoot, uuid.MustParse("dc48710b-0e2e-419f-94fb-178c7fc1370b")))

	w := &bytes.Buffer{}
	binaryEncoder := NewBinaryEncoder(eResource, w, map[string]any{
		BINARY_OPTION_ID_ATTRIBUTE:        true,
		BINARY_OPTION_NAMESPACE_ATTRIBUTE: false,
	})
	binaryEncoder.EncodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	//os.WriteFile(path.Join(testdataPath,"library.complex.id.bin", w.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "library.complex.id.bin"))
	assert.Nil(t, err)
	assert.Equal(t, bytes, w.Bytes())
}

func TestBinaryEncoder_ComplexBig(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	// load resource
	xmlProcessor := ecore.NewXMLProcessor(ecore.XMLProcessorPackages([]ecore.EPackage{ePackage}))
	eResource := xmlProcessor.LoadWithOptions(ecore.NewURI(path.Join(testdataPath, "library.complex.big.xml")), nil)
	require.NotNil(t, eResource)
	require.True(t, eResource.IsLoaded())
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	require.True(t, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	w := &bytes.Buffer{}
	binaryEncoder := NewBinaryEncoder(eResource, w, map[string]any{})
	binaryEncoder.EncodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	//os.WriteFile(path.Join(testdataPath,"library.complex.big.bin", w.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "library.complex.big.bin"))
	assert.Nil(t, err)
	assert.Equal(t, bytes, w.Bytes())
}

func TestBinaryEncoder_SimpleWithDataTypeList(t *testing.T) {
	// load package
	ePackage := loadPackage("library.datalist.ecore")
	require.NotNil(t, ePackage)

	// load resource
	xmlProcessor := ecore.NewXMLProcessor(ecore.XMLProcessorPackages([]ecore.EPackage{ePackage}))
	eResource := xmlProcessor.LoadWithOptions(ecore.NewURI(path.Join(testdataPath, "library.datalist.xml")), nil)
	require.NotNil(t, eResource)
	require.True(t, eResource.IsLoaded())
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	require.True(t, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	w := &bytes.Buffer{}
	binaryEncoder := NewBinaryEncoder(eResource, w, map[string]any{})
	binaryEncoder.EncodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	//os.WriteFile(path.Join(testdataPath,"library.datalist.bin", w.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "library.datalist.bin"))
	assert.Nil(t, err)
	assert.Equal(t, bytes, w.Bytes())
}

func TestBinaryEncoder_Maps(t *testing.T) {
	// load package
	ePackage := loadPackage("emap.ecore")
	require.NotNil(t, ePackage)

	// load resource
	xmlProcessor := ecore.NewXMLProcessor(ecore.XMLProcessorPackages([]ecore.EPackage{ePackage}))
	eResource := xmlProcessor.LoadWithOptions(ecore.NewURI(path.Join(testdataPath, "emap.xml")), nil)
	require.NotNil(t, eResource)
	require.True(t, eResource.IsLoaded())
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	require.True(t, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	w := &bytes.Buffer{}
	binaryEncoder := NewBinaryEncoder(eResource, w, map[string]any{})
	binaryEncoder.EncodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	// os.WriteFile(path.Join(testdataPath,"emap.bin", w.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "emap.bin"))
	assert.Nil(t, err)
	assert.Equal(t, bytes, w.Bytes())
}

func TestBinaryEncoder_AllTypes(t *testing.T) {
	// load package
	ePackage := loadPackage("alltypes.ecore")
	require.NotNil(t, ePackage)

	// load resource
	xmlProcessor := ecore.NewXMLProcessor(ecore.XMLProcessorPackages([]ecore.EPackage{ePackage}))
	eResource := xmlProcessor.LoadWithOptions(ecore.NewURI(path.Join(testdataPath, "alltypes.xml")), nil)
	require.NotNil(t, eResource)
	require.True(t, eResource.IsLoaded())
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	require.True(t, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	w := &bytes.Buffer{}
	binaryEncoder := NewBinaryEncoder(eResource, w, map[string]any{})
	binaryEncoder.EncodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	//os.WriteFile(path.Join(testdataPath,"alltypes.bin", w.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "alltypes.bin"))
	assert.Nil(t, err)
	assert.Equal(t, bytes, w.Bytes())
}

func BenchmarkBinaryEncoderLibraryComplexBig(b *testing.B) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(b, ePackage)

	// load resource
	xmlProcessor := ecore.NewXMLProcessor(ecore.XMLProcessorPackages([]ecore.EPackage{ePackage}))
	eResource := xmlProcessor.LoadWithOptions(ecore.NewURI(path.Join(testdataPath, "library.complex.big.xml")), nil)
	require.NotNil(b, eResource)
	require.True(b, eResource.GetWarnings().Empty(), diagnosticError(eResource.GetWarnings()))

	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		binaryEncoder := NewBinaryEncoder(eResource, &buffer, nil)
		binaryEncoder.EncodeResource()
		require.True(b, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	}
}
