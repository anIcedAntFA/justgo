package exercises

// Exercise 3: Word-frequency counter.
//
// Count how often each space-separated word appears in text, then return the
// counts ordered by frequency (highest first). Map iteration order is random, so
// you MUST sort to get a deterministic result: extract the entries into a slice
// and sort them. Break ties alphabetically by word so the output is stable.

// WordCount pairs a word with how many times it occurred.
type WordCount struct {
	Word  string
	Count int
}

// WordFrequency splits text on spaces, tallies each word in a map, then returns
// the tally sorted by Count descending, ties broken by Word ascending.
//
// Example: "the quick the fox the" →
//
//	[]WordCount{{"the", 3}, {"fox", 1}, {"quick", 1}}
//
// TODO:
//  1. counts := map[string]int{} and tally with strings.Fields(text).
//  2. Move entries into a []WordCount.
//  3. slices.SortFunc by Count desc, then Word asc.
func WordFrequency(text string) []WordCount {
	// TODO: implement
	return nil
}
