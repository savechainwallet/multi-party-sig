package hash

import (
	"math/big"
	"testing"

	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
)

func TestHash_WriteAny(t *testing.T) {
	var err error

	a := func(v interface{}) {
		h := New(nil)
		err = h.WriteAny(v)
		if err != nil {
			t.Error(err)
		}
	}
	b := func(vs ...interface{}) {
		h := New(nil)
		for _, v := range vs {
			err = h.WriteAny(v)
			if err != nil {
				t.Error(err)
			}
		}
	}

	a(big.NewInt(35))
	a(curve.NewIdentityPoint().ScalarBaseMult(curve.NewScalarRandom()))
	a([]byte{1, 4, 6})

	b(big.NewInt(35), []byte{1, 4, 6})
}
