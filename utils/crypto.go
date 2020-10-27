package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

const salt string = "*$salt@*"

// MD5 encrypt
func MD5(text string) string {
	hash := md5.New()
	text = salt + text + salt
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// SHA256 encrypt
func SHA256(text string) string {
	hash := sha256.New()
	text = salt + text + salt
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// GetRandomString
func GetRandomString(l int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytesGenerate := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytesGenerate[r.Intn(len(bytesGenerate))])
	}
	return result
}

// PKCS5Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt aes encrypt
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt aes decrypt
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
