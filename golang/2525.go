package main

import (
	"fmt"
	"math"
)

func categorizeBox(length int, width int, height int, mass int) string {
	v := length * width * height
	ten4 := int(math.Pow10(4))
	if v >= int(math.Pow10(9)) || length >= ten4 || width >= ten4 || height >= ten4 {
		if mass >= 100 {
			return "Both"
		} else {
			return "Bulky "
		}
	} else if mass >= 100 {
		return "Heavy"
	} else {
		return "Neither"
	}

}
func main() {
	fmt.Printf(categorizeBox(1000, 35, 700, 300))
}
