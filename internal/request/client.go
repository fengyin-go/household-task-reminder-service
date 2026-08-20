package request

import "context"

type Client struct{ worker *Worker }

func NewClient(worker *Worker) *Client { return &Client{worker: worker} }

func (c *Client) Send(ctx context.Context, id string) error {
	return c.worker.Persist(ctx, id)
}
