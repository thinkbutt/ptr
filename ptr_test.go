package ptr

import (
	"testing"
)

func TestDeref(t *testing.T) {
	t.Run("ValueIsCorrect", func(t *testing.T) {
		i := -5
		if  Deref(&i) != i {
			t.Fatalf("invalid deref(%d) %T", i, i)
		}
	})
}
