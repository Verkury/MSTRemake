package main

import (
	"fmt"
	"github.com/Verkury/MSTRemake/cryp"
)

func main() {
	mp := cryp.MakeMapS(234152)
	mp1 := cryp.MakeMap()
	message := "Hello, world!"
	message1 := "Привет, мир!"
	message2 := "Пример достаточно длинного текста для проверки большего количества символов на корректность 1234567890 !\"№;:??*(()_)"

	NewMessage := cryp.Encrypt(mp, message)
	NewMessage1 := cryp.Encrypt(mp, message1)
	NewMessage2 := cryp.Encrypt(mp1, message)
	NewMessage3 := cryp.Encrypt(mp1, message1)
	NewMessage4 := cryp.Encrypt(mp, message2)
	NewMessage5 := cryp.Encrypt(mp1, message2)

	fmt.Println(message, "->", NewMessage, "->", cryp.Decrypt(mp, NewMessage))
	fmt.Println(message1, "->", NewMessage1, "->", cryp.Decrypt(mp, NewMessage1))
	fmt.Println(message, "->", NewMessage2, "->", cryp.Decrypt(mp1, NewMessage2))
	fmt.Println(message1, "->", NewMessage3, "->", cryp.Decrypt(mp1, NewMessage3))
	fmt.Println(message2, "->", NewMessage4, "->", cryp.Decrypt(mp, NewMessage4))
	fmt.Println(message2, "->", NewMessage5, "->", cryp.Decrypt(mp1, NewMessage5))
}