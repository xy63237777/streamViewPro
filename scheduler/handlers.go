package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"streamViewPro/scheduler/dbops"
)

func vidDelRecHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	vid := params.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(writer, 400, "video id should not be empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(writer, 500, "Internal server error")
		return
	}
	sendResponse(writer, 200, "ok")
	return
}