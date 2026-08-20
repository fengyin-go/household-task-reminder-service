package payload

type Producer struct {
	pool  *Pool
	cache *Cache
}

func NewProducer(pool *Pool, cache *Cache) *Producer { return &Producer{pool: pool, cache: cache} }

func (p *Producer) Submit(id, message string) []byte {
	body := p.pool.Copy(message)
	p.cache.Store(id, body)
	return body
}
