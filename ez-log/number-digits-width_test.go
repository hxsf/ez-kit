package ez_log

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_uintLength(t *testing.T) {
	count := 100000
	testsUint := make([]uint, count)
	answers := make([]int, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		n := r.Uint64()
		testsUint[i] = uint(n)
		answers[i] = len(strconv.FormatUint(n, 10))
	}
	t.Run("uint", func(t *testing.T) {
		for i, tt := range testsUint {
			if !assert.Equal(t, answers[i], uintLength(tt)) {
				t.FailNow()
			}
		}
	})
	testsInt := make([]int, count)
	for i := 0; i < count; i++ {
		n := r.Int63()
		testsInt[i] = int(n)
		answers[i] = len(strconv.FormatInt(n, 10))
	}
	t.Run("int", func(t *testing.T) {
		for i, tt := range testsInt {
			if !assert.Equal(t, answers[i], intLength(tt)) {
				t.FailNow()
			}
		}
	})
}
