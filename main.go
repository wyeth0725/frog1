package main

import "fmt"

func abs(a int) int{
	if a < 0{
		return -a
	}
	return a
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	h := make([]int, n)
	for i := 0; i < n; i++{
		fmt.Scan(&h[i])
	}
	
	costs := make([]int, n)
	costs[0] = 0
	costs[1] = abs(h[1] - h[0])

	for i := 2; i < n; i++{
		c1 := costs[i - 1] + abs(h[i] - h[i - 1])
		c2 := costs[i - 2] + abs(h[i] - h[i - 2])
		costs[i] = min(c1, c2)
	}

	fmt.Println(costs[n - 1])
}
