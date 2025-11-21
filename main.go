package main

import (
	"fmt"

	"github.com/Verkury/MSTRemake/cryp"
)

func main() {
	cfg := cryp.MakeMap()
	i := 0
	num := 0
	num1 := 0
	for n, v := range cfg {
		fmt.Println(n, "-", string(v))
		i++
		if (len(n) == 8) {
			num++
		} else {
			num1++
		}
	}
	fmt.Println(i, num, num1, num1+num)
}