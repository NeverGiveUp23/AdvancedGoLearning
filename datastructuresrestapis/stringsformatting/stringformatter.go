package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)


func main() {
	banner("Go", 6)
	banner("G💀", 6)

	hello := "Hello" // len(hello) -> 5 (5 ASCII chars = 5 bytes)
	chineseWord := "世界"  // len(chineseWord) -> 6 (2 Chinese chars = 6 bytes, 3 bytes each)
	emoji := "🚀" // len(emoji) -> 4 (1 emoji = 4 bytes)


	fmt.Println(len(hello))
	fmt.Println(len(chineseWord))
	fmt.Println(len(emoji))
	
	s := "G💀"

	fmt.Println("len:", len(s))
	fmt.Println("s[1]:", s[1])


	// strings are -> UTF-8 encoded
	// len, s[] -> are byte (uint8)
	// for loops are ->  rune (int32)

	// Range auto handles runes

	for i, c := range s {
		fmt.Printf("%v at %d\n", i, c)
	}

	for i, v := range emoji {
		fmt.Printf("%v is %d", i, v)
	}

	fmt.Println("")
	banner(chineseWord, 6)

	// you can convert it to a rune first
	runes := []rune(chineseWord)
	fmt.Printf("%c\n", runes[0]) // prints the rune at index 0 -> 世
   
	countChars(emoji)
fmt.Println(safeCharAt(chineseWord, 1))
}

// always use utf8.RuneCountInString for char count 
func countChars(char string) int {
	return utf8.RuneCountInString(char)
}

// safe string operations are important
func safeCharAt(s string, pos int) (rune, bool) {
	runes := []rune(s)
	if pos < 0 || pos >= len(runes) {
		return 0, false
	}

	return runes[pos], true
}



func banner(test string, width int) {
	// BUG: len() is in bytes
	// padding := (width - len(test) / 2)

	// utf8.RuneCountInString returns the amount of runes in a string
	padding := (width - utf8.RuneCountInString(test)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(test)
}
