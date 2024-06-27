package iteration

const repeatCount = 5

func Repeat(char string, times int) string {
  var repeated string

  if times <= 0 {
    times = repeatCount
  }

  for i := 0; i < times; i++ {
    repeated += char
  }

  return repeated
}
