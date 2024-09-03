// The random number generator.
package machine

import (
	"math/rand"
	"time"
)

type Random interface {
	// Next returns the next random number 1 <= x <= n, for any value 1 <= n <= 32767
	Next() uint16
}

func NewPredictable(seed uint16) Random {
	if seed < 1000 {
		return &predictableRandom{last: 0, top: seed}
	}

	return &randomSeed{rand.New(rand.NewSource(int64(seed)))}
}

func NewRandom() Random {
	return &randomSeed{rand.New(rand.NewSource(time.Now().UnixMilli()))}
}

type predictableRandom struct {
	last uint16
	top  uint16
}

func (p *predictableRandom) Next() uint16 {
	next := p.last + 1
	if next > p.top {
		p.last = 1
	} else {
		p.last = next
	}
	return next
}

type randomSeed struct {
	gen *rand.Rand
}

func (r *randomSeed) Next() uint16 {
	// Intn(n) returns a random number in the range [0, n)
	return uint16(r.gen.Intn(32767) + 1)
}
