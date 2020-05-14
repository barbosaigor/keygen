package keygen

import "testing"

func TestBase64Symbols(t *testing.T) {
	t.Log(string(Base64Symbols))
	if len(Base64Symbols) != 64 {
		t.Errorf("Base64Symbols: Invalid length, expected %d but got %d", 64, len(Base64Symbols))
	}
}
