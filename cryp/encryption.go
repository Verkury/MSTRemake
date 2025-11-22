package cryp

import (
	"fmt"
)

var symbolsList string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя!% &*()+-="


func MakeMap() (map[string] []rune, map[string]string) {
	configDefaultIn := make(map[string][]rune)
	configDefaultFrom := make(map[string]string)
	runes := []rune(symbolsList)

	for n, v := range runes {
		binaryKey := decimalToBinary(n)
		configDefaultIn[binaryKey] = []rune{v}
		configDefaultFrom[string(v)] = binaryKey
	}
	return configDefaultIn, configDefaultFrom
}

func MakeMapS(seed int) (map[string][]rune, map[string]string) {
	configDefaultIn := make(map[string][]rune)
	configDefaultFrom := make(map[string]string)
	shift := seed % 128
	runes := []rune(symbolsList)

	for n, v := range runes {
		newPos := (n+ shift) % 128
		binaryKey := decimalToBinary(newPos)
		configDefaultIn[binaryKey] = []rune{v}
		configDefaultFrom[string(v)] = binaryKey
	}
	return configDefaultIn, configDefaultFrom
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