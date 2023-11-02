package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("create a new bill name ", reader)
	b := newBill(name)

	fmt.Println("Created the bill -", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add an item, s - save a bill, t - add a tip) ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name ", reader)
		price, _ := getInput("Item price ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err !=  nil {
			fmt.Println("price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, p)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved file - ", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount($) ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err !=  nil {
			fmt.Println("tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("item added - ", tip)
		promptOptions(b)
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}
	
}

func workingWithStructsTuts() bill {
	myBill := newBill("wario's bill")
	myBill.updateTip(12)

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("coffee", 3.50)
	myBill.addItem("carrot pie", 6.49)

	return myBill
}

func main()  {
	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill.format())
}