package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"streamViewPro/api/defs"
)


func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/",homeHandler)
	router.POST("/",homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	router.POST("/upload/:vid-id", proxyHandler)
	return router
}





func main() {
	r := RegisterHandler()
	err := http.ListenAndServe(":8080", r)
	defs.CheckErrorForExit(err)
}