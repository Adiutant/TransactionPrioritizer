package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"testTaskSec/model"
)

func main() {
	file, err := os.Open("data/transactions.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	readTx, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var transactions []model.Transaction
	for i := 1; i < len(readTx); i++ {
		transactions = append(transactions, model.Transaction{
			ID:              readTx[i][0],
			Amount:          readTx[i][1],
			BankName:        "default",
			BankCountryCode: readTx[i][2],
		})
	}

	res, err := model.Prioritize(transactions)
	if err != nil {
		return
	}
	fmt.Print(res)
}
