// Sample code @vnitc
package main

import (
	"fmt"
	"unsafe"
)

func foo() *int {
	t := 2
	return &t
}

type MyInt64 int64

func (mi MyInt64) dump() {
	fmt.Println(mi)
}

func main() {
	// Go can handle kanji and more :)
	fmt.Println("合コンにいこう！")

	x := foo()
	fmt.Println(*x)

	// t := 2
	// fmt.Println(t++)
	// fmt.Println(z := 3)

	// var t MyInt64 = 10
	// fmt.Println(t)
	// var t2 int64 = t

	fmt.Println(unsafe.Sizeof(struct{}{}))
}
