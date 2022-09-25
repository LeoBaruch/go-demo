package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello, demo!")
	fmt.Println(morestrings.ReverseRunes("!og, olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
