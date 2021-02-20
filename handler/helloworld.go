package handler

import (
	"io"
	"net/http"
)

//HelloWorldHandler is Show "hello"
func HelloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello")
}
