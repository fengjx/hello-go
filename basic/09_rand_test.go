package basic

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func randomInt64(min int64, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min+1) + min
}

func TestMathRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(randomInt64(1, 10))
	}
}

func TestUrandom(t *testing.T) {
	for i := 0; i < 4; i++ {
		n, _ := crand.Int(crand.Reader, big.NewInt(1000))
		t.Log(n.Int64())
	}
}
