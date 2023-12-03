package utils

import (
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateToken(t *testing.T) {
	tests := []validation.test[string]{
		{
			testName:         "empty token",
			in:               "",
			shouldThrowError: true,
		},
		{
			testName:         "invalid token",
			in:               "not a token",
			shouldThrowError: true,
		},
	}

	for _, tc := range tests {
		if tc.shouldThrowError {
			assert.Error(t, ValidateToken(tc.in))
		} else {
			assert.NoError(t, ValidateToken(tc.in))
		}
	}
}
