// Main demonstrates a race condition when the two go routines try to operate on the same shared global variable count.
package main

var count int

func race() {
	count++
}

func main() {
	go race()
	go race()
}
