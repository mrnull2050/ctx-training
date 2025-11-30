package trainingwithvalue

import "context"

func Service(ctx context.Context) {
	println("service start")
	Repo(ctx)
}
