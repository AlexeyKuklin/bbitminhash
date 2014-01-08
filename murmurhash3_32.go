package main

import (
	"bytes"
	"encoding/binary"
)

func Murmurhash3_32(key []byte, seed uint32) uint32 {
	length := len(key)
	if length == 0 {
		return 0
	}
	var h = seed
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	nblocks := length / 4
	var k uint32
	buf := bytes.NewBuffer(key)
	for i := 0; i < nblocks; i++ {
		binary.Read(buf, binary.LittleEndian, &k)
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
		h = (h << 13) | (h >> (32 - 13))
		h = (h * 5) + 0xe6546b64
	}

	k = 0
	tail := nblocks * 4
	switch length & 3 {
	case 3:
		k ^= uint32(key[tail+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[tail+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[tail])
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
	}

	h ^= uint32(length)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
