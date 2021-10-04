package transport

import (
	"context"

	"github.com/go-kit/log"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/manosriram/docker-kubernetes/endpoints"
	"github.com/manosriram/docker-kubernetes/pb"
)

type gRPCServer struct {
	Greeter gt.Handler
	pb.UnimplementedGreetServiceServer
}

func decodeGreetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GreetRequest)
	return endpoints.GreetRequest{Name: req.Name}, nil
}

func encodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GreetResponse)
	return &pb.GreetResponse{Greet: resp.Greeting}, nil
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.GreetServiceServer {
	return &gRPCServer{
		Greeter: gt.NewServer(
			endpoints.Greet,
			decodeGreetRequest,
			encodeResponse,
		),
	}
}

func (s *gRPCServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	_, resp, err := s.Greeter.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GreetResponse), nil
}
