package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Can't  continue, sorry.")

	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us: ", randomdata.PhoneNumber())

	for {

		presentOptions()
		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)

		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. you can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New Amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return

		}

	}

}
