package lab3

import "testing"

func TestPolynomForm(t *testing.T) {
	type args struct {
		baseNum string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check build polynom",
			args: args{baseNum: "10011"},
			want: "x^4+x+1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PolynomForm(tt.args.baseNum); got != tt.want {
				t.Errorf("PolynomForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
