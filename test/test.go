package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/demdxx/gocast"
)

func encrypt(key, iv, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	paddedPlaintext := pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
func decrypt(key, iv []byte, ciphertext string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	decryptedData := make([]byte, len(decodedCiphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decryptedData, decodedCiphertext)
	return pkcs7Unpadding(decryptedData), nil
}

// 使用PKCS7填充方式对数据进行填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 对使用PKCS7填充方式的数据进行去填充
func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
func main() {
	key := []byte("wUNDf2LyG3NqzQpVArqmasdf") // 16字节密钥
	iv := []byte("1234567890ABCDEF")          // 16字节IV偏移量
	plaintext := []byte("Hello, World!")
	ciphertext, err := encrypt(key, iv, plaintext)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	tmp := hex.EncodeToString([]byte(ciphertext))
	fmt.Println("Ciphertext:", tmp)
	b, _ := hex.DecodeString(tmp)
	bs := gocast.ToString(b)
	decryptedText, err := decrypt(key, iv, bs)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted Text:", string(decryptedText))
}
