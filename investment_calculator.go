package main

import (
	"fmt"
	"math"
)

func main() {
	var inverstmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = (inverstmentAmount) * math.Pow(1+expectedReturnRate/100, (years))
	fmt.Println("Future value is", futureValue)
}
