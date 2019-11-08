package main

import (
	"fmt"
	"strings"
)

const vocal string = "aiueo"
const consonant string = "bcdfghijklmnpqrstvwxyz"

func main() {
	CountWord("omama")
}

// CountWord : menghitung huruf mati dan hidup
/*
	@Param string
	@Return Println
*/
func CountWord(word string) {
	splitWord := strings.Split(word, "")
	fmt.Println(splitWord)
}
