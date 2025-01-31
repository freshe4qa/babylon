package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/babylonchain/babylon/x/epoching/simulation"
)

func TestParamChanges(t *testing.T) {
	t.Skip("TODO: support changing epoch interval on-the-fly")

	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"epoching/EpochInterval", "EpochInterval", "82", "epoching"},
	}

	paramChanges := simulation.ParamChanges(r)

	require.Len(t, paramChanges, 1)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
