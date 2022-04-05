package main

import "fmt"

// dinh nghia 1 interface
type nguyenam interface {
	tim() []rune
}

type mystring string

//kieu du lieu mystring trien khai interface nguyenam
func (ms mystring) tim() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := mystring("ngoxuan")
	var n nguyenam
	n = name
	fmt.Printf("vowels are %c", n.tim())
}
