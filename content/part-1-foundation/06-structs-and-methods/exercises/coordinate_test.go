package exercises

import "testing"

func TestPlaceCache(t *testing.T) {
	t.Skip("Chapter 06 exercise: implement PlaceCache, then delete this Skip")

	cache := PlaceCache{}
	cache.Set(Coordinate{Lat: 10.5, Lng: 20.5}, "Home")

	// A separately-created Coordinate with the same field values must hit the
	// same entry — struct value equality as a map key.
	got, ok := cache.Get(Coordinate{Lat: 10.5, Lng: 20.5})
	if !ok {
		t.Fatal("expected a cache hit for an equal Coordinate")
	}
	if got != "Home" {
		t.Errorf("Get(...) = %q, want %q", got, "Home")
	}

	// A different Coordinate must miss.
	if _, ok := cache.Get(Coordinate{Lat: 0, Lng: 0}); ok {
		t.Error("expected a cache miss for an absent Coordinate")
	}
}
