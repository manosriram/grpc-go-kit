package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/manosriram/docker-kubernetes/service"
)

type Endpoints struct {
	Greet endpoint.Endpoint
}

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Greeting string `json:"greeting"`
}

func makeGreetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GreetRequest)
		return GreetResponse{
			Greeting: s.Greet(ctx, req.Name),
		}, nil
	}
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Greet: makeGreetEndpoint(s),
	}
}
