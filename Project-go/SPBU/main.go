package main

import (
	"SPBU/helper"
	"fmt"
)

func main() {

	var userOption string
	var userBuy float32
	var user bool = true
	var cost float32

loop:

	for user {
		helper.GreetUser()
		fmt.Print("pilih menu (1 - 4) : ")
		fmt.Scanln(&userOption)

		switch userOption {
		case "1":
			//opsi premium
			fmt.Print("ingin beli berapa : ")
			fmt.Scanln(&userBuy)

			cost = helper.CostPremium(userBuy)
			helper.IsSuccess(userBuy, cost)

			user = false

		case "2":
			//opsi pertalite
			fmt.Print("ingin beli berapa : ")
			fmt.Scanln(&userBuy)

			cost = helper.CostPertalite(userBuy)
			helper.IsSuccess(userBuy, cost)

			user = false
		case "3":
			//opsi pertamax
			fmt.Print("ingin beli berapa : ")
			fmt.Scanln(&userBuy)

			cost = helper.CostPertamax(userBuy)
			helper.IsSuccess(userBuy, cost)

			user = false
		case "4":
			//opsi solar
			fmt.Print("ingin beli berapa : ")
			fmt.Scanln(&userBuy)

			cost = helper.CostSolar(userBuy)
			helper.IsSuccess(userBuy, cost)

			user = false

		case "5":
			fmt.Println("mohon maaf atas ketidaknyamanan\nselamt jalan")
			user = false

		default:
			fmt.Println("tidak ada dalam menu !!!")
			continue loop

		}
	}

}
