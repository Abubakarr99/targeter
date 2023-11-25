package main

import (
	"context"
	"targeter/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}
