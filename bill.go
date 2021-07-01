package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"
	"strconv"
)

type Bill struct {
	name string
	items map[string] float64
	tip float64
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt + ":")
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func newBill(name string) Bill {
	return Bill {
		name: name,
		items: map[string]float64{ },
		tip: 0,
	}
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("What is your name?", reader)
	newBill := newBill(name)

	return newBill
}

func (bill *Bill) updateName(name string) {
	bill.name = name
}

func (bill *Bill) updateTip(tip float64) {
	bill.tip = tip
}

func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *Bill) removeItem(name string) {
	delete(bill.items, name)
}

func (bill *Bill) format() string {
	lineItemFormat := "%-25v ...$%0.2f\n"
	formatedString := fmt.Sprintf("Bill Breakdown for %v: \n", bill.name)
	formatedString += (strings.Repeat("-", 35) + "\n")
	total := 0.0
	for item, cost := range bill.items {
		formatedString += fmt.Sprintf(lineItemFormat, item + ":", cost)
		total += cost
	}
	formatedString += (strings.Repeat("-", 35) + "\n")
	tip := total * bill.tip / 100.0
	formatedString += fmt.Sprintf(lineItemFormat, "Tip (" + fmt.Sprintf("%0.1f", bill.tip) + "%): ", tip)
	total += tip
	formatedString += fmt.Sprintf(lineItemFormat, "Total:", total)

	return formatedString
}

func (bill *Bill) save() {
	data := []byte(bill.format())
	err := ioutil.WriteFile("bills/bill_for_" + bill.name + ".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Saved bill!")
}

func promptOptions(bill *Bill) {
	optionPrompt := "Choose option (a - add item, r - remove item, t - add tip, s - save bill, e - exit)"
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput(optionPrompt, reader)
	for option != "e" {
		switch option {
		case "a":
			name, _ := getInput("What did you buy?", reader)
			priceInput, _ := getInput("How much did it cost?", reader)
			price, priceErr := strconv.ParseFloat(priceInput, 64)
			if priceErr == nil {
				bill.addItem(name, price)
			} else {
				fmt.Println("Invalid price!")
			}
		case "r":
			name, _ := getInput("What do you want to remove?", reader)
			bill.removeItem(name)
		case "t":
			tipInput, _ := getInput("How much do you want to tip?", reader)
			tip, tipErr := strconv.ParseFloat(tipInput, 64)
			if tipErr == nil {
				bill.updateTip(tip)
			} else {
				fmt.Println("Invalid tip!")
			}
		case "s":
			fmt.Println("Saving bill....")
			bill.save()
		default:
			fmt.Println(option, "is an invalid option. Try again...")
		}

		option, _ = getInput(optionPrompt, reader)
	}
}