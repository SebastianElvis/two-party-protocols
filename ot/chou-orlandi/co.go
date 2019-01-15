package co

import (
	"crypto/rand"
	"math/big"
	mathrand "math/rand"

	mycrypt "github.com/SebastianElvis/two-party-protocols/common/crypto"

	"github.com/SebastianElvis/two-party-protocols/common/ff"
)

var (
	p *big.Int
	g *ff.FieldElement
)

func Init() {
	var err error
	p, err = rand.Prime(rand.Reader, 128)
	if err != nil {
		panic(err)
	}
	g, err = ff.NewFieldElement(3, p.Int64())
	if err != nil {
		panic(err)
	}
}

func Gen() (*ff.FieldElement, error) {
	n := mathrand.Intn(int(p.Int64()) - 1)
	elem, err := ff.NewFieldElement(int64(n), p.Int64())
	return elem, err
}

func A(a *ff.FieldElement) (*ff.FieldElement, error) {
	return ff.Pow(g, a.N().Int64())
}

func B(c bool, b, A *ff.FieldElement) (*ff.FieldElement, error) {
	var B *ff.FieldElement
	gb, err := ff.Pow(g, b.N().Int64())
	if err != nil {
		return nil, err
	}
	if c {
		B, err = ff.Mul(A, gb)
		if err != nil {
			return nil, err
		}
	} else {
		B = gb
	}
	return B, nil
}

func E(a, A, B *ff.FieldElement, m0, m1 []byte) ([]byte, []byte, error) {
	Ba, err := ff.Mul(B, a)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	k0 := mycrypt.CreateHash(Ba.String())

	BA, err := ff.Mul(B, A.Inv())
	if err != nil {
		return []byte{}, []byte{}, err
	}
	BAa, err := ff.Pow(BA, a.N().Int64())
	if err != nil {
		return []byte{}, []byte{}, err
	}
	k1 := mycrypt.CreateHash(BAa.String())

	e0 := mycrypt.Encrypt(m0, k0)
	e1 := mycrypt.Encrypt(m1, k1)
	return e0, e1, nil
}

func D(e0, e1 []byte, c bool, A, b *ff.FieldElement) (string, error) {
	var msg string
	Ab, err := ff.Pow(A, b.N().Int64())
	if err != nil {
		return "", err
	}
	k := mycrypt.CreateHash(Ab.String())
	if c {
		msg = mycrypt.Decrypt(e1, k)
	} else {
		msg = mycrypt.Decrypt(e0, k)
	}
	return msg, nil
}
