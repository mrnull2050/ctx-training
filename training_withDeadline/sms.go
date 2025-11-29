package trainingwithdeadline

import (
	"context"
	"fmt"
	"time"
)

func SMS(ctx context.Context){
	select{
	case<-time.After(1 * time.Second):
		fmt.Println("SMS send")
	case<-ctx.Done():
		fmt.Println("deadline reach")
	}
}