package core

// import (
// 	"math/rand"
// )

// type Prober interface{ Next() int }

// type Linear struct{ k, m int }

// func (p *Linear) Init(start, mod int) { p.k, p.m = start, mod }
// func (p *Linear) Next() int           { p.k = (p.k + 1) % p.m; return p.k }

// type Random struct{ m int }

// func (p *Random) Init(mod int) { p.m = mod }
// func (p *Random) NextFrom(start int) func() int {
// 	seen := map[int]bool{}
// 	return func() int {
// 		for {
// 			j := rand.Intn(p.m)
// 			if !seen[j] {
// 				seen[j] = true
// 				return j
// 			}
// 		}
// 	}
// }

/////////////////////////////// THIS IS FOR DEMONSTRATION PURPOSES ONLY //////////////////////////////////////////

type probeFunc func(start, mod int, isFree func(int) bool) (final int, probes int)

func LinearProbe(start, mod int, isFree func(int) bool) (final int, probes int) {
	i := start
	probes = 1
	for attempts := 0; attempts < mod; attempts++ {
		if isFree(i) {
			return i, probes
		}
		i = (i + 1) % mod
		probes++
	}
	return -1, probes // table full
}
