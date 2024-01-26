package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123")
	F1(ctx)
}

func F1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	F2(ctx)
}
func F2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	F3(ctx)
}
func F3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}
