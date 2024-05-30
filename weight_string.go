package main

import "fmt"

/**
a = 1
b = 2
c = 3
d = 4

bb = 4
ccc = 9
d = 4

*/

func WeightedStrings(word string, queries []int) []string {

	// create weight each character from a - z
	alphabetWeight := make(map[string]int)
	counter := 1
	for alp := 'a'; alp <= 'z'; alp++ {
		alphabetWeight[string(alp)] = counter
		counter++
	}

	// create weight each character from the input
	weightFound := make(map[string]int)
	for i := 0; i < len(word); i++ {
		weightFound[string(word[i])] = alphabetWeight[string(word[i])]

		// count weight for repeated characters
		count := 1
		concatenated := string(word[i])
		for i+1 < len(word) && word[i] == word[i+1] {
			concatenated = concatenated + string(word[i+1])
			count++
			i++
		}
		if count > 1 {
			weightFound[concatenated] = alphabetWeight[string(word[i])] * count
		}

	}

	result := make([]string, len(queries))
	for key, query := range queries {
		if isFound(query, weightFound) {
			result[key] = "Yes"
		} else {
			result[key] = "No"
		}
	}

	fmt.Println(result)
	return result
}

func isFound(query int, weightFound map[string]int) bool {
	for _, value := range weightFound {
		if value == query {
			return true
		}
	}
	return false
}
