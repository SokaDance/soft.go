package library

import (
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/assert"
)

func diagnosticError(errors ecore.EList) string {
	if errors.Empty() {
		return ""
	} else {
		return errors.Get(0).(ecore.EDiagnostic).GetMessage()
	}
}

func TestSerializationLoadSimpleXML(t *testing.T) {
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())

	fileURI := &url.URL{Path: "testdata/library.simple.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.Load()
	assert.True(t, resource.IsLoaded())
	assert.True(t, resource.GetErrors().Empty(), diagnosticError(resource.GetErrors()))
	assert.True(t, resource.GetWarnings().Empty(), diagnosticError(resource.GetWarnings()))
}

func TestSerializationLoadOwnerXML(t *testing.T) {
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())

	fileURI := &url.URL{Path: "testdata/library.owner.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.Load()
	assert.True(t, resource.IsLoaded())
	assert.True(t, resource.GetErrors().Empty(), diagnosticError(resource.GetErrors()))
	assert.True(t, resource.GetWarnings().Empty(), diagnosticError(resource.GetWarnings()))
}

func TestSerializationLoadComplexXML(t *testing.T) {
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())

	fileURI := &url.URL{Path: "testdata/library.complex.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.Load()
	assert.True(t, resource.IsLoaded())
	assert.True(t, resource.GetErrors().Empty(), diagnosticError(resource.GetErrors()))
	assert.True(t, resource.GetWarnings().Empty(), diagnosticError(resource.GetWarnings()))
}

func TestSerializationLoadSaveSimpleXML(t *testing.T) {
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())

	fileURI := &url.URL{Path: "testdata/library.simple.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.Load()

	resource.SetURI(&url.URL{Path: "testdata/library.simple.output.xml"})
	resource.Save()
}

func TestSerializationLoadSaveXML(t *testing.T) {
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())

	fileURI := &url.URL{Path: "testdata/library.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.Load()

	var strbuff strings.Builder
	resource.SaveWithWriter(&strbuff)

	bytes, err := ioutil.ReadFile("testdata/library.xml")
	assert.Nil(t, err)
	assert.Equal(t, strings.ReplaceAll(string(bytes), "\r\n", "\n"), strings.ReplaceAll(strbuff.String(), "\r\n", "\n"))
}

func TestSerializationSaveXmlFile(t *testing.T) {
	// create a library model with a single employee
	root := GetFactory().CreateDocumentRoot()
	library := GetFactory().CreateLibrary()
	root.SetLibrary(library)

	employee := GetFactory().CreateEmployee()
	employee.SetFirstName("First Name")
	employee.SetLastName("Last Name")
	employee.SetAddress("adress")
	library.GetEmployees().Add(employee)

	// save library model with a resource
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())
	fileURI := &url.URL{Path: "testdata/simple.output.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.GetContents().Add(root)
	resource.Save()

	bytesInput, errInput := ioutil.ReadFile("testdata/simple.input.xml")
	assert.Nil(t, errInput)
	bytesOutput, errOutput := ioutil.ReadFile("testdata/simple.output.xml")
	assert.Nil(t, errOutput)
	assert.Equal(t, strings.ReplaceAll(string(bytesInput), "\r\n", "\n"), strings.ReplaceAll(string(bytesOutput), "\r\n", "\n"))

}

func TestDeepOperations(t *testing.T) {
	ecore.GetPackageRegistry().RegisterPackage(GetPackage())

	fileURI := &url.URL{Path: "testdata/library.xml"}
	resourceFactory := ecore.GetResourceFactoryRegistry().GetFactory(fileURI)
	resource := resourceFactory.CreateResource(fileURI)
	resource.Load()

	eObject := resource.GetContents().Get(0).(ecore.EObject)
	eCopyObject := ecore.Copy(eObject)
	assert.True(t, ecore.Equals(eObject, eCopyObject))
}
