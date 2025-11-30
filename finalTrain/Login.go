package finaltrain

import (
	"context"
	"fmt"
	"time"
)



func Login(ctx context.Context) {
	fmt.Println("login start")
	
	select {
	case <-time.After(1 * time.Second):
		LoadData(ctx)
	case <-ctx.Done():
		fmt.Println("time out ", ctx.Err())
	}
}
