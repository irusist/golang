package main

import "fmt"

func testPanic() {
	panic("this is panic")
}

func testRecover(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()

	f()
	return
}

func main() {
	result := testRecover(testPanic)
	fmt.Println("result is:%b", result)
}