package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func DebugRequest(request *http.Request) {
	buf, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Printf("Error reading request body: %s", err.Error())
		return
	}

	log.Printf("Request body: %s", string(buf))

	reader := ioutil.NopCloser(bytes.NewBuffer(buf))
	request.Body = reader
}