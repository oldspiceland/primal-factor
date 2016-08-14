package primal

import (
	"github.com/oldspiceland/primal-factor/primal"
	"testing"
)

func TestIsPrime(t *testing.T) {
	primedTest := 13577
	isPrime := primal.IsPrime(primedTest)
	if isPrime != true {
		t.Error("Expected true, got ", isPrime)
	}
}
