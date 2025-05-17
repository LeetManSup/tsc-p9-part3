package service

import (
	"testing"
	"tsc-p9-part3/mocks"

	"github.com/stretchr/testify/require"
)

func TestService_Notify(t *testing.T) {
	mock := &mocks.MockEmailSender{}

	s := NewService(mock)

	err := s.Notify("user@example.com", "Hello", "Test body")
	require.NoError(t, err)

	require.Len(t, mock.Calls, 1)
	require.Contains(t, mock.Calls[0], "user@example.com")
}
