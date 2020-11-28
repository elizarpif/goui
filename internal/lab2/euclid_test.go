package lab2

import "testing"

func TestModIn(t *testing.T) {
	type args struct {
		a int
		m int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "check inverse",
			args: args{
				a: 11,
				m: 256,
			},
			want: 163,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ModIn(tt.args.a, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ModIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinPow(t *testing.T) {
	type args struct {
		a int
		n, m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "check",
			args: args{
				a: 3,
				n: 2,
				m: 256,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinPow(tt.args.a, tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("BinPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
