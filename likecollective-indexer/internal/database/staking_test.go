package database

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestPoolSharePercentage(t *testing.T) {
	tenLike := uint256.MustFromDecimal("10000000000000000000")
	twentyLike := uint256.MustFromDecimal("20000000000000000000")

	cases := []struct {
		name   string
		staked *uint256.Int
		total  *uint256.Int
		want   string
	}{
		{"empty pool", uint256.NewInt(0), uint256.NewInt(0), "0"},
		{"small values", uint256.NewInt(1), uint256.NewInt(3), "0.33"},
		// 1e19 and 2e19 both exceed math.MaxInt64
		{"values above int64", tenLike, twentyLike, "0.5"},
		{"whole pool above int64", twentyLike, twentyLike, "1"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := PoolSharePercentage(c.staked, c.total); got != c.want {
				t.Errorf("PoolSharePercentage(%s, %s) = %s, want %s", c.staked, c.total, got, c.want)
			}
		})
	}
}
