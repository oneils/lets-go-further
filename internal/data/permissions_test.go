package data

import "testing"

func TestName(t *testing.T) {
	t.Run("should include", func(t *testing.T) {
		perms := Permissions{"perm1", "perm2"}

		result := perms.Include("perm2")

		if !result {
			t.Fatal("Expected perms contains perm2")
		}
	})

	t.Run("should not include", func(t *testing.T) {
		perms := Permissions{"perm1", "perm2"}

		result := perms.Include("perm3")

		if result {
			t.Fatal("Expected perms contains perm2")
		}
	})
}
