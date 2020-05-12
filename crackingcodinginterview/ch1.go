package crackingcodinginterview

func isUniqueChars(str string) bool {
	if len(str) > 256 {
		return false
	}

	var charDisc [4]int64

	for i := 0; i < len(str); i++ {

		if (1<<(str[i]%64))&charDisc[getSector(str[i])] > 0 {
			return false
		}

		charDisc[getSector(str[i])] |= 1 << (str[i] % 64)
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
