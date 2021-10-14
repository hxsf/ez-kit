package ez_log

import (
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"
)

func BenchmarkLogger(b *testing.B) {
	ezLogger := NewStdCombineLogger(InfoLevel, NewRawFormatter(), &nopWriteSyncer{})
	var count = b.N
	cases := make([]string, count)
	for i := 0; i < count; i++ {
		cases[i] = RandString(20 + RandomInt(1000))
	}
	b.Run("log-no-category", func(b *testing.B) {
		for i := 0; i < count; i++ {
			ezLogger.Info(nil, String(cases[i]))
		}
		_ = ezLogger.Sync()
	})
	b.Run("log-category", func(b *testing.B) {
		cates := make([]*Category, count)
		for i := 0; i < count; i++ {
			cates[i] = &Category{
				Category:    RandString(5 + RandomInt(15)),
				SubCategory: RandString(5 + RandomInt(15)),
				Module:      RandString(5 + RandomInt(15)),
				Filter1:     RandString(5 + RandomInt(15)),
				Filter2:     RandString(5 + RandomInt(15)),
			}
		}
		for i := 0; i < count; i++ {
			ezLogger.Info(cates[i], String(cases[i]))
		}
		_ = ezLogger.Sync()
	})
}

type nopWriteSyncer struct{}

func (nop *nopWriteSyncer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (nop *nopWriteSyncer) Sync() error {
	return nil
}

func TestStdLogger(t *testing.T) {
	mf := newMemFile(" ", EOL)
	ezLogger := NewStdCombineLogger(InfoLevel, NewRawFormatter(), mf)
	var count = 1000
	cases := make([]string, count*2)
	for i := 0; i < count*2; i++ {
		cases[i] = RandString(20 + RandomInt(1000))
	}
	cates := make([]*Category, count)
	for i := 0; i < count; i++ {
		cates[i] = &Category{
			Category:    RandString(5 + RandomInt(15)),
			SubCategory: RandString(5 + RandomInt(15)),
			Module:      RandString(5 + RandomInt(15)),
			Filter1:     RandString(5 + RandomInt(15)),
			Filter2:     RandString(5 + RandomInt(15)),
		}
	}
	for i := 0; i < count; i++ {
		ezLogger.Info(cates[i], String(cases[i*2]), String(cases[i*2+1]))
	}
	_ = ezLogger.Sync()
	if count != mf.GetCount(EOL) {
		t.Fatalf("EOL count expect %d, got %d\n", count, mf.GetCount(EOL))
	}
	if count != mf.GetCount(EOL) {
		t.Fatalf("SEP count expect %d, got %d\n", count*9, mf.GetCount(" "))
	}
}

type memFile struct {
	keywords []string
	counts   map[string]int

	writeLock sync.Mutex
}

func newMemFile(keywords ...string) *memFile {
	return &memFile{keywords: keywords, counts: make(map[string]int, len(keywords))}
}

func (m *memFile) Write(p []byte) (n int, err error) {
	m.writeLock.Lock()
	for _, keyword := range m.keywords {
		m.counts[keyword] += strings.Count(string(p), keyword)
	}
	m.writeLock.Unlock()
	return len(p), nil
}

func (m *memFile) GetCount(keyword string) int {
	return m.counts[keyword]
}

func (m *memFile) Sync() error {
	return nil
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// [0, max) random
func RandomInt(max int) int {
	return r.Intn(max)
}
