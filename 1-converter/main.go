package main

import (
	"fmt"
	"strings"
)

const (
	usdToEurRate = 0.89936034
	usdToRubRate = 90.999385
)

type ratesMap = map[string]map[string]float64
type currMap = map[string]float64

func main() {
	rates := initRatesMap()
	sourceCurrency := getSourceCurrency(rates)
	amount := getExchangeAmount()
	target, rate := getTargetCurrencyAndRate(sourceCurrency)
	exchangeResult := getExchangeResult(rate, amount)
	fmt.Printf("Вы получите %.2f %s\n", exchangeResult, target)
}

func initRatesMap() ratesMap {
	usdMap := currMap{"EUR": usdToEurRate, "RUB": usdToRubRate}
	eurMap := currMap{"USD": 1.0 / usdToEurRate, "RUB": (1.0 / usdToEurRate) * usdToRubRate}
	rubMap := currMap{"USD": 1.0 / usdToRubRate, "EUR": (1.0 / usdToRubRate) * usdToEurRate}
	resultMap := ratesMap{"USD": usdMap, "EUR": eurMap, "RUB": rubMap}
	return resultMap
}

func getSourceCurrency(rates ratesMap) currMap {
	var source string
	promptCurr := []string{}
	for k := range rates {
		promptCurr = append(promptCurr, k)
	}
	for {
		fmt.Printf("Введите исходную валюту (%s): ", strings.Join(promptCurr, ","))
		fmt.Scan(&source)
		m, ok := rates[source]
		if ok {
			return m
		}
	}
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

func getTargetCurrencyAndRate(sourceCur currMap) (string, float64) {
	var target string
	promptCurr := []string{}
	for k := range sourceCur {
		promptCurr = append(promptCurr, k)
	}
	for {
		fmt.Printf("Введите целевую валюту (%s): ", strings.Join(promptCurr, ","))
		fmt.Scan(&target)
		m, ok := sourceCur[target]
		if ok {
			return target, m
		}
	}
}

func getExchangeResult(rate float64, amount float64) float64 {
	return rate * amount
}
