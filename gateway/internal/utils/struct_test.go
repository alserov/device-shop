package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type testDecodeStruct struct {
	String  string `json:"string"`
	Integer int    `json:"integer"`

	MustError bool
}

func testString(s *testDecodeStruct) error {
	if len(s.String) < 3 {
		return errors.New("too short")
	}

	if len(s.String) > 5 {
		return errors.New("too long")
	}

	return nil
}

func testInt(s *testDecodeStruct) error {
	if s.Integer < 0 {
		return errors.New("too small")
	}

	if s.Integer > 5 {
		return errors.New("too big")
	}

	return nil
}

func TestDecode(t *testing.T) {
	tests := []testDecodeStruct{
		{
			String:    "",
			Integer:   3,
			MustError: true,
		},
		{
			String:    "abc",
			Integer:   -1,
			MustError: true,
		},
		{
			String:    "abc",
			Integer:   3,
			MustError: false,
		},
	}

	for _, tc := range tests {
		b, _ := json.Marshal(tc)
		reader := bytes.NewReader(b)
		req, _ := http.NewRequest(http.MethodPost, "/", reader)

		decoded, err := Decode[testDecodeStruct](req, testInt, testString)

		switch tc.MustError {
		case true:
			assert.Error(t, err)
		case false:
			assert.NoError(t, err)
			assert.Equal(t, tc, *decoded)
		}
	}
}
