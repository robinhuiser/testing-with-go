package random

import (
	"math/rand"
	"testing"
)

func TestPick(t *testing.T) {
	// Trick is to fix the seed in a test
	// so it becomes repeatable
	var seed int64 = 153599603691326400
	r := rand.New(rand.NewSource(seed))
	arg := make([]int, 25)
	for i := 0; i < 25; i++ {
		arg[i] = r.Int()
	}

	got := Pick(arg)
	for _, v := range arg {
		if got == v {
			return
		}
	}
	t.Errorf("Pick(seed=%d) = %d; not in slice", seed, got)
}
