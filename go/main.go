package main

import (
	"clojure-go/go/lang"
	"fmt"
)

func main() {
	fmt.Println("Fuck!")
	v := lang.CreateVector(1)
	fmt.Println(v.String())
	y := v.Cons(3)
	fmt.Println(y.String())

}
