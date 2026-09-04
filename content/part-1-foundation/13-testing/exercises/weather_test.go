package exercises

import (
	"errors"
	"testing"
)

// fakeWeatherAPI is a hand-written fake WeatherAPI: it returns canned values
// instead of calling a real service, so the test is fast and deterministic.
type fakeWeatherAPI struct {
	temp float64
	err  error
}

func (f fakeWeatherAPI) GetTemp(string) (float64, error) {
	return f.temp, f.err
}

var errAPIDown = errors.New("weather api down")

func TestWeatherServiceIsHot(t *testing.T) {
	t.Skip("Chapter 13 exercise: implement IsHot, then delete this Skip")

	tests := []struct {
		name    string
		api     fakeWeatherAPI
		want    bool
		wantErr error
	}{
		{"hot", fakeWeatherAPI{temp: 35}, true, nil},
		{"exactly 30 is not hot", fakeWeatherAPI{temp: 30}, false, nil},
		{"cold", fakeWeatherAPI{temp: 12}, false, nil},
		{"api error propagates", fakeWeatherAPI{err: errAPIDown}, false, errAPIDown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &WeatherService{API: tt.api}

			got, err := svc.IsHot("Hanoi")
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("IsHot err = %v; want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("IsHot = %v; want %v", got, tt.want)
			}
		})
	}
}
