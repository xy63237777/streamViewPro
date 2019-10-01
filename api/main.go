package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,request *http.Request) {
	validateUserSession(request)
	m.r.ServeHTTP(w, request)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler  {
	m := middleWareHandler{}
	m.r = r
	return m
}


func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}



func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
