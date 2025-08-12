package main

import (
	"fmt"
	"unicode/utf8"
)


func analyzeString(s string) {
	fmt.Printf("\nString: %q:\n", s)
	fmt.Printf("len() (bytes): %d\n", len(s))
	fmt.Printf("utf8.RuneCountInString(): %d", utf8.RuneCountInString(s))
	fmt.Printf("Bytes %v\n", []byte(s))
	fmt.Printf("Runes: %v\n", []rune(s))
}

func demonstrateIterations(s string) {
	fmt.Printf("\nString: %q\n", s)

	// wrong way - iteration by bytes
	fmt.Println("Byte iteration (WRONG for multi-byte chars)")
	for i := 0; i < len(s); i++ {
		fmt.Printf(" [%d]: %c (byte: %d\n)", i, s[i], s[i])
	}

	// Right way - iterating by Runes
	fmt.Println("Rune iteration (CORRECT):")
	for i, r := range []rune(s) {
		fmt.Printf("[%d]: %c (rune: %d, byte: %d)\n", i, r, r, utf8.RuneLen(r))
	}
	

	// Alternative - convert to rune first
	fmt.Println("Converting to rune first:")
	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf(" [%d]: %c (rune: %d)\n", i, r, r)
	}
}


func workWithRunes(){
	var r1 = 'A'
	var r2 rune = '中'
	var r3 rune = '🚀'
	
	fmt.Printf("Rune: 'A': value=%d, bytes needed=%d\n", r1, utf8.RuneLen(r1))
	fmt.Printf("Rune: '中': value=%d, bytes needed=%d\n", r2, utf8.RuneLen(r2))
	fmt.Printf("Rune: '🚀': value=%d, bytes needed=%d\n", r3, utf8.RuneLen(r3))


	//building strings from runes

	runes := []rune{'H', 'E', 'L', 'L', 'O'}
	str := string(runes)
	fmt.Printf("string from runes %q\n", str)


	// converting between string and []rune
	original := "Go语言"
	runeSlice := []rune(original)
	backToString := string(runeSlice)

	fmt.Printf("Original: %q, Back: %q, Equal: %t\n", original, backToString, original == backToString )
}



func PracticalExamples() {
	text := "Programming in Go is fun"
	
	// safe string truncation
	fmt.Printf("original: %q (len=%d, runes=%d)\n", text, len(text), utf8.RuneCountInString(text))


	// WRONG WAY to truncate -can break utf-8
	if len(text) > 20 {
		truncateWrong := text[:20] // might cut the middle of UTF-8 sequence
		fmt.Printf("Wrong Truncation: %q\n", truncateWrong)
	}


	// Right way to truncate
	truncateRight := safeSubstring(text, 0, 20)
	fmt.Printf("Safe truncation: %q\n", truncateRight)
}


func safeSubstring(s string, start, maxRunes int) string {
	runes := []rune(s)

	if start >= len(runes) {
		return ""
	}
	
	end := start + maxRunes
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}



