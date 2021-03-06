package dto

import (
	"github.com/magleff/galaxy-challenge/common"
)

type StatusPlanet struct {
	ID       int16   `json:"id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	OwnerID  int16   `json:"owner"`
	Units    int16   `json:"units"`
	MaxUnits int16   `json:"mu"`
	Growth   int16   `json:"gr"`
	Category string  `json:"class"`
}

func (p StatusPlanet) Copy() StatusPlanet {
	return StatusPlanet{
		ID:       p.ID,
		X:        p.X,
		Y:        p.Y,
		OwnerID:  p.OwnerID,
		Units:    p.Units,
		MaxUnits: p.MaxUnits,
		Growth:   p.Growth,
		Category: p.Category,
	}
}

func FilterStatusPlanets(toFilter []StatusPlanet, predicate func(StatusPlanet) bool) []StatusPlanet {
	filtered := make([]StatusPlanet, 0)
	for _, planet := range toFilter {
		if predicate(planet) {
			filtered = append(filtered, planet)
		}
	}
	return filtered
}

func FilterOwnPlanets(toFilter []StatusPlanet) []StatusPlanet {
	return FilterStatusPlanets(toFilter, func(planet StatusPlanet) bool {
		return planet.OwnerID == common.PLAYER_OWNER_ID
	})
}

func FilterEnemyPlanets(toFilter []StatusPlanet) []StatusPlanet {
	return FilterStatusPlanets(toFilter, func(planet StatusPlanet) bool {
		return planet.OwnerID != common.PLAYER_OWNER_ID
	})
}

func FilterPlanetsByPlayerID(toFilter []StatusPlanet, playerID int16) []StatusPlanet {
	return FilterStatusPlanets(toFilter, func(planet StatusPlanet) bool {
		return planet.OwnerID == playerID
	})
}

func GetByID(toFilter []StatusPlanet, id int16) StatusPlanet {
	filtered := FilterStatusPlanets(toFilter, func(planet StatusPlanet) bool {
		return planet.ID == id
	})
	if len(filtered) == 1 {
		return filtered[0]
	} else {
		return StatusPlanet{}
	}
}
