package cache

import "time"

type Cache struct {
	items map[string]Item
}

type Item struct {
	value    string
	deadline *time.Time
}

func NewCache() Cache {
	items := make(map[string]Item)
	return Cache{items}
}

func (c *Cache) Get(key string) (string, bool) {
	item, ok := c.items[key]
	if !ok || c.items[key].deadline != nil && c.items[key].deadline.Before(time.Now()) {
		return "", false
	}
	return item.value, ok
}

func (c *Cache) Put(key, value string) {
	c.items[key] = Item{value, nil}
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0)
	for k := range c.items {
		if c.items[k].deadline != nil && c.items[k].deadline.Before(time.Now()) {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.items[key] = Item{value, &deadline}
}
