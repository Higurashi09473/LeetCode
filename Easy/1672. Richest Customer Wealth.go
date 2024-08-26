package main

func maximumWealth(accounts [][]int) int {
	var bank, maxbank int
	for i := range accounts {
		for j := range accounts[i] {
			bank += accounts[i][j]
		}
		if bank > maxbank {
			maxbank = bank
		}
		bank = 0
	}
	return maxbank
}
