package auth

import (
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {

	password := "test_password123"
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password for testing: %v", err)
	}

	t.Run("successfull comparison", func(t *testing.T) {
		match := ComparePasswords(string(hashed_password), []byte(password))
		if !match {
			t.Errorf("Expected passwords to match, but they did not")
		}
	})

	t.Run("failed comparison", func(t *testing.T) {
		wrong_password := "thisiswrong"

		match := ComparePasswords(string(hashed_password), []byte(wrong_password))
		if match {
			t.Errorf("Expected passwords to not match, but they did")
		}
	})

	t.Run("bad hash input", func(t *testing.T) {
		bad_hash := "randomstringfortest"

		match := ComparePasswords(bad_hash, []byte(password))

		if match {
			t.Errorf("Expected bad hash to result in a failed match, but it succeeded")
		}
	})
}
