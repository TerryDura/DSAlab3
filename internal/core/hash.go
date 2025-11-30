package core

import (
	"fmt"
	"hash/fnv"
)

// asciiCat exactly as given in the PDF
func asciiCat(key string, i, j int) int64 {
	if i < 1 || j > 16 || i > j {
		return 0
	}
	s := ""
	for p := i - 1; p <= j-1; p++ {
		s += fmt.Sprintf("%d", int(key[p]))
	}
	var x int64
	fmt.Sscan(s, &x)
	return x
}

// === MY HASH (given; intentionally poor) ===
func MyHash(key string, mod int) int {
	a := asciiCat(key, 1, 1) + asciiCat(key, 5, 5)
	b := asciiCat(key, 3, 4)
	c := asciiCat(key, 5, 6)
	h := float64(a)/517.0 + float64(b)/217.0 + float64(c)/256.0
	if h < 0 {
		h = -h
	}
	return int(h) % mod
}

// === YOUR HASH (best choice – FNV-1a) ===
func YourHash(key string, mod int) int {
	h := fnv.New32a()
	h.Write([]byte(key)) // key is exactly 16 bytes
	return int(h.Sum32() % uint32(mod))
}
