// Package main implements a profit calculator that computes business metrics
// including Earnings Before Tax (EBT), Profit, and Profit Ratio.
// The program provides an interactive command-line interface and persists
// calculation results to a file.
package main

import (
	"errors"
	"fmt"
	"os"
)

// dataFile is the name of the file where calculation results are stored
const dataFile = "profit_data.txt"

// main is the entry point of the program. It:
// 1. Prompts the user for revenue, expense, and tax rate
// 2. Calculates business metrics
// 3. Saves results to file
// 4. Displays results to the user
func main() {
	//revenue, expense, tax_rate := inputData()
	revenue, err := getUserInput("Please enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expense, err := getUserInput("Please enter Expense: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Please enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expense, taxRate)

	writeDataToFile(ebt, profit, ratio)

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}

// writeDataToFile writes the calculation results (EBT, profit, and ratio) to a file.
// The data is formatted as a string with each metric on a new line.
// The file is created if it doesn't exist, or overwritten if it does.
//
// Parameters:
//   - ebt: Earnings Before Tax
//   - profit: Net profit after tax
//   - ratio: Profit ratio (EBT/Profit)
func writeDataToFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(dataFile, []byte(data), 0644)
}

// getUserInput prompts the user for input and validates the entered value.
// It ensures the input is a positive number.
//
// Parameters:
//   - userText: The prompt message to display to the user
//
// Returns:
//   - float64: The validated user input
//   - error: An error if the input is invalid (negative or zero)
func getUserInput(userText string) (float64, error) {
	var userInput float64
	fmt.Print(userText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("input can't be negative")
	}

	return userInput, nil
}

// func inputData() (revenue, expense, tax_rate float64) {
// 	fmt.Print("Plese enter Revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Plese enter Expense: ")
// 	fmt.Scan(&expense)

// 	fmt.Print("Plese enter Tax Rate: ")
// 	fmt.Scan(&tax_rate)

// 	return
// }

// calculateProfit computes key business metrics based on revenue, expense, and tax rate.
//
// Parameters:
//   - revenue: Total business income
//   - expense: Total business expenses
//   - taxRate: Tax rate as a percentage (e.g., 10 for 10%)
//
// Returns:
//   - ebt: Earnings Before Tax (revenue - expense)
//   - profit: Net profit after tax (ebt * (1 - taxRate/100))
//   - ratio: Profit ratio (ebt / profit)
func calculateProfit(revenue, expense, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - (taxRate / 100))
	ratio = ebt / profit

	return
}
