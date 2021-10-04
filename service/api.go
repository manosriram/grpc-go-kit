package service

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
)

// Function Definition of ProtoBuf APIs are done in this file.

type service struct {
	logger log.Logger
}

type Service interface {
	Greet(ctx context.Context, name string) string
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Greet(ctx context.Context, name string) string {
	return fmt.Sprintf("Hello, %v", name)
}
