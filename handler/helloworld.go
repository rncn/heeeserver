package handler

import (
	"io"
	"net/http"
)

//HelloWorldHandler is Show "unko"
func HelloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "unko")
}
