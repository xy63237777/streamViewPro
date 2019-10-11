package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"streamViewPro/api/dbops"
	"streamViewPro/api/defs"
	"streamViewPro/api/session"
	"streamViewPro/api/utils"
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
	bytes, _ := ioutil.ReadAll(request.Body)
	log.Println(bytes)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(bytes,ubody); err != nil {
		log.Println(err)
		sendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}
	uname := params.ByName("username")
	log.Println("Login url name : ", uname)
	log.Println("Login body name : ", ubody.Username)
	if uname != ubody.Username {
		sendErrorResponse(writer, defs.ErrorNotAuthUser)
		return
	}
	log.Println(ubody.Username)
	pwd, err := dbops.GetUserCredential(ubody.Username)
	log.Println("Login pwd : ", pwd)
	log.Println("Login body pwd : ",ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(writer, defs.ErrorNotAuthUser)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{Success:true, SessionId:id}
	if resp,err := json.Marshal(si); err != nil {
		sendErrorResponse(writer,defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(writer,string(resp),http.StatusOK)
	}
}

func ShowComments(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if !ValidateUser(writer, request) {
		return
	}
	vid := params.ByName("vid-id")
	comments, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Println("Error in ShowComments : ", err)
		sendErrorResponse(writer, defs.ErrorDBError)
		return
	}

	cms := &defs.Comments{Comments: comments}
	if resp, err := json.Marshal(cms); err != nil {
		sendErrorResponse(writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(writer,string(resp), http.StatusOK)
	}
}

func PostComment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if !ValidateUser(writer, request) {
		return
	}
	reqBody, _ := ioutil.ReadAll(request.Body)
	newComment := &defs.NewComment{}
	if err := json.Unmarshal(reqBody, newComment); err != nil {
		log.Println(err)
		sendErrorResponse(writer, defs.ErrorRequestBodyParseFailed)
		return
	}
	vid := params.ByName("vid-id")
	if err := dbops.AddNewComments(vid, newComment.AuthorId,newComment.Content); err != nil {
		log.Println("Error in Post Comment : ", err)
		sendErrorResponse(writer, defs.ErrorDBError)
	} else {
		sendNormalResponse(writer, "ok",http.StatusCreated)
	}
}

func DeleteVideo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if !ValidateUser(writer, request) {
		return
	}
	vid := params.ByName("vid-id")
	err := dbops.DeleteVideoInfo(vid)
	if err != nil {
		log.Println("Error in DeleteVideo : ", err)
		sendErrorResponse(writer, defs.ErrorDBError)
		return
	}
}

func ListAllVideos(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if !ValidateUser(writer, request) {
		return
	}
	uname := params.ByName("username")
	vs, err := dbops.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Println("Error in ListAllVideos : ",err)
		sendErrorResponse(writer, defs.ErrorDBError)
		return
	}
	videosInfo := &defs.VideosInfo{Videos:vs}
	if resp,err := json.Marshal(videosInfo); err != nil {
		sendErrorResponse(writer,defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(writer, string(resp), http.StatusOK)
	}
}

func AddNewVideo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if !ValidateUser(writer, request) {
		log.Println("UnAthorized user ")
		return
	}
	res, _ := ioutil.ReadAll(request.Body)
	newVideo := &defs.NewVideo{}
	if err := json.Unmarshal(res,newVideo); err != nil {
		log.Println(err)
		sendErrorResponse(writer,defs.ErrorRequestBodyParseFailed)
		return
	}
	vi, err := dbops.AddNewVideo(newVideo.AuthorId, newVideo.Name)
	if err != nil {
		log.Println("Error in AddNewVideo : ", err)
		sendErrorResponse(writer, defs.ErrorDBError)
		return
	}
	if resp,err := json.Marshal(vi); err != nil {
		sendErrorResponse(writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(writer, string(resp), http.StatusCreated)
	}
}

func GetUserInfo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if !ValidateUser(writer, request) {
		log.Println("UnAthorized user ")
		return
	}
	uname := params.ByName("username")
	user, err := dbops.GetUser(uname)
	if err != nil {
		log.Println("error in GetUserInfo : ", err)
		sendErrorResponse(writer,defs.ErrorDBError)
		return
	}
	ui := &defs.UserInfo{Id: user.Id}
	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(writer, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(writer, string(resp), http.StatusOK)
	}
}