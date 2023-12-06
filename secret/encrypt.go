// Package secret encrypt.go
package secret

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"github.com/forgoer/openssl"
	"github.com/yimoka/go/utils"
)

// Encrypt 加密字符串
func Encrypt(str string, secret string) (string, error) {
	src := []byte(str)
	key := []byte(secret)
	dst, err := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}

// Decrypt 解密字符串
func Decrypt(str string, secret string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	key := []byte(secret)
	dst, err := openssl.AesECBDecrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}

// IrreversibleEncrypt 不可逆加密
func IrreversibleEncrypt(str string, secret string) (string, string) {
	nonce := utils.RandomStr(8)
	return nonce, irreversible(str, nonce, secret)
}

func irreversible(str string, nonce string, secret string) string {
	m := sha256.New()
	m.Write([]byte(nonce))
	m.Write([]byte(str))
	m.Write([]byte(secret))
	return hex.EncodeToString(m.Sum(nil))
}

// VerifyIrreversible 验证不可逆加密
func VerifyIrreversible(cipher string, str string, nonce string, secret string) bool {
	if cipher == "" || str == "" || nonce == "" || secret == "" {
		return false
	}
	return cipher == irreversible(str, nonce, secret)
}
