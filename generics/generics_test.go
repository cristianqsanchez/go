package generics

import "testing"

func TestAssertFunction(t *testing.T) {
  t.Run("asserting on integers", func(t *testing.T) {
    AssertEqual(t, 1, 1)
    AssertNotEqual(t, 1, 2)
  })

  t.Run("asserting on strings", func(t *testing.T) {
    AssertEqual(t, "hello", "hello")
    AssertNotEqual(t, "hello", "bye")
  })
}

func TestStack(t *testing.T) {
  t.Run("integer stack", func(t *testing.T) {
    myStackOfInts := new(Stack[int])

    AssertTrue(t, myStackOfInts.IsEmpty())

    myStackOfInts.Push(123)
    AssertFalse(t, myStackOfInts.IsEmpty())

    myStackOfInts.Push(456)
    value, _ := myStackOfInts.Pop()
    AssertEqual(t, value, 456)
    value, _ = myStackOfInts.Pop()
    AssertEqual(t, value, 123)
    AssertTrue(t, myStackOfInts.IsEmpty())
  })

  t.Run("string stack", func(t *testing.T) {
    stackOfStrings := new(Stack[string])

    AssertTrue(t, stackOfStrings.IsEmpty())

    stackOfStrings.Push("123")
    AssertFalse(t, stackOfStrings.IsEmpty())

    stackOfStrings.Push("456")
    value, _ := stackOfStrings.Pop()
    AssertEqual(t, value, "456")
    value, _ = stackOfStrings.Pop()
    AssertEqual(t, value, "123")
    AssertTrue(t, stackOfStrings.IsEmpty())
  })
}
