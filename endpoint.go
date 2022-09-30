package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addRequest struct {
	Name string `json:"name"`
}

type removeRequest struct {
	ID int `json:"id"`
}

type addResponse struct {
	err error
}

type removeResponse struct {
	err error
}

type getAllResponse struct {
	payload []model
	err     error
}

func makeAddEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		input := request.(addRequest)
		err = s.add(input.Name)
		return &addResponse{
			err: err,
		}, nil
	}
}
func makeRemoveEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		input := request.(removeRequest)
		err = s.remove(input.ID)
		return &removeResponse{
			err: err,
		}, nil
	}
}

func makeGetAllEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := s.getAll()
		return &getAllResponse{
			payload: res,
			err:     err,
		}, nil
	}
}
