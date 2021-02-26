// This is just a short program for (my) educational purposes.
// It doesn't do anything useful, I'm just learning Go.
// This should be a Godoc comment.
package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
)

/*	is this a multiline comment?,
	good
*/

func main() {
	// for now just do it sequentially.
	fmt.Println("start here")
	fileContents := readMyFile("dracula.txt")
	fmt.Printf("%.50s...\n", fileContents)
	// split on word boundary
	tokens := splitString(fileContents)
	dict := createDictionary(tokens)
	fmt.Printf("Value for \"the\": %d\n", dict["the"])
	sortedDict := sortDictionary(dict)
	fmt.Printf("Value for \"the\": %d\n", sortedDict["the"])
	fmt.Println("fin")

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

	type ttuple struct {
		key   string
		value int
	}

	// make a slice with capacity equal to the size of the dictionary
	fmt.Printf("Dict has %d entries\n", len(dict))
	var words = make([]ttuple, len(dict))

	// populate the slice with the word,count pair
	iterations := 0
	for k, v := range dict {
		iterations++
		var entry = ttuple{k, v}
		words = append(words, entry)
	}
	fmt.Printf("Iterated through %d entries\n", iterations)
	// now sort the slice, most common words first
	// but how do I reverse it, sort.Sort(sort.Reverse(sort.Slice))) is showing an error
	// and I don't want to change my "less" function to actually return "more"

	// looks like I have to implement the Interface
	//TODO: take out the hack and implement full interface to allow use
	// of sort.Reverse
	fmt.Printf("!Words has %d entries\n", len(words))
	sort.Slice(words, func(i, j int) bool { return words[i].value > words[j].value })
	fmt.Printf("!Words has %d entries\n", len(words))

	sortedDict := make(map[string]int)
	iterations = 0
	for i, entry := range words {
		//fmt.Printf("Pos: %d, key: %s, value: %d\n", i, entry.key, entry.value)
		iterations++
		sortedDict[entry.key] = i
	}
	fmt.Printf("Iterated through %d entries\n", iterations)
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

func splitString(source string) []string {
	re := regexp.MustCompile("/\\w+|\\s+|[^\\s\\w]+/g")
	return re.Split(source, -1)
}
