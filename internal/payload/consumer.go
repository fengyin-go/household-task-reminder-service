package payload

type Consumer struct{ cache *Cache }

func NewConsumer(cache *Cache) *Consumer { return &Consumer{cache: cache} }

func (c *Consumer) Read(id string) []byte {
	value := c.cache.Load(id)
	copy := append([]byte(nil), value...)
	if copy == nil {
		return []byte{}
	}
	return copy
}
