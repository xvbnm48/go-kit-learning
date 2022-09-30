package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func makeHandler(s Service) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeErr),
	}

	addHandler := httptransport.NewServer(
		makeAddEndpoints(s),
		decodeAddRequest,
		encodeAddResponse,
		options...,
	)

	removeHandler := httptransport.NewServer(
		makeRemoveEndpoints(s),
		decodeRemoveRequest,
		encodeRemoveResponse,
		options...,
	)
	getAllHandler := httptransport.NewServer(
		makeGetAllEndpoints(s),
		decodeGetAllRequest,
		encodeGetAllResponse,
		options...,
	)

	r := chi.NewRouter()
	r.Route("/items", func(r chi.Router) {
		r.Get("/", getAllHandler.ServeHTTP)
		r.Post("/", addHandler.ServeHTTP)
		r.Delete("/{ID}", removeHandler.ServeHTTP)
	})

	return r
}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrBadRequest
	}

	return request, nil
}

func decodeRemoveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		return nil, ErrInvalidId
	}

	return removeRequest{
		ID: id,
	}, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func encodeAddResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*addResponse)
	return json.NewEncoder(w).Encode(res.err)
}

func encodeGetAllResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*getAllResponse)
	if res != nil {
		return json.NewEncoder(w).Encode(res.err)
	}

	return json.NewEncoder(w).Encode(res.payload)
}

func encodeRemoveResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*removeResponse)
	return json.NewEncoder(w).Encode(res.err)
}

func encodeErr(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case ErrInvalidId:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var ErrBadRequest = errors.New("bad request")
var ErrInvalidId = errors.New("invalid id")
