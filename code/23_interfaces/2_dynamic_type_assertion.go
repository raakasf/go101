package main

import "fmt"

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"
	// Now the dynamic type of y is "string".
	var w Writer
	var ok bool

	// Type DummyWriter implements both
	// Writer and interface{}.
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// The dynamic type of y is "string",
	// which doesn't implement Writer.
	w, ok = y.(Writer)
	fmt.Println(w, ok) // nil false
	w = y.(Writer)     // will panic
}
