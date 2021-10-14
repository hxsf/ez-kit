package buffer_pool

import (
	"github.com/hxsf/ez-kit/ez-buffer"
)

var (
	_pool = ez_buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = _pool.Get
)
