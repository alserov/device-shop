package validation

import (
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test[T any] struct {
	testName         string
	in               T
	shouldThrowError bool /*if test should return an error*/
}

func TestCheckCollection(t *testing.T) {
	tests := []test[collection.ChangeCollectionReq]{
		{
			testName: "empty user uuid",
			in: collection.ChangeCollectionReq{
				DeviceUUID: "uuid",
				UserUUID:   "",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty device uuid",
			in: collection.ChangeCollectionReq{
				DeviceUUID: "",
				UserUUID:   "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "valid",
			in: collection.ChangeCollectionReq{
				DeviceUUID: "uuid",
				UserUUID:   "uuid",
			},
			shouldThrowError: false,
		},
	}

	for _, tc := range tests {
		if tc.shouldThrowError {
			assert.Error(t, CheckCollection(&tc.in))
		} else {
			assert.NoError(t, CheckCollection(&tc.in))
		}
	}
}

func TestCheckCreateDevice(t *testing.T) {
	tests := []test[device.CreateDeviceReq]{
		{
			testName: "invalid price",
			in: device.CreateDeviceReq{
				Title:        "title",
				Description:  "description",
				Price:        0,
				Manufacturer: "manufacturer",
				Amount:       1,
			},
			shouldThrowError: true,
		},
		{
			testName: "invalid amount",
			in: device.CreateDeviceReq{
				Title:        "title",
				Description:  "description",
				Price:        1,
				Manufacturer: "manufacturer",
				Amount:       0,
			},
			shouldThrowError: true,
		},
		{
			testName: "empty title",
			in: device.CreateDeviceReq{
				Title:        "",
				Description:  "description",
				Price:        1,
				Manufacturer: "manufacturer",
				Amount:       1,
			},
			shouldThrowError: true,
		},
		{
			testName: "empty description",
			in: device.CreateDeviceReq{
				Title:        "title",
				Description:  "",
				Price:        1,
				Manufacturer: "manufacturer",
				Amount:       1,
			},
			shouldThrowError: true,
		},
		{
			testName: "empty manufacturer",
			in: device.CreateDeviceReq{
				Title:        "title",
				Description:  "description",
				Price:        1,
				Manufacturer: "",
				Amount:       1,
			},
			shouldThrowError: true,
		},
		{
			testName: "valid",
			in: device.CreateDeviceReq{
				Title:        "title",
				Description:  "description",
				Price:        1,
				Manufacturer: "manufacturer",
				Amount:       1,
			},
			shouldThrowError: false,
		},
	}

	for _, tc := range tests {
		if tc.shouldThrowError {
			assert.Error(t, CheckCreateDevice(&tc.in))
		} else {
			assert.NoError(t, CheckCreateDevice(&tc.in))
		}
	}
}

func TestCheckUpdateDevice(t *testing.T) {
	tests := []test[device.UpdateDeviceReq]{
		{
			testName: "invalid price",
			in: device.UpdateDeviceReq{
				Price:       0,
				Title:       "title",
				Description: "description",
				UUID:        "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty title",
			in: device.UpdateDeviceReq{
				Price:       1,
				Title:       "",
				Description: "description",
				UUID:        "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty description",
			in: device.UpdateDeviceReq{
				Price:       1,
				Title:       "title",
				Description: "",
				UUID:        "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty uuid",
			in: device.UpdateDeviceReq{
				Price:       1,
				Title:       "title",
				Description: "description",
				UUID:        "",
			},
			shouldThrowError: true,
		},
		{
			testName: "valid",
			in: device.UpdateDeviceReq{
				Price:       1,
				Title:       "title",
				Description: "description",
				UUID:        "uuid",
			},
			shouldThrowError: false,
		},
	}

	for _, tc := range tests {
		if tc.shouldThrowError {
			assert.Error(t, CheckUpdateDevice(&tc.in), tc.testName)
		} else {
			assert.NoError(t, CheckUpdateDevice(&tc.in), tc.testName)
		}
	}
}
