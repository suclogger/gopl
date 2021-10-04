// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
// only deal with files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	reverseIndex := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		return
	}
	for _, filename := range files {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return
		}
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			counts[line]++
			child, ok := reverseIndex[line]
			if !ok {
				child := make(map[string]int)
				child[filename]++
				reverseIndex[line] = child
			} else {
				child[filename]++
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for k := range reverseIndex[line] {
				fmt.Printf("%s\t", k)
			}
			fmt.Printf("\n")
		}
	}
}
