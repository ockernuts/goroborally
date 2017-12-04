package tiles

import (
	"ockernuts/goroborally/models"
)

type Accessor interface {
	GetDirection() models.Direction
	GetType() models.TileType
}

func NewAccessor(tile *models.Tile) Accessor {
	return &accessor{ tile: tile}
}

type accessor struct {
	tile *models.Tile
}
func (a *accessor) GetDirection() models.Direction {
	return a.tile.Direction
}

func (a *accessor) GetType() models.TileType {
	return a.tile.Type
}
