package utils

import (
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
	"io"
	"net/http"
)

func RequestToPBMessage[T interface{}](req io.ReadCloser, w http.ResponseWriter) (*T, error) {
	bytes, err := io.ReadAll(req)
	if err != nil {
		return nil, err
	}
	var r protoiface.MessageV1
	if err = proto.Unmarshal(bytes, r); err != nil {
		return nil, err
	}
	if err = models.Validate(&r); err != nil {
		responser.UserError(w, err.Error())
		return nil, err
	}

	res := r.(T)
	return &res, nil
}
