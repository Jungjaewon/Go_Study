package main

import (
	"sort"
)

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {

	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	// sort libaray is diverse and int and int32 are different.
	var idx int32 = 0
	var cnt int32 = 0
	var total_price int32 = 0
	var pcnt int32
	var ptotal_price int32
	_ = ptotal_price // unsed variable

	for {
		if total_price < k {
			pcnt = cnt
			ptotal_price = total_price

			total_price += prices[idx]
			idx += 1
			cnt += 1

			if total_price > k {
				return pcnt
			}
		}
	}

}

func main() {

}
