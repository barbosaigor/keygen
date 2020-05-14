package keygen

// KeyGenerator implements operations for generating unique keys
type KeyGenerator struct {
	// Symbols are the alphabet used to generate the keys
	Symbols []rune
	// data is the current holding key
	data []uint32
}

// NewWithCustomSymbols creates a Key generator that is going to use a set of symbols
// to generate the keys. If the symbols is not provided then it will use base 64 characters.
// Symbols must be unique among the slice.
func NewWithCustomSymbols(symbols []rune, size uint32) *KeyGenerator {
	if symbols == nil {
		symbols = Base64Symbols
	} else if len(symbols) == 0 {
		return nil
	} else {
		for i, prvSymbol := range symbols {
			if i == len(symbols)-1 {
				break
			}
			for _, s := range symbols[1+i:] {
				if prvSymbol == s {
					return nil
				}
			}
		}
	}
	return &KeyGenerator{Symbols: symbols, data: make([]uint32, size)}
}

// New creates a Key generator that is going to use base 64 symbols
// to generate the keys.
func New(size uint32) *KeyGenerator {
	return &KeyGenerator{Symbols: Base64Symbols, data: make([]uint32, size)}
}

// Keys generates a slice of n unique keys
func (k *KeyGenerator) Keys(n int) []string {
	keys := make([]string, n)
	keys[0] = k.Current()
	for i := 1; i < n; i++ {
		keys[i] = k.Next()
	}
	return keys
}

// Current returns the current key as a string
func (k KeyGenerator) Current() string {
	key := make([]rune, cap(k.data))
	// Take the correct symbol position for a Data element,
	// and place the correct rune into the key
	for i, p := range k.data {
		key[i] = k.Symbols[p]
	}
	return string(key)
}

// Next generates the posterior key and returns it.
// The key is generated such as a sum of two numbers, though it will be using the current symbols.
func (k *KeyGenerator) Next() string {
	for i := cap(k.data) - 1; i >= 0; i-- {
		// If achieved the last element of the alphabet, then
		// get back to the first element.
		// And increment to the next position of the key.
		if k.data[i] == uint32(len(k.Symbols)-1) {
			k.data[i] = 0
		} else {
			k.data[i]++
			break
		}
	}
	return k.Current()
}
