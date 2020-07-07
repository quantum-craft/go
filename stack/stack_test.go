package stack

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {
	f, _ := os.Open("../data/QuickSortNumbers.txt")
	defer f.Close()

	data := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i, ok := strconv.Atoi(line)

		if ok == nil {
			data = append(data, i)
		}

	}

	stack := MakeStack()

	for i := 0; i < len(data); i++ {
		stack.Push(data[i])
	}

	for i := len(data) - 1; i >= 0; i-- {
		d := stack.Pop()

		if d != data[i] {
			t.Error("Stack error !")
		}
	}
}
