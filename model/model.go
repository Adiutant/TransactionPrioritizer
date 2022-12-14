package model

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
)

type Interface interface {
	// Len - количество элементов в коллекции.
	Len() int
	// Less сообщает должен ли элемента с индексом i
	// быть отсортированным перед элементом с индексом j.
	Less(i, j int) bool
	// Swap меняет местами элементы с индексами i и j.
	Swap(i, j int)
}
type Transaction struct {
	// a UUID of transaction
	ID string
	// in USD, typically a value between "0.01" and "1000" USD.
	Amount string
	// bank name, e.g. "Bank of Scotland"
	BankName string
	// a 2-letter country code of where the bank is located
	BankCountryCode string
}
type Transactions struct {
	txList    []Transaction
	latencies map[string]int
}

func (tx Transactions) Len() int {
	return len(tx.txList)
}
func (tx Transactions) Less(i, j int) bool {
	firstAmount, err := strconv.ParseFloat(tx.txList[i].Amount, 32)
	if err != nil {
		return false
	}
	secondAmount, err := strconv.ParseFloat(tx.txList[j].Amount, 32)
	if err != nil {
		return false
	}
	return firstAmount/float64(tx.latencies[tx.txList[i].BankCountryCode]) < secondAmount/float64(tx.latencies[tx.txList[j].BankCountryCode])
}

type FraudDetectionResult struct {
	TransactionID string
	IsFraudulent  bool
}

func (tx Transactions) Swap(i, j int) {
	tx.txList[i], tx.txList[j] = tx.txList[j], tx.txList[i]
}

type FraudDetectionResults []FraudDetectionResult

func sum(array []Transaction) float64 {
	result := float64(0)
	for _, v := range array {
		amountVal, _ := strconv.ParseFloat(v.Amount, 32)
		result += amountVal
	}
	return result
}
func Prioritize(tx []Transaction) ([]Transaction, error) {
	jsonBytes, err := ioutil.ReadFile("data/api_latencies.json")
	if err != nil {
		return nil, err
	}
	var lat map[string]int
	json.Unmarshal(jsonBytes, &lat)
	transactions := Transactions{
		txList:    tx,
		latencies: lat,
	}
	sort.Sort(transactions)
	time := 0
	resultSlice := make([]Transaction, 0)
	for i := len(transactions.txList) - 1; i >= 0 && time < 1000; i-- {
		time += transactions.latencies[transactions.txList[i].BankCountryCode]
		resultSlice = append(resultSlice, transactions.txList[i])
	}
	return resultSlice, nil

}

//func processTransactions(tx []Transaction) ([]FraudDetectionResults, error) {
//	results := make(FraudDetectionResults, len(tx))
//	for i := range tx {
//		results[i] = FraudDetectionResult{
//			TransactionID: tx[i].ID,
//			IsFraudulent:  processTransaction[i],
//		}
//	}
//	return results, nil
//}
