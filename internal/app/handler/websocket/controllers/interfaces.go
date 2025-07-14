package controllers

import (
	"context"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type Request interface {
	Request() *pb.Request
	Ctx() context.Context
	SendResponse(ctx context.Context, response *pb.Response) error
}
