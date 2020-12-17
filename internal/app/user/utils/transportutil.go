package utils

import (
	"blogger-kit/internal/pkg/requests"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func DecodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return nil, err
	}
	var reg requests.RegisterRequest
	if err = json.Unmarshal(body, &reg); err != nil {
		log.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	if reg.Username == "" || reg.Password == "" || reg.Email == "" {
		return nil, ErrorBadRequest
	}

	return &requests.RegisterRequest{
		Username: reg.Username,
		Email:    reg.Email,
		Password: reg.Password,
	}, nil
}

func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return nil, err
	}
	var reg requests.LoginRequest
	if err = json.Unmarshal(body, &reg); err != nil {
		log.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	if reg.Email == "" || reg.Password == "" {
		return nil, ErrorBadRequest
	}
	return &requests.LoginRequest{
		Email:    reg.Email,
		Password: reg.Password,
	}, nil
}

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
