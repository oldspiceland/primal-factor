package primal

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	primedTest := 13577
	isPrime := IsPrime(primedTest)
	if isPrime != true {
		t.Error("Expected true, got ", isPrime)
	}
}
