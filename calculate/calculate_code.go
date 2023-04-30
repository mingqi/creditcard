package calculate

import "example/hello/tx"

func Calculate(categoryoutput map[string][]tx.Tx) map[string]float32 {
	catPercMap := make(map[string]float32)
	var totalamount float32
	for _, value := range categoryoutput {
		for _, tx := range value {
			totalamount = totalamount + -tx.Amount
		}
	}
	for key, value := range categoryoutput {
		var totalcategory float32
		for _, tx := range value {
			totalcategory = totalcategory + -tx.Amount
		}
		categoryPercent := totalcategory / totalamount * 100
		catPercMap[key] = categoryPercent
	}
	return catPercMap
}
