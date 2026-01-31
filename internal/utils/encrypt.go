package utils

import (
	"log"
	"maps"
	"strings"
)

var chars = map[string]struct{}{
	"a": {},
	"b": {},
	"c": {},
	"d": {},
	"e": {},
	"f": {},
	"g": {},
	"h": {},
	"i": {},
	"j": {},
	"k": {},
	"l": {},
	"m": {},
	"n": {},
	"o": {},
	"p": {},
	"q": {},
	"r": {},
	"s": {},
	"t": {},
	"u": {},
	"v": {},
	"w": {},
	"x": {},
	"y": {},
	"z": {},
}

func getAlphabets() []rune {
	alphabets := make([]rune, 0)
	for i := 97; i <= 122; i++ {
		char := rune(i)
		alphabets = append(alphabets, char)
	}

	return alphabets
}

func GenerateKey() string {
	keys := make([]string, 26)
	copyChars := make(map[string]struct{}, len(chars))

	maps.Copy(copyChars, chars)
	for k := range copyChars {
		keys = append(keys, k)
		delete(copyChars, k)
	}

	key := strings.Join(keys, "")
	return key
}

func Encrypt(key, text string) string {
	alphabets := getAlphabets()
	splitKey := strings.Split(key, "")
	splitText := strings.Split(text, "")

	result := []rune{}
	mapping := make(map[rune]rune)
	for i := 0; i < len(alphabets); i++ {
		mapping[alphabets[i]] = []rune(splitKey[i])[0]
	}

	for _, val := range splitText {
		char := []rune(val)[0]
		var txt rune
		e, ok := mapping[char]
		txt = e
		if !ok {
			code := int(char)
			if code >= 65 && code <= 90 {
				code = code + 32            // to lowercase
				temp := mapping[rune(code)] // find matched key
				code = int(temp) - 32       // to uppercase
				txt = rune(code)
			} else {
				txt = char
			}
		}

		result = append(result, txt)
	}
	return string(result)
}

func Decrypt(key, encrypted string) string {
	alphabets := getAlphabets()

	splitKey := strings.Split(key, "")
	splitEncrypted := strings.Split(encrypted, "")

	dictionary := make(map[rune]rune)
	for i := 0; i < len(alphabets); i++ {
		char := []rune(splitKey[i])[0]
		dictionary[char] = alphabets[i]
	}

	var result []rune
	var txt rune
	for _, val := range splitEncrypted {
		char := []rune(val)[0]
		ch, ok := dictionary[char]
		txt = ch
		if !ok {
			code := int(char)
			log.Printf("%v, %d", string(char), code)
			if code >= 65 && ch <= 90 {
				code = code + 32
				temp := dictionary[rune(code)]
				code = int(temp) - 32
				txt = rune(code)
			} else {
				txt = char
			}
		}
		result = append(result, txt)
	}
	return string(result)
}
