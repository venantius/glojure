package main

import (
	"bufio"
	"fmt"
	"os"
)

// A file for playing with things i don't understand in go

func main() {

	r, _ := os.Open("derp.txt")
	s := bufio.NewReader(r)
	ch, x, _ := s.ReadRune()
	for s.ReadRune() {
		fmt.Println("\"")
		fmt.Println(s.Text())
	}
}
