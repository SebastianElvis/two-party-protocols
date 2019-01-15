package mod

import (
	"errors"
	"math/big"

	"github.com/frrad/euler"
)

// w = a^2 mod N
// know w, N, solve a
// a= w // 2 mod N
func ModSqrtMultiple(w, N *big.Int) (*big.Int, *big.Int, error) {
	x, y, work := euler.SqrtMod(w.Int64(), N.Int64())
	if !work {
		return nil, nil, errors.New("no modsqrt")
	}
	return big.NewInt(x), big.NewInt(y), nil
}
