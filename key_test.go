package keygen

import "testing"

var symbols = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

func newKey() *Key {
	return NewWithCustomSymbols(symbols, 5)
}

func TestNew(t *testing.T) {
	key := newKey()
	if len(key.Symbols) != len(symbols) {
		t.Errorf("New has an incorrect length of symbols, expected %v but got %v", len(symbols), key.Current())
	}
	if cap(key.Data) != 5 {
		t.Errorf("New has an incorrect capacity of data, expected %v but got %v", 5, cap(key.Data))
	}
}

func TestCurrent(t *testing.T) {
	key := newKey()
	expectedKey := "aaaaa"
	if key.Current() != expectedKey {
		t.Errorf("Current fail to generate the correct key, expected %v but got %v", expectedKey, key.Current())
	}
}

func TestNext(t *testing.T) {
	key := newKey()
	key.Next()
	expectedKey := "aaaab"
	if key.Current() != expectedKey {
		t.Errorf("Next fail to generate the correct key, expected %v but got %v", expectedKey, key.Current())
	}
	for i := 0; i < 5; i++ {
		key.Next()
	}
	expectedKey = "aaaag"
	if key.Current() != expectedKey {
		t.Errorf("Next fail to generate the correct key, expected %v but got %v", expectedKey, key.Current())
	}
	key.Next()
	expectedKey = "aaaba"
	if key.Current() != expectedKey {
		t.Errorf("Next fail to generate the correct key, expected %v but got %v", expectedKey, key.Current())
	}
	for i := 0; i < 3000; i++ {
		key.Next()
	}
	expectedKey = "bbfce"
	if key.Current() != expectedKey {
		t.Errorf("Next fail to generate the correct key, expected %v but got %v", expectedKey, key.Current())
	}
}

func TestKeys(t *testing.T) {
	key := newKey()
	ks := key.Keys(3)
	if len(ks) != 3 {
		t.Errorf("Keys invalid length, expected %v but got %v", 3, len(ks))
	}
	expectedKey := "aaaaa"
	if ks[0] != expectedKey {
		t.Errorf("Keys first element incorrect key, expected %v but got %v", expectedKey, key.Current())
	}
	expectedKey = string(expectedKey[:len(expectedKey)-1]) + "b"
	if ks[1] != expectedKey {
		t.Errorf("Keys second element incorrect key, expected %v but got %v", expectedKey, key.Current())
	}
	expectedKey = string(expectedKey[:len(expectedKey)-1]) + "c"
	if ks[2] != expectedKey {
		t.Errorf("Keys second element incorrect key, expected %v but got %v", expectedKey, key.Current())
	}
}
