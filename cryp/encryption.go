package cryp

import (
	"fmt"
)

type MP struct{
	mapIn map[string]rune
	mapFrom map[string]string
}

var symbolsList string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
    "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя" +
    "0123456789" + // Цифры
    "!@#$%^&*()_+-=[]{}|;:,.<>?/~`\"'\\ " +
	"¡¢£¤¥¦§¨©ª«¬®¯°±²³´µ¶·¸¹º»¼½¾¿" +
	"ҐЂЃЄЅІЇЈЉЊЋЌЎЏґђѓєѕіїјљњћќўџ" +
    "─│┌┐└┘├┤┬┴┼▀▄█▌▐░▒▓⌠■∙√≈≤≥" + 
	"ñòóôõö÷øùúû"

func MakeMap() MP {
	configDefaultIn := make(map[string]rune)
	configDefaultFrom := make(map[string]string)
	runes := []rune(symbolsList)

	for n, v := range runes {
		binaryKey := decimalToBinary(n)
		configDefaultIn[binaryKey] = v
		configDefaultFrom[string(v)] = binaryKey
	}
	return MP{configDefaultIn, configDefaultFrom}
}

func MakeMapS(seed int) MP {
	configDefaultIn := make(map[string]rune)
	configDefaultFrom := make(map[string]string)
	shift := seed % 256
	runes := []rune(symbolsList)

	for n, v := range runes {
		newPos := (n+ shift) % 256
		binaryKey := decimalToBinary(newPos)
		configDefaultIn[binaryKey] = v
		configDefaultFrom[string(v)] = binaryKey
	}
	return MP{configDefaultIn, configDefaultFrom}
}

func decimalToBinary(n int) string {
	binary := ""
	for i := 7; i >= 0; i-- {
		if n&(1<<i) != 0 {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	return binary
}

func crypt(mp MP, message string) string{
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
	for i := len(binaryString) - 1; i >= 0; i-- {
		reversedBinary += string(binaryString[i])
	}
	
	// Convert reversed binary back to characters (group by 8 bits)
	for i := 0; i < len(reversedBinary); i += 8 {
		end := i + 8
		if end > len(reversedBinary) {
			end = len(reversedBinary)
		}
		binaryKey := reversedBinary[i:end]
		
		// Pad with zeros if needed
		for len(binaryKey) < 8 {
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

func Encrypt(mp MP, message string) string {
    runes := []rune(message)
    if len(runes) == 0 {
        return crypt(mp, message)
    }
    
    last := runes[len(runes)-1]
    for i := len(runes) - 1; i > 0; i-- {
        runes[i] = runes[i-1]
    }
    runes[0] = last
    
    return crypt(mp, string(runes))
}

func Decrypt(mp MP, encrypted string) string {
    decrypted := crypt(mp, encrypted)
    
    runes := []rune(decrypted)
    if len(runes) == 0 {
        return decrypted
    }
    
    first := runes[0]
    for i := 0; i < len(runes)-1; i++ {
        runes[i] = runes[i+1]
    }
    runes[len(runes)-1] = first
    
    return string(runes)
}

func Check() {
	localmp := MakeMap()
	fmt.Println("Testing...")
	
	runes := []rune(symbolsList)
	test1 := len(runes) == 256
	
	testMessage := "Hello, World!"
	encrypted := Encrypt(localmp, testMessage)
	decrypted := Decrypt(localmp, encrypted)
	test2 := testMessage == decrypted
	
	fmt.Println("Length test:", test1, "(", len(runes), "symbols )")
	fmt.Println("Crypto test:", test2)
	fmt.Println("Original:", testMessage)
	fmt.Println("Encrypted:", encrypted)
	fmt.Println("Decrypted:", decrypted)

	if test1 && test2 {
		fmt.Println("All tests passed!")
	} else {
		fmt.Println("Some tests failed!")
	}
}