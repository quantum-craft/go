package crackingcodinginterview

import (
	"fmt"
	"testing"
)

func TestIsUniqueChars(t *testing.T) {
	fmt.Println(isUniqueChars("apple"))                                                // false
	fmt.Println(isUniqueChars("google"))                                               // false
	fmt.Println(isUniqueChars("facebok"))                                              // true
	fmt.Println(isUniqueChars("abcdefghijklmnopqrstuvwxyz"))                           // true
	fmt.Println(isUniqueChars("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")) // true
	fmt.Println(isUniqueChars("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYz")) // false
}
