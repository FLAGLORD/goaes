package goaes

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAes(t *testing.T) {
	Convey("AES Test", t, func() {
		key, _ := NewAesKey(AES256)
		Convey("PKCS#7Padding Test", func() {
			text := "hello"
			expected := append([]byte(text), bytes.Repeat([]byte{byte(11)}, 11)...)
			paddedText, err := PKCS7Pad([]byte(text), 16)
			if err != nil {
				t.Error(err)
				return
			}
			So(paddedText, ShouldResemble, expected)

			unpadText, err := PKCS7UnPad(paddedText, 16)
			if err != nil {
				t.Error(err)
				return
			}
			So(unpadText, ShouldResemble, []byte(text))
		})
		Convey("Simple String", func() {
			plainText := "hello, world"
			cipherText, err := Encrypt(key, []byte(plainText))
			if err != nil {
				t.Error(err)
				return
			}
			actual, err := Decrypt(key, cipherText)
			if err != nil {
				t.Error(err)
				return
			}
			So(string(actual), ShouldEqual, plainText)
		})
		Convey("Long string that exceeds 16 bytes(128 bits)", func() {
			plainText := `
Encrypt-then-MAC:

Provides integrity of Ciphertext. Assuming the MAC shared secret has not been compromised, 
we ought to be able to deduce whether a given ciphertext is indeed authentic or has been forged; 
for example, in public-key cryptography anyone can send you messages. EtM ensures you only read 
valid messages.`
			cipherText, err := Encrypt(key, []byte(plainText))
			if err != nil {
				t.Error(err)
				return
			}
			actual, err := Decrypt(key, cipherText)
			if err != nil {
				t.Error(err)
				return
			}
			So(string(actual), ShouldEqual, plainText)
		})
	})
}
