package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(cap(arr))  // 5
	fmt.Println(cap(&arr)) // 5
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(slice)) // 5
	str := "12345"
	ch := make(chan string, 10)
	ch <- str
	ch <- str
	ch <- str
	ch <- str
	ch <- str
	fmt.Println(cap(ch)) // 5
}
