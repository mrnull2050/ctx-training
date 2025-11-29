package trainingwithtimeout

import (
	"context"
	"fmt"
	"time"
)

func Paying(ctx context.Context) {
	fmt.Println("paying start")
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("paying done")
	case <-ctx.Done():
		fmt.Println("paying err", ctx.Err())

	}
}
