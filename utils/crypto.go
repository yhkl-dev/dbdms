package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
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
