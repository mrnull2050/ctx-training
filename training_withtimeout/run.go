package trainingwithtimeout

import (
	"context"
	"time"
)

func RUN() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	time.Sleep(5 * time.Second)
	go Check(ctx)
	time.Sleep(2 * time.Second)
	go Paying(ctx)
	time.Sleep(2 * time.Second)

	go SmsEmail(ctx)
	time.Sleep(2 * time.Second)

	time.Sleep(5 * time.Second)
}
