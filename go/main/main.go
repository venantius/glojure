package main

import (
	"fmt"
	"glojure/go/lang"
	// "murmur3"
)

func main() {
	v := lang.CreateVector(1, 2)
	// fmt.Println(v.String())
	// y := v.Cons(3).(*lang.PersistentVector)
	z := v.AssocN(0, "other").(*lang.PersistentVector)
	fmt.Println(z)

	fmt.Println(lang.HashString("bananas in pajamas"))

}
