// Package ff implements finite field arithmetic over any field (usually prime)
package ff

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// FieldElement creates a finite field element n over p
type FieldElement struct {
	n big.Int // n is in {0..p-1}
	p big.Int // Fp
}

func (f *FieldElement) N() *big.Int {
	return &f.n
}

func (f *FieldElement) P() *big.Int {
	return &f.p
}

// NewFieldElement takes an element and p an order
func NewFieldElement(n int64, p int64) (*FieldElement, error) {

	if n >= p || n < 0 {
		return nil, errors.New("n must be a positive integer: " + strconv.FormatInt(n, 10))
	}

	bigN := big.NewInt(n)
	bigP := big.NewInt(p)

	Fp := FieldElement{*bigN, *bigP}

	return &Fp, nil

}

// Equal is a method to compare two Field Elements and return a bool
func (f *FieldElement) Equal(e FieldElement) bool {

	if e.n.Cmp(&f.n) == 0 && e.p.Cmp(&f.p) == 0 {

		return true
	}

	return false
}

func (f *FieldElement) Inv() *FieldElement {
	inv := GCD(f.n.Int64(), f.p.Int64())
	return &FieldElement{
		n: *big.NewInt(inv),
		p: f.p,
	}
}

// NotEqual is the opposite of Equal
func (f *FieldElement) NotEqual(e FieldElement) bool {
	if f.Equal(e) {
		return false
	}

	return true
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("%s%s", f.p.Text(10), f.n.Text(10))
}

// Print will print the element in a cool way
func (f *FieldElement) Print() {
	fmt.Printf("FieldElement_%s(%s)", f.p.Text(10), f.n.Text(10))
}

// Add will add f and e and return their sum
func Add(e, f *FieldElement) (*FieldElement, error) {

	// check that e and f are over the same p
	if e.p.Cmp(&f.p) != 0 {
		return nil, errors.New("Both elements should be on the same p")
	}

	sum := big.NewInt(0)
	sum.Add(&e.n, &f.n)

	z := big.NewInt(0)
	z.Mod(sum, &e.p)

	return &FieldElement{*z, e.p}, nil

}

// Sub will substract f from e and return their dff
func Sub(e, f *FieldElement) (*FieldElement, error) {

	// check that e and f are over the same p
	if e.p.Cmp(&f.p) != 0 {
		return nil, errors.New("Both elements should be on the same p")
	}

	red := big.NewInt(0)
	red.Sub(&e.n, &f.n)

	z := big.NewInt(0)
	z.Mod(red, &e.p)

	return &FieldElement{*z, e.p}, nil

}

// Mul will multiply two field elements
func Mul(e, f *FieldElement) (*FieldElement, error) {

	// check that e and f are over the same p
	if e.p.Cmp(&f.p) != 0 {
		return nil, errors.New("Both elements should be on the same p")
	}

	mul := big.NewInt(0)

	mul.Mul(&e.n, &f.n)

	z := big.NewInt(0)

	z.Mod(mul, &e.p)

	return &FieldElement{*z, e.p}, nil
}

// Pow will return the power of a field element to a power
func Pow(e *FieldElement, power int64) (*FieldElement, error) {

	// power reduction
	//p := power % (e.p - 1)

	bPower := big.NewInt(0)
	bPower.Sub(&e.p, big.NewInt(1))

	bP := big.NewInt(0)
	bP.Mod(big.NewInt(power), bPower)
	// n**(p-1) % p
	//n := math.Pow(float64(e.n), float64(p))

	bN := big.NewInt(0)
	bN.Exp(&e.n, bP, &e.p)

	return &FieldElement{*bN, e.p}, nil

}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
