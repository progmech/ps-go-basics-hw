package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	opAvg = "AVG"
	opSum = "SUM"
	opMed = "MED"
	sep   = ","
)

var ops = map[string]func(...float64) float64{
	opAvg: getAvg,
	opSum: getSum,
	opMed: getMed,
}

func main() {
	oper := getOperation()
	numbers := getNumbers()
	result := ops[oper](numbers...)
	fmt.Printf("Результат операции %s: %f\n", oper, result)
}

func getAvg(numbers ...float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum = sum + val
	}
	return sum / float64(len(numbers))
}

func getSum(numbers ...float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum = sum + val
	}
	return sum
}

func getMed(numbers ...float64) float64 {
	slices.Sort(numbers)
	lens := len(numbers)
	midIdx := lens / 2
	if lens%2 == 0 {
		return (numbers[midIdx] + numbers[midIdx-1]) / 2
	}
	return numbers[midIdx]
}

func getOperation() string {
	var oper string
	for oper != opAvg && oper != opSum && oper != opMed {
		fmt.Printf("Введите операцию (%s, %s, %s): ", opAvg, opSum, opMed)
		fmt.Scan(&oper)
	}
	return oper
}

func getNumbers() []float64 {
	var numberString string
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&numberString)
	return stringToSliceOfInt(numberString)
}

func stringToSliceOfInt(numberString string) []float64 {
	sliceOfStrings := strings.Split(numberString, sep)
	sliceOfInt := []float64{}
	for _, v := range sliceOfStrings {
		floatValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		sliceOfInt = append(sliceOfInt, floatValue)
	}
	return sliceOfInt
}
