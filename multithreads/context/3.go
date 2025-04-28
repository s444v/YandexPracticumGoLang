package main

import (
	"context"
	"fmt"
)

func action(ctx context.Context) {
	fmt.Println(ctx.Value("func"), ctx.Value("uid"), ctx.Value("user"))
}

func main() {
	ctx := context.WithValue(context.Background(), "func", "main")

	action(context.WithValue(ctx, "uid", 701))
}
