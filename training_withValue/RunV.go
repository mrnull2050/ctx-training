package trainingwithvalue

import "context"

type CtxKey string

const (
	UserId CtxKey = "UserId"
	ReqId  CtxKey = "ReqId"
)

func RunV() {
	ctx := context.WithValue(context.Background(), ReqId, "A-099")
	ctx = context.WithValue(ctx, UserId, "Amiryasin")
	Service(ctx)
}
