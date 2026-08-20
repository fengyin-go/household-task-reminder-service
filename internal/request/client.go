package request

import "context"

type Client struct{ worker *Worker }

func NewClient(worker *Worker) *Client { return &Client{worker: worker} }

func (c *Client) Send(ctx context.Context, id string) error {
	if !Active(ctx) {
		return context.Canceled
	}
	return c.worker.Persist(ctx, id)
}
