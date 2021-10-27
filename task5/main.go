package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i, spell(i))
	}
	 i := 88
	 fmt.Println(i, spell(i))
}

func pow(i int, p int) int {
	return int(math.Pow(1000, float64(p)))
}

func spell(n int) string {
	to19 := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
		",Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	tens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	if n == 0 {
		return ""
	}
	if n < 20 {
		return to19[n-1]
	}
	if n < 100 {
		return tens[n/10-2] + " " + spell(n%10)
	}
	if n < 1000 {
		return to19[n/100-1] + " Hundred " + spell(n%100)
	}

	for idx, w := range []string{"Thousand", "Million", "Billion"} {
		p := idx + 1
		if n < pow(1000, (p+1)) {
			return spell(n/pow(1000, p)) + " " + w + " " + spell(n%pow(1000, p))
		}
	}

	return "error"
}