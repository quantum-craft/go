package huffman

import "testing"

func TestHuffmanEncoding(t *testing.T) {
	file := "../data/huffman.txt"

	n := Encoding(file)

	min, max := Iterate(&n)

	if min != 9 {
		t.Error("HuffmanEncoding error !")
	}

	if max != 19 {
		t.Error("HuffmanEncoding error !")
	}
}
