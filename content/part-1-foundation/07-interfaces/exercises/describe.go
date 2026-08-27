package exercises

import (
	"fmt"
	"sort"
	"strings"
)

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
	switch x := v.(type) {
	case string:
		return fmt.Sprintf("%q", x)
	case int:
		return fmt.Sprintf("%d", x)
	case float64:
		return fmt.Sprintf("%v", x)
	case bool:
		if x {
			return "yes"
		}
		return "no"
	case []any:
		parts := make([]string, len(x))
		for i, el := range x {
			parts[i] = Describe(el)
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		parts := make([]string, len(keys))
		for i, k := range keys {
			parts[i] = k + ": " + Describe(x[k])
		}
		return "{" + strings.Join(parts, ", ") + "}"
	default:
		return fmt.Sprintf("%v (%T)", x, x)
	}
}

// Describe(v any)
// │
// └── type switch
//     │
//     ├── string
//     │     └── quote
//     │
//     ├── int
//     │     └── format number
//     │
//     ├── float64
//     │     └── format number
//     │
//     ├── bool
//     │     ├── true  → yes
//     │     └── false → no
//     │
//     ├── []any
//     │     │
//     │     ├── Describe(element 1)
//     │     ├── Describe(element 2)
//     │     └── ...
//     │
//     ├── map[string]any
//     │     │
//     │     ├── extract keys
//     │     ├── sort keys
//     │     ├── Describe(value 1)
//     │     ├── Describe(value 2)
//     │     └── ...
//     │
//     └── default
//           └── %v (%T)
