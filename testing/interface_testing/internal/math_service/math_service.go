package mathService

import "github.com/geraldofigueiredo/practice-go/testing/interface_testing/pkg/math"

type MathServiceInterface interface {
	CheckNonZeroSum(a, b int32) (int32, error)
}

type MathService struct {
	Math math.MathInterface
}

func (m *MathService) CheckNonZeroSum(a, b int) (int, error) {
	return m.Math.NonZeroSum(a, b)
}
