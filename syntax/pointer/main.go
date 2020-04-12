package main

import (
	"fmt"
)

func main() {
	n := 100
	// 値渡し
	// コピーされるので、元のnに変化はない
	returnValue := increment(n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Value of n address is %v\n", &n)
	fmt.Printf("Return Value of increment is %d\n", returnValue)
	// 参照渡し
	// 戻り値を受け取らなくても渡した変数が書き換わる
	incrementWithPointer(&n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Value of n address is %v\n", &n)
}

func increment(n int) int {
	fmt.Printf("Value of n address in increment() is %v\n", &n)
	return n + 1
}
func incrementWithPointer(n *int) {
	fmt.Printf("Value of n address in incrementWithPointer() is %v\n", n)
	*n++
}
