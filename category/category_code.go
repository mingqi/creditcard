package category

import "example/hello/tx"

func Tocategory(txlist []tx.Tx) map[string][]tx.Tx {
	txmap := make(map[string][]tx.Tx)
	for _, t_x := range txlist {
		if _, ok := txmap[t_x.Category]; !ok {
			txmap[t_x.Category] = []tx.Tx{}
		}
		txmap[t_x.Category] = append(txmap[t_x.Category], t_x)
	}
	return txmap
}