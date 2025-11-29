package trainingwithtimeout

import (
	"context"
	"fmt"
	"time"
)

func SmsEmail(ctx context.Context) {
	fmt.Println("send email start....")
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("sms send ok")
	case <-ctx.Done():
		fmt.Println("email err : ", ctx.Err())
	}
}
