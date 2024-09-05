package check

import (
	"os"
	"testing"
)

func TestValidFile(t *testing.T) {
	fileData, err := os.ReadFile("./standard.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	type args struct {
		bannerFileData []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "standard.txt",
			args: args{fileData},
			want: "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidFile(tt.args.bannerFileData); got != tt.want {
				t.Errorf("ValidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
