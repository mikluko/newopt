package newopt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var testErr = errors.New("the parson had a dog")

func TestNew(t *testing.T) {
	type T struct {
		A int
	}
	WithA := func(a int) Option[T] {
		return func(t T) T {
			t.A = a
			return t
		}
	}
	v := New[T](WithA(42))
	require.Equal(t, 42, v.A)
}

func TestNewE(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		type T struct {
			A int
		}
		WithA := func(a int) OptionE[T] {
			return func(t T) (T, error) {
				t.A = a
				return t, nil
			}
		}
		v, err := NewE[T](WithA(42))
		require.NoError(t, err)
		require.Equal(t, 42, v.A)
	})
	t.Run("error", func(t *testing.T) {
		type T struct {
			A int
		}
		WithA := func(a int) OptionE[T] {
			return func(t T) (T, error) {
				return t, testErr
			}
		}
		_, err := NewE[T](WithA(42))
		require.Error(t, err)
		require.Equal(t, testErr, err)
	})
}

func TestNewP(t *testing.T) {
	type T struct {
		A int
	}
	WithA := func(a int) Option[*T] {
		return func(t *T) *T {
			t.A = a
			return t
		}
	}
	v := NewP[T](WithA(42))
	require.Equal(t, 42, v.A)
}

func TestNewPE(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		type T struct {
			A int
		}
		WithA := func(a int) OptionE[*T] {
			return func(t *T) (*T, error) {
				t.A = a
				return t, nil
			}
		}
		v, err := NewPE[T](WithA(42))
		require.NoError(t, err)
		require.Equal(t, 42, v.A)
	})
	t.Run("error", func(t *testing.T) {
		type T struct {
			A int
		}
		WithA := func(a int) OptionE[*T] {
			return func(t *T) (*T, error) {
				return nil, testErr
			}
		}
		v, err := NewPE[T](WithA(42))
		require.Error(t, err)
		require.Equal(t, testErr, err)
		require.Nil(t, v)
	})
}
