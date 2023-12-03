package converter

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test[T any] struct {
	in T
}

func TestDeviceToRepoStruct(t *testing.T) {
	tests := []test[pb.Device]{
		{
			in: pb.Device{
				UUID:         "uuid",
				Title:        "title",
				Description:  "description",
				Price:        1,
				Manufacturer: "manu",
				Amount:       1,
			},
		},
	}

	for _, tc := range tests {
		convertedStr := DeviceToRepoStruct(&tc.in)

		assert.Equal(t, tc.in.UUID, convertedStr.UUID)
		assert.Equal(t, tc.in.Title, convertedStr.Title)
		assert.Equal(t, tc.in.Manufacturer, convertedStr.Manufacturer)
		assert.Equal(t, tc.in.Amount, convertedStr.Amount)
	}
}
