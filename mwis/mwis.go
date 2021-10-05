package mwis

// Mwis calculates the maximum-weight independent set
func Mwis(nums []int) []int {
	dp := make([]int, len(nums))
	dp[0] = 0
	dp[1] = nums[1]

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp
}

// Reconstruct recovers the maximum-weight independent set from cache
func Reconstruct(cache []int, weights []int) (included []bool) {
	included = make([]bool, len(cache))

	i := len(included) - 1
	for i >= 2 {
		if cache[i-1] > cache[i-2]+weights[i] {
			i = i - 1
		} else {
			included[i] = true
			i = i - 2
		}
	}

	if !included[2] {
		included[1] = true
	}

	return included
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
