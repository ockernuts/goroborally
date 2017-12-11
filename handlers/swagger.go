package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	BasePath           string = "/api"
	ForceHttpsRequests bool
)

func GetSwaggerJson(w http.ResponseWriter, r *http.Request) {

	input, err := ioutil.ReadFile("swagger.json")
	if err != nil {
		error := fmt.Sprintf("Could not read swagger.json, err = %v", err)
		w.Write([]byte(error))
		println(error)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(input, &jsonMap)
	if err != nil {
		error := fmt.Sprintf("Could not parse swagger/cd-vdcm-swagger.json as json, err = %v", err)
		w.Write([]byte(error))
		println(error)
		return
	}
	// replace the basePath from in the swagger.json file to cope with reverse-proxy added url parts
	jsonMap["basePath"] = BasePath

	if ForceHttpsRequests {
		log.Println("forcing schemes to https")
		schemas := make([]string, 1)
		schemas[0] = "https"
		jsonMap["schemes"] = schemas
	}

	var result []byte
	result, err = json.Marshal(jsonMap)

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
