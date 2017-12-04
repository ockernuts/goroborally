package boards

import (
	"errors"
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ockernuts/goroborally/rest"
	"ockernuts/goroborally/models"
	"path"
	"github.com/go-openapi/swag"
)

// BoardProvider Interface abstracting how we get to a board.
type BoardProvider interface {
	GetBoardByName(name string) (*models.Board, rest.Error)
	GetBoardNames() ([]string, rest.Error)
}

// NewBoardProviderFromFile is the default BoardProvider provider implmentation.
// read them back from a directory of board files in json format (following the REST API format)
func NewBoardProviderFromFile(path string) BoardProvider {
	return &boardFileReader{ path: path}
}

type boardFileReader struct {
	path string
}

func (p *boardFileReader) validate(board *models.Board ) error {
	if board.Width == nil {
		return errors.New("Missing board.width field")
	}
	if board.Height == nil {
		return errors.New("Missing board.height field")
	}
	if board.Name == nil {
		return errors.New("Missing board.name field")
	}
	if board.Description == nil {
		return errors.New("Missing board.description field")
	}
	if board.Tiles == nil {
		return errors.New("Missing board.tiles field")
	}
    return nil
}

func (p *boardFileReader) GetBoardByName(name string) (*models.Board, rest.Error) {
	bytes, err := ioutil.ReadFile(path.Join(p.path, name + ".json"))
	if err != nil {
		return nil, rest.NewRestError(http.StatusNotFound, fmt.Errorf("Could not find board with name %s, error %v", name, err))
	}
	board := models.Board{}
	err = json.Unmarshal(bytes, &board)
	if err != nil {
		return nil, rest.NewRestError(http.StatusInternalServerError, fmt.Errorf("Cannot parse json of board %s, error %v", name, err))
	}

	err = p.validate(&board)
	if err != nil {
		return nil, rest.NewRestError(http.StatusInternalServerError, fmt.Errorf("Error validating parsed json of board %s, %s", name, err.Error()))
	}

	// Shallow copy
	var expandedBoard models.Board = board
	width := swag.Int32Value(board.Width)
	heigth := swag.Int32Value(board.Height)
	size :=  width * heigth
	expandedBoard.Tiles = make([]*models.Tile, size )
	autoDirections := []models.Direction { models.DirectionUp, models.DirectionLeft, models.DirectionDown, models.DirectionRight }
	plainType := models.TileTypePlain
	
	// Initialize board with plain tiles
	for i := int32(0); i < size; i++ {
	   x:= i % width
	   y:= i / width
	   plainTile := models.Tile { Direction : autoDirections[i % 4] , Type: plainType ,  X : &x , Y: &y }
       expandedBoard.Tiles[i] = &plainTile
	}

	// Overrule with read-in tiles. 
	for _, tile := range board.Tiles {
		position := swag.Int32Value(tile.X) + swag.Int32Value(tile.Y)* width
		expandedBoard.Tiles[position] = tile
	}


	return &expandedBoard, nil
}

func (p *boardFileReader) GetBoardNames() ([]string, rest.Error) {
	files, err := ioutil.ReadDir(p.path)
    if err != nil {
		return nil, rest.NewRestError(http.StatusInternalServerError, fmt.Errorf("No boards could be found in path %s, error %v", p.path, err))
	}
	
	result := make([]string, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		name := strings.TrimRight(f.Name(), ".json")
		result = append(result, name)
	}
	return result, nil
}
