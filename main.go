package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func intFromWord(word string) int {
	resultInt := 0
	for _, rr := range word {
		if strings.Contains("1234567890", string(rr)) {
			resultInt, _ = strconv.Atoi(string(rr))
			break
		}
	}
	return resultInt
}

// func Order(sentence string) string {
// 	if sentence == "" {
// 		return ""
// 	}

// 	intSli := make([]int, 0) // len, cap 1
// 	sentenceSli := strings.Split(sentence, " ")
// 	for _, word := range sentenceSli {
// 		// extract the number from the `word`
// 		intSli = append(intSli, intFromWord(word))
// 	}

// 	sort.Ints(intSli)
// 	strSli := make([]string, 0)
// 	for _, myInt := range intSli {
// 		chosenWord := ""
// 		for _, str := range sentenceSli {
// 			if strings.Contains(str, strconv.Itoa(myInt)) {
// 				chosenWord = str
// 				break
// 			}
// 		}
// 		strSli = append(strSli, chosenWord)
// 	}

// 	return strings.Join(strSli, " ")
// }

type byInt []string

func (s byInt) Len() int {
	return len(s)
}

func (s byInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byInt) Less(i, j int) bool {
	// customized logic
	return intFromWord(s[i]) < intFromWord(s[j])
}

func Order(sentence string) string {
	sentenceSli := strings.Split(sentence, " ")
	sort.Sort(byInt(sentenceSli))
	return strings.Join(sentenceSli, " ")
}

func main() {
	fmt.Printf(">" + Order("is2 Thi1s T4est 3a") + "\n")
	// fmt.Printf(">" + Order2("is2 Thi1s T4est 3a") + "\n")
	// fmt.Println("a" > "b")
}
