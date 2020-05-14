# Keygen

Keygen is a lightweight library, which implements a faster unique key generator.
It is possible to change the symbols used to generate the key, as well as the key size. 
This library is not recommended for security algorithms, but is well suitable for  
applications that need a faster key generator. 
For better performance there is the cache package, 
which implements an in-memory cache containing a set of unique keys.

*Installation*
```bash
go get github.com/barbosaigor/keygen
```

## Usage

_Next_ generates a new key and returns it.
The key is generated such as a sum of two numbers, though it will be using the current symbols.  
```golang
// Creates a key generator that generates a 7 length key  
// 'New' will use base 64 symbols for the key
kgen := keygen.New(7)  
// Next generates a new 7 length key  
key := kgen.Next()  
```  

_Keys_ generates a set of N unique keys.  
```golang
// Creates a key generator that generates a 13 length key  
kgen := keygen.New(13)  
keys := kgen.Keys(2000)
```  

_NewWithCustomSymbols_ creates a Key generator that is going to use a set of symbols 
to generate the keys. If the symbols is not provided then it will use base 64 characters such as 'New'.  
```golang
symbols := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
kgen := keygen.NewWithCustomSymbols(symbols, 5)
keys := key.Keys(2000)
```  

## Cache
_Key_ returns a key from the cache. 
Before returning a key, if the cache is empty then it is going to calculate new keys as well.  
```golang
symbols := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
keySize := 5
cacheCapacity := 5000
// If the symbols are not provided (nil), then cache will use base 64 symbols instead.  
kgenCache := cache.New(symbols, keySize, cacheCapacity)
key := kgenCache.Key()
```  

