package agentbugfix

import "testing"

func TestDivision(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 8, b: 2, want: 4},
		{name: "with remainder truncates", a: 7, b: 2, want: 3},
		{name: "negative dividend", a: -9, b: 3, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := division(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("division(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
