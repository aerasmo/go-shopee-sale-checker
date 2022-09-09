package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	if int(now.Month()) == now.Day() {
		fmt.Println("There is a sale today.")
	} else {
		fmt.Println("There is no sale today.")
	}
}
