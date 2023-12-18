package secret

import (
	"fmt"
	"testing"
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
