package gotest1

import "fmt"

// say Hi to someone
func SayHi(name string) string {
	return fmt.Sprintf("Hi,gotest1  version 1.0.1" +
		"test1 %s", name)
}