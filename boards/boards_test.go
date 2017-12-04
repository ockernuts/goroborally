package boards

import (
	"ockernuts/goroborally/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func Test_boardFileReader_GetBoardByName(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test")
	board, restError := reader.GetBoardByName("island")
	assert.NotNil(t, board )
	assert.Equal(t, int32(12), *board.Height)
	assert.Equal(t, int32(12), *board.Width)
	assert.Equal(t, 144, len(board.Tiles))
	assert.Nil(t, restError)
	boardAccessor := NewAccessor(board)
	// Check some tiles
	tileAccessor := boardAccessor.GetTile(0,0)
	assert.Equal(t, models.TileTypePlain, tileAccessor.GetType())
	tileAccessor = boardAccessor.GetTile(0,2)
	assert.Equal(t, models.TileTypeRepair, tileAccessor.GetType())
}

func Test_boardFileReader_GetBoardByName_NoWidthField(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test/badfiles")
	board, restError := reader.GetBoardByName("missing-width")
	assert.Nil(t, board )
	assert.NotNil(t, restError)
	assert.Contains(t, restError.Error(), "board.width")
	
}

func Test_boardFileReader_GetBoardByName_NoHeightField(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test/badfiles")
	board, restError := reader.GetBoardByName("missing-height")
	assert.Nil(t, board )
	assert.NotNil(t, restError)
	assert.Contains(t, restError.Error(), "board.height")
}


func Test_boardFileReader_GetBoardByName_NoNameField(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test/badfiles")
	board, restError := reader.GetBoardByName("missing-name")
	assert.Nil(t, board )
	assert.NotNil(t, restError)
	assert.Contains(t, restError.Error(), "board.name")
}



func Test_boardFileReader_GetBoardByName_NoDescriptionField(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test/badfiles")
	board, restError := reader.GetBoardByName("missing-description")
	assert.Nil(t, board )
	assert.NotNil(t, restError)
	assert.Contains(t, restError.Error(), "board.description")
}

func Test_boardFileReader_GetBoardByName_NoTilesField(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test/badfiles")
	board, restError := reader.GetBoardByName("missing-tiles")
	assert.Nil(t, board )
	assert.NotNil(t, restError)
	assert.Contains(t, restError.Error(), "board.tiles")
}


func Test_boardFileReader_GetBoardByName_NotExistingFileGives404(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test")
	board, restError := reader.GetBoardByName("not-existing")
	assert.Nil(t,board)
	assert.Equal(t, http.StatusNotFound, restError.GetResultCode())
}
func Test_boardFileReader_GetBoardByName_BadFileGives500(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test")
	board, restError := reader.GetBoardByName("badfile")
	assert.Nil(t,board)
	assert.Equal(t, http.StatusInternalServerError, restError.GetResultCode())
}


func Test_boardFilterReader_GetBoardNames(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test")
	boardNames, restError := reader.GetBoardNames()
	assert.Nil(t, restError)
	assert.Equal(t, []string{ "badfile", "island"}, boardNames)
}

func Test_boardFilterReader_GetBoardNames_BadPath(t *testing.T) {
	reader := NewBoardProviderFromFile("./not-existing-files")
	boardNames, restError := reader.GetBoardNames()
	assert.NotNil(t, restError)
	assert.Nil(t,boardNames)
}