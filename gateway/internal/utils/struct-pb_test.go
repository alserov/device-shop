package utils

import (
	"bytes"
	"encoding/json"
	"github.com/alserov/device-shop/gateway/pkg/models"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/runtime/protoiface"
	"io"
	"net/http"
	"testing"
)

type test struct {
	PBMessage protoiface.MessageV1
	Struct    interface{}
	IsValid   bool
}

func TestRequestToPBMessage(t *testing.T) {
	tests := []test{
		{
			PBMessage: &pb.Device{},
			Struct: models.Device{
				Manufacturer: "test",
				UUID:         uuid.New().String(),
				Title:        "test",
				Description:  "test",
				Price:        1.0,
				Amount:       1,
			},
			IsValid: true,
		},
	}

	for _, ts := range tests {
		marshalled, err := json.Marshal(ts.Struct)
		if err != nil {
			t.Errorf("failed to encode struct\terror: %v", err)
			continue
		}
		body := bytes.NewBuffer(marshalled)

		req := &http.Request{
			Body: io.NopCloser(body),
		}

		_, err = RequestToPBMessage[models.Device, pb.Device](req, DeviceToPB)
		if err != nil && ts.IsValid {
			t.Errorf("failed to transform struct to pb message\terror: %v", err)
			continue
		}
	}
}
