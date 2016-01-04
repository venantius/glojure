package main

import (
	"clojure-go/go/lang"
	"fmt"
)

func main() {
	v := lang.CreateVector(1, 2)
	fmt.Println(v.String())
	y := v.Cons(3)
	fmt.Println(y.String())
	z := v.AssocN(0, "other")
	fmt.Println(z.String())
	fmt.Println(z)

	fmt.Println("asdf"[3:])

}
