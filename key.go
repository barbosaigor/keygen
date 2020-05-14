package keygen

// Key implements operations for generating unique keys
type Key struct {
	// Symbols are the alphabet used to generate the keys
	Symbols []rune
	// Data is the current holding key
	Data []uint32
}

// NewWithCustomSymbols creates a Key generator that is going to use a set of symbols
// to generate the keys. If the symbols is not provided then it will use base 64 characters
func NewWithCustomSymbols(symbols []rune, size uint32) *Key {
	if symbols == nil {
		symbols = Base64Symbols
	}
	return &Key{Symbols: symbols, Data: make([]uint32, size)}
}

// New creates a Key generator that is going to use base 64 symbols
// to generate the keys
func New(size uint32) *Key {
	return &Key{Symbols: Base64Symbols, Data: make([]uint32, size)}
}

// Keys generates a slice of n unique keys
func (k *Key) Keys(n int) []string {
	keys := make([]string, n)
	keys[0] = k.Current()
	for i := 1; i < n; i++ {
		keys[i] = k.Next()
	}
	return keys
}

// Current returns the current key as a string
func (k Key) Current() string {
	key := make([]rune, cap(k.Data))
	// Take the correct symbol position for a Data element,
	// and place the correct rune into the key
	for i, p := range k.Data {
		key[i] = k.Symbols[p]
	}
	return string(key)
}

// Next generates the posterior key and returns it.
// The key is generated such as a sum of two numbers, though it will be using the current symbols.
func (k *Key) Next() string {
	for i := cap(k.Data) - 1; i >= 0; i-- {
		// If achieved the last element of the alphabet, then
		// get back to the first element.
		// And increment to the next position of the key.
		if k.Data[i] == uint32(len(k.Symbols)-1) {
			k.Data[i] = 0
		} else {
			k.Data[i]++
			break
		}
	}
	return k.Current()
}
