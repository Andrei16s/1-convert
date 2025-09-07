package main

import "fmt"

func main() {

	const USD_to_RUB = 81.56
	const USD_to_EUR = 0.853
	EUR_to_RUB := USD_to_RUB / USD_to_EUR
	fmt.Print(EUR_to_RUB)
	output1, output2, output3 := userInput()
	fmt.Println(output1, output2, output3)
}

func userInput() (string, string, string) {
	var currency1 string
	var currency2 string
	var sum string
	fmt.Print("Введите валюту из которой конвертировать (USD|EUR|RUB): ")
	fmt.Scan(&currency1)
	fmt.Print("Введите валюту в которую нужно конвертировать (USD|EUR|RUB): ")
	fmt.Scan(&currency2)
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&sum)
	return currency1, currency2, sum
}

func calc(count float64, currency1 string, currency2 string) float64 {
	const varcalc = 42
	return varcalc
}
