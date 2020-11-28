package lab3

import (
	"reflect"
	"testing"
)

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

func TestNewBinaryPolynom(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want *BinaryPolynom
	}{
		{
			name: "check build",
			args: args{num: 10},
			want: &BinaryPolynom{numbers: []int{0,1,0,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryPolynom(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryPolynom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryPolynom_String(t *testing.T) {
	type fields struct {
		numbers []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "check",
			fields: fields{numbers: []int{0,1,0,1}},
			want: "x+x^3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bp := &BinaryPolynom{
				numbers: tt.fields.numbers,
			}
			if got := bp.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}