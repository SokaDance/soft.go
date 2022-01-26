// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func discardEGenericType() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEGenericTypeAsEGenericType(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Equal(t, o, o.asEGenericType())
}

func TestEGenericTypeStaticClass(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Equal(t, GetPackage().GetEGenericType(), o.EStaticClass())
}

func TestEGenericTypeFeatureCount(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Equal(t, EGENERIC_TYPE_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEGenericTypeEClassifierGet(t *testing.T) {
	o := newEGenericTypeImpl()

	// get default value
	assert.Nil(t, o.GetEClassifier())

	// initialize object with a mock value
	mockValue := &MockEClassifier{}
	o.eClassifier = mockValue

	// events
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// set object resource
	mockResourceSet := new(MockEResourceSet)
	mockResource := new(MockEResource)
	o.ESetInternalResource(mockResource)

	// get non resolved value
	mockValue.On("EIsProxy").Return(false).Once()
	assert.Equal(t, mockValue, o.GetEClassifier())
	mock.AssertExpectationsForObjects(t, mockValue, mockAdapter, mockResource, mockResourceSet)

	// get a resolved value
	mockURI := NewURI("test:///file.t")
	mockResolved := &MockEClassifier{}
	mockResolved.On("EProxyURI").Return(nil).Once()
	mockResource.On("GetResourceSet").Return(mockResourceSet).Once()
	mockResourceSet.On("GetEObject", mockURI, true).Return(mockResolved).Once()
	mockValue.On("EIsProxy").Return(true).Once()
	mockValue.On("EProxyURI").Return(mockURI).Twice()
	mockAdapter.On("NotifyChanged", mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == RESOLVE && notification.GetFeatureID() == EGENERIC_TYPE__ECLASSIFIER && notification.GetOldValue() == mockValue && notification.GetNewValue() == mockResolved
	})).Once()
	assert.Equal(t, mockResolved, o.GetEClassifier())
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResolved, mockAdapter, mockResource, mockResourceSet)
}

func TestEGenericTypeEClassifierSet(t *testing.T) {
	o := newEGenericTypeImpl()
	v := &MockEClassifier{}
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetEClassifier(v)
	mockAdapter.AssertExpectations(t)
}

func TestEGenericTypeELowerBoundGet(t *testing.T) {
	o := newEGenericTypeImpl()
	// get default value
	assert.Nil(t, o.GetELowerBound())

	// get initialized value
	v := &MockEGenericType{}
	o.eLowerBound = v
	assert.Equal(t, v, o.GetELowerBound())
}

func TestEGenericTypeELowerBoundSet(t *testing.T) {
	o := newEGenericTypeImpl()

	// add listener
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue1 := &MockEGenericType{}
	mockValue1.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ELOWER_BOUND, mock.Anything).Return(nil).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetELowerBound(mockValue1)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue1)

	// second value
	mockValue2 := &MockEGenericType{}
	mockValue1.On("EInverseRemove", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ELOWER_BOUND, nil).Return(nil).Once()
	mockValue2.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ELOWER_BOUND, nil).Return(nil).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetELowerBound(mockValue2)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue1, mockValue2)

}

func TestEGenericTypeELowerBoundBasicSet(t *testing.T) {
	o := newEGenericTypeImpl()

	// add listener
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue := &MockEGenericType{}
	mockNotifications := new(MockENotificationChain)
	mockNotifications.On("Add", mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == SET && notification.GetFeatureID() == EGENERIC_TYPE__ELOWER_BOUND
	})).Return(true).Once()
	o.basicSetELowerBound(mockValue, mockNotifications)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockNotifications)
}

func TestEGenericTypeERawTypeGet(t *testing.T) {
	o := newEGenericTypeImpl()

	// get default value
	assert.Nil(t, o.GetERawType())

	// initialize object with a mock value
	mockValue := &MockEClassifier{}
	o.eRawType = mockValue

	// events
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// set object resource
	mockResourceSet := new(MockEResourceSet)
	mockResource := new(MockEResource)
	o.ESetInternalResource(mockResource)

	// get non resolved value
	mockValue.On("EIsProxy").Return(false).Once()
	assert.Equal(t, mockValue, o.GetERawType())
	mock.AssertExpectationsForObjects(t, mockValue, mockAdapter, mockResource, mockResourceSet)

	// get a resolved value
	mockURI := NewURI("test:///file.t")
	mockResolved := &MockEClassifier{}
	mockResolved.On("EProxyURI").Return(nil).Once()
	mockResource.On("GetResourceSet").Return(mockResourceSet).Once()
	mockResourceSet.On("GetEObject", mockURI, true).Return(mockResolved).Once()
	mockValue.On("EIsProxy").Return(true).Once()
	mockValue.On("EProxyURI").Return(mockURI).Twice()
	mockAdapter.On("NotifyChanged", mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == RESOLVE && notification.GetFeatureID() == EGENERIC_TYPE__ERAW_TYPE && notification.GetOldValue() == mockValue && notification.GetNewValue() == mockResolved
	})).Once()
	assert.Equal(t, mockResolved, o.GetERawType())
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResolved, mockAdapter, mockResource, mockResourceSet)
}

func TestEGenericTypeETypeArgumentsGet(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.NotNil(t, o.GetETypeArguments())
}

func TestEGenericTypeETypeParameterGet(t *testing.T) {
	o := newEGenericTypeImpl()
	// get default value
	assert.Nil(t, o.GetETypeParameter())

	// get initialized value
	v := &MockETypeParameter{}
	o.eTypeParameter = v
	assert.Equal(t, v, o.GetETypeParameter())
}

func TestEGenericTypeETypeParameterSet(t *testing.T) {
	o := newEGenericTypeImpl()
	v := &MockETypeParameter{}
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetETypeParameter(v)
	mockAdapter.AssertExpectations(t)
}

func TestEGenericTypeEUpperBoundGet(t *testing.T) {
	o := newEGenericTypeImpl()
	// get default value
	assert.Nil(t, o.GetEUpperBound())

	// get initialized value
	v := &MockEGenericType{}
	o.eUpperBound = v
	assert.Equal(t, v, o.GetEUpperBound())
}

func TestEGenericTypeEUpperBoundSet(t *testing.T) {
	o := newEGenericTypeImpl()

	// add listener
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue1 := &MockEGenericType{}
	mockValue1.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__EUPPER_BOUND, mock.Anything).Return(nil).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetEUpperBound(mockValue1)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue1)

	// second value
	mockValue2 := &MockEGenericType{}
	mockValue1.On("EInverseRemove", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__EUPPER_BOUND, nil).Return(nil).Once()
	mockValue2.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__EUPPER_BOUND, nil).Return(nil).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetEUpperBound(mockValue2)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue1, mockValue2)

}

func TestEGenericTypeEUpperBoundBasicSet(t *testing.T) {
	o := newEGenericTypeImpl()

	// add listener
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue := &MockEGenericType{}
	mockNotifications := new(MockENotificationChain)
	mockNotifications.On("Add", mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == SET && notification.GetFeatureID() == EGENERIC_TYPE__EUPPER_BOUND
	})).Return(true).Once()
	o.basicSetEUpperBound(mockValue, mockNotifications)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockNotifications)
}

func TestEGenericTypeIsInstanceOperation(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Panics(t, func() { o.IsInstance(nil) })
}

func TestEGenericTypeEGetFromID(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetETypeArguments(), FromAnyList[EGenericType](o.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, true)))
}

func TestEGenericTypeESetFromID(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := &MockEClassifier{}
		o.ESetFromID(EGENERIC_TYPE__ECLASSIFIER, v)
		assert.Equal(t, v, o.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false))
	}
	{
		mockValue := &MockEGenericType{}
		mockValue.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ELOWER_BOUND, mock.Anything).Return(nil).Once()
		o.ESetFromID(EGENERIC_TYPE__ELOWER_BOUND, mockValue)
		assert.Equal(t, mockValue, o.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		// list with a value
		mockList := &MockEList[EGenericType]{}
		mockValue := &MockEGenericType{}
		mockIterator := &MockEIterator[EGenericType]{}
		mockList.On("Iterator").Return(mockIterator).Once()
		mockIterator.On("HasNext").Return(true).Once()
		mockIterator.On("Next").Return(mockValue).Once()
		mockIterator.On("HasNext").Return(false).Once()
		mockValue.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ETYPE_ARGUMENTS, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, ToAnyList[EGenericType](mockList))
		// checks
		assert.Equal(t, 1, o.GetETypeArguments().Size())
		assert.Equal(t, mockValue, o.GetETypeArguments().Get(0))
		mock.AssertExpectationsForObjects(t, mockList, mockIterator, mockValue)
	}
	{
		v := &MockETypeParameter{}
		o.ESetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, v)
		assert.Equal(t, v, o.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false))
	}
	{
		mockValue := &MockEGenericType{}
		mockValue.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__EUPPER_BOUND, mock.Anything).Return(nil).Once()
		o.ESetFromID(EGENERIC_TYPE__EUPPER_BOUND, mockValue)
		assert.Equal(t, mockValue, o.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false))
		mock.AssertExpectationsForObjects(t, mockValue)
	}

}

func TestEGenericTypeEIsSetFromID(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EGENERIC_TYPE__ECLASSIFIER))
	assert.False(t, o.EIsSetFromID(EGENERIC_TYPE__ELOWER_BOUND))
	assert.False(t, o.EIsSetFromID(EGENERIC_TYPE__ERAW_TYPE))
	assert.False(t, o.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	assert.False(t, o.EIsSetFromID(EGENERIC_TYPE__ETYPE_PARAMETER))
	assert.False(t, o.EIsSetFromID(EGENERIC_TYPE__EUPPER_BOUND))
}

func TestEGenericTypeEUnsetFromID(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(EGENERIC_TYPE__ECLASSIFIER)
		assert.Nil(t, o.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false))
	}
	{
		o.EUnsetFromID(EGENERIC_TYPE__ELOWER_BOUND)
		assert.Nil(t, o.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false))
	}
	{
		o.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		v := o.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false)
		assert.NotNil(t, v)
		l := v.(EList[any])
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(EGENERIC_TYPE__ETYPE_PARAMETER)
		assert.Nil(t, o.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false))
	}
	{
		o.EUnsetFromID(EGENERIC_TYPE__EUPPER_BOUND)
		assert.Nil(t, o.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false))
	}
}

func TestEGenericTypeEInvokeFromID(t *testing.T) {
	o := newEGenericTypeImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
}

func TestEGenericTypeEBasicInverseRemove(t *testing.T) {
	o := newEGenericTypeImpl()
	{
		mockObject := new(MockEObject)
		mockNotifications := new(MockENotificationChain)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		mockObject := &MockEGenericType{}
		o.EBasicInverseRemove(mockObject, EGENERIC_TYPE__ELOWER_BOUND, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		// initialize list with a mock object
		mockObject := &MockEGenericType{}
		mockObject.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ETYPE_ARGUMENTS, mock.Anything).Return(nil).Once()

		l := o.GetETypeArguments()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, EGENERIC_TYPE__ETYPE_ARGUMENTS, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := &MockEGenericType{}
		o.EBasicInverseRemove(mockObject, EGENERIC_TYPE__EUPPER_BOUND, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}