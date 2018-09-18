package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = "/"
		}

		fmt.Println(s)
	*/

	/*
		i := len(os.Args)
		for i >= 0 {
			fmt.Println(i)
			i--
		}
	*/

	/*
		s, sep := "", ""
		for index, arg := range os.Args[1:] {
			s += sep + strconv.FormatInt(int64(index), 10) + arg
			sep = "/"
		}

		fmt.Println(s)
	*/

	fmt.Println(strings.Join(os.Args[1:], "/"))
	fmt.Println(os.Args[1:])

	/// Three forms of variable declarations
	/*
		s := "" // Short declaration, only inside functions, not for package level variables **
		var s = "" // Rarely used, except when declaring multiple variables
		var s string // Relies on default initialization **
		var s string = "" // Necesary when variable type isn't the same type than the initialization value

		/// ** : most used in practice
	*/

}
