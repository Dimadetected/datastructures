package main

import (
	"fmt"

	"github.com/goavengers/go-datastructure/tree"
)

func main() {
	t := tree.NewNode(8)
	t.Insert(1)
	t.Insert(2)
	t.Insert(11)
	t.Insert(9)
	t.PrintInOrder()
	fmt.Println("=====")
	t.Insert(4)
	t.Insert(6)
	t.Insert(3)
	t.PrintInOrder()
}
