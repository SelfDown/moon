package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"moon/model"
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
	//dataSourceName := "./test.sqlite"
	dataSourceName := "./test/test.sqlite"
	db, _ := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{
		CreateBatchSize: 100,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	id := "91d4abd1-14c9-4705-9de7-1518da99842711"
	name := "二级服务分类，service 一般 xxx.xx。第一级表示项目，    第二级表示具体服务."
	var o int32
	o = 22
	c1 := model.CollectDocImportant{
		CollectDocID:   &id,
		DocImportantID: id,
		Name:           &name,
		OrderIndex:     &o,
	}
	l2 := make([]model.CollectDocImportant, 1)
	l2[0] = c1
	fieldNames2 := make([]string, 0)
	fieldNames2 = append(fieldNames2, "collect_doc_id", "doc_important_id", "name", "order_index")
	db.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(fieldNames2),
	}).Create(l2)

	//key := []byte("wUNDf2LyG3NqzQpVArqmasdf") // 16字节密钥
	//iv := []byte("1234567890ABCDEF")          // 16字节IV偏移量
	//plaintext := []byte("Hello, World!")
	//ciphertext, err := encrypt(key, iv, plaintext)
	//if err != nil {
	//	fmt.Println("Encryption error:", err)
	//	return
	//}
	//tmp := hex.EncodeToString([]byte(ciphertext))
	//fmt.Println("Ciphertext:", tmp)
	//b, _ := hex.DecodeString(tmp)
	//bs := gocast.ToString(b)
	//decryptedText, err := decrypt(key, iv, bs)
	//if err != nil {
	//	fmt.Println("Decryption error:", err)
	//	return
	//}
	//fmt.Println("Decrypted Text:", string(decryptedText))
}
