package cryp

import (
	"fmt"
	"math/rand"
)

type MP struct{
	mapIn map[string]rune
	mapFrom map[string]string
}

type AdvancedMP struct {
	MP
	salt string 
}

var symbolsList string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
    "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя" +
    "0123456789" + 
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

func generateRandomSalt(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func MakeAdvancedMap(seed int, customSalt string) AdvancedMP {
	if customSalt == "" {
		customSalt = generateRandomSalt(16)
	}

	baseMP := MakeMapS(seed)

	return AdvancedMP{
		MP: baseMP,
		salt: customSalt,
	}
}

func convertStringToBits(amp AdvancedMP, message string) string {
	// Convert message to binary
	var binaryString string
    for _, v := range message {
        if binary, exists := amp.mapFrom[string(v)]; exists {
            binaryString += binary
        } else {
            binaryString += decimalToBinary(int(v))
        }
    }
    
    return binaryString
}

func reverseBits(binaryString string) string {
    var reversedBinary string
    for i := len(binaryString) - 1; i >= 0; i-- {
        reversedBinary += string(binaryString[i])
    }
    return reversedBinary
}

func convertBitsToString(amp AdvancedMP, message string) string {
	var result string
	
	for i := 0; i < len(message); i += 8 {
		end := i + 8
		if end > len(message) {
			end = len(message)
		}
		binaryKey := message[i:end]
		
		// Pad with zeros if needed
		for len(binaryKey) < 8 {
			binaryKey += "0"
		}
		
		if charSlice, exists := amp.mapIn[binaryKey]; exists {
			result += string(charSlice)
		} else {
			val := 0
			for _, bit := range binaryKey {
				val = val*2 + int(bit-'0')
			}
			result += string(rune(val))
		}
	}
	return result
}

func AdvancedEncrypt(amp AdvancedMP, message string) string {
	runes := []rune(message)
    if len(runes) == 0 {
        return ""
    }
    
    last := runes[len(runes)-1]
    for i := len(runes) - 1; i > 0; i-- {
        runes[i] = runes[i-1]
    }
    runes[0] = last

	Binary := convertStringToBits(amp, string(runes) + amp.salt)
	reversedBinary := reverseBits(Binary)

	// New approach: interleave first and last 4-bit chunks
	var processedBinary string
	bits := reversedBinary
	length := len(bits)
	
	for i := 0; i < length/2; i += 4 {
		// Take 4 bits from the beginning
		if i+4 <= length/2 {
			processedBinary += bits[i:i+4]
		} else {
			processedBinary += bits[i:length/2]
		}
		
		// Take 4 bits from the end
		endStart := length - i - 4
		if endStart >= length/2 {
			processedBinary += bits[endStart:length-i]
		} else if endStart < length/2 && length/2 < length-i {
			processedBinary += bits[length/2:length-i]
		}
	}
    return convertBitsToString(amp, processedBinary)
}

func AdvancedDecrypt(amp AdvancedMP, encrypted string) string {
	Binary := convertStringToBits(amp, encrypted)

	length := len(Binary)
	var firstHalf, secondHalf , three string

	for i := 0; i < length; i += 8 {
        if i+4 <= length {
            firstHalf += Binary[i:i+4]
        } else {
            firstHalf += Binary[i:]
        }
        
        if i+8 <= length {
            secondHalf += Binary[i+4:i+8]
        } else if i+4 < length {
            secondHalf += Binary[i+4:]
        }
    }
	for i:=0; i < length/2; i += 4 {
		if i+4 <= length/2 {
			three += reverseBits(secondHalf[i:i+4])
		} else {
			three += reverseBits(secondHalf[i:])
		}
	}

	reversedBinary := reverseBits(firstHalf + reverseBits(three))

	decryptedWithSalt := convertBitsToString(amp, reversedBinary)

	runesWithSalt := []rune(decryptedWithSalt)
    if len(runesWithSalt) <= len(amp.salt) {
        return ""
    }
    
    decrypted := string(runesWithSalt[:len(runesWithSalt)-len(amp.salt)])

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

func AdvancedCheck() {
    fmt.Println("Testing Advanced Encryption...")
    
    seed := 42
    salt := "test_salt_123"
    amp := MakeAdvancedMap(seed, salt)
    
    testMessage := "Hello, World! 123"
    encrypted := AdvancedEncrypt(amp, testMessage)
    decrypted := AdvancedDecrypt(amp, encrypted)
    
    fmt.Println("Original:", testMessage)
    fmt.Println("Encrypted:", encrypted)
    fmt.Println("Decrypted:", decrypted)
    fmt.Println("Test passed:", testMessage == decrypted)
    
    if testMessage == decrypted {
        fmt.Println("Advanced encryption test passed!")
    } else {
        fmt.Println("Advanced encryption test failed!")
    }
}