package main

import "testTaskSec/model"

func main() {
	_, err := model.Prioritize(make([]model.Transaction,0))
	if err != nil {
		return
	}
}
