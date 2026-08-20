package payload

type Consumer struct{ cache *Cache }

func NewConsumer(cache *Cache) *Consumer { return &Consumer{cache: cache} }

func (c *Consumer) Read(id string) []byte { return []byte(c.cache.Load(id)) }
