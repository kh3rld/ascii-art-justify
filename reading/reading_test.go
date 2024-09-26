package reading

import (
	"os"
	"reflect"
	"testing"
)

func TestReading(t *testing.T) {
	tempFile, err := os.Create("temporary.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove("temporary.txt")

	_, err = tempFile.WriteString("line 1\nline 2\nline 3\n")
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tempFile.Close()

	type args struct {
		bannerFile string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "starting with default",
			args: args{"temporary.txt"},
			want: []string{"line 1", "line 2", "line 3", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reading(tt.args.bannerFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reading() = %v, want %v", got, tt.want)
			}
		})
	}
}
