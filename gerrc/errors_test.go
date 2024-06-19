package gerrc

import (
	"errors"
	"fmt"
	"testing"

	errorsmod "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"
)

func TestBasics(t *testing.T) {
	threeArbitraryErrors := []error{
		ErrCancelled,
		ErrUnknown,
		ErrInvalidArgument,
	}

	t.Run("stdlib wrapping", func(t *testing.T) {
		require.True(t, errors.Is(fmt.Errorf("foo: %w", fmt.Errorf("bar: %w", threeArbitraryErrors[0])), threeArbitraryErrors[0]))
		require.True(t, errorsmod.IsOf(fmt.Errorf("foo: %w", fmt.Errorf("bar: %w", threeArbitraryErrors[0])), threeArbitraryErrors[0]))
	})
	t.Run("sdk wrapping", func(t *testing.T) {
		require.True(t, errors.Is(errorsmod.Wrap(errorsmod.Wrap(threeArbitraryErrors[0], "foo"), "bar"), threeArbitraryErrors[0]))
		require.True(t, errorsmod.IsOf(errorsmod.Wrap(errorsmod.Wrap(threeArbitraryErrors[0], "foo"), "bar"), threeArbitraryErrors[0]))
	})
	t.Run("stdlib joining", func(t *testing.T) {
		err0 := threeArbitraryErrors[0]
		err1 := fmt.Errorf("foo %w %w", threeArbitraryErrors[1], err0)
		err2 := fmt.Errorf("bar %w %w", threeArbitraryErrors[2], err1)

		require.True(t, errors.Is(err2, threeArbitraryErrors[0]))
		require.True(t, errors.Is(err2, threeArbitraryErrors[1]))
		require.True(t, errors.Is(err2, threeArbitraryErrors[2]))

		require.True(t, errorsmod.IsOf(err2, threeArbitraryErrors[0]))
		require.True(t, errorsmod.IsOf(err2, threeArbitraryErrors[1]))
		require.True(t, errorsmod.IsOf(err2, threeArbitraryErrors[2]))
	})
	t.Run("sdk joining", func(t *testing.T) {
		err0 := threeArbitraryErrors[0]
		err1 := errorsmod.Wrap(errors.Join(err0, threeArbitraryErrors[1]), "foo")
		err2 := errorsmod.Wrap(errors.Join(err1, threeArbitraryErrors[2]), "bar")

		require.True(t, errors.Is(err2, threeArbitraryErrors[0]))
		require.True(t, errors.Is(err2, threeArbitraryErrors[1]))
		require.True(t, errors.Is(err2, threeArbitraryErrors[2]))

		require.True(t, errorsmod.IsOf(err2, threeArbitraryErrors[0]))
		require.True(t, errorsmod.IsOf(err2, threeArbitraryErrors[1]))
		require.True(t, errorsmod.IsOf(err2, threeArbitraryErrors[2]))
	})
}
