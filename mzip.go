// This is just a short program for (my) educational purposes.
// It doesn't do anything useful, I'm just learning Go.
// This should be a Godoc comment.
package main

import (
	"fmt"
	"os"
)

/*	is this a multiline comment?,
	good
*/

func main() {
	fmt.Println("start here")

	fileContents := readMyFile("dracula.txt")

	fmt.Printf("%.25s", fileContents)

}

// Reads the complete contents of the file, and returns it as a string
func readMyFile(filePath string) string {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
