package helpers

import (
	"context"
	"math/rand"
	"time"
)

func Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Millisecond * time.Duration(rand.Int63n(2000))):
		return nil
	}
}
