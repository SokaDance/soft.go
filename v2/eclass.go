// Code generated by soft.generator.go. DO NOT EDIT.



package ecore



// EClass is the representation of the model object 'EClass'
type EClass interface {
    EClassifier

    IsSuperTypeOf(EClass) bool
    GetFeatureCount() int
    GetEStructuralFeature(int) EStructuralFeature
    GetEStructuralFeatureFromName(string) EStructuralFeature
    GetFeatureID(EStructuralFeature) int
    GetOperationCount() int
    GetEOperation(int) EOperation
    GetOperationID(EOperation) int
    GetOverride(EOperation) EOperation
    GetFeatureType(EStructuralFeature) EClassifier
    
    IsAbstract() bool
    SetAbstract( bool )
    
    IsInterface() bool
    SetInterface( bool )
    
    
    GetEStructuralFeatures() EList[EStructuralFeature]
    
    GetEAttributes() EList[EAttribute]
    
    GetEReferences() EList[EReference]
    
    GetESuperTypes() EList[EClass]
    
    GetEOperations() EList[EOperation]
    
    GetEContainmentFeatures() EList[EStructuralFeature]
    
    GetECrossReferenceFeatures() EList[EStructuralFeature]
    
    GetEAllAttributes() EList[EAttribute]
    
    GetEAllReferences() EList[EReference]
    
    GetEAllContainments() EList[EReference]
    
    GetEAllOperations() EList[EOperation]
    
    GetEAllStructuralFeatures() EList[EStructuralFeature]
    
    GetEAllSuperTypes() EList[EClass]
    
    GetEIDAttribute() EAttribute
    
    
    // Start of user code EClass
    // End of user code
}

