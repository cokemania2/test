// 크로아티아 알파벳.go
// https://www.acmicpc.net/problem/2941

package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a*(b-1) + 1)
}
