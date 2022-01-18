// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

// ETypedElement is the representation of the model object 'ETypedElement'
type ETypedElement interface {
	ENamedElement

	IsOrdered() bool
	SetOrdered(bool)

	IsUnique() bool
	SetUnique(bool)

	GetLowerBound() int
	SetLowerBound(int)

	GetUpperBound() int
	SetUpperBound(int)

	IsMany() bool

	IsRequired() bool

	GetEType() EClassifier
	SetEType(EClassifier)
	UnsetEType()

	// Start of user code ETypedElement
	// End of user code
}
