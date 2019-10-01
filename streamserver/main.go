package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,request *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many Request")
		return
	}

	m.r.ServeHTTP(w, request)
	defer m.l.ReleaseConn()
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}



func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)

	return router
}



func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r,2)
	http.ListenAndServe(":9000", mh)
}