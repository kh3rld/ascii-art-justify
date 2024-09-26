package printart

import (
	"testing"
)

func TestPrintArt(t *testing.T) {
	type args struct {
		bannerFileSlice []string
		inputString     string
		alignFlag       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with left alignment",
			args: args{
				bannerFileSlice: mockBannerFileSlice(), // Mock banner data
				inputString:     "Hello",
				alignFlag:       "--align=left",
			},
		},
		{
			name: "Test with right alignment",
			args: args{
				bannerFileSlice: mockBannerFileSlice(),
				inputString:     "Hello",
				alignFlag:       "--align=right",
			},
		},
		{
			name: "Test with center alignment",
			args: args{
				bannerFileSlice: mockBannerFileSlice(),
				inputString:     "Hello",
				alignFlag:       "--align=center",
			},
		},
		{
			name: "Test with justify alignment",
			args: args{
				bannerFileSlice: mockBannerFileSlice(),
				inputString:     "Hello World",
				alignFlag:       "--align=justify",
			},
		},
		{
			name: "Test with unprintable sequence",
			args: args{
				bannerFileSlice: mockBannerFileSlice(),
				inputString:     "\\a",
				alignFlag:       "",
			},
		},
		{
			name: "Test with newline character",
			args: args{
				bannerFileSlice: mockBannerFileSlice(),
				inputString:     "\\n",
				alignFlag:       "--align=left",
			},
		},
		{
			name: "Test with tab character",
			args: args{
				bannerFileSlice: mockBannerFileSlice(),
				inputString:     "\\t",
				alignFlag:       "--align=left",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintArt(tt.args.bannerFileSlice, tt.args.inputString, tt.args.alignFlag)
		})
	}
}

func mockBannerFileSlice() []string {
	// Mocking ASCII art for simplicity. This should match the characters' positions.
	return []string{
		"      ", " _  ", " | |", " | |", " | |", " | |", " | |", " |_|", // 'H'
		"      ", "     ", "  | ", "  | ", "  | ", "  | ", "  | ", "  | ", // 'e'
		"      ", "     ", "  | ", "  | ", "  | ", "  | ", "  | ", "  | ", // 'l'
		"      ", "     ", "  | ", "  | ", "  | ", "  | ", "  | ", "  | ", // 'l'
		"      ", " ___ ", " / _\\", "| | ", " | |", " | |", "|_| ", "     ", // 'o'
		// Repeat for other characters...
	}
}
