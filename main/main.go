package main

import (
	"fmt"
)

func minTransfers(transactions [][]int) int {
	debt := make(map[int]int)

	for i := range transactions {
		debt[transactions[i][0]] -= transactions[i][2]
		debt[transactions[i][1]] += transactions[i][2]
	}

	balance := make([]int, 0)

	for _, v := range debt {
		if v != 0 {
			balance = append(balance, v)
		}
	}

	return dfs(0, balance)
}

func dfs(i int, balance []int) int {
	if i == len(balance) {
		return 0
	}

	if balance[i] == 0 {
		return dfs(i+1, balance)
	}

	minTrans := maxInt

	for j := i + 1; j < len(balance); j++ {
		if balance[i]*balance[j] < 0 {
			temp := balance[j]

			balance[j] += balance[i]
			minTrans = min(minTrans, 1+dfs(i+1, balance))
			balance[j] = temp

			if balance[i]+balance[j] == 0 {
				break
			}
		}
	}

	return minTrans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

func main() {
	balance := []int{-8, -6, -5, -4, -2, -1, -1, 2, 3, 3, 3, 4, 6, 6}

	fmt.Println(dfs(0, balance))
}
