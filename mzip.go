// This is just a short program for (my) educational purposes.
// It doesn't do anything useful, I'm just learning Go.
// This should be a Godoc comment.
package main

import (
	"fmt"
	"os"
	"regexp"
)

/*	is this a multiline comment?,
	good
*/

func main() {
	fmt.Println("start here")

	fileContents := readMyFile("dracula.txt")
	fmt.Printf("%.50s...", fileContents)
	// split on word boundary
	tokens := splitString(fileContents)
	createDictionary(tokens)
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

	return dict

}

func sortDictionary(dict map[string]int) map[string]int {
	return dict
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
