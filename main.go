package main

import (
	"context"
	"github.com/Abubakarr99/targeter/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}
