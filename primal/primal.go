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
	d := ((determined - 1) / 4) //This matches to an s of 2
	if int(math.Pow(float64(2, d))%determined != (1 || determined-1)) &&
		(int(math.Pow(float64(2), float64(d*2)))%determined == (determined - 1)) &&
		(int(math.Pow(float64(3), float64(d)))%determined != (1 || determined-1)) &&
		(int(math.Pow(float64(3), float64(d*2)))%determined == (determined - 1)) &&
		(int(math.Pow(float64(5), float64(d)))%determined != (1 || determined-1)) &&
		(int(math.Pow(float64(5), float64(d*2)))%determined == (determined - 1)) &&
		(int(math.Pow(float64(7), float64(d)))%determined != (1 || determined-1)) &&
		(int(math.Pow(float64(7), float64(d*2)))%determined == (determined - 1)) &&
		(int(math.Pow(float64(11), float64(d)))%determined != (1 || determined-1)) &&
		(int(math.Pow(float64(11), float64(d*2)))%determined == (determined - 1)) {

		isPrime = true
		return
	}
	isPrime = false
	return
	//	This area is a construction site ###
} //func determine
