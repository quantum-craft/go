package crackingcodinginterview

import (
	"github.com/quantum-craft/go/sorting"
)

func checkPermutation2(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Statistics := make([]int, 256)

	for i := 0; i < len(str1); i++ {
		str1Statistics[str1[i]]++
	}

	for i := 0; i < len(str2); i++ {
		str1Statistics[str2[i]]--
		if str1Statistics[str2[i]] < 0 {
			return false
		}
	}

	return true
}

func checkPermutation(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Slice := make([]byte, len(str1))
	str2Slice := make([]byte, len(str2))

	for i := 0; i < len(str1); i++ {
		str1Slice[i] = str1[i]
		str2Slice[i] = str2[i]
	}

	sorting.QuickSortByte(str1Slice)
	sorting.QuickSortByte(str2Slice)

	for i := 0; i < len(str1Slice); i++ {
		if str1Slice[i] != str2Slice[i] {
			return false
		}
	}

	return true
}

func isUniqueChars(str string) bool {
	if len(str) > 256 {
		return false
	}

	var charStatistics [4]int64

	for i := 0; i < len(str); i++ {

		if (1<<(str[i]%64))&charStatistics[getSector(str[i])] > 0 {
			return false
		}

		charStatistics[getSector(str[i])] |= 1 << (str[i] % 64)
	}

	return true
}

func getSector(r byte) int {
	if r < 64 {
		return 0
	} else if r >= 64 && r < 128 {
		return 1
	} else if r >= 128 && r < 192 {
		return 2
	} else {
		return 3
	}
}
