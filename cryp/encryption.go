package cryp

import (
	"fmt"
)

type MP struct{
	mapIn map[string][]rune
	mapFrom map[string]string
}

var symbolsList string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя!,. *()+-="

func MakeMap() MP {
	configDefaultIn := make(map[string][]rune)
	configDefaultFrom := make(map[string]string)
	runes := []rune(symbolsList)

	for n, v := range runes {
		binaryKey := decimalToBinary(n)
		configDefaultIn[binaryKey] = []rune{v}
		configDefaultFrom[string(v)] = binaryKey
	}
	return MP{configDefaultIn, configDefaultFrom}
}

func MakeMapS(seed int) MP {
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
	return MP{configDefaultIn, configDefaultFrom}
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

func Encrypt(mp MP, message string) string{
	var result string
	
	// Convert message to binary
	var binaryString string
	for _, v := range message {
		if binary, exists := mp.mapFrom[string(v)]; exists {
			binaryString += binary
		} else {
			binaryString += decimalToBinary(int(v))
		}
	}
	
	// Reverse binary string
	var reversedBinary string
	for _, r := range binaryString {
		reversedBinary = string(r) + reversedBinary
	}
	
	// Convert reversed binary back to characters (group by 8 bits)
	for i := 0; i < len(reversedBinary); i += 7 {
		end := i + 7
		if end > len(reversedBinary) {
			end = len(reversedBinary)
		}
		binaryKey := reversedBinary[i:end]
		
		// Pad with zeros if needed
		for len(binaryKey) < 7 {
			binaryKey += "0"
		}
		
		if charSlice, exists := mp.mapIn[binaryKey]; exists {
			result += string(charSlice)
		} else {
			// If no mapping found, use the original character code
			// Convert binary to decimal and then to rune
			val := 0
			for _, bit := range binaryKey {
				val = val*2 + int(bit-'0')
			}
			result += string(rune(val))
		}
	}
	return result
}

