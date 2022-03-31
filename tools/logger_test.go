package tools

import "testing"

func Test_fmtArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "basic",
			args: []string{"a", "b", "c"},
			want: "a b c",
		},
		{
			name: "empty",
			args: []string{},
			want: "",
		},
		{
			name: "nil",
			want: "",
		},
		{
			name: "has space",
			args: []string{"a", "b c", "d"},
			want: `a "b c" d`,
		},
		{
			name: "has tab",
			args: []string{"a", "b\tc", "d"},
			want: "a \"b\tc\" d",
		},
		{
			name: "has newline",
			args: []string{"a", "b\nc", "d"},
			want: "a \"b\nc\" d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmtArgs(tt.args); got != tt.want {
				t.Errorf("fmtArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
