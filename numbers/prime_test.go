package numbers

import "testing"

func TestIsPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{4}, false},
		{"case 2", args{5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.num); got != tt.want {
				t.Errorf("Case %v IsPrime() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
