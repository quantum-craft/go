package huffman

import (
	"testing"
)

func TestHuffmanEncodingWithQueue2(t *testing.T) {
	file := "../data/huffman.txt"

	root := EncodingWithQueue2(file)

	min, max := Iterate2(root)

	if min != 9 {
		t.Error("HuffmanEncoding error !")
	}

	if max != 19 {
		t.Error("HuffmanEncoding error !")
	}
}

func TestHuffmanEncodingWithQueue(t *testing.T) {
	file := "../data/huffman.txt"

	n := EncodingWithQueue(file)

	min, max := Iterate(&n)

	if min != 9 {
		t.Error("HuffmanEncoding error !")
	}

	if max != 19 {
		t.Error("HuffmanEncoding error !")
	}
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
