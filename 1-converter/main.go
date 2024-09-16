package main

import "fmt"

func main() {
	usdToEur, usdToRub := getUserInput()
	eurToRub := (1.0 / usdToEur) * usdToRub
	fmt.Println(eurToRub)
}

func getUserInput() (float64, float64) {
	var usdToEur float64
	var usdToRub float64
	fmt.Print("Введите курс USD к EUR: ")
	fmt.Scan(&usdToEur)
	fmt.Print("Введите курс USD к RUB: : ")
	fmt.Scan(&usdToRub)
	return usdToEur, usdToRub
}

func convert(val int, fromCurr string, toCurr string) {
	return
}
