package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Println("a ==", a[:cap(a)]) // a == [1 2 3 4 5]
	fmt.Println("s ==", s[:cap(s)]) // s == [2 3 4 5]

	c := append(s, 6)
	fmt.Println("c ==", c[:cap(c)]) // c == [2 3 6 5]

	d := append(a, 6)
	fmt.Println("d ==", d[:cap(d)]) // d == [1 2 3 6 5 6 0 0 0 0]
}
