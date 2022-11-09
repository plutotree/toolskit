package randnum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	r, err := Rand(10, 20, 20)

	assert.Nil(t, err)
	assert.Equal(t, 20, len(r))

	for _, v := range r {
		if v < 10 || v >= 20 {
			t.Errorf("%d is not in [10, 20)", v)
		}
	}
}

func TestRandPerm(t *testing.T) {
	r, err := RandPerm(10, 20, 5)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(r))
	for i := 0; i < len(r); i++ {
		assert.Equal(t, 10, len(r[i]))

		for _, v := range r[i] {
			if v < 10 || v >= 20 {
				t.Errorf("%d is not in [10, 20)", v)
			}
			fmt.Printf("%d,", v)
		}
		fmt.Printf("\n")
	}
}

func TestRandPick(t *testing.T) {
	r, err := RandPick(10, 20, 2, 5)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(r))
	for i := 0; i < len(r); i++ {
		assert.Equal(t, 2, len(r[i]))

		for _, v := range r[i] {
			if v < 10 || v >= 20 {
				t.Errorf("%d is not in [10, 20)", v)
			}
			fmt.Printf("%d,", v)
		}
		fmt.Printf("\n")
	}
}
