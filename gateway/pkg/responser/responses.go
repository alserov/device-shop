package responser

import (
	"encoding/json"
	"github.com/alserov/device-shop/gateway/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
)

type WithError struct {
	Error string `json:"error"`
}

type responser struct {
	w   http.ResponseWriter
	log *slog.Logger
}

type Responser interface {
	UserError(msg string)
	ServerError()
	Value(msg interface{})
	Data(data H)
	NotAllowed()
	OK()

	HandleServiceError(err error, action string, log *slog.Logger)
}

func NewResponser(w http.ResponseWriter) Responser {
	return &responser{
		w: w,
	}
}

func (r *responser) OK() {
	r.w.WriteHeader(http.StatusOK)
}

func (r *responser) UserError(msg string) {
	r.w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(r.w).Encode(&WithError{
		Error: msg,
	})
}

func (r *responser) NotAllowed() {
	r.w.WriteHeader(http.StatusMethodNotAllowed)
}

func (r *responser) ServerError() {
	r.w.WriteHeader(http.StatusInternalServerError)
}

type H map[string]interface{}

func (r *responser) Data(data H) {
	r.w.WriteHeader(http.StatusOK)
	json.NewEncoder(r.w).Encode(data)
}

func (r *responser) Value(msg interface{}) {
	r.w.WriteHeader(http.StatusOK)
	json.NewEncoder(r.w).Encode(msg)
}

func (r *responser) HandleServiceError(err error, action string, log *slog.Logger) {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.Internal:
			log.Error("failed to execute grpc function", logger.Error(err, action))
			r.ServerError()
		default:
			r.UserError(st.Message())
		}
		return
	}
	log.Error("unexpected error", logger.Error(err, action))
	r.ServerError()
}
