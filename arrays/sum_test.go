package arrays

import (
	"testing"
  "slices"
)

func TestSum(t *testing.T) {
  t.Run("collection of 5 numbers", func(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    if got != want {
      t.Errorf("\033[33;1m got %d want %d, given %v \033[0m", got, want, numbers)
    }
  })
}

func TestSumAll(t *testing.T) {
  got := SumAll([]int{1, 2}, []int{0, 9})
  want := []int{3, 9}

  if !slices.Equal(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}
