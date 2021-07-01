package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Order Calculator!")
	bill := createBill()
	bill.updateTip(10.0)
	promptOptions(&bill)
	fmt.Println(bill.format())
}