package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

const (
	INDEX_TEMPLATE = "handlers/templates/index.html"
)

func DefaultPageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.New("index.html").ParseFiles(INDEX_TEMPLATE)
	if err != nil {
		error := fmt.Sprintf("Error, could not parse %s, err = %v", INDEX_TEMPLATE, err)
		w.Write([]byte(error))
		println(error)
		return
	}

	var nothing struct{}
	if err := t.Execute(w, nothing); err != nil {
		error := fmt.Sprintf("Error executing template 'index.html', err = %v ", err)
		w.Write([]byte(error))
		println(error)
		return
	}
}


func checkFilePresent(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("%s could not be read. Please execute roborally server in its correct workingdirectory.\n", filename)
	}

}

func init() {
	checkFilePresent(INDEX_TEMPLATE)
}
