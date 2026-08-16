package exercises

// IsPrime reports whether n is a prime number. Any n < 2 is not prime.
//
// TODO: implement with a for loop. Trial division up to √n is more than enough.
func IsPrime(n int) bool {
	return false // TODO: replace with the real implementation
}

// Classify categorises n with a TAGLESS switch, evaluated top-down so the first
// matching case wins:
//
//	divisible by BOTH 3 and 5 → "fizzbuzz"
//	prime                     → "prime"
//	even                      → "even"
//	otherwise                 → "odd"
//
// TODO: implement this. Order matters — 2 is prime AND even, and must come out
// as "prime" because the prime case is checked before the even case.
func Classify(n int) string {
	return "" // TODO: replace with the real implementation
}
