package helper

import (
	"fmt"
	"time"
)

func ValidateBuy(userBuy float32) (bool, bool) {

	var maxBuy bool = userBuy <= float32(1000000)
	var minBuy bool = userBuy >= float32(1000)

	return maxBuy, minBuy

}

func GreetUser() {
	fmt.Println("=====Selamat datang di SPBU =====")
	fmt.Println("silahkan pilih bbm yang ingin dibeli")
	fmt.Println("1. Premium")
	fmt.Println("2. Pertalite")
	fmt.Println("3. Pertamax")
	fmt.Println("4. Solar")
	fmt.Println("5. exit")
}

func CostPertalite(userInput float32) float32 {
	var price float32 = 7650
	cost := userInput / price

	return cost

}

func CostPertamax(userInput float32) float32 {
	var price float32 = 9000
	cost := userInput / price

	return cost

}

func CostPremium(userInput float32) float32 {
	var price float32 = 6000
	cost := userInput / price

	return cost

}

func CostSolar(userInput float32) float32 {
	var price float32 = 5000
	cost := userInput / price

	return cost

}

func printNote(userBuy float32, cost float32) {

	fmt.Println("nota akan segera dicetak")
	time.Sleep(3 * time.Second)
	fmt.Println("==================================")
	fmt.Printf("beli : %v \n", userBuy)
	fmt.Printf("value : %.2f L \n", cost)
	fmt.Println("terimakasih selamat jalan")
	fmt.Println("==================================")

}

func IsSuccess(userBuy float32, cost float32) {

	maxBuy, minBuy := ValidateBuy(userBuy)

	if maxBuy && minBuy {
		printNote(userBuy, cost)

	} else {

		if !maxBuy {
			fmt.Println("pembelian melebihi batas maximal (1.000.000)")
		}
		if !minBuy {
			fmt.Println("belum mencapai minimum pembelian (1.000)")
		}
	}
}
