package co

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCO(t *testing.T) {
	Init()
	a, err := Gen()
	assert.NoError(t, err)
	b, err := Gen()
	assert.NoError(t, err)
	t.Log("a", a, "b", b)

	A, err := A(a)
	assert.NoError(t, err)
	B, err := B(false, b, A)
	assert.NoError(t, err)
	t.Log("A", A, "B", B)

	e0, e1, err := E(a, A, B, []byte("fuck"), []byte("you"))
	assert.NoError(t, err)

	msg, err := D(e0, e1, false, A, b)
	assert.NoError(t, err)

	t.Log(msg)
}
