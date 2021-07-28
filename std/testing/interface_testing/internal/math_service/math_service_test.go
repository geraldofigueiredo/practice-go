package mathService

import (
	"errors"
	"testing"

	"github.com/geraldofigueiredo/practice-go/testing/interface_testing/pkg/math"
	"github.com/geraldofigueiredo/practice-go/testing/interface_testing/pkg/math/mock"
	"github.com/golang/mock/gomock"
)

func TestMathService_CheckNonZeroSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMath := mock.NewMockMathInterface(ctrl)

	mockMath.EXPECT().NonZeroSum(gomock.Eq(1), gomock.Eq(2)).Return(3, nil)
	mockMath.EXPECT().NonZeroSum(gomock.Eq(-5), gomock.Eq(5)).Return(0, errors.New(""))

	type fields struct {
		Math math.MathInterface
	}
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{name: "non zero sum", fields: fields{mockMath}, args: args{1, 2}, want: 3, wantErr: false},
		{name: "zero sum", fields: fields{mockMath}, args: args{-5, 5}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathService{
				Math: tt.fields.Math,
			}
			got, err := m.CheckNonZeroSum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathService.CheckNonZeroSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MathService.CheckNonZeroSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
