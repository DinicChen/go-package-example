package main

import (
	"fmt"
)

func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()

	fun()
}

func say(s string) {
	fmt.Println(s)
}

func main() {
	say("Hello")

	Try(
		func() {
			panic("World")
		},
		func(e interface{}) {
			fmt.Println("catch", e)
		},
	)

	say("end")
}

// Hello
// catch World
// end
