package main

import (
	"fmt"
)

const (
	usdToEurRate = 0.89936034
	usdToRubRate = 90.999385
	usdAbbr      = "USD"
	eurAbbr      = "EUR"
	rubAbbr      = "RUB"
)

func main() {
	sourceCurrency := getSourceCurrency()
	amount := getExchangeAmount()
	targetCurrency := getTargetCurrency(sourceCurrency)
	exchangeResult := getExchangeResult(sourceCurrency, amount, targetCurrency)
	fmt.Printf("Вы получите %.2f %s\n", exchangeResult, targetCurrency)
}

func getSourceCurrency() string {
	var cur string
	for cur != usdAbbr && cur != eurAbbr && cur != rubAbbr {
		fmt.Printf("Введите исходную валюту (%s, %s или %s): ", usdAbbr, eurAbbr, rubAbbr)
		fmt.Scan(&cur)
	}
	return cur
}

func getExchangeAmount() float64 {
	var cur float64
	for {
		fmt.Printf("Введите сумму исходной валюты: ")
		_, err := fmt.Scan(&cur)
		if err != nil {
			fmt.Printf("Ошибка ввода: %s\n", err)
			continue
		}
		if cur <= 0 {
			fmt.Println("Сумма должна быть больше нуля.")
			continue
		}
		return cur
	}
}

func getTargetCurrency(sourceCur string) string {
	var cur string
	restCur1, restCur2 := getRestCurr(sourceCur)
	for cur != restCur1 && cur != restCur2 {
		fmt.Printf("Введите целевую валюту (%s или %s): ", restCur1, restCur2)
		fmt.Scan(&cur)
	}
	return cur
}

func getRestCurr(sourceCur string) (string, string) {
	switch sourceCur {
	case usdAbbr:
		return eurAbbr, rubAbbr
	case eurAbbr:
		return usdAbbr, rubAbbr
	default:
		return eurAbbr, usdAbbr
	}
}

func getExchangeResult(sourceCur string, amount float64, targetCur string) float64 {
	switch sourceCur {
	case usdAbbr:
		if targetCur == eurAbbr {
			return calcUsdToEur(amount)
		} else {
			return calcUsdToRub(amount)
		}
	case eurAbbr:
		if targetCur == usdAbbr {
			return calcEurToUsd(amount)
		} else {
			return calcEurToRub(amount)
		}
	default:
		if targetCur == usdAbbr {
			return calcRubToUsd(amount)
		} else {
			return calcRubToEur(amount)
		}
	}
}

func calcUsdToEur(amount float64) float64 {
	return usdToEurRate * amount
}

func calcUsdToRub(amount float64) float64 {
	return usdToRubRate * amount
}

func calcEurToUsd(amount float64) float64 {
	return amount / usdToEurRate
}

func calcEurToRub(amount float64) float64 {
	return (1.0 / usdToEurRate) * usdToRubRate * amount
}

func calcRubToUsd(amount float64) float64 {
	return amount / usdToRubRate
}
func calcRubToEur(amount float64) float64 {
	return (1.0 / usdToRubRate) * usdToEurRate * amount
}
