package finaltrain

import (
	"context"
	"fmt"
	"time"
)

func SendWelcome(ctx context.Context) {
	fmt.Println("sending sms...")
	userName := ctx.Value(UN)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("welcome to application", userName)
	case <-ctx.Done():
		fmt.Println("time out")
	}
}
