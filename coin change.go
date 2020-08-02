func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)

	for i, v := range dp {
		if i != 0 && v == 0 {
			continue
		}
		// i denotes current amount of money morphing from 0 to amount step by 1
		// v denotes minimum number of coins to reach the current i
		for _, coin := range coins {
			temp := i + coin
			if temp > amount {
				continue
			}
			if dp[temp] > v+1 || dp[temp] == 0 {
				dp[temp] = v+1
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}