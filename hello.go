// go run ./main.go

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/wcharczuk/go-chart/v2"
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
	/*
		fmt.Println("-----------------------------")
		txList, err := tolist("/Users/xulei/creditcard/may.csv")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("-----------------------------")
		for _, t := range txList {
			fmt.Println("TX:")
			fmt.Println(t.amount)
			fmt.Println(t.payee)
			fmt.Println(t.category)
		}
		category_output := tocategory(txList)
		fmt.Println(calculate(category_output))
	*/
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)

}

type Tx struct {
	payee    string  // ""
	amount   float32 // 0
	category string  // ""
}

func tolist(filepath string) ([]Tx, error) {
	var txlist []Tx
	var firstline []string
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
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
						return nil, err
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
	return txlist, nil
}
func tocategory(txlist []Tx) map[string][]Tx {
	txmap := make(map[string][]Tx)
	for _, tx := range txlist {
		if _, ok := txmap[tx.category]; !ok {
			txmap[tx.category] = []Tx{}
		}
		txmap[tx.category] = append(txmap[tx.category], tx)
	}
	return txmap
}
func calculate(categoryoutput map[string][]Tx) map[string]float32 {
	catPercMap := make(map[string]float32)
	var totalamount float32
	for _, value := range categoryoutput {
		for _, tx := range value {
			totalamount = totalamount + -tx.amount
		}
	}
	for key, value := range categoryoutput {
		var totalcategory float32
		for _, tx := range value {
			totalcategory = totalcategory + -tx.amount
		}
		categoryPercent := totalcategory / totalamount * 100
		catPercMap[key] = categoryPercent
		fmt.Println(totalamount)
	}
	return catPercMap
}
