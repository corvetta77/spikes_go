package models

var (
	//hot enough for beach
	beachVacationThreshold float64 = 22

	// cold enough for beach
	skiVacationThreshold float64 = -2
)

type city struct {
	name        string
	tempC       float64
	hasBeach    bool
	hasMountain bool
}

type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64

	BeachVacationReady() bool
	SkiVacationReady() bool
}

// NewCity func that creates new City instance with given name
func NewCity(n string, t float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

// Name returns the city name
func (c city) Name() string {
	return c.name
}

// TempC returns the city temperature in Celsius
func (c city) TempC() float64 {
	return c.tempC
}

// TempF returns the city temperature converted to Fahrenheit
func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c city) BeachVacationReady() bool {
	return c.hasBeach && c.tempC > beachVacationThreshold
}

func (c city) SkiVacationReady() bool {
	return c.hasMountain && c.tempC < skiVacationThreshold
}
