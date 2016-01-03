package main

import (
	"fmt"
)

func f() {
	defer func() {
		fmt.Println("f() defer")
	}()

	fmt.Println("f() before panic")
	panic(0)

	// 此后的代码不会被执行
	fmt.Println("after panic")
	defer func() {
		fmt.Println("f() defer agin")
	}()
}

func g() {
	defer func() {
		fmt.Println("g() defer")
	}()

	f()
	// 此后的代码不会被执行
	fmt.Println("after call f()")
}

func main() {
	g()
}

// f() before panic
// f() defer
// g() defer
// panic: 0

// goroutine 1 [running]:
// main.f()
//	/root/goTest/src/go-package-example/builtin/example/panic.go:13 +0x136
// main.g()
//	/root/goTest/src/go-package-example/builtin/example/panic.go:27 +0x3d
// main.main()
//	/root/goTest/src/go-package-example/builtin/example/panic.go:33 +0x14
// exit status 2
