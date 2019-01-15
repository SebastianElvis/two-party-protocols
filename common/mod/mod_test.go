package mod

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModSqrtMultiple(t *testing.T) {
	w := big.NewInt(4)
	N := big.NewInt(23)
	x, y, err := ModSqrtMultiple(w, N)
	assert.NoError(t, err)
	t.Log(x.Int64(), y.Int64())
}
