package trainingwithtimeout

import (
	"context"
	"fmt"
	"time"
)

func Check(ctx context.Context) {
	fmt.Println("checking start...")
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Checking ok")
	case <-ctx.Done():
		fmt.Println("checking err", ctx.Err())
	}
}
