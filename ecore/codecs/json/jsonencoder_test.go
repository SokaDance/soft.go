package json

import (
	"bytes"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONEncoder_EncodeResourceSimple(t *testing.T) {
	eResourceSet := ecore.NewEResourceSetImpl()
	ePackageResource, ePackage := loadTestPackage(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "library.simple.ecore")))
	require.NotNil(t, ePackage)
	require.NotNil(t, ePackageResource)
	eModelResource, eModel := loadTestModel(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "library.simple.xml")), false)
	require.NotNil(t, eModel)
	require.NotNil(t, eModelResource)

	buffer := &bytes.Buffer{}
	encoder := NewJSONEncoder(eModelResource, buffer, nil)
	encoder.EncodeResource()

	bytes, err := os.ReadFile(path.Join(testdataPath, "library.simple.json"))
	assert.Nil(t, err)
	assert.Equal(t, strings.ReplaceAll(string(bytes), "\r\n", "\n"), strings.ReplaceAll(buffer.String(), "\r\n", "\n"))
}

func TestJSONEncoder_EncodeResourceComplex(t *testing.T) {
	eResourceSet := ecore.NewEResourceSetImpl()
	ePackageResource, ePackage := loadTestPackage(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "library.complex.ecore")))
	require.NotNil(t, ePackage)
	require.NotNil(t, ePackageResource)
	eModelResource, eModel := loadTestModel(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "library.complex.xml")), false)
	require.NotNil(t, eModel)
	require.NotNil(t, eModelResource)

	buffer := &bytes.Buffer{}
	encoder := NewJSONEncoder(eModelResource, buffer, nil)
	encoder.EncodeResource()

	bytes, err := os.ReadFile(path.Join(testdataPath, "library.complex.json"))
	assert.Nil(t, err)
	assert.Equal(t, strings.ReplaceAll(string(bytes), "\r\n", "\n"), strings.ReplaceAll(buffer.String(), "\r\n", "\n"))
}

func TestJSONEncoder_EncodeObject(t *testing.T) {
	eResourceSet := ecore.NewEResourceSetImpl()
	// load packages & models
	eShopPackageResource, eShopPackage := loadTestPackage(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "shop.ecore")))
	require.NotNil(t, eShopPackage)
	require.NotNil(t, eShopPackageResource)
	eShopModelResource, eShopModel := loadTestModel(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "shop.xml")), true)
	require.NotNil(t, eShopModel)
	require.NotNil(t, eShopModelResource)

	eOrdersPackageResource, eOrdersPackage := loadTestPackage(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "orders.ecore")))
	require.NotNil(t, eOrdersPackageResource)
	require.NotNil(t, eOrdersPackage)
	eOrdersModelResource, eOrdersModel := loadTestModel(t, eResourceSet, ecore.NewURI(path.Join(testdataPath, "orders.xml")), true)
	require.NotNil(t, eOrdersModelResource)
	require.NotNil(t, eOrdersModel)

	codecOptions := map[string]any{JSON_OPTION_ID_ATTRIBUTE_NAME: "id"}
	var buffer bytes.Buffer
	jsonEncoder := NewJSONEncoder(eOrdersModelResource, &buffer, codecOptions)
	require.NoError(t, jsonEncoder.EncodeObject(eOrdersModel))

	//os.WriteFile("testdata/orders.json", buffer.Bytes(), 0644)

	bytes, err := os.ReadFile(path.Join(testdataPath, "orders.json"))
	assert.Nil(t, err)
	assert.Equal(t, strings.ReplaceAll(string(bytes), "\r\n", "\n"), strings.ReplaceAll(buffer.String(), "\r\n", "\n"))
}
