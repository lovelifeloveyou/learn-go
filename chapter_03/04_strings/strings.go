package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "You弄啥嘞！"
	fmt.Println(s, len(s))                                   // You弄啥嘞！ 15
	fmt.Println(s, "Rune count=", utf8.RuneCountInString(s)) // You弄啥嘞！ Rune count= 7
	fmt.Println("--------------分割线----------------")
	fmt.Println("[]byte(s) = ", []byte(s))
	fmt.Println("[]rune(s) = ", []rune(s))
	fmt.Println()
	fmt.Println("--------------分割线----------------")
	for i, ch := range s {
		// fmt.Printf("(%d, %d, %c) ", i, ch, ch)
		fmt.Printf("(%d, %x) ", i, ch) // (0, 59) (1, 6f) (2, 75) (3, 5f04) (6, 5565) (9, 561e) (12, ff01)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		// fmt.Printf("(%d, %x, %c) ", i, ch, ch)
		fmt.Printf("(%d, %x) ", i, ch) // (0, 59) (1, 6f) (2, 75) (3, 5f04) (4, 5565) (5, 561e) (6, ff01)
	}
	fmt.Println()
	fmt.Println("--------------分割线----------------")
	myBytes := []byte(s)
	for len(myBytes) > 0 {
		ch, size := utf8.DecodeRune(myBytes)
		myBytes = myBytes[size:]
		fmt.Printf("(%c, %d) ", ch, size)
	}
}
