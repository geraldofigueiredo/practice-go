package math

type MathInterface interface {
	NonZeroSum(a, b int) (int, error)
}

type Math struct{}
