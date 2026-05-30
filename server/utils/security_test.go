package utils

import "testing"

func TestAESCrypto(t *testing.T) {
	key := []byte("07DF4DCDE2003621337EA86747159A1B")
	data := `{"uid":3,"username":"springbear"}`

	cipherText, err := AesGCMEncrypt([]byte(data), key)
	if err != nil {
		t.Errorf("AesGCMEncrypt failed" + err.Error())
		return
	}
	t.Log(cipherText)

	plainText, err := AesGCMDecrypt(cipherText, key)
	if err != nil {
		t.Errorf("AesGCMDecrypt failed" + err.Error())
	}
	t.Log(string(plainText))
}
