package utils

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
	"net/http"
)

var (
	ErrInvalidUUID = errors.New("the provided ID is not a valid UUID")
)

func ParseToUUID(value string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(value)
	if err != nil {
		return uuid.UUID{}, ErrInvalidUUID
	}

	return parsed, nil
}

func RespondWithJSON(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if response != nil {
		_ = json.NewEncoder(w).Encode(response)
	}
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, map[string]string{"message": message})
}

func DecodeJSONBody(body io.ReadCloser, e interface{}) error {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(e); err != nil {
		return errors.New("invalid JSON body")
	}

	return nil
}
