package pwd

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

const (
	saltPassword    = "qkhPaGA13HocW3GAEWwb"
	defaultPassword = "123456"
)

func GeneratePassword(str string) (password string) {
	// md5
	m := md5.New()
	m.Write([]byte(str))
	mByte := m.Sum(nil)

	// hmac
	h := hmac.New(sha256.New, []byte(saltPassword))
	h.Write(mByte)
	password = hex.EncodeToString(h.Sum(nil))

	return
}

func ResetPassword() (password string) {
	m := md5.New()
	m.Write([]byte(defaultPassword))
	mStr := hex.EncodeToString(m.Sum(nil))

	password = GeneratePassword(mStr)

	return
}
