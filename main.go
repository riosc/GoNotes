package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*


//dup V2
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, error := os.Open(arg)
			if error != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", error)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
*/

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, filename)
			}
		}

	}
}
