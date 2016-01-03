package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))  // 5
	fmt.Println(len(&arr)) // 5
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice)) // 5
	str := "12345"
	fmt.Println(len(str)) // 5
	ch := make(chan string, 10)
	ch <- str
	ch <- str
	ch <- str
	ch <- str
	ch <- str
	fmt.Println(len(ch)) // 5
}
