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

	fmt.Println(message, "->", cryp.Encrypt(mp, message), "->", cryp.Encrypt(mp, cryp.Encrypt(mp, message)))
	fmt.Println(message1, "->", cryp.Encrypt(mp, message1), "->", cryp.Encrypt(mp, cryp.Encrypt(mp, message1)))
	fmt.Println(message, "->", cryp.Encrypt(mp1, message), "->", cryp.Encrypt(mp1, cryp.Encrypt(mp1, message)))
	fmt.Println(message1, "->", cryp.Encrypt(mp1, message1), "->", cryp.Encrypt(mp1, cryp.Encrypt(mp1, message1)))
	fmt.Println("А теперь попытка расшифровать что-то с некоректным сидом")
	fmt.Println(message, "->", cryp.Encrypt(mp, message), "->", cryp.Encrypt(mp, cryp.Encrypt(mp1, message)))
	fmt.Println("Брух сиды можно стакать. Мощьность шифрования в абосолюте = 128^128")
}