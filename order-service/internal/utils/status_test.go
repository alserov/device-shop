package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	testName string
	code     int32
	status   string
}

func TestStatusCodeToString(t *testing.T) {
	tests := []test{
		{
			testName: "creating",
			code:     CREATING_CODE,
			status:   CREATING,
		},
		{
			testName: "pending",
			code:     PENDING_CODE,
			status:   PENDING,
		},
		{
			testName: "canceled",
			code:     CANCELED_CODE,
			status:   CANCELED,
		},
		{
			testName: "delivering",
			code:     DELIVERING_CODE,
			status:   DELIVERING,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, StatusCodeToString(tc.code), tc.status)
	}
}

func TestStatusToCode(t *testing.T) {
	tests := []test{
		{
			testName: "creating",
			code:     CREATING_CODE,
			status:   CREATING,
		},
		{
			testName: "pending",
			code:     PENDING_CODE,
			status:   PENDING,
		},
		{
			testName: "canceled",
			code:     CANCELED_CODE,
			status:   CANCELED,
		},
		{
			testName: "delivering",
			code:     DELIVERING_CODE,
			status:   DELIVERING,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, StatusToCode(tc.status), tc.code)
	}
}
