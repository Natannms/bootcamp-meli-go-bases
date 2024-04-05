package main

import (
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("mommento do error: %v", time.Now())
	fmt.Println("Error: ", err)
}
