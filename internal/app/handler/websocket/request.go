package websocket

import (
	"context"
	"fmt"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type Request struct {
	session *Session

	req *pb.Request
}

func NewRequest(session *Session, req *pb.Request) *Request {
	return &Request{session: session, req: req}
}

func (r *Request) Request() *pb.Request {
	return r.req
}

func (r *Request) Ctx() context.Context {
	return r.session.ctx()
}

func (r *Request) SendResponse(ctx context.Context, response *pb.Response) error {
	response.Id = r.req.GetId()

	if err := r.session.sendResponse(ctx, response); err != nil {
		return fmt.Errorf("error sending response. %w", err)
	}

	return nil
}
