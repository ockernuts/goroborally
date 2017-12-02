package boards

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func Test_boardFileReader_GetBoardByName(t *testing.T) {
	reader := NewBoardProviderFromFile("./files-test")
	board, restError := reader.GetBoardByName("island")
	assert.NotNil(t, board )
	assert.Nil(t, restError)
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