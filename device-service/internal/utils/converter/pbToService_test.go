package converter

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type test[T any] struct {
	in T
}

func TestCreateDeviceToRepoStruct(t *testing.T) {
	tests := []test[pb.CreateDeviceReq]{
		{
			in: pb.CreateDeviceReq{
				Amount:       1,
				Title:        "title",
				Description:  "description",
				Price:        1,
				Manufacturer: "manu",
			},
		},
	}

	for _, tc := range tests {
		convertedStruct := CreateDeviceToRepoStruct(&tc.in)

		assert.Equal(t, strings.ToLower(tc.in.Title), convertedStruct.Title)
		assert.Equal(t, strings.ToLower(tc.in.Description), convertedStruct.Description)
		assert.Equal(t, strings.ToLower(tc.in.Manufacturer), convertedStruct.Manufacturer)
		assert.Equal(t, tc.in.Amount, convertedStruct.Amount)
		assert.Equal(t, tc.in.Price, convertedStruct.Price)
	}
}

func TestUpdateDeviceToRepoStruct(t *testing.T) {
	tests := []test[pb.UpdateDeviceReq]{
		{
			in: pb.UpdateDeviceReq{
				UUID:        "uuid",
				Title:       "title",
				Description: "description",
				Price:       1,
			},
		},
	}

	for _, tc := range tests {
		convertedStruct := UpdateDeviceToRepoStruct(&tc.in)

		assert.Equal(t, tc.in.UUID, convertedStruct.UUID)
		assert.Equal(t, strings.ToLower(tc.in.Title), convertedStruct.Title)
		assert.Equal(t, strings.ToLower(tc.in.Description), convertedStruct.Description)
		assert.Equal(t, tc.in.Price, convertedStruct.Price)
	}
}
