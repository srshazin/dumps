package main

import (
	"fmt"
)

type MyWriter struct {
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	println(string(p))
	return 0, nil
}

func test() {

	writer := &MyWriter{}

	fmt.Fprintln(writer, "hello world")

}
