package main

import (
	"fmt"
)

//START OMIT
func main() {
	a := []*int{new(int), new(int)}
	fmt.Println(a)
	b := a[:1]
	fmt.Println(b) // [&0]

	// second element is not garbage collected, because it's *still* accessible
	c := b[:2] //[&0[
	fmt.Println(c)
}

//END OMIT
