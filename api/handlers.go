package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	io.WriteString(writer, "create User Handler")
}

func Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	uname := params.ByName("user_name")
	io.WriteString(writer, uname)
}