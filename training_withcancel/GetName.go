package trainingwithcancel

import (
	"context"
	"fmt"
	"time"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	go GetName(ctx)
	cancel()
	time.Sleep(1 * time.Second)

}

func GetName(ctx context.Context) {
	fmt.Println("what is yout name")
	var name string
	fmt.Println(&name)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("your name is ", name)
	case <-ctx.Done():
		fmt.Println("err", ctx.Err())
	}
}
