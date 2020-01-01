package main

import "fmt"

type Result struct {
	collections [][]byte
	current     []byte
}

type Character struct {
	value  byte
	weight int
}

func contains(arr []*Character, v byte) (*Character, bool) {
	for _, c := range arr {
		if c.value == v {
			return c, true
		}
	}
	return nil, false
}

func getChars(str string) []*Character {
	chars := make([]*Character, 0)
	for _, s := range str {
		if c, ok := contains(chars, byte(s)); ok {
			c.weight++
		} else {
			chars = append(chars, &Character{byte(s), 1})
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
