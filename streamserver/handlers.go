package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func uploadHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	vid := params.ByName("vid-id")
	vl := VIDEO_DER + vid
	file, err := os.Open(vl)
	defer file.Close()
	if err != nil {
		sendErrorResponse(writer, http.StatusInternalServerError, "Internal Error")
		return
	}
	writer.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(writer,request,"",time.Now(), file)

}

func streamHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	log.Println("streamHandler invoke...")
	request.Body = http.MaxBytesReader(writer, request.Body, MAX_UPLOAD_SIZE)
	if err := request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil{
		sendErrorResponse(writer, http.StatusBadRequest,"file is too big")
		return
	}
	//可以拿到header的头部可以做一些文件的验证
	file, _, err := request.FormFile("file") //<form name = "file"
	if err != nil {
		log.Println("Error when try open file: ", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Internal Error")
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Read file error: ", err)
		sendErrorResponse(writer, http.StatusInternalServerError,"Internal Error")
	}
	fileName := params.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DER+fileName, bytes, 0666)
	if err != nil {
		log.Println("Writer file error: ", err)
		sendErrorResponse(writer, http.StatusInternalServerError, "Internal Error upload File Error")
	}
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer,"upload file success")
}