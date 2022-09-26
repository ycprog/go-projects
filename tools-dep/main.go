package main

import "fmt"

//go:generate ./bin/mockery --name IceCream
type IceCream interface {
	Price() int
}

func main() {
	fmt.Println("Hello Ice Cream.")
}
