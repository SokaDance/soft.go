package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasicEDataTypeListAccessors(t *testing.T) {
	{
		list := NewBasicEDataTypeList[string](nil, 1, false)
		assert.Equal(t, nil, list.GetNotifier())
		assert.Equal(t, nil, list.GetFeature())
		assert.Equal(t, 1, list.GetFeatureID())
	}
	{
		mockOwner := &MockEObjectInternal{}
		list := NewBasicEDataTypeList[int](mockOwner, 1, true)
		assert.Equal(t, mockOwner, list.GetNotifier())
		assert.Equal(t, 1, list.GetFeatureID())
		mockClass := &MockEClass{}
		mockFeature := &MockEStructuralFeature{}
		mockClass.On("GetEStructuralFeature", 1).Return(mockFeature)
		mockOwner.On("EClass").Return(mockClass)
		assert.Equal(t, mockFeature, list.GetFeature())
		mock.AssertExpectationsForObjects(t, mockOwner, mockClass, mockFeature, mockClass)
	}
}
