package cache

import "github.com/barbosaigor/keygen"

// Cache represents an in-memory cache containing unique keys.
// For the cache size is recommended to set up a large enough size, to avoid
// calculating the keys every time.
type Cache struct {
	key *keygen.KeyGenerator
	// Keys has a in-memory set of keys.
	Keys []string
	// Length represent the current number of elements.
	Length int
}

// New creates a Cache using a custom set of symbols or
// base64 if not provided.
func New(symbols []rune, keySize uint32, cacheCap int) *Cache {
	return &Cache{key: keygen.NewWithCustomSymbols(symbols, keySize), Keys: make([]string, cacheCap), Length: 0}
}

// Fill inserts new keys up until the max capacity of Keys.
func (c *Cache) Fill() {
	for i := c.Length; i < cap(c.Keys); i++ {
		c.Keys[i] = c.key.Next()
	}
	c.Length = cap(c.Keys)
}

// Key returns a key from the cache.
// Before returning a key, if the cache is empty then it is going to calculate new keys as well.
func (c *Cache) Key() string {
	if c.Length == 0 {
		c.Fill()
	}
	c.Length--
	return c.Keys[c.Length]
}
