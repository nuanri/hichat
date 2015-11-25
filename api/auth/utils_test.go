package auth

import "testing"

func Test_Encrypt(t *testing.T) {
	clear_pass := "123456"
	en_pass := EncryptPassword(clear_pass)
	if !CheckPassword(en_pass, clear_pass) {
		t.Error("not match")
	}
}
