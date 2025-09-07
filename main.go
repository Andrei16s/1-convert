package main

import "fmt"

func main() {

	const USD_to_RUB = 81.56
	const USD_to_EUR = 0.853
	EUR_to_RUB := USD_to_RUB / USD_to_EUR
	fmt.Print(EUR_to_RUB)
}
