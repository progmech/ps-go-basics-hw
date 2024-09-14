package main

import "fmt"

func main() {
	const usdToEur = 0.9029
	const usdToRub = 90.1361 // 99.7572
	eurToRub := (1.0 / usdToEur) * usdToRub
	fmt.Println(eurToRub)
}
