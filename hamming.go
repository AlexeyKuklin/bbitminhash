package main

import (
	"math/big"
)

func popCount(bits uint32) uint32 {
	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

func bigPopCount(bits *big.Int) uint32 {
	result := uint32(0)
	for _, v := range bits.Bytes() {
		result += popCount(uint32(v))
	}
	return result
}

func HammingDist(a, b *big.Int) uint32 {
	var t = big.NewInt(0)
	t.Xor(a, b)
	return bigPopCount(t)
}
