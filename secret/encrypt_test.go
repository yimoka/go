package secret

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/forgoer/openssl"
)

func TestIrreversibleEncrypt(t *testing.T) {
	str := "12311231"
	secret := "12345678"
	nonce := "gEtryqKc"
	enc := irreversible(str, nonce, secret)
	fmt.Printf("enc %+v \n", enc)

	b := VerifyIrreversible(enc, str, nonce, secret)
	fmt.Printf("b %+v \n", b)
}
func TestEncrypt(t *testing.T) {
	str := "12311231"
	secret := "12345678"
	enc, err := Encrypt(str, secret)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}

	t.Log(enc)

	// Verify the encryption
	decoded, err := base64.StdEncoding.DecodeString(enc[16:])
	if err != nil {
		t.Errorf("Error decoding: %v", err)
	}

	decrypted, err := openssl.AesCBCDecrypt(decoded, []byte(secret), []byte(enc[:16]), openssl.PKCS7_PADDING)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}

	if string(decrypted) != str {
		t.Errorf("Decryption failed. Expected: %s, Got: %s", str, string(decrypted))
	}
}

func TestDecrypt(t *testing.T) {
	str := "12311231"
	secret := "12345678"
	enc, err := Encrypt(str, secret)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}

	dec, err := Decrypt(enc, secret)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}

	x, xE := Decrypt("yrjjpmNnBJuzTTNQH2yQ+uOfY2hcaSVJVtWirmRYlkLJbZWYb5baCvQsLn+XzhNcly10qf1IqRIDbmzviE/qjWcIaabAWxwZgUOquODfTobZMl0CG+vEUvEMryY0XRybNOB2ATElLt4tMp97cuV/7/LZxtNKbJW9tl809CYaZbt2UkyoXavS9bnsctVHPu5GgUltqJR+VcNopX8v1OVGtKqyA+BYR3lG3cLHWJgjIoMMskICjIDE7UaiwyoPhEv3reXtWO7blK2VY98P", "XulybtoHUqGPEsxsHfpanXizo1w7Lqwd")
	fmt.Printf("x %v \n", x)
	fmt.Printf("x %v \n", xE)

	if dec != str {
		t.Errorf("Decryption failed. Expected: %s, Got: %s", str, dec)
	}
}
func TestECBEncrypt(t *testing.T) {
	str := "12311231"
	secret := "12345678"
	enc, err := ECBEncrypt(str, secret)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}

	t.Log(enc)

	// Verify the encryption
	decoded, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		t.Errorf("Error decoding: %v", err)
	}

	decrypted, err := openssl.AesECBDecrypt(decoded, []byte(secret), openssl.PKCS7_PADDING)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}

	if string(decrypted) != str {
		t.Errorf("Decryption failed. Expected: %s, Got: %s", str, string(decrypted))
	}
}
func TestECBDecrypt(t *testing.T) {
	str := "12311231"
	secret := "12345678"
	enc, err := ECBEncrypt(str, secret)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}

	dec, err := ECBDecrypt(enc, secret)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}

	if dec != str {
		t.Errorf("Decryption failed. Expected: %s, Got: %s", str, dec)
	}
}
