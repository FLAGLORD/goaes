# goaes
an implementation of AES algorithm using golang.

In my implementation, padding uses the [PKCS#7](https://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS#5_and_PKCS#7), and mode of operation uses the [CBC](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Cipher_block_chaining_(CBC)). You can replace these with yours.



中文用户可以在[我的博客](https://flaglord.com/2022/12/26/%E4%BD%BF%E7%94%A8-Golang-%E5%AE%9E%E7%8E%B0-AES-%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95/)中了解实现思路。

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

## Reference
[What is block cipher](https://www.techtarget.com/searchsecurity/definition/block-cipher)
