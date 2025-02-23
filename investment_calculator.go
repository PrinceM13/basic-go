package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var inverstmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&inverstmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the annual interest rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := (inverstmentAmount) * math.Pow(1+expectedReturnRate/100, (years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)
}
