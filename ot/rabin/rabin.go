package ot

import (
	"crypto/rand"
	mathrand "math/rand"

	"math/big"

	mycrypt "github.com/SebastianElvis/two-party-protocols/common/crypto"
	"github.com/SebastianElvis/two-party-protocols/common/mod"
)

// Alice generates 2 big random primes p, q
func PrimeGen() (*big.Int, *big.Int, error) {
	// 1024 bits
	// 1024/8 = 128
	p, err := rand.Prime(nil, 1024)
	if err != nil {
		return nil, nil, err
	}
	q, err := rand.Prime(nil, 1024)
	if err != nil {
		return nil, nil, err
	}
	return p, q, nil
}

// Alice encrypts the msg by p, q
func EncMsg(msg []byte, p *big.Int, q *big.Int) []byte {
	return mycrypt.Encrypt(msg, p.String()+q.String())
}

// Bob generates random a, and w = a^2 mod N
func RandGen(N *big.Int) (*big.Int, *big.Int, error) {
	maxBytes := make([]byte, 128)
	for i, _ := range maxBytes {
		maxBytes[i] = 255
	}
	a, err := rand.Int(nil, new(big.Int).SetBytes(maxBytes))
	if err != nil {
		return nil, nil, err
	}
	w := new(big.Int).Exp(a, big.NewInt(2), N)
	return a, w, nil
}

// Alice computs the possible square root of w: {x, -x, y, -y}, chooses one randomly, cand sends it to Bob
func ChooseRandSqrtW(w, N *big.Int) (*big.Int, error) {
	x, y, err := mod.ModSqrtMultiple(w, N)
	if err != nil {
		return nil, err
	}
	arr := []*big.Int{x, big.NewInt(-x.Int64()), y, big.NewInt(-y.Int64())}
	rand := mathrand.Intn(3)
	return arr[rand], nil
}

// TODO
// Bob receives possibleSqrt, checks if possibleSqrt == a or -a
// If yes, he can factor N to p, q and decrypt
// Else, he cannot
func TryDecrypt(possibleSqrt *big.Int, a *big.Int, N *big.Int, cipher []byte) ([]byte, bool) {
	if new(big.Int).Sub(possibleSqrt, a).Int64() == 0 || new(big.Int).Add(possibleSqrt, a).Int64() == 0 {
		// cannot factor
		// wtf!!!!!!!!!!
		return []byte{}, false
	} else {
		// can factor
		// wtf!!!!!!!!!!
		return []byte{}, true
	}
}
