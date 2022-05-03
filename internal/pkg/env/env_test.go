package env

import "testing"

func TestIsLocal(t *testing.T) {
	t.Run("fo", func(t *testing.T) {
		if !Env("local").IsLocal() {
			t.Fatalf("foo %d", 4)
		}
	})
}
