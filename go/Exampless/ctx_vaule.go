package main

import (
	"context"
	"fmt"
)

type key int

const userIDKey key = iota

func main() {
	// 创建一个根 context
	ctx := context.Background()

	// 在根 context 基础上创建一个带有用户 ID 的 context
	ctxWithUserID := context.WithValue(ctx, userIDKey, 123)

	// 在子 context 中获取用户 ID
	userID, ok := ctxWithUserID.Value(userIDKey).(int)
	if !ok {
		fmt.Println("User ID not found in context")
	} else {
		fmt.Println("User ID:", userID)
	}
}
