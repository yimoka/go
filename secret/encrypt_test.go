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

	x, xE := Decrypt("0MJnqN3zHvxCKPSozAedr2AfBJ7XA/+JfT9hLQ==", "bL7b9BlqQPKJ3HVa")
	fmt.Printf("x %v \n", x)
	fmt.Printf("x %v \n", xE)

	if dec != str {
		t.Errorf("Decryption failed. Expected: %s, Got: %s", str, dec)
	}
}
