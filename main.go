package main

import "fmt"

type Result struct {
	collections [][]byte
	current     []Character
}

type Character struct {
	value  byte
	weight int
}

func contains(arr []Character, v byte) (int, bool) {
	for index, c := range arr {
		if c.value == v {
			return index, true
		}
	}
	return 0, false
}

func getChars(str string) []Character {
	chars := make([]Character, 0)
	for _, s := range str {
		if charIndex, ok := contains(chars, byte(s)); ok {
			chars[charIndex].weight++
		} else {
			chars = append(chars, Character{byte(s), 1})
		}
	}
	return chars
}

func main() {
	chars := getChars("aabc")
	for _, c := range chars {
		fmt.Println(c)
	}
}
