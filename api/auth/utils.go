package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func GeneratePwd(salt string, pwd string) string {

	data := []byte(salt + pwd)
	Spwd := sha256.Sum256(data)
	spwd := string(Spwd[:])

	return fmt.Sprintf("%x", spwd)
}

func GenerateSalt() string {

	buf := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	salt := sha256.Sum256(buf)
	x := fmt.Sprintf("%x", salt)
	return x[:12]
}

func EncryptPassword(clear string) string {
	salt := GenerateSalt()
	return salt + "$" + GeneratePwd(salt, clear)
}

func CheckPassword(en_pass string, clear_pass string) bool {
	fmt.Println("en_pass = ", en_pass)
	fmt.Println("clear_pass = ", clear_pass)
	L := strings.Split(en_pass, "$")
	salt := L[0]
	en := L[1]
	return GeneratePwd(salt, clear_pass) == en
}
