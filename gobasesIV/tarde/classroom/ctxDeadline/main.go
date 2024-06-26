package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	deadline := time.Now().Add(time.Second * 5)

	ctx, _ = context.WithDeadline(ctx, deadline)

	// context.WithTimeout(ctx, time.Second*5)
  
	<-ctx.Done()
	fmt.Println(ctx.Err().Error())
}
