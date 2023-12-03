package responser

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type WithError struct {
	Error string `json:"error"`
}

func UserError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&WithError{
		Error: msg,
	})
}

func NotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func ServerError(w http.ResponseWriter, lg *slog.Logger, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	lg.Error(err.Error())
}

type H map[string]interface{}

func Data(w http.ResponseWriter, msg H) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func Value(w http.ResponseWriter, msg interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
