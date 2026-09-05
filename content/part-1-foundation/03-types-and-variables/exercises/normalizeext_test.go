package exercises

import "testing"

func TestNormalizeExt(t *testing.T) {
	t.Skip("Chapter 03 exercise: implement NormalizeExt, then delete this Skip")

	cases := []struct {
		name string
		in   string
		want string
	}{
		{"simple upper", "IMG_1234.JPG", ".jpg"},
		{"double extension", "archive.tar.gz", ".gz"},
		{"no extension", "README", ""},
		{"dotfile", ".gitignore", ""},
		{"mixed case", "Report.Pdf", ".pdf"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NormalizeExt(tc.in); got != tc.want {
				t.Errorf("NormalizeExt(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
