package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	// fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter the number of years: ")
	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	// fmt.Print("Enter the annual interest rate: ")
	outputText("Enter the annual interest rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, years, expectedReturnRate)
	// futureValue := (investmentAmount) * math.Pow(1+expectedReturnRate/100, (years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value is %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future real value is %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
	fv := (investmentAmount) * math.Pow(1+expectedReturnRate/100, (years))
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
