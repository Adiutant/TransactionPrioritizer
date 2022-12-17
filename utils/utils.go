package utils

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"testTaskSec/model"
	"time"
)

func Prioritize(tx []model.Transaction, totalTime time.Duration) ([]model.Transaction, error) {
	jsonBytes, err := ioutil.ReadFile("data/api_latencies.json")
	if err != nil {
		return nil, err
	}
	var lat map[string]int
	json.Unmarshal(jsonBytes, &lat)
	transactions := model.Transactions{
		TxList:    tx,
		Latencies: lat,
	}
	sort.Sort(transactions)
	currentTime := 0
	resultSlice := make([]model.Transaction, 0)
	for i := len(transactions.TxList) - 1; i >= 0; i-- {
		if int64(currentTime+transactions.Latencies[transactions.TxList[i].BankCountryCode]) > totalTime.Milliseconds() {
			break
		}
		currentTime += transactions.Latencies[transactions.TxList[i].BankCountryCode]
		resultSlice = append(resultSlice, transactions.TxList[i])
	}
	return resultSlice, nil

}
