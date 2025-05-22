package bin

import (
	"bytes"
	"io"
	"os"
	"path"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBinaryDecoder_Invalid(t *testing.T) {
	// file
	f, err := os.Open(path.Join(testdataPath, "invalid.bin"))
	require.Nil(t, err)

	mockErrors := ecore.NewMockEList(t)
	mockResource := ecore.NewMockEResource(t)
	mockResource.EXPECT().GetURI().Return(ecore.NewURI(path.Join(testdataPath, "invalid.bin")))
	mockResource.EXPECT().GetErrors().Return(mockErrors).Once()
	mockErrors.EXPECT().Add(mock.Anything).Return(true).Once()
	binaryDecoder := NewBinaryDecoder(mockResource, f, nil)
	binaryDecoder.DecodeResource()
}

func TestBinaryDecoder_Complex(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	//
	uri := ecore.NewURI(path.Join(testdataPath, "library.complex.bin"))
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// file
	f, err := os.Open(uri.String())
	require.Nil(t, err)

	binaryDecoder := NewBinaryDecoder(eResource, f, nil)
	binaryDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	// retrieve document root class , library class & library name attribute
	eDocumentRootClass, _ := ePackage.GetEClassifier("DocumentRoot").(ecore.EClass)
	assert.NotNil(t, eDocumentRootClass)
	eDocumentRootLibraryFeature, _ := eDocumentRootClass.GetEStructuralFeatureFromName("library").(ecore.EReference)
	assert.NotNil(t, eDocumentRootLibraryFeature)
	eLibraryClass, _ := ePackage.GetEClassifier("Library").(ecore.EClass)
	assert.NotNil(t, eLibraryClass)
	eLibraryNameAttribute, _ := eLibraryClass.GetEStructuralFeatureFromName("name").(ecore.EAttribute)
	assert.NotNil(t, eLibraryNameAttribute)

	// check library name
	eDocumentRoot, _ := eResource.GetContents().Get(0).(ecore.EObject)
	assert.NotNil(t, eDocumentRoot)
	eLibrary, _ := eDocumentRoot.EGet(eDocumentRootLibraryFeature).(ecore.EObject)
	assert.NotNil(t, eLibrary)
	assert.Equal(t, "My Library", eLibrary.EGet(eLibraryNameAttribute))

	// book class and attributes
	eLibraryBooksRefeference, _ := eLibraryClass.GetEStructuralFeatureFromName("books").(ecore.EReference)
	assert.NotNil(t, eLibraryBooksRefeference)
	eBookClass, _ := ePackage.GetEClassifier("Book").(ecore.EClass)
	require.NotNil(t, eBookClass)
	eBookTitleAttribute, _ := eBookClass.GetEStructuralFeatureFromName("title").(ecore.EAttribute)
	require.NotNil(t, eBookTitleAttribute)
	eBookDateAttribute, _ := eBookClass.GetEStructuralFeatureFromName("publicationDate").(ecore.EAttribute)
	require.NotNil(t, eBookDateAttribute)
	eBookCategoryAttribute, _ := eBookClass.GetEStructuralFeatureFromName("category").(ecore.EAttribute)
	require.NotNil(t, eBookCategoryAttribute)
	eBookAuthorReference, _ := eBookClass.GetEStructuralFeatureFromName("author").(ecore.EReference)
	require.NotNil(t, eBookAuthorReference)

	// retrive book
	eBooks, _ := eLibrary.EGet(eLibraryBooksRefeference).(ecore.EList)
	assert.NotNil(t, eBooks)
	eBook := eBooks.Get(0).(ecore.EObject)
	require.NotNil(t, eBook)

	// check book name
	assert.Equal(t, "Title 0", eBook.EGet(eBookTitleAttribute))

	// check book date
	date, _ := eBook.EGet(eBookDateAttribute).(*time.Time)
	require.NotNil(t, date)
	assert.Equal(t, time.Date(2015, time.September, 6, 4, 24, 46, 0, time.UTC), date.UTC())

	// check book category
	category := eBook.EGet(eBookCategoryAttribute)
	assert.Equal(t, 2, category)

	// check author
	author := eBook.EGet(eBookAuthorReference).(ecore.EObject)
	require.NotNil(t, author)

	eWriterClass, _ := ePackage.GetEClassifier("Writer").(ecore.EClass)
	require.NotNil(t, eWriterClass)
	eWriterNameAttribute := eWriterClass.GetEStructuralFeatureFromName("firstName").(ecore.EAttribute)
	require.NotNil(t, eWriterNameAttribute)
	authorName := author.EGet(eWriterNameAttribute).(string)
	assert.Equal(t, "First Name 0", authorName)
}

func TestBinaryDecoder_ComplexWithID(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	//
	uri := ecore.NewURI(path.Join(testdataPath, "library.complex.id.bin"))
	idManager := ecore.NewUUIDManager()
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResource.SetObjectIDManager(idManager)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// file
	f, err := os.Open(uri.String())
	require.Nil(t, err)

	binaryDecoder := NewBinaryDecoder(eResource, f, nil)
	binaryDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	// retrieve document root class , library class & library name attribute
	eDocumentRootClass, _ := ePackage.GetEClassifier("DocumentRoot").(ecore.EClass)
	require.NotNil(t, eDocumentRootClass)
	eDocumentRootLibraryFeature, _ := eDocumentRootClass.GetEStructuralFeatureFromName("library").(ecore.EReference)
	require.NotNil(t, eDocumentRootLibraryFeature)

	// check ids for document root and library
	eDocumentRoot, _ := eResource.GetContents().Get(0).(ecore.EObject)
	require.NotNil(t, eDocumentRoot)
	eLibrary, _ := eDocumentRoot.EGet(eDocumentRootLibraryFeature).(ecore.EObject)
	require.NotNil(t, eLibrary)
	assert.Equal(t, uuid.MustParse("dc48710b-0e2e-419f-94fb-178c7fc1370b"), idManager.GetID(eDocumentRoot))
	assert.Equal(t, uuid.MustParse("75aa92db-b419-4259-93c4-0e542d33aa35"), idManager.GetID(eLibrary))

}

func TestBinaryDecoder_SimpleWithEDataTypeList(t *testing.T) {
	// load package
	ePackage := loadPackage("library.datalist.ecore")
	require.NotNil(t, ePackage)

	//
	uri := ecore.NewURI(path.Join(testdataPath, "library.datalist.bin"))
	idManager := ecore.NewUUIDManager()
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResource.SetObjectIDManager(idManager)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// file
	f, err := os.Open(uri.String())
	require.Nil(t, err)

	binaryDecoder := NewBinaryDecoder(eResource, f, nil)
	binaryDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	// retrieve library class & library name attribute
	libraryClass, _ := ePackage.GetEClassifier("Library").(ecore.EClass)
	require.NotNil(t, libraryClass)
	libraryBooksFeature := libraryClass.GetEStructuralFeatureFromName("books")
	require.NotNil(t, libraryBooksFeature)
	bookClass, _ := ePackage.GetEClassifier("Book").(ecore.EClass)
	require.NotNil(t, bookClass)
	bookContentsFeature := bookClass.GetEStructuralFeatureFromName("contents")
	require.NotNil(t, bookContentsFeature)

	require.Equal(t, 1, eResource.GetContents().Size())
	eLibrary, _ := eResource.GetContents().Get(0).(ecore.EObject)
	require.NotNil(t, eLibrary)

	eBooks, _ := eLibrary.EGet(libraryBooksFeature).(ecore.EList)
	require.NotNil(t, eBooks)
	require.Equal(t, 4, eBooks.Size())

	eLastBook, _ := eBooks.Get(3).(ecore.EObject)
	require.NotNil(t, eLastBook)
	eContents, _ := eLastBook.EGet(bookContentsFeature).(ecore.EList)
	require.NotNil(t, eContents)
	assert.Equal(t, 3, eContents.Size())
	assert.Equal(t, "c31", eContents.Get(0))

}

func TestBinaryDecoder_ComplexBig(t *testing.T) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	//
	uri := ecore.NewURI(path.Join(testdataPath, "library.complex.big.bin"))
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// file
	f, err := os.Open(uri.String())
	require.Nil(t, err)

	binaryDecoder := NewBinaryDecoder(eResource, f, nil)
	binaryDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
}

func TestBinaryDecoder_Maps(t *testing.T) {
	// load package
	ePackage := loadPackage("emap.ecore")
	require.NotNil(t, ePackage)

	//
	uri := ecore.NewURI(path.Join(testdataPath, "emap.bin"))
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// file
	f, err := os.Open(uri.String())
	require.Nil(t, err)

	binaryDecoder := NewBinaryDecoder(eResource, f, nil)
	binaryDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	eMapTestClass, _ := ePackage.GetEClassifier("EMapTest").(ecore.EClass)
	require.NotNil(t, eMapTestClass)
	eMapTestKeyToValueReference, _ := eMapTestClass.GetEStructuralFeatureFromName("keyToValue").(ecore.EReference)
	require.NotNil(t, eMapTestKeyToValueReference)
	eMapTestKeyToIntReference, _ := eMapTestClass.GetEStructuralFeatureFromName("keyToInt").(ecore.EReference)
	require.NotNil(t, eMapTestKeyToIntReference)
	eKeyTypeClass, _ := ePackage.GetEClassifier("KeyType").(ecore.EClass)
	require.NotNil(t, eKeyTypeClass)
	eKeyTypeNameAttribute, _ := eKeyTypeClass.GetEStructuralFeatureFromName("name").(ecore.EAttribute)
	require.NotNil(t, eKeyTypeNameAttribute)
	eValueTypeClass, _ := ePackage.GetEClassifier("ValueType").(ecore.EClass)
	require.NotNil(t, eValueTypeClass)
	eValueTypeNameAttribute, _ := eValueTypeClass.GetEStructuralFeatureFromName("name").(ecore.EAttribute)
	require.NotNil(t, eValueTypeNameAttribute)

	mapTest := eResource.GetContents().Get(0).(ecore.EObject)
	require.Equal(t, eMapTestClass, mapTest.EClass())
	mapKeyToValue, _ := mapTest.EGet(eMapTestKeyToValueReference).(ecore.EMap)
	require.NotNil(t, mapKeyToValue)
	assert.Equal(t, 5, mapKeyToValue.Size())
	mapKeyToInt, _ := mapTest.EGet(eMapTestKeyToIntReference).(ecore.EMap)
	require.NotNil(t, mapKeyToInt)
	assert.Equal(t, 5, mapKeyToInt.Size())
}

func TestBinaryDecoder_AllTypes(t *testing.T) {

	ePackage := loadPackage("alltypes.ecore")
	require.NotNil(t, ePackage)

	// retrive library class & library name attribute
	objectClass, _ := ePackage.GetEClassifier("Object").(ecore.EClass)
	require.NotNil(t, objectClass)

	enumType := ePackage.GetEClassifier("EnumCategory").(ecore.EEnum)
	require.NotNil(t, enumType)

	objectF32Attribute := objectClass.GetEStructuralFeatureFromName("f32")
	require.NotNil(t, objectF32Attribute)

	objectF64Attribute := objectClass.GetEStructuralFeatureFromName("f64")
	require.NotNil(t, objectF64Attribute)

	objectStringAttribute := objectClass.GetEStructuralFeatureFromName("str")
	require.NotNil(t, objectStringAttribute)

	objectI8Attribute := objectClass.GetEStructuralFeatureFromName("i8")
	require.NotNil(t, objectI8Attribute)

	objectI16Attribute := objectClass.GetEStructuralFeatureFromName("i16")
	require.NotNil(t, objectI16Attribute)

	objectI32Attribute := objectClass.GetEStructuralFeatureFromName("i32")
	require.NotNil(t, objectI32Attribute)

	objectI64Attribute := objectClass.GetEStructuralFeatureFromName("i64")
	require.NotNil(t, objectI64Attribute)

	objectIntAttribute := objectClass.GetEStructuralFeatureFromName("i")
	require.NotNil(t, objectIntAttribute)

	objectBytesAttribute := objectClass.GetEStructuralFeatureFromName("bytes")
	require.NotNil(t, objectBytesAttribute)

	objectBoolAttribute := objectClass.GetEStructuralFeatureFromName("b")
	require.NotNil(t, objectBoolAttribute)

	objectEnumAttribute := objectClass.GetEStructuralFeatureFromName("e")
	require.NotNil(t, objectEnumAttribute)

	//
	uri := ecore.NewURI(path.Join(testdataPath, "alltypes.bin"))
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// file
	f, err := os.Open(uri.String())
	require.Nil(t, err)

	binaryDecoder := NewBinaryDecoder(eResource, f, nil)
	binaryDecoder.DecodeResource()
	require.True(t, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))

	eObject := eResource.GetContents().Get(0).(ecore.EObject)
	require.Equal(t, float32(3.0), eObject.EGet(objectF32Attribute))
	require.Equal(t, float64(4.0), eObject.EGet(objectF64Attribute))
	require.Equal(t, "str", eObject.EGet(objectStringAttribute))
	require.Equal(t, byte('b'), eObject.EGet(objectI8Attribute))
	require.Equal(t, int16(2), eObject.EGet(objectI16Attribute))
	require.Equal(t, int32(1), eObject.EGet(objectI32Attribute))
	require.Equal(t, int64(0), eObject.EGet(objectI64Attribute))
	require.Equal(t, int(-1), eObject.EGet(objectIntAttribute))
	require.Equal(t, []byte("bytes"), eObject.EGet(objectBytesAttribute))
	require.Equal(t, true, eObject.EGet(objectBoolAttribute))
	require.Equal(t, enumType.GetDefaultValue(), eObject.EGet(objectEnumAttribute))
}

func BenchmarkBinaryDecoderLibraryComplexBig(b *testing.B) {
	// load package
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(b, ePackage)

	// create resource
	uri := ecore.NewURI(path.Join(testdataPath, "library.complex.big.bin"))
	eResource := ecore.NewEResourceImpl()
	eResource.SetURI(uri)
	eResourceSet := ecore.NewEResourceSetImpl()
	eResourceSet.GetResources().Add(eResource)
	eResourceSet.GetPackageRegistry().RegisterPackage(ePackage)

	// get file content
	content, err := os.ReadFile(uri.String())
	require.Nil(b, err)
	require.Nil(b, err)
	r := bytes.NewReader(content)

	for i := 0; i < b.N; i++ {
		_, err = r.Seek(0, io.SeekStart)
		require.Nil(b, err)
		binaryDecoder := NewBinaryDecoder(eResource, r, nil)
		binaryDecoder.DecodeResource()
		require.True(b, eResource.GetErrors().Empty(), diagnosticError(eResource.GetErrors()))
	}
}
