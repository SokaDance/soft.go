package ecore

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dynamicSQLEObjectImpl struct {
	DynamicEObjectImpl
	sqlID int64
}

func newDynamicSQLEObjectImpl() *dynamicSQLEObjectImpl {
	o := new(dynamicSQLEObjectImpl)
	o.SetInterfaces(o)
	o.Initialize()
	return o
}

func (o *dynamicSQLEObjectImpl) SetSqlID(sqlID int64) {
	o.sqlID = sqlID
}

func (o *dynamicSQLEObjectImpl) GetSqlID() int64 {
	return o.sqlID
}

type dynamicSQLFactory struct {
	EFactoryExt
}

func newDynamicSQLFactory() *dynamicSQLFactory {
	eFactory := new(dynamicSQLFactory)
	eFactory.SetInterfaces(eFactory)
	eFactory.Initialize()
	return eFactory
}

func (eFactory *dynamicSQLFactory) Create(eClass EClass) EObject {
	if eFactory.GetEPackage() != eClass.GetEPackage() || eClass.IsAbstract() {
		panic("The class '" + eClass.GetName() + "' is not a valid classifier")
	}
	if IsMapEntry(eClass) {
		eEntry := NewDynamicEMapEntryImpl()
		eEntry.SetEClass(eClass)
		return eEntry
	} else {
		eObject := newDynamicSQLEObjectImpl()
		eObject.SetEClass(eClass)
		return eObject
	}
}

func TestSQLStore_Constructor(t *testing.T) {
	s, err := NewSQLStore("testdata/library.store.sqlite", NewURI(""), nil, nil, nil)
	require.Nil(t, err)
	require.NotNil(t, s)
	require.Nil(t, s.Close())
}

func closeFile(f *os.File, reported *error) {
	if err := f.Close(); *reported == nil {
		*reported = err
	}
}

func copyFile(src, dest string) (err error) {
	if err = os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return
	}

	d, err := os.Create(dest)
	if err != nil {
		return
	}
	defer closeFile(d, &err)

	s, err := os.Open(src)
	if err != nil {
		return
	}
	defer closeFile(s, &err)

	_, err = io.Copy(d, s)
	return
}

func TestSQLStore_SetSingleValue_Int(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Lendable").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("copies")
	require.NotNil(t, eFeature)

	// database
	dbPath := filepath.Join(t.TempDir(), "library.store.sqlite")
	err := copyFile("testdata/library.store.sqlite", dbPath)
	require.Nil(t, err)

	// store
	s, err := NewSQLStore(dbPath, NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(3)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	s.Set(mockObject, eFeature, -1, 5)
}

func TestSQLStore_GetSingleValue_Int(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Lendable").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("copies")
	require.NotNil(t, eFeature)

	// store
	s, err := NewSQLStore("testdata/library.store.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(3)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	v := s.Get(mockObject, eFeature, -1)
	assert.Equal(t, 4, v)
}

func TestSQLStore_GetSingleValue_Enum(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Book").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("category")
	require.NotNil(t, eFeature)

	// store
	mockPackageRegitry := NewMockEPackageRegistry(t)
	s, err := NewSQLStore("testdata/library.store.sqlite", NewURI(""), nil, mockPackageRegitry, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	// Mystery == 0
	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(4)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	mockPackageRegitry.EXPECT().GetPackage("http:///org/eclipse/emf/examples/library/library.ecore/1.0.0").Return(ePackage).Once()
	v := s.Get(mockObject, eFeature, -1)
	assert.Equal(t, 0, v)

	// Biography == 2
	mockObject.EXPECT().GetSqlID().Return(int64(3)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	v = s.Get(mockObject, eFeature, -1)
	assert.Equal(t, 2, v)
}

func TestSQLStore_GetSingleValue_String_Null(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("name")
	require.NotNil(t, eFeature)

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(1)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()

	s, err := NewSQLStore("testdata/library.owner.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	v := s.Get(mockObject, eFeature, -1)
	assert.Equal(t, "", v)
}

func TestSQLStore_GetSingleValue_Object_Nil(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("ownerPdg")
	require.NotNil(t, eFeature)

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(2)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()

	s, err := NewSQLStore("testdata/library.store.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	v := s.Get(mockObject, eFeature, -1)
	assert.Nil(t, v)
}

func TestSQLStore_GetSingleValue_Object(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)
	ePackage.SetEFactoryInstance(newDynamicSQLFactory())

	eLibraryClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eLibraryClass)

	eLibraryOwnerFeature := eLibraryClass.GetEStructuralFeatureFromName("ownerPdg")
	require.NotNil(t, eLibraryOwnerFeature)

	ePersonClass, _ := ePackage.GetEClassifier("Person").(EClass)
	require.NotNil(t, ePersonClass)

	ePersonAdressAttribute, _ := ePersonClass.GetEStructuralFeatureFromName("address").(EAttribute)
	require.NotNil(t, ePersonAdressAttribute)

	ePersonFirstNameAttribute, _ := ePersonClass.GetEStructuralFeatureFromName("firstName").(EAttribute)
	require.NotNil(t, ePersonFirstNameAttribute)

	mockObject := NewMockSQLObject(t)
	mockPackageRegitry := NewMockEPackageRegistry(t)
	s, err := NewSQLStore("testdata/library.owner.sqlite", NewURI(""), nil, mockPackageRegitry, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	mockObject.EXPECT().GetSqlID().Return(int64(1)).Once()
	mockObject.EXPECT().EClass().Return(eLibraryClass).Once()
	mockPackageRegitry.EXPECT().GetPackage("http:///org/eclipse/emf/examples/library/library.ecore/1.0.0").Return(ePackage).Once()
	v, _ := s.Get(mockObject, eLibraryOwnerFeature, -1).(SQLObject)
	require.NotNil(t, v)
	assert.Equal(t, ePersonClass, v.EClass())
	assert.Equal(t, int64(2), v.GetSqlID())
}

func TestSQLStore_GetSingleValue_Reference(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Book").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("author")
	require.NotNil(t, eFeature)

	mockPackageRegitry := NewMockEPackageRegistry(t)
	mockObject := NewMockSQLObject(t)

	s, err := NewSQLStore("testdata/library.complex.sqlite", NewURI(""), nil, mockPackageRegitry, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	mockObject.EXPECT().GetSqlID().Return(int64(3)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	v, _ := s.Get(mockObject, eFeature, -1).(EObject)
	require.NotNil(t, v)
	assert.True(t, v.EIsProxy())
	assert.Equal(t, "#//@library/@writers.0", v.(EObjectInternal).EProxyURI().String())
}

func TestSQLStore_SetListValue(t *testing.T) {
	ePackage := loadPackage("library.datalist.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Book").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("contents")
	require.NotNil(t, eFeature)

	// database
	dbPath := filepath.Join(t.TempDir(), "library.datalist.sqlite")
	err := copyFile("testdata/library.datalist.sqlite", dbPath)
	require.Nil(t, err)

	// create store
	s, err := NewSQLStore(dbPath, NewURI(""), nil, nil, nil)
	require.Nil(t, err)
	require.NotNil(t, s)
	defer s.Close()

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(5)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	s.Set(mockObject, eFeature, 1, "c4")

	// load db and retrieve new value
}

func TestSQLStore_GetListValue(t *testing.T) {
	ePackage := loadPackage("library.datalist.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Book").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("contents")
	require.NotNil(t, eFeature)

	// create store
	s, err := NewSQLStore("testdata/library.datalist.sqlite", NewURI(""), nil, nil, nil)
	require.Nil(t, err)
	require.NotNil(t, s)
	defer s.Close()

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(5)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()
	s.Get(mockObject, eFeature, 1)

	// load db and retrieve new value
}

func TestSQLStore_IsSet_SingleValue_NotSet(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("ownerPdg")
	require.NotNil(t, eFeature)

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(2)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()

	s, err := NewSQLStore("testdata/library.store.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	isSet := s.IsSet(mockObject, eFeature)
	assert.False(t, isSet)
}

func TestSQLStore_IsSet_SingleValue_Set(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("ownerPdg")
	require.NotNil(t, eFeature)

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(1)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()

	s, err := NewSQLStore("testdata/library.owner.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	isSet := s.IsSet(mockObject, eFeature)
	assert.True(t, isSet)
}

func TestSQLStore_IsSet_ManyValue_NotSet(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("books")
	require.NotNil(t, eFeature)

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(1)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()

	s, err := NewSQLStore("testdata/library.owner.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	isSet := s.IsSet(mockObject, eFeature)
	assert.False(t, isSet)
}

func TestSQLStore_IsSet_ManyValue_Set(t *testing.T) {
	ePackage := loadPackage("library.complex.ecore")
	require.NotNil(t, ePackage)

	eClass, _ := ePackage.GetEClassifier("Library").(EClass)
	require.NotNil(t, eClass)

	eFeature := eClass.GetEStructuralFeatureFromName("books")
	require.NotNil(t, eFeature)

	mockObject := NewMockSQLObject(t)
	mockObject.EXPECT().GetSqlID().Return(int64(2)).Once()
	mockObject.EXPECT().EClass().Return(eClass).Once()

	s, err := NewSQLStore("testdata/library.complex.sqlite", NewURI(""), nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, s)
	defer s.Close()

	isSet := s.IsSet(mockObject, eFeature)
	assert.True(t, isSet)
}
