// go run ./main.go

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Age    int
	Name   string
	Height int
}

func findPeople(name string, people []Person) Person {
	var theperson Person
	for i := 0; i < 10; i = i + 1 {
		if people[i].Name == name {
			theperson = people[i]
			break
		}
	}
	return theperson
}
func main() {
	csvFile, err := os.Open("/Users/xulei/creditcard/may.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		fmt.Println(line[0])
		//t := Tx{
		//payee:    "fff",
		//amount:   10,
		//category: "house",
	}
	// fmt.Println(tolist("/Users/xulei/creditcard/may.csv"))
	fmt.Println("-----------------------------")
	txList := tolist("/Users/xulei/creditcard/may.csv")
	for _, t := range txList {
		fmt.Println("TX:")
		fmt.Println(t.amount)
		fmt.Println(t.payee)
		fmt.Println(t.category)
	}
}

type Tx struct {
	payee    string  // ""
	amount   float32 // 0
	category string  // ""
}

func tolist(filepath string) []Tx {
	var txlist []Tx
	var firstline []string
	csvFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		firstline = line
		break
	}
	for _, line := range csvLines {
		tx := Tx{}

		for i := 0; i < len(line); i = i + 1 {
			if firstline[i] == "Payee" {
				if line[i] != "Payee" {
					tx.payee = line[i]
				}
			}
			if firstline[i] == "Amount" {
				if line[i] != "Amount" {
					f, err := strconv.ParseFloat(line[i], 32)
					if err != nil {
						fmt.Println(err)
						continue
					}
					tx.amount = float32(f)
				}
			}
			if firstline[i] == "Category" {
				if line[i] != "Category" {
					tx.category = line[i]
				}
			}
		}
		if tx.category != "" {
			txlist = append(txlist, tx)
		}

	}
	return txlist
}
