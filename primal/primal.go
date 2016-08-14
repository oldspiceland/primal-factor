package primal

import (
	"math"
)

//IsPrime takes an int, primed, and returns a bool, isPrime
func IsPrime(primed int) (isPrime bool) {
	isComp := fermat(primed) //first use fermat to check for compositeness
	if isComp == true {
		isPrime = false
		return
	} //if isComp == true
	isPrime = determine(primed)
	return
} //func IsPrime

func fermat(fermatted int) (isComp bool) {
	if fermatted > 1 {
		little := (int(math.Pow(float64(2), (float64(fermatted-1)))) % fermatted)
		if little != 1 {
			isComp = true
			return
		} //if little != 1
		isComp = false
		return
	} //if fermatted > 1
	return
} //func fermat

func determine(determined int) (isPrime bool) {
	//TODO: This area is a construction site ###
	baseArray := [5]int{2, 3, 5, 7, 11}
	primeArray := [5]bool{false, false, false, false, false}
	d := ((determined - 1) / 4) //This matches to an s of 2
	d2 := d * 2

	for i := 0; i < 5; i++ {
		if int(math.Pow(float64(baseArray[i]), float64(d))) != (1 || determined-1) {
			if int(math.Pow(float64(baseArray[i]), float64(d2))) == determined-1 {
				primeArray[i] = true
			}
		}
	}
	for i := 0; i < 5; i++ {
		if !primeArray[i] {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
	//	This area is a construction site ###
} //func determine
