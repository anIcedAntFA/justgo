package exercises

// Describe renders a value differently depending on its dynamic type, using a
// type switch. Rules:
//
//	"hi"                   → `"hi"`          (string, quoted)
//	42                     → "42"            (int)
//	3.5                    → "3.5"           (float64)
//	true / false           → "yes" / "no"    (bool)
//	[]any{1, "x"}          → `[1, "x"]`      (slice, each element described)
//	map[string]any{"k": 1} → "{k: 1}"        (map, keys sorted, values described)
//	anything else          → "5 (uint)"      ("%v (%T)")
//
// TODO: implement with a type switch (switch x := v.(type)). Recurse for []any
// and map[string]any. Sort the map keys so the output is deterministic. Add the
// imports you need (fmt, sort, strings).
func Describe(v any) string {
	return "" // TODO: replace
}
