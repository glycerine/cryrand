package cryrand

import (
	"encoding/binary"

	cr "crypto/rand"
	mr "math/rand"
)

var mathRandSrc *mr.Rand

func init() {
	seed := CryptoRandInt64()
	mathRandSrc = mr.New(mr.NewSource(seed))
}

func MathRandInt64() int64 {
	// generate one rand for the sign, xor with a 2nd.
	return (mathRandSrc.Int63() << 1) ^ mathRandSrc.Int63()
}

// Use crypto/rand to get an random int64.
func CryptoRandInt64() int64 {
	b := make([]byte, 8)
	_, err := cr.Read(b)
	if err != nil {
		panic(err)
	}
	r := int64(binary.LittleEndian.Uint64(b))
	return r
}

func CryptoRandBytes(n int) []byte {
	b := make([]byte, n)
	_, err := cr.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func CryptoRandNonNegInt(n int64) int64 {
	x := CryptoRandInt64()
	if x < 0 {
		x = -x
	}
	return x % n
}

var ch = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

func RandomString(n int) string {
	s := make([]byte, n)
	m := int64(len(ch))
	for i := 0; i < n; i++ {
		r := CryptoRandInt64()
		if r < 0 {
			r = -r
		}
		k := r % m
		a := ch[k]
		s[i] = a
	}
	return string(s)
}
