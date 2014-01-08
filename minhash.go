package main

import (
	"math"
	"math/big"
)

const bitMask = uint32(0x1)
const hashFuncNum = 128

func minHash(data []string, seed uint32) uint32 {
	m := uint32(math.MaxUint32)
	for k := range data {
		h := Murmurhash3_32([]byte(data[k]), seed)
		if h < m {
			m = h
		}
	}
	return m
}

func Signature(data []string) *big.Int {
	sigBig := big.NewInt(0)
	for i := 0; i < hashFuncNum; i++ {
		sigBig.SetBit(sigBig, i, uint(minHash(data, uint32(i))&bitMask))
	}
	return sigBig
}

func MinhashFromSignature(sig1, sig2 *big.Int) float32 {
	var dist = HammingDist(sig1, sig2)
	return 1.0 - 2.0*float32(dist)/float32(hashFuncNum)
}

func Minhash(v1, v2 []string) float32 {
	return MinhashFromSignature(Signature(v1), Signature(v2))
}
