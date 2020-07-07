package huffman

import (
	"fmt"
	"testing"
)

func TestHuffmanEncoding2Stacks(t *testing.T) {
	file := "../data/huffmanSmall.txt"

	n := Encoding2Stacks(file)

	min, max := Iterate(&n)

	fmt.Println(min)
	fmt.Println(max)

	// if min != 9 {
	// 	t.Error("HuffmanEncoding error !")
	// }

	// if max != 19 {
	// 	t.Error("HuffmanEncoding error !")
	// }
}

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
