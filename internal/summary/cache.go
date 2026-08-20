package summary

type Cache struct{ entries map[string][]string }

func NewCache() *Cache { return &Cache{entries: map[string][]string{}} }

// Store 写入一份独立快照，避免调用方后续原地修改源切片时把旧统计串改掉。
func (c *Cache) Store(id string, tags []string) {
	snapshot := make([]string, len(tags))
	copy(snapshot, tags)
	c.entries[id] = snapshot
}

func (c *Cache) Count(id string) int { return len(c.entries[id]) }

func (c *Cache) First(id string) string {
	if len(c.entries[id]) == 0 {
		return ""
	}
	return c.entries[id][0]
}
