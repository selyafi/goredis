package client

import "context"

type Client struct {
}

func NewClient() *Client {
	return &Client{}

}

func (c *Client) Set(ctx context.Context, key string, val string) {

}
