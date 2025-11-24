package main

import (
	"fmt"
	"github.com/Verkury/MSTRemake/cryp"
)

func main() {

	cryp.Check()

	cryp.AdvancedCheck()

	mp := cryp.MakeMapS(234152)
	mp1 := cryp.MakeMap()
	message := "Hello, world!"
	message1 := "Привет, мир!"
	NewMessage := cryp.Encrypt(mp, message)
	NewMessage1 := cryp.Encrypt(mp, message1)
	NewMessage2 := cryp.Encrypt(mp1, message)
	NewMessage3 := cryp.Encrypt(mp1, message1)

	fmt.Println(message, "->", NewMessage, "->", cryp.Decrypt(mp, NewMessage))
	fmt.Println(message1, "->", NewMessage1, "->", cryp.Decrypt(mp, NewMessage1))
	fmt.Println(message, "->", NewMessage2, "->", cryp.Decrypt(mp1, NewMessage2))
	fmt.Println(message1, "->", NewMessage3, "->", cryp.Decrypt(mp1, NewMessage3))

	amp := cryp.MakeAdvancedMap(4444, "") // We nead 2 args, but 2-nd arg can be clear

	AMPMessage := cryp.AdvancedEncrypt(amp, message)
	decMessage := cryp.AdvancedDecrypt(amp, AMPMessage)

	fmt.Print("\n\nAdvanced level\n\n")

	fmt.Println(message, "->", AMPMessage, "->", decMessage)
}