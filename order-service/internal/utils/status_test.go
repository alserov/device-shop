package utils

import (
	"github.com/alserov/device-shop/order-service/internal/utils/status"
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
			code:     status.CREATING_CODE,
			status:   status.CREATING,
		},
		{
			testName: "pending",
			code:     status.PENDING_CODE,
			status:   status.PENDING,
		},
		{
			testName: "canceled",
			code:     status.CANCELED_CODE,
			status:   status.CANCELED,
		},
		{
			testName: "delivering",
			code:     status.DELIVERING_CODE,
			status:   status.DELIVERING,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, status.StatusCodeToString(tc.code), tc.status)
	}
}

func TestStatusToCode(t *testing.T) {
	tests := []test{
		{
			testName: "creating",
			code:     status.CREATING_CODE,
			status:   status.CREATING,
		},
		{
			testName: "pending",
			code:     status.PENDING_CODE,
			status:   status.PENDING,
		},
		{
			testName: "canceled",
			code:     status.CANCELED_CODE,
			status:   status.CANCELED,
		},
		{
			testName: "delivering",
			code:     status.DELIVERING_CODE,
			status:   status.DELIVERING,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, status.StatusToCode(tc.status), tc.code)
	}
}
