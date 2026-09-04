package exercises

// WeatherAPI is the dependency WeatherService needs — the seam a test replaces
// with a fake so no real network call happens. This is the "accept interfaces"
// lesson from Chapter 7, applied to make code testable.
type WeatherAPI interface {
	// GetTemp returns the current temperature in Celsius for city.
	GetTemp(city string) (float64, error)
}

// WeatherService is the code under test. It holds a WeatherAPI, not a concrete
// HTTP client, so tests can inject a fake.
type WeatherService struct {
	API WeatherAPI
}

// IsHot reports whether it is hotter than 30°C in city.
//
// TODO:
//  1. Call s.API.GetTemp(city).
//  2. If it returns an error, return false and that error.
//  3. Otherwise return whether the temperature is greater than 30.
func (s *WeatherService) IsHot(city string) (bool, error) {
	return false, nil // TODO
}
