// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

// EFactory is the representation of the model object 'EFactory'
type EFactory interface {
	EModelElement

	Create(EClass) EObject
	CreateFromString(EDataType, string) any
	ConvertToString(EDataType, any) string

	GetEPackage() EPackage
	SetEPackage(EPackage)

	// Start of user code EFactory
	// End of user code
}
