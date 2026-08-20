package payload

import "sync"

type Pool struct{ buffers sync.Pool }

func NewPool() *Pool {
	return &Pool{buffers: sync.Pool{New: func() any { return make([]byte, 0, 64) }}}
}

func (p *Pool) Copy(value string) []byte {
	buf := p.buffers.Get().([]byte)[:0]
	buf = append(buf, value...)
	p.buffers.Put(buf[:0])
	return buf
}
