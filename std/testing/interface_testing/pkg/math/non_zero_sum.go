package math

import (
	"errors"
	"fmt"
)

func (m *Math) NonZeroSum(a, b int) (int, error) {
	fmt.Println("Real dependency")
	total := a + b
	if total == 0 {
		return 0, errors.New("zero sum")
	}

	return total, nil
}
