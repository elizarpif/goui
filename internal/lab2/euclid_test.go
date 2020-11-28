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