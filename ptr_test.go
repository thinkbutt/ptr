package ptr

import (
	"testing"
)

func TestDeref(t *testing.T) {
	t.Run("ValueIsCorrect", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			i := -5
			if Deref(&i) != i {
				t.Fatalf("invalid deref(%d) %T", i, i)
			}
		})
	})
	t.Run("EmptyValue", func(t *testing.T) {
		t.Run("String", func(t *testing.T) {
			var s *string
			val := Deref(s)
			if val != "" {
				t.Fatalf("invalid deref (%T): %v", s, val)
			}
		})

		t.Run("Int", func(t *testing.T) {
			var i int
			val := Deref(&i)
			if val != 0 {
				t.Fatalf("invalid deref (%T): %v", i, val)
			}
		})
	})
}

func TestFrom(t *testing.T) {
	t.Run("CreatesIntPointer", func(t *testing.T) {
		i := 42
		ptr := From(i)
		if ptr == nil {
			t.Fatal("From() returned nil")
		}
		if *ptr != i {
			t.Errorf("*From(%v) = %v; want %v", i, *ptr, i)
		}
	})

	t.Run("CreatesStringPointer", func(t *testing.T) {
		s := "test"
		ptr := From(s)
		if ptr == nil {
			t.Fatal("From() returned nil")
		}
		if *ptr != s {
			t.Errorf("string pointer creation failed")
		}
	})
}
