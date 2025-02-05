package client

import (
	"context"
	"time"
)

// Client is a DNS client impl.
type Client interface {
	ExpireTime(ctx context.Context, domain string) (time.Time, error)
}
