package exercises

// Duration counts a number of whole seconds. Implementing String makes it
// satisfy fmt.Stringer, so fmt prints it in the human-readable form.
type Duration int

// String renders the duration as hours, minutes, and seconds:
//
//	3661 → "1h 1m 1s"   (hours present → show h, m, s)
//	  90 → "1m 30s"     (no hours → show m, s)
//	  45 → "45s"        (under a minute → show s only)
//	3600 → "1h 0m 0s"
//	   0 → "0s"
//
// TODO: implement with fmt.Sprintf and add the import yourself.
func (d Duration) String() string {
	return "" // TODO: replace
}
