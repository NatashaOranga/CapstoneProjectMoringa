package main

//declares an executable program

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//imports packages

// account struct representing bank acc
type Account struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

// deposit(add) money function
func (a *Account) Deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}
	a.Balance += amount
	return true
}

// withdraw(remove) money function
func (a *Account) Withdraw(amount float64) bool {
	if amount <= 0 || amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}

// shows acc details
func (a *Account) DisplayInfo() {
	fmt.Printf("\nAccount: %s\nHolder: %s\nBalance: %.2f\n",
		a.AccountNumber, a.HolderName, a.Balance)
}

func main() {
	//map to store all accounts: account number -> Account pointer
	accounts := make(map[string]*Account)
	reader := bufio.NewReader(os.Stdin)

	//main program loop
	for {
		fmt.Println("\n=== Bank Management System ===")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Check Balance")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			//create account
			fmt.Print("Enter account number: ")
			accNum, _ := reader.ReadString('\n')
			accNum = strings.TrimSpace(accNum)

			fmt.Print("Enter holder name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter initial balance: ")
			balStr, _ := reader.ReadString('\n')
			balance, _ := strconv.ParseFloat(strings.TrimSpace(balStr), 64)

			accounts[accNum] = &Account{accNum, name, balance}
			fmt.Println("Account created successfully!")

		case "2":
			//deposit money
			fmt.Print("Enter account number: ")
			accNum, _ := reader.ReadString('\n')
			accNum = strings.TrimSpace(accNum)

			if acc, exists := accounts[accNum]; exists {
				fmt.Print("Enter amount: ")
				amtStr, _ := reader.ReadString('\n')
				amount, _ := strconv.ParseFloat(strings.TrimSpace(amtStr), 64)

				if acc.Deposit(amount) {
					fmt.Printf("Deposited %.2f successfully!\n", amount)
					acc.DisplayInfo()
				} else {
					fmt.Println("Invalid amount")
				}
			} else {
				fmt.Println("Account not found")
			}

		case "3":
			//withdraw money
			fmt.Print("Enter account number: ")
			accNum, _ := reader.ReadString('\n')
			accNum = strings.TrimSpace(accNum)

			if acc, exists := accounts[accNum]; exists {
				fmt.Print("Enter amount: ")
				amtStr, _ := reader.ReadString('\n')
				amount, _ := strconv.ParseFloat(strings.TrimSpace(amtStr), 64)

				if acc.Withdraw(amount) {
					fmt.Printf("Withdrew %.2f successfully!\n", amount)
					acc.DisplayInfo()
				} else {
					fmt.Println("Insufficient funds or invalid amount")
				}
			} else {
				fmt.Println("Account not found")
			}

		case "4":
			//check balance
			fmt.Print("Enter account number: ")
			accNum, _ := reader.ReadString('\n')
			accNum = strings.TrimSpace(accNum)

			if acc, exists := accounts[accNum]; exists {
				acc.DisplayInfo()
			} else {
				fmt.Println("Account not found")
			}

		case "5":
			fmt.Println("Thank you for using our bank!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
