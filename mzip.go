// This is just a short program for (my) educational purposes.
// It doesn't do anything useful, I'm just learning Go.
// This should be a Godoc comment.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
)

/*	is this a multiline comment?,
	good
*/
type book struct {
	Dictionary map[int]string `json:"dictionary"`
	Text       []int          `json:"text"`
}

func main() {
	// for now just do it sequentially.
	fmt.Println("start here")
	fileContents := readMyFile("dracula.txt")
	fmt.Printf("%.50s...\n", fileContents)
	// create an array split on word boundary
	tokens := splitString(fileContents)
	// from this array, create a dictionary of unique
	// tokens, with a count of occurences for each
	dict := createDictionary(tokens)
	fmt.Printf("Value for \"the\": %d\n", dict["the"])
	// now create a new dictionary, so the most common words
	// have the shortest numbers
	lookupMap := sortDictionary(dict)
	fmt.Printf("Value for \"the\": %d\n", lookupMap["the"])

	// now create an array by looking up the value for
	// each word
	encoded := encode(lookupMap, tokens)
	// when we save, we're going to need the reverse dictionary
	// to map the numbers back to the words.
	reverseDict := createReverseDict(lookupMap)
	// the book is the reverse dictionary plus the int array
	book := book{reverseDict, encoded}
	// write it to a json file
	writeMyFileAsJSON("intermediate.json", book)

	//end of part 1

	fmt.Println("fin")

}

func createReverseDict(lookupMap map[string]int) map[int]string {
	var result = make(map[int]string)
	for key, value := range lookupMap {
		result[value] = key
	}
	return result
}

func encode(lookupMap map[string]int, tokens []string) []int {

	// create a slice same length as input token slice
	// array length would need to be known before hand
	var result = make([]int, len(tokens))

	for i, element := range tokens {
		// convert each word to corresponding index
		result[i] = lookupMap[element]
	}

	return result
}

// create a dictionary of tokens, with the count of
// occurrences as the value
func createDictionary(tokens []string) map[string]int {
	dict := make(map[string]int)

	for _, element := range tokens {
		count := dict[element]
		count++
		dict[element] = count
		// fmt.Printf("%s: %d\n", element, count)
	}
	fmt.Printf("Tokens has %d entries\n", len(tokens))
	fmt.Printf("Dict has %d entries\n", len(dict))
	return dict

}

// Sort the dictionary by value, descending
// this will allow us to allocate the most common words
// the lowest value
// except there's no such thing as a sorted dictionary in go
// so we'll have to create an array of tuples
func sortDictionary(dict map[string]int) map[string]int {

	type wordCount struct {
		key   string
		value int
	}

	// make a slice with capacity equal to the size of the dictionary
	fmt.Printf("Dict has %d entries\n", len(dict))
	var words = make([]wordCount, 0, len(dict))

	// populate the slice with the word,count pair
	for k, v := range dict {
		var entry = wordCount{k, v}
		words = append(words, entry)
	}

	// now sort the slice, most common words first
	// but how do I reverse it, sort.Sort(sort.Reverse(sort.Slice))) is showing an error
	// and I don't want to change my "less" function to actually return "more"

	// looks like I have to implement the Interface
	//TODO: take out the hack and implement full interface to allow use
	// of sort.Reverse

	// Where less means earlier in sort
	sort.Slice(words, func(i, j int) bool { return words[i].value > words[j].value })

	sortedDict := make(map[string]int)
	for i, entry := range words {
		//fmt.Printf("Pos: %d, key: %s, value: %d\n", i, entry.key, entry.value)
		sortedDict[entry.key] = i
	}

	return sortedDict

}

// Reads the complete contents of the file, and returns it as a string
func readMyFile(filePath string) string {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func writeMyFileAsJSON(filePath string, contents book) {
	jsonifiedBook, err := json.Marshal(contents)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filePath, jsonifiedBook, 0666)
	if err != nil {
		panic(err)
	}
}

func splitString(source string) []string {
	// ha ha, this doesn't work as I'd like it in go
	// as it doesn't return the boundaries in the array.
	re := regexp.MustCompile("/\\w+|\\s+|[^\\s\\w]+/g")
	return re.Split(source, -1)
}
