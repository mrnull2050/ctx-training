package finaltrain

import (
	"context"
	"fmt"
	"time"
)

func LoadData(ctx context.Context) {
	fmt.Println("Loading data start")
	userName := ctx.Value(UN)
	Passwrod := ctx.Value(PW)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("load data sesseccfuly", "user name", userName, "password", Passwrod)
	case <-ctx.Done():
		fmt.Println("request time out", ctx.Err())
	}
}
