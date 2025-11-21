package cryp

import "fmt"

var symbolsList string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя!% &*()+-="


func MakeMap() map[string][]rune {
	configDefault := make(map[string][]rune)
	runes := []rune(symbolsList)
	for n, v := range runes {
		configDefault[decimalToBinary(n)] = []rune{v}
	}
	return configDefault
}

func MakeMapS(seed int) map[string][]rune {
	configDefault := make(map[string][]rune)
	shift := seed % 128
	runes := []rune(symbolsList)

	for n, v := range runes {
		newPos := (n+ shift) % 128
		configDefault[decimalToBinary(newPos)] = []rune{v}
	}
	return configDefault
}

func decimalToBinary(n int) string {
	if n == 0 {
		return "0000000"
	}

	var binary string
	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		n = n / 2
	}
	for{
		if (len(binary) < 7) {
			binary = fmt.Sprintf("%s%s", "0", binary)
		}
		if (len(binary) >= 7) {
			break
		}
	}
	return binary
}