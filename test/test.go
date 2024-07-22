package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	source := "wghis"
	fmt.Println("原字符：", source)
	key := "wUNDf2LyG3NqzQpVArqm"
	//encryptCode := AESEncryptECBStr(source, key)
	encryptCode := "8D5A970B144B9763F27BD8E306B8655F"
	fmt.Println("密文：", encryptCode)
	decryptCode := AESDecryptECBStr(encryptCode, key)
	fmt.Println("解密：", decryptCode)
}

func AESEncryptECBStr(source string, keys string) string {
	// 字符串转换成切片
	src := []byte(source)
	key := []byte(keys)
	cipher, _ := aes.NewCipher(generateKeys(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	//return encrypted
	encryptstr := strings.ToUpper(hex.EncodeToString(encrypted))
	return encryptstr
}

func AESDecryptECBStr(encrypteds string, keys string) string {
	// 字符串转换成切片
	//encrypted := []byte(encrypteds)
	encrypted, _ := hex.DecodeString(encrypteds)
	key := []byte(keys)

	cipher, _ := aes.NewCipher(generateKeys(key))
	decrypted := make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return string(decrypted[:trim])
}

func generateKeys(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
