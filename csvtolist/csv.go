package csvtolist

import (
	"encoding/csv"
	"example/hello/tx"
	"fmt"
	"os"
	"strconv"
)

func Tolist(filepath string) ([]tx.Tx, error) {
	var txlist []tx.Tx
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
		tx := tx.Tx{}

		for i := 0; i < len(line); i = i + 1 {
			if firstline[i] == "Payee" {
				if line[i] != "Payee" {
					tx.Payee = line[i]
				}
			}
			if firstline[i] == "Amount" {
				if line[i] != "Amount" {
					f, err := strconv.ParseFloat(line[i], 32)
					if err != nil {
						return nil, err
					}
					tx.Amount = float32(f)
				}
			}
			if firstline[i] == "Category" {
				if line[i] != "Category" {
					tx.Category = line[i]
				}
			}
		}
		if tx.Category != "" {
			txlist = append(txlist, tx)
		}

	}
	return txlist, nil
}
