package solutionone

import (
	"fmt"
	"strings"
)

const vocal string = "aiueo"
const consonant string = "bcdfghijklmnpqrstvwxyz"

// CountWord : menghitung huruf mati dan hidup
/*
	@Param string
	@Return totalVocal, totalConsonant
*/
func CountWord(word string) (int, int) {
	splitWord := strings.Split(word, "")
	var totalVocal int
	var totalConsonant int
	var testedWord string

	for _, item := range splitWord {
		if strings.ContainsAny(item, vocal) {
			if !strings.ContainsAny(item, testedWord) {
				totalVocal++
				testedWord += item
			}
		}
		if strings.ContainsAny(item, consonant) {
			totalConsonant++
		}
	}

	fmt.Println("Huruf hidup: ", totalVocal)
	fmt.Println("Huruf mati: ", totalConsonant)

	return totalVocal, totalConsonant
}
