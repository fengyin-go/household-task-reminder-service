package payload

type Cache struct{ values map[string][]byte }

func NewCache() *Cache { return &Cache{values: map[string][]byte{}} }

func (c *Cache) Store(id string, body []byte) {
	c.values[id] = append([]byte(nil), body...)
}

func (c *Cache) Load(id string) string {
	return string(append([]byte(nil), c.values[id]...))
}
