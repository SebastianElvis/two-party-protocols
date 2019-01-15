// Testing
package ff

import (
	"math/big"
	"testing"
)

func TestFieldElementFactory(t *testing.T) {

	result, err := NewFieldElement(2, 31)
	caseN := big.NewInt(2)
	caseP := big.NewInt(31)

	if result.n.Cmp(caseN) != 0 && result.p.Cmp(caseP) != 0 || err != nil {
		t.Fatalf("Expected %d,%d but got %v", caseN, caseP, result)
	}

}

func TestFieldElementEqual(t *testing.T) {

	a := FieldElement{*big.NewInt(15), *big.NewInt(31)}
	b := FieldElement{*big.NewInt(15), *big.NewInt(31)}

	result := a.Equal(b)

	if result != true {
		t.Fatalf("Expected true, but got %v", result)
	}

}

func TestFieldElementNotEqual(t *testing.T) {

	a := FieldElement{*big.NewInt(15), *big.NewInt(31)}
	b := FieldElement{*big.NewInt(7), *big.NewInt(31)}

	result := a.NotEqual(b)

	if result != true {
		t.Fatalf("Expected true, but got %v", result)
	}

}

func TestFieldElementAdd(t *testing.T) {

	a := FieldElement{*big.NewInt(17), *big.NewInt(31)}
	b := FieldElement{*big.NewInt(21), *big.NewInt(31)}

	result, err := Add(a, b)

	if err != nil {
		t.Fatalf("Error : %v", err)
	}

	truth := FieldElement{*big.NewInt(7), *big.NewInt(31)}

	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}

}

func TestFieldElementSub(t *testing.T) {
	a := FieldElement{*big.NewInt(29), *big.NewInt(31)}
	b := FieldElement{*big.NewInt(4), *big.NewInt(31)}

	result, err := Sub(a, b)

	if err != nil {
		t.Fatalf("Error : %v", err)
	}

	truth := FieldElement{*big.NewInt(25), *big.NewInt(31)}

	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}
}

func TestFieldElementMul(t *testing.T) {

	a := FieldElement{*big.NewInt(24), *big.NewInt(31)}
	b := FieldElement{*big.NewInt(19), *big.NewInt(31)}

	result, err := Mul(a, b)

	if err != nil {
		t.Fatalf("Error : %v", err)
	}

	truth := FieldElement{*big.NewInt(22), *big.NewInt(31)}

	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}

}

func TestFieldElementPow(t *testing.T) {

	a := FieldElement{*big.NewInt(17), *big.NewInt(31)}

	result, err := Pow(a, 3)

	if err != nil {
		t.Fatalf("Error : %v", err)
	}

	truth := FieldElement{*big.NewInt(15), *big.NewInt(31)}

	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}

}
