package boards

import (
	"github.com/go-openapi/swag"
	"ockernuts/goroborally/boards/tiles"
	"ockernuts/goroborally/models"
)

type Accessor interface {
	GetTile(x int, y int ) tiles.Accessor
}

func NewAccessor(board *models.Board) Accessor {
	return &accessor{ board : board}
}

type accessor struct {
	board *models.Board
}

func (a *accessor) GetTile(x int, y int ) tiles.Accessor {
	width := int(swag.Int32Value(a.board.Width))
	position := x + width*y
	tile := a.board.Tiles[position]
	if tile == nil {
		return nil
	}
	return tiles.NewAccessor(tile)
}