package ot

import (
	"crypto/rand"
	"math/big"

	mycrypt "github.com/SebastianElvis/two-party-protocols/common/crypto"
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

// TODO
// Alice computs the possible square root of w: {x, -x, y, -y}, chooses one randomly, cand sends it to Bob
func ChooseRandSqrtW(w *big.Int) *big.Int {

}

// TODO
// Bob receives possibleSqrt, checks if possibleSqrt == a or -a
// If yes, he can factor N to p, q and decrypt
// Else, he cannot
func TryDecrypt(possibleSqrt *big.Int, N *big.Int, cipher []byte) ([]byte, bool) {

}
