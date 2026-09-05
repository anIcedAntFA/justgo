package exercises

// NormalizeExt returns a file's extension, lower-cased and including the leading
// dot, or "" when the name has no extension. Only the final extension counts.
// This is how `gorg` turns a filename into a rule key.
//
//	NormalizeExt("IMG_1234.JPG")   → ".jpg"
//	NormalizeExt("archive.tar.gz") → ".gz"
//	NormalizeExt("README")         → ""
//	NormalizeExt(".gitignore")     → ""    (a leading dot is a dotfile, not an ext)
//	NormalizeExt("Report.Pdf")     → ".pdf"
//
// TODO: implement using the strings package (strings.LastIndex, strings.ToLower).
// Edge case: a dot at index 0 (or no dot at all) means "no extension".
func NormalizeExt(name string) string {
	// TODO: implement
	return ""
}
