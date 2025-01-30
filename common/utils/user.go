package utils

import (
	"context"
	"encoding/json"
)

func GetUserId(ctx context.Context) int64 {
	userId, _ := ctx.Value("user_id").(json.Number).Int64()

	return userId
}
