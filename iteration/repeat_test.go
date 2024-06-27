package iteration

import (
	"fmt"
	"strings"
	"testing"
)
                      
func TestRepeat(t *testing.T) {
  t.Run("repeat character 5 times by default", func(t *testing.T) {
    repeated := Repeat("a", 0)
    expected := strings.Repeat("a", 5)
    assertCorrectMessage(t, repeated, expected)
  })

  t.Run("repeat character specific times", func(t *testing.T) {
    repeated := Repeat("z", 8)
    expected := strings.Repeat("z", 8)
    assertCorrectMessage(t, repeated, expected)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("expected %q but got %q", got, want)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 0)
  }
}

func ExampleRepeat() {
  repeated := Repeat("x", 0)
  fmt.Println(repeated)
  // Output: xxxxx
}
