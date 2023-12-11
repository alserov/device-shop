package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	password := "my_secret_password"

	hashed, err := HashPassword(password)
	require.NoError(t, err)

	err = CheckPassword(password, hashed)
	require.NoError(t, err)
}
