package trainingwithdeadline

import (
	"context"
	"fmt"
	"time"
)

func ValidateUser(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("validate user done")
	case <-ctx.Done():
		fmt.Println("Deadline reach", ctx.Err())
	}
}
