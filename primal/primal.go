package primal

import (
	"math"
)

//IsPrime takes an int, primed, and returns a bool, isPrime
func IsPrime(primed int) (isPrime bool) {
	isComp := fermat(primed) //first use fermat to check for compositeness
	if isComp == true {
		isPrime == false
		return
	} //if isComp == true
	isPrime = determine(primed)
} //func IsPrime

func fermat(fermatted int) (isComp bool) {
	if fermatted > 1 {
		little := (int(math.Pow(2, (fermatted - 1)))) % fermatted
		if little != 1 {
			isComp = true
			return
		} //if little != 1
		isComp = false
		return
	} //if fermatted > 1
} //func fermat

func determine(determined int) (isPrime bool) {
	//TODO: This area is a construction site ###
	d := ((determined - 1) / 4) //This matches to an s of 2
	if (int(math.Pow(2, d))%determined != (1 || determined-1)) &&
		(int(math.Pow(2, d*2))%determined == (determined - 1)) &&
		(int(math.Pow(3, d))%determined != (1 || determined-1)) &&
		(int(math.Pow(3, d*2))%determined == (determined - 1)) &&
		(int(math.Pow(5, d))%determined != (1 || determined-1)) &&
		(int(math.Pow(5, d*2))%determined == (determined - 1)) &&
		(int(math.Pow(7, d))%determined != (1 || determined-1)) &&
		(int(math.Pow(7, d*2))%determined == (determined - 1)) &&
		(int(math.Pow(11, d))%determined != (1 || determined-1)) &&
		(int(math.Pow(11, d*2))%determined == (determined - 1)) {

		isPrime = true
		return
	} else {
		isPrime = false
		return
	}
	//	This area is a construction site ###
} //func determine
