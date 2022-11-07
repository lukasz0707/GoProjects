package main

import (
	"fmt"
	"math"
)

// Given group of n number of people (assuming that there is equal propability of birth for every day of the year).
// Whats the propability that at least 2 people of the group have birthday on the same day.
// (Dont count 29 of February and assume year have 365 days).
// complexity O(n)
func sameBirthdayLoop(numberOfPeople float64) float64 {
	value1 := 1.0
	for i := 0.0; i < numberOfPeople; i++ {
		value1 *= 365.0 - i
	}
	return 1.0 - (value1 / (math.Pow(365.0, numberOfPeople)))
}

// Given propability return number of people
func sameBirthdayLog(propability float64) float64 {

	return math.Ceil(math.Sqrt(2 * 365 *
		math.Log(1/(1-propability))))
}

func main() {
	count := 50.
	for i := 0.; i < count; i++ {
		propability := sameBirthdayLoop(i)
		fmt.Println("Num of ppl Var:", i)
		fmt.Println("Propability:", propability)
		fmt.Println("Num of PPl:", sameBirthdayLog(propability))
		fmt.Println()
	}
}
