package client

import (
	"context"
	"fmt"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	conn *websocket.Conn
}

func NewClient(ctx context.Context, addr string) (*Client, error) {
	cfg, err := websocket.NewConfig(addr, "https://test-client.com")
	if err != nil {
		return nil, fmt.Errorf("error make config. %w", err)
	}

	conn, err := cfg.DialContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error make client. %w", err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Request(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	requestBytes, err := proto.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshal request. %w", err)
	}

	if _, err := c.conn.Write(requestBytes); err != nil {
		return nil, fmt.Errorf("error write request. %w", err)
	}

	var responseBytes []byte

	if err := websocket.Message.Receive(c.conn, &responseBytes); err != nil {
		return nil, fmt.Errorf("error receive response. %w", err)
	}

	var response pb.Response

	if err := proto.Unmarshal(responseBytes, &response); err != nil {
		return nil, fmt.Errorf("error unmarshal response. %w", err)
	}

	return &response, nil
}

func (c *Client) Close() {
	_ = c.conn.WriteClose(0)
}
