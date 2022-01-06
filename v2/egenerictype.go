// Code generated by soft.generator.go. DO NOT EDIT.



package ecore



// EGenericType is the representation of the model object 'EGenericType'
type EGenericType interface {
    EObject

    IsInstance(any) bool
    
    GetEUpperBound() EGenericType
    SetEUpperBound( EGenericType )
    
    GetETypeArguments() EList[EGenericType]
    
    GetERawType() EClassifier
    
    GetELowerBound() EGenericType
    SetELowerBound( EGenericType )
    
    GetETypeParameter() ETypeParameter
    SetETypeParameter( ETypeParameter )
    
    GetEClassifier() EClassifier
    SetEClassifier( EClassifier )
    
    
    // Start of user code EGenericType
    // End of user code
}
