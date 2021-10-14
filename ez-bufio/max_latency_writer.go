package ez_bufio

import (
	"io"
	"sync"
	"time"
)

type maxLatencyWriter struct {
	dst     writeFlusher
	latency time.Duration

	mu           sync.Mutex // protects Write + Flush
	t            *time.Timer
	flushPending bool
}

func (m *maxLatencyWriter) Flush() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.flushPending {
		return nil
	}
	m.flushPending = false
	if m.t != nil {
		m.t.Stop()
		m.t = nil
	}
	return m.dst.Flush()
}

func (m *maxLatencyWriter) Write(p []byte) (n int, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	n, err = m.dst.Write(p)
	if m.latency < 0 {
		m.dst.Flush()
		return
	}
	if m.flushPending {
		return
	}
	if m.t == nil {
		m.t = time.AfterFunc(m.latency, m.delayedFlush)
	} else {
		m.t.Reset(m.latency)
	}
	m.flushPending = true
	return
}
func (m *maxLatencyWriter) delayedFlush() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.flushPending { // if stop was called but AfterFunc already started this goroutine
		return
	}
	m.dst.Flush()
	m.flushPending = false
}

func (m *maxLatencyWriter) stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.flushPending = false
	if m.t != nil {
		m.t.Stop()
		m.dst.Flush()
	}
}
func (m *maxLatencyWriter) Stop() { m.stop() }
func (m *maxLatencyWriter) Start() {
	m.flushPending = true
	m.t = time.AfterFunc(m.latency, m.delayedFlush)
}

type writeFlusher interface {
	io.Writer
	Flush() error
}

var defaultLatency = time.Microsecond * 5

func NewMaxLatencyWriter(wf writeFlusher, latency time.Duration) MaxLatencyWriter {
	if latency == 0 {
		latency = defaultLatency
	}
	return &maxLatencyWriter{
		dst:     wf,
		latency: latency,
	}
}

type MaxLatencyWriter interface {
	io.Writer
	Flush() error
	Stop()
}
