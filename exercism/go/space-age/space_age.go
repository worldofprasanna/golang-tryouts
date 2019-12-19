// Package space contains the functionality to calculate age
package space

// Planet struct is used to hold the planet name and its corresponding orbital period
type Planet struct {
	name string
	orbitalPeriod float64
}

const earthOrbitSeconds = 31557600

var planets = []Planet {
	{
		name: "Mercury",
		orbitalPeriod: 0.2408467,
	},
	{
		name: "Venus",
		orbitalPeriod: 0.61519726,
	},
	{
		name: "Earth",
		orbitalPeriod: 1,
	},
	{
		name: "Mars",
		orbitalPeriod: 1.8808158,
	},
	{
		name: "Jupiter",
		orbitalPeriod: 11.862615,
	},
	{
		name: "Saturn",
		orbitalPeriod: 29.447498,
	},
	{
		name: "Uranus",
		orbitalPeriod: 84.016846,
	},
	{
		name: "Neptune",
		orbitalPeriod: 164.79132,
	},
}

// Age function is used to compute the age of the person in the given planet
func Age(seconds float64, planetName string) float64 {
	var givenPlanetOrbital Planet
	for _, planetInfo := range planets {
		if planetInfo.name == planetName {
			givenPlanetOrbital = planetInfo
			break
		}
	}
	return seconds / (givenPlanetOrbital.orbitalPeriod * earthOrbitSeconds)
}


