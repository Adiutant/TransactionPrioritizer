package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"testTaskSec/model"
	"testTaskSec/utils"
	"time"
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
	totalTime := time.Millisecond * 1000
	res, err := utils.Prioritize(transactions, totalTime)
	if err != nil {
		return
	}
	//fmt.Println(res)
	fmt.Printf("Can process %v USD in %d msec\n", model.Sum(res), totalTime.Milliseconds())

	totalTime = time.Millisecond * 50
	res, err = utils.Prioritize(transactions, totalTime)
	if err != nil {
		return
	}
	//fmt.Println(res)
	fmt.Printf("Can process %v USD in %d msec\n", model.Sum(res), totalTime.Milliseconds())

	totalTime = time.Millisecond * 60
	res, err = utils.Prioritize(transactions, totalTime)
	if err != nil {
		return
	}
	//fmt.Println(res)
	fmt.Printf("Can process %v USD in %d msec\n", model.Sum(res), totalTime.Milliseconds())

	totalTime = time.Millisecond * 90
	res, err = utils.Prioritize(transactions, totalTime)
	if err != nil {
		return
	}
	//fmt.Println(res)
	fmt.Printf("Can process %v USD in %d msec\n", model.Sum(res), totalTime.Milliseconds())
}
