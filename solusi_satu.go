package main

import (
	"fmt"
	"strings"
)

const hurufHidup string = "aiueo"
const hurufMati string = "bcdfghijklmnpqrstvwxyz"

func main() {
	Hitung("omama")
}

// Hitung : menghitung huruf mati dan hidup
/*
	@Param string
	@Return Println
*/
func Hitung(huruf string) {
	splitHuruf := strings.Split(huruf, "")
	fmt.Println(splitHuruf)
}
