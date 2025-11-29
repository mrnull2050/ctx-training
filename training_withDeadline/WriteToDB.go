package trainingwithdeadline

import (
	"context"
	"fmt"
	"time"
)

func WriteToDB(ctx context.Context) {
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("write to db")
	case <-ctx.Done():
		fmt.Println("deadline reach", ctx.Err())
	}
}
