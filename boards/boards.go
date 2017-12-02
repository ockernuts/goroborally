package boards

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ockernuts/goroborally/rest"
	"ockernuts/goroborally/models"
	"path"
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
	return &board, nil
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
