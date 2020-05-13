package crackingcodinginterview

import (
	"fmt"
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	fmt.Println(checkPermutation("google", "ggoole"))   // true
	fmt.Println(checkPermutation("google", "gggole"))   // false
	fmt.Println(checkPermutation("google", "facebook")) // false
	fmt.Println(checkPermutation("elgoog", "ooggel"))   // true

	fmt.Println(checkPermutation2("google", "ggoole"))   // true
	fmt.Println(checkPermutation2("google", "gggole"))   // false
	fmt.Println(checkPermutation2("google", "facebook")) // false
	fmt.Println(checkPermutation2("elgoog", "ooggel"))   // true
	fmt.Println(checkPermutation2("elgoog", "oogGel"))   // false
}

func TestIsUniqueChars(t *testing.T) {
	fmt.Println(isUniqueChars("apple"))                                                // false
	fmt.Println(isUniqueChars("google"))                                               // false
	fmt.Println(isUniqueChars("facebok"))                                              // true
	fmt.Println(isUniqueChars("abcdefghijklmnopqrstuvwxyz"))                           // true
	fmt.Println(isUniqueChars("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")) // true
	fmt.Println(isUniqueChars("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYz")) // false
}
