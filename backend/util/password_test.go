package util

import "testing"

func Test_generateSalt(t *testing.T) {
	t.Run("salt is longer than 10 chars", func(t *testing.T) {
		salt := GenerateSalt()
		if len(salt) < 10 {
			t.Error("Salt is too short")
		}
	})
}

func Test_passwordHashing(t *testing.T) {
	t.Run("doesn't accept empty password", func(t *testing.T) {
		if _, err := HashPassword(""); err == nil {
			t.Error("Didn't return err on empty password")
		}
	})
}
