package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	len := copy(dst, src[4:])
	fmt.Println(len, dst) // 1 [5 0 0]
	len = copy(dst, src[0:])
	fmt.Println(len, dst) // 3 [1 2 3]
	len = copy(dst, src)
	fmt.Println(len, dst) // 3 [1 2 3]
}
