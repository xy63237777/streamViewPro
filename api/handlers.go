package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"streamViewPro/api/dbops"
	"streamViewPro/api/defs"
	"streamViewPro/api/session"
)

func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bytes, _ := ioutil.ReadAll(request.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(bytes, ubody); err != nil {
		sendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(writer, defs.ErrorDBError)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true, SessionId:id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(writer,defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(writer, string(resp), 201)
	}
}

func Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	uname := params.ByName("user_name")
	io.WriteString(writer, uname)
}