package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s = `The quick brown Fox jumps over 
	the lazy dog. Fox? Yes, the fox is quick––very quick!
	123 123 123. The dog, however, is not... not at all.`
	fmt.Println(freqCounter(s))
}

// better to use FieldsFunc but i ignored it for some reason :)
func freqCounter(str string) map[string]int {
	ret := make(map[string]int)

	var word []rune
	for _, r := range str {
		if unicode.IsLetter(r) {
			word = append(word, unicode.ToLower(r))
		} else {
			if word != nil {
				ret[string(word)]++
				word = nil
			}
		}
	}
	if word != nil {
		ret[string(word)]++
		word = nil
	}

	return ret
}
