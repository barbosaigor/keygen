package keygen

// Base64Symbols has a set of symbols used for base64 encoding
var Base64Symbols []rune

func init() {
	// Base64 symbols
	Base64Symbols = make([]rune, 0, 64)
	for i := 'a'; i <= 'z'; i++ {
		Base64Symbols = append(Base64Symbols, i)
	}
	for i := 'A'; i <= 'Z'; i++ {
		Base64Symbols = append(Base64Symbols, i)
	}
	for i := '0'; i <= '9'; i++ {
		Base64Symbols = append(Base64Symbols, i)
	}
	Base64Symbols = append(Base64Symbols, '+', '/')
}
