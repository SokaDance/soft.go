package sqlite

import (
	"testing"

	"github.com/masagroup/soft.go/ecore"
	"github.com/stretchr/testify/require"
)

func TestNewSQLEncoder(t *testing.T) {
	codec := &SQLCodec{}
	mockResource := ecore.NewMockEResource(t)
	mockResource.EXPECT().GetURI().Return(ecore.NewURI("")).Twice()
	mockResource.EXPECT().GetObjectIDManager().Return(nil).Once()
	require.NotNil(t, codec.NewEncoder(mockResource, nil, nil))
}

func TestNewSQLDecoder(t *testing.T) {
	codec := &SQLCodec{}
	mockResource := ecore.NewMockEResource(t)
	mockResource.EXPECT().GetResourceSet().Return(nil).Once()
	mockResource.EXPECT().GetURI().Return(ecore.NewURI("")).Once()
	mockResource.EXPECT().GetObjectIDManager().Return(nil).Once()
	require.NotNil(t, codec.NewDecoder(mockResource, nil, nil))
}
