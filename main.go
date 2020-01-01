package main

import "fmt"

type Result struct {
	collections [][]byte
	current     []byte
}

func (r *Result) setChar(index int, c Character) {
	if len(r.current)-1 < index {
		r.current = append(r.current, c.v)
	} else {
		r.current[index] = c.v
	}
}

func (r *Result) storeCurrent() {
	r.collections = append(r.collections, r.current)
}

func (r *Result) String() string {
	return string(r.current)
}

type Character struct {
	v byte
	w int
}

func getChars(str string) []Character {
	chars := make([]Character, 0)
	for _, s := range str {
		if charIndex, ok := contains(chars, byte(s)); ok {
			chars[charIndex].w++
		} else {
			chars = append(chars, Character{byte(s), 1})
		}
	}
	return chars
}

var result = &Result{current: make([]byte, 0)}

func main() {
	chars := getChars("aabc")
	// cChars := &CurrentChars{Chars: chars}
	perm(chars, 0, 0)
	fmt.Println(result)

}

func newChars(chars []Character, currentIndex int) ([]Character, int, bool) {
	// index already on final string.
	if currentIndex == len(chars) - 1 {
		return []Character{}, 0, false
	}

	for index := range chars {
		if index == currentIndex {
			chars[index].w--
		}
	}

	for index, v := range chars {
		if v.w > 0 {
			currentIndex = index
			return chars, currentIndex, true
		}
	}

	return []Character{}, 0, false
}

func perm(chars []Character, currentIndex, recursionDepth int) {
	fmt.Println(currentIndex, chars)
	result.setChar(recursionDepth, chars[currentIndex])
	chars, currentIndex, successful := newChars(chars, currentIndex)
	if !successful {
		return
	}
	recursionDepth++
	perm(chars, currentIndex, recursionDepth)
}
