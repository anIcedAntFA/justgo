package exercises

import "fmt"

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
	totalSeconds := d
	hours := totalSeconds / 3600
	remainingSeconds := totalSeconds % 3600
	minutes := remainingSeconds / 60
	seconds := remainingSeconds % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	} else {
		return fmt.Sprintf("%ds", seconds)
	}
}

// seconds = duration
// hours = seconds / 3600
// remaining = seconds % 3600
// minutes = remaining / 60
// seconds = remaining % 60
//
// rules
// hours > 0 -> h + m + s
// hours == 0 && minutes > 0 -> m + s
// hours == 0 && minutes == 0 && seconds > 0 -> s
//
// 1. lấy total seconds từ d
// 2. tính hours
// 3. tính remaining seconds
// 4. tính minutes
// 5. tính seconds
// 6. nếu hours > 0 → format h/m/s
// 7. nếu minutes > 0 → format m/s
// 8. nếu không → format s

// caller
//   │
//   │ Duration(3661)
//   ▼
// fmt.Println
//   │
//   │ consumer
//   ▼
// "Does this value have
// String() string?"
//   │
//   ▼
// Stringer
//   ▲
//   │
//   │ yes
//   │
// Duration
//   │
//   ▼
// Duration.String()
//   │
//   ▼
// "1h 1m 1s"
//   │
//   ▼
// terminal

// ┌──────────────────────────────────────────────┐
// │                  CONTRACT                    │
// │                                              │
// │  fmt.Stringer                                │
// │  String() string                             │
// └──────────────────────▲───────────────────────┘
//                        │
//                        │ satisfies
//                        │
// ┌──────────────────────┴───────────────────────┐
// │             CONCRETE TYPE                    │
// │                                              │
// │  Duration                                    │
// │  String() string { ... }                     │
// └──────────────────────▲───────────────────────┘
//                        │
//                        │ capability consumed
//                        │
// ┌──────────────────────┴───────────────────────┐
// │                 CONSUMER                     │
// │                                              │
// │  fmt.Println(Duration(3661))                 │
// │                                              │
// │  fmt formatting logic                        │
// └──────────────────────────────────────────────┘
