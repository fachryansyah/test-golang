package main

import (
	"fmt"
	"strings"
)

const hurufHidup string = "aiueo"
const hurufMati string = "bcdfghijklmnpqrstvwxyz"

func main() {
	hitung("omama")
}

func hitung(huruf string) {
	splitHuruf := strings.Split(huruf, "")
	fmt.Println(splitHuruf)
}
