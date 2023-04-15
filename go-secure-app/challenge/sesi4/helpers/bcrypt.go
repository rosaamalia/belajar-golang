package helpers

import "golang.org/x/crypto/bcrypt"

// hashing password sebelum disimpan
func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

// membandingkan password yg diinput dgn yg sudah di-hash
func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}