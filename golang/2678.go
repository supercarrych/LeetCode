package main

import "fmt"

func countSeniors(details []string) int {
	res := 0
	for _, s := range details {
		fmt.Printf("v%\n", s)
		fmt.Printf("v%\n", s[11])
		fmt.Printf("v%\n", s[11]&15)
		fmt.Printf("v%\n", s[11]&10)
		if s[11]&15*10+s[12]&15 > 60 {
			res++
		}

	}

	return res
}

func main() {
	countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"})
}
