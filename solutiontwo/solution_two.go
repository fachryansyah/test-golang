package solutiontwo

import (
	"fmt"
	"sort"
	"strings"
)

const vocal string = "aiueo"
const consonant string = "bcdfghijklmnpqrstvwxyz"

// SortWord : mengurutkan huruf mati dan hidup
/*
	@Param string
	@Return result string
*/
func SortWord(word string) string {
	fmt.Println(word)

	var vocalWordGroup []string
	var consonantWordGroup []string
	var mergingGroupOfWord []string
	var result string

	splitWord := strings.Split(word, "")
	for _, item := range splitWord {
		if strings.ContainsAny(item, vocal) {
			vocalWordGroup = append(vocalWordGroup, item)
		}
		if strings.ContainsAny(item, consonant) {
			consonantWordGroup = append(consonantWordGroup, item)
		}
	}

	sort.Strings(vocalWordGroup)
	sort.Strings(consonantWordGroup)

	mergingGroupOfWord = append(mergingGroupOfWord, vocalWordGroup...)
	mergingGroupOfWord = append(mergingGroupOfWord, consonantWordGroup...)

	for _, item := range mergingGroupOfWord {
		result += item
	}

	return result
}
