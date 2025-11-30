package trainingwithvalue

import (
	"context"
	"fmt"
)

func Repo(ctx context.Context) {
	userID := ctx.Value(UserId)
	reqID := ctx.Value(ReqId)
	fmt.Println("Repo : ", reqID, " UserID  : ", userID)
}
