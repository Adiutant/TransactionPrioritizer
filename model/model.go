package model

import (
	"github.com/shopspring/decimal"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
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
	TxList    []Transaction
	Latencies map[string]int
}

func (tx Transactions) Len() int {
	return len(tx.TxList)
}
func (tx Transactions) Less(i, j int) bool {
	firstAmount, err := decimal.NewFromString(tx.TxList[i].Amount)
	if err != nil {
		return false
	}
	secondAmount, err := decimal.NewFromString(tx.TxList[j].Amount)
	if err != nil {
		return false
	}
	return firstAmount.Div(decimal.NewFromInt(int64(tx.Latencies[tx.TxList[i].BankCountryCode]))).LessThan(secondAmount.Div(decimal.NewFromInt(int64(tx.Latencies[tx.TxList[j].BankCountryCode]))))
}

type FraudDetectionResult struct {
	TransactionID string
	IsFraudulent  bool
}

func (tx Transactions) Swap(i, j int) {
	tx.TxList[i], tx.TxList[j] = tx.TxList[j], tx.TxList[i]
}

type FraudDetectionResults []FraudDetectionResult

func Sum(array []Transaction) string {
	result := decimal.NewFromInt(0)
	for _, v := range array {
		amountVal, _ := decimal.NewFromString(v.Amount)
		result = result.Add(amountVal)
	}
	return result.String()
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
