package signals_test

import (
	"context"
	"testing"

	"github.com/joetifa2003/signals"
)

func TestBatch(t *testing.T) {
	ctx := context.Background()

	x := signals.New(1)
	y := signals.New(2)

	runs := 0
	signals.Effect(ctx, func(ctx context.Context) {
		runs++
		_ = x.Get(ctx) + y.Get(ctx)
	})

	signals.Batch(ctx, func(ctx context.Context) {
		x.Set(ctx, 5)
		y.Set(ctx, 10)
	})

	if runs != 2 {
		t.Errorf("expected 2 runs, got %d", runs)
	}
}
