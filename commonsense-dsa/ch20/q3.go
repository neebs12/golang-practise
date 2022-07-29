package main

import "fmt"

func GetGreatestProfit(prices []int) int {
	buyPrice := prices[0]
	sellPrice := prices[0]
	profit := sellPrice - buyPrice
	for _, price := range prices {
		if price < buyPrice {
			buyPrice = price
			sellPrice = price // reset!
			// but difference is not changed!

			// if the current price is higher than the sellPrice, reassg sellPrice
			// -- and calculate the localProfit
			// -- if the localProfit is greater than the store `profit`, reassg `profit`
		} else if sellPrice < price {
			sellPrice = price
			localProfit := sellPrice - buyPrice
			if localProfit > profit {
				profit = localProfit
			}
		}
	}
	return profit
}

func main() {
	prices := []int{10, 7, 5, 8, 11, 2, 6}
	fmt.Println(GetGreatestProfit(prices))
}
