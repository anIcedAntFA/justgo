package exercises

// Coordinate is a comparable struct (both fields are comparable), so two
// separately-created Coordinates with equal Lat/Lng are the SAME map key.
type Coordinate struct {
	Lat float64
	Lng float64
}

// PlaceCache maps a Coordinate to a place name. It relies on struct value
// equality: Get with a freshly-built Coordinate finds what Set stored under an
// equal one.
type PlaceCache map[Coordinate]string

// Set stores name at coord.
//
// TODO: implement (one line — write into the map).
func (pc PlaceCache) Set(coord Coordinate, name string) {
	// TODO: replace
}

// Get returns the name stored at coord and whether it was present, mirroring the
// comma-ok map lookup.
//
// TODO: implement with a comma-ok map read.
func (pc PlaceCache) Get(coord Coordinate) (string, bool) {
	return "", false // TODO: replace
}
