package main

import "fmt"

type Result struct {
	collections [][]byte
	current     []byte
}

func (r *Result) setChar(index int, c Character) {
	r.current = append(r.current, c.v)
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

type CurrentChars struct {
	CurrentIndex int
	Chars        []Character
}

func (c *CurrentChars) val() Character {
	return c.Chars[c.CurrentIndex]
}

func (c *CurrentChars) next() bool {
	if c.CurrentIndex == len(c.Chars)-1 {
		return false
	}
	for index := range c.Chars {
		if index == c.CurrentIndex {
			c.Chars[index].w--
		}
	}

	for index, v := range c.Chars {
		if v.w > 0 {
			c.CurrentIndex = index
			return true
		}
	}
	return false
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

var result = &Result{}

func main() {
	chars := getChars("aabc")
	cChars := &CurrentChars{Chars: chars}
	perm(cChars, 0)

	// fmt.Println(cChars)
	// fmt.Println(cChars.next())

	// fmt.Println(cChars)
	// fmt.Println(cChars.next())

	// fmt.Println(cChars)
	// fmt.Println(cChars.next())

	// fmt.Println(cChars)

	// fmt.Println(cChars.next())
	// fmt.Println(cChars)
	// perm(chars, 0)
	// fmt.Println(chars)
}

func perm(cChars *CurrentChars, recursionDepth int) {
	result.setChar(recursionDepth, cChars.val())
	fmt.Println(result)

	recursionDepth++
	cChars.next()
	result.setChar(recursionDepth, cChars.val())
	fmt.Println(result)

	recursionDepth++
	cChars.next()
	result.setChar(recursionDepth, cChars.val())
	fmt.Println(result)

	recursionDepth++
	cChars.next()
	result.setChar(recursionDepth, cChars.val())
	fmt.Println(result)
	
}
