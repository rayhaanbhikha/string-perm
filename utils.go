package main

func contains(arr []Character, v byte) (int, bool) {
	for index, c := range arr {
		if c.v == v {
			return index, true
		}
	}
	return 0, false
}
