package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/manosriram/docker-kubernetes/endpoints"
	"github.com/manosriram/docker-kubernetes/pb"
	"github.com/manosriram/docker-kubernetes/service"
	transport "github.com/manosriram/docker-kubernetes/transports"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	greetService := service.NewService(logger)
	greetEndpoint := endpoints.MakeEndpoints(greetService)
	grpcServer := transport.NewGRPCServer(greetEndpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterGreetServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	request := pb.GreetRequest{
		Name: "Mano",
	}
	response, err := grpcServer.Greet(context.TODO(), &request)
	fmt.Println(response)
	level.Error(logger).Log("exit", <-errs)
}
