package main

import (
	"fmt"

	"github.com/Verkury/MSTRemake/cryp"
)

func main() {
	fmt.Println("Default")
	cfg := cryp.MakeMap()
	for n, v := range cfg {
		fmt.Println(n, "-", string(v))
	}

	fmt.Println("On seed 234234")
	cfg1 := cryp.MakeMapS(234234)
	for n, v := range cfg1 {
		fmt.Println(n, "-", string(v))
	}

}