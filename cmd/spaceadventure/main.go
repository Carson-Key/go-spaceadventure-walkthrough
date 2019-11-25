package main

import "github.com/carson-key/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: []spaceadventure.Planet{
				spaceadventure.Plant{
					"Earth",
					"Home to life"
				}
			},
		}
	)
}

func mockPlanets() spaceadventure.Plant[] {
	return []spaceadventure.Planet{
		Planet{"Earth", "Life Planet"},
		Planet{"Mars", "Desolate Planet"}
	}
}
