package main

import (
	"fmt"
)

// Invoice struct
type Invoice struct {
	To     string
	Amount int
}

var invoices = []Invoice{
	{To: "John Doe", Amount: 120000},
	{To: "John Doe", Amount: 80000},
	{To: "Jane Doe", Amount: 100000},
}

func groupInvoiceByPerson(invoices []Invoice) []Invoice {
	// Make temporary invoice object that holds information about Person and its index for newInvoice
	tempInvoice := make(map[string]int)
	newInvoice := []Invoice{}
	for _, invoiceDetail := range invoices {
		personDetail := invoiceDetail.To
		amountDetail := invoiceDetail.Amount
		// Check if person already added into newInvoice
		value, isExist := tempInvoice[personDetail]
		if isExist {
			newInvoice[value].Amount += amountDetail
		} else {
			// Add new person and its amountDetail into newInvoice. Record its index information into tempInvoice
			currentLength := len(newInvoice)
			newInvoice = append(newInvoice, invoiceDetail)
			tempInvoice[personDetail] = currentLength
		}
	}
	return newInvoice
}

func main() {
	newInvoice := groupInvoiceByPerson(invoices)
	fmt.Printf("%+v\n", newInvoice)
}
