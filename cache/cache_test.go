package cache

import "testing"

func TestNew(t *testing.T) {
	capacity := 1000
	c := New(nil, 5, capacity)
	if c.Length != 0 {
		t.Errorf("New: cache has an invalid length, expected %d but got %d", 0, c.Length)
	}
	if c.key == nil {
		t.Errorf("New: cache has an nil key, expected an key allocation")
	}
	if cap(c.Keys) != capacity {
		t.Errorf("New: cache.Keys has an invalid capacity, expected %d but got %d", capacity, cap(c.Keys))
	}
}

func TestFill(t *testing.T) {
	capacity := 1000
	c := New(nil, 5, capacity)
	c.Fill()
	if c.Length != capacity {
		t.Errorf("Fill: cache has an invalid length, expected %d but got %d", capacity, c.Length)
	}
	for i, k := range c.Keys {
		if k == "" {
			t.Errorf("Fill: got an empty key at %d position", i)
		}
	}
	// Just a sample
	if c.Keys[0] == c.Keys[1] {
		t.Errorf("Fill: the keys should be unique")
	}
}

func TestKey(t *testing.T) {
	c := New(nil, 5, 1000)
	c.Fill()
	k := c.Key()
	if k != "aaapO" {
		t.Errorf("Key: invalid key, expected %s but got %s", "aaapO", k)
	}
	k = c.Key()
	if k != "aaapN" {
		t.Errorf("Key: invalid key, expected %s but got %s", "aaapN", k)
	}
	k = c.Key()
	if k != "aaapM" {
		t.Errorf("Key: invalid key, expected %s but got %s", "aaapM", k)
	}
	c.Fill()
	if c.Length != 1000 {
		t.Errorf("Key: invalid length, expected %d but got %d", 1000, c.Length)
	}
	prvLen := c.Length
	for i := c.Length / 2; i < prvLen; i++ {
		c.Key()
	}
	if c.Length != 500 {
		t.Errorf("Key: invalid length, expected %d but got %d", 500, c.Length)
	}
	prvLen = c.Length
	for i := 0; i < prvLen; i++ {
		c.Key()
	}
	if c.Length != 0 {
		t.Errorf("Key: invalid length, expected %d but got %d", 0, c.Length)
	}
	c.Key()
	if c.Length != 999 {
		t.Errorf("Key: invalid length, expected %d but got %d", 999, c.Length)
	}
}
