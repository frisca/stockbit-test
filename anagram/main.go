package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	tamps := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		// process anagram untuk setiap string dalam array
		word := ProcessAnagram(strs[i])

		// word dimasukkan kedalam array dengan key nya bertipe string dari hasil join tersebut
		tamps[word] = append(tamps[word], strs[i])
	}

	for _, tamp := range tamps {
		fmt.Println(tamp)
	}
}

// function anagram is process function check anagram and group words ...
func ProcessAnagram(word string) string {
	var res string

	// string dipisahkan dengan tanda ""
	splitWord := strings.Split(word, "")

	// string disort/diurutkan
	sort.Strings(splitWord)

	// string yang sudah diurutkan/disort digabung
	res = strings.Join(splitWord, "")

	return res
}
