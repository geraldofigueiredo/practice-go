package mocks

type MockMathService struct {
	NonZeroSumFn func(a, b int32) (int32, error)
}

func (m *MockMathService) NonZeroSum(a, b int32) (int32, error) {
	return m.NonZeroSumFn(a, b)
}
