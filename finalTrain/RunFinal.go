package finaltrain

import (
	"context"
	"time"
)

type CtxKey string

const (
	UN CtxKey = "UserName"
	PW CtxKey = "PassWord"
)

func RunFinal() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.WithValue(context.Background(), UN, "ADMIN")
	ctx = context.WithValue(ctx, PW, "1234")
	defer cancel()

	Login(ctx)

	SendWelcome(ctx)
	time.Sleep(2 * time.Second)

}
