package ecore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqlDecoder_DecodeResource(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	// create resource & resourceset
	uri := NewURI("testdata/library.complex.sqlite")
	eResource := NewEResourceImpl()
	eResource.SetURI(uri)
	eResourceSet := NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	r, err := os.Open(uri.String())
	require.NoError(t, err)
	defer r.Close()

	sqlDecoder := NewSQLReaderDecoder(r, eResource, nil)
	sqlDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
}

func TestSqlDecoder_EMaps(t *testing.T) {
	// load package
	ePackage := loadPackage("emap.ecore")
	require.NotNil(t, ePackage)

	// create resource & resourceset
	sqlURI := NewURI("testdata/emap.sqlite")
	sqlResource := NewEResourceImpl()
	sqlResource.SetURI(sqlURI)

	eResourceSet := NewEResourceSetImpl()
	eResourceSet.GetResources().Add(sqlResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	sqlReader, err := os.Open(sqlURI.String())
	require.NoError(t, err)
	defer sqlReader.Close()

	sqlDecoder := NewSQLReaderDecoder(sqlReader, sqlResource, nil)
	sqlDecoder.DecodeResource()
	require.True(t, sqlResource.GetErrors().Empty(), diagnosticError(sqlResource.GetErrors()))

}

func TestSqlDecoder_SimpleNoIDs_NoObjectIDManager(t *testing.T) {
	// load package
	ePackage := loadPackage("library.simple.ecore")
	require.NotNil(t, ePackage)

	// create resource & resourceset
	sqlURI := NewURI("testdata/library.simple.sqlite")
	sqlResource := NewEResourceImpl()
	sqlResource.SetURI(sqlURI)

	eResourceSet := NewEResourceSetImpl()
	eResourceSet.GetResources().Add(sqlResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	sqlReader, err := os.Open(sqlURI.String())
	require.NoError(t, err)
	defer sqlReader.Close()

	sqlDecoder := NewSQLReaderDecoder(sqlReader, sqlResource, nil)
	sqlDecoder.DecodeResource()
	require.True(t, sqlResource.GetErrors().Empty(), diagnosticError(sqlResource.GetErrors()))
}

func TestSqlDecoder_SimpleNoIDs(t *testing.T) {
	// load package
	ePackage := loadPackage("library.simple.ecore")
	require.NotNil(t, ePackage)

	// create resource & resourceset
	sqlURI := NewURI("testdata/library.simple.sqlite")
	sqlResource := NewEResourceImpl()
	sqlResource.SetURI(sqlURI)

	eResourceSet := NewEResourceSetImpl()
	eResourceSet.GetResources().Add(sqlResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	sqlReader, err := os.Open(sqlURI.String())
	require.NoError(t, err)
	defer sqlReader.Close()

	sqlDecoder := NewSQLReaderDecoder(sqlReader, sqlResource, nil)
	sqlDecoder.DecodeResource()
	require.True(t, sqlResource.GetErrors().Empty(), diagnosticError(sqlResource.GetErrors()))
}

func TestSqlDecoder_SimpleWithIDs(t *testing.T) {
	// load package
	ePackage := loadPackage("library.simple.ecore")
	require.NotNil(t, ePackage)

	// create resource & resourceset
	sqlURI := NewURI("testdata/library.simple.ids.sqlite")
	sqlResource := NewEResourceImpl()
	sqlResource.SetURI(sqlURI)

	eResourceSet := NewEResourceSetImpl()
	eResourceSet.GetResources().Add(sqlResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	sqlReader, err := os.Open(sqlURI.String())
	require.NoError(t, err)
	defer sqlReader.Close()

	sqlDecoder := NewSQLReaderDecoder(sqlReader, sqlResource, map[string]any{SQL_OPTION_ID_ATTRIBUTE_NAME: "esyncID"})
	sqlDecoder.DecodeResource()
	require.True(t, sqlResource.GetErrors().Empty(), diagnosticError(sqlResource.GetErrors()))
}
