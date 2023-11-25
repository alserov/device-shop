package utils

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test[T any] struct {
	testName         string
	in               T
	shouldThrowError bool /*if test should return an error*/
}

func TestCheckCollection(t *testing.T) {
	tests := []test[pb.ChangeCollectionReq]{
		{
			testName: "empty user uuid",
			in: pb.ChangeCollectionReq{
				DeviceUUID: "uuid",
				UserUUID:   "",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty device uuid",
			in: pb.ChangeCollectionReq{
				DeviceUUID: "",
				UserUUID:   "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "valid",
			in: pb.ChangeCollectionReq{
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
	tests := []test[pb.CreateDeviceReq]{
		{
			testName: "invalid price",
			in: pb.CreateDeviceReq{
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
			in: pb.CreateDeviceReq{
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
			in: pb.CreateDeviceReq{
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
			in: pb.CreateDeviceReq{
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
			in: pb.CreateDeviceReq{
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
			in: pb.CreateDeviceReq{
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
	tests := []test[pb.UpdateDeviceReq]{
		{
			testName: "invalid price",
			in: pb.UpdateDeviceReq{
				Price:       0,
				Title:       "title",
				Description: "description",
				UUID:        "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty title",
			in: pb.UpdateDeviceReq{
				Price:       1,
				Title:       "",
				Description: "description",
				UUID:        "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty description",
			in: pb.UpdateDeviceReq{
				Price:       1,
				Title:       "title",
				Description: "",
				UUID:        "uuid",
			},
			shouldThrowError: true,
		},
		{
			testName: "empty uuid",
			in: pb.UpdateDeviceReq{
				Price:       1,
				Title:       "title",
				Description: "description",
				UUID:        "",
			},
			shouldThrowError: true,
		},
		{
			testName: "valid",
			in: pb.UpdateDeviceReq{
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
