package trainingwithdeadline

import (
	"context"
	"time"
)

func RunAppD() {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ValidateUser(ctx)
	WriteToDB(ctx)
	SMS(ctx)
	time.Sleep(100 * time.Millisecond)
}
