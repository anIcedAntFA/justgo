package exercises

// HumanSize formats a byte count as a short human-readable string using base-1024
// units: B, KB, MB, GB, TB. Whole bytes are printed as an integer; larger units
// get exactly one decimal place. This is the kind of formatting `gorg stats`
// prints per category.
//
//	HumanSize(512)     → "512 B"
//	HumanSize(1024)    → "1.0 KB"
//	HumanSize(1536)    → "1.5 KB"
//	HumanSize(1048576) → "1.0 MB"
//
// TODO: implement. Hints:
//   - If n < 1024, return strconv.FormatInt(n, 10) + " B".
//   - Otherwise divide by 1024.0 (a float64!) repeatedly, advancing the unit,
//     while the value stays >= 1024.
//   - Format the number with strconv.FormatFloat(v, 'f', 1, 64).
func HumanSize(n int64) string {
	// TODO: implement
	return ""
}
