package grains
import "fmt"

func Square(n int) (uint64, error) {
  if n > 64 || n < 1 {
 		return 0, fmt.Errorf("Input[%d] invalid. Input should be between 1 and 64 (inclusive)", n)
  }
  return 1 << (uint16(n) - 1), nil
}

func Total() uint64 {
  return (1 << 64) - 1
}
