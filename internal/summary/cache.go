package summary

type Cache struct{ entries map[string][]string }

func NewCache() *Cache { return &Cache{entries: map[string][]string{}} }

func (c *Cache) Store(id string, tags []string) {
	c.entries[id] = append([]string(nil), tags...)
}

func (c *Cache) Count(id string) int { return len(c.entries[id]) }

func (c *Cache) First(id string) string {
	if len(c.entries[id]) == 0 {
		return ""
	}
	return c.entries[id][0]
}
