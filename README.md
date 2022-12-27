# goaes
an implementation of AES algorithm using golang.

In my implementation, padding uses the [PKCS#7](https://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS#5_and_PKCS#7), and mode of operation uses the [CBC](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Cipher_block_chaining_(CBC)). You can replace these with yours.

## Usage

If you don't have a aes key, you can create one as follows:

```go
// create a 256-bit key
key, _ := NewAesKey(AES256) 
// you can also create 128-bit, 192-bit
...
```

Encrypt:

```go
plainText := "hello, world"
cipherText, err := Encrypt(key, []byte(plainText))
```

Decrypt:

```go
plainText, err := Decrypt(key, cipherText)
```



You can also copy code in `aes.go` to your file *as it will not import new dependency*.
