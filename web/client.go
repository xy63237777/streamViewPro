package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient *http.Client

func init()  {
	httpClient = &http.Client{}
}

func doRequest(b *ApiBody, writer http.ResponseWriter, request *http.Request)  {
	switch b.Method {
	case http.MethodGet:
		doRequest0(http.MethodGet,b.Url,nil,request,writer)
	case http.MethodPost:
		doRequest0(http.MethodPost,b.Url,bytes.NewBuffer([]byte(b.ReqBody)),request,writer)
	case http.MethodDelete:
		doRequest0(http.MethodDelete,b.Url,nil,request,writer)
	default:
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Bad api request")
	}
}

func normalResponse(writer http.ResponseWriter, response *http.Response)  {
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, string(re))
		return
	}
	writer.WriteHeader(response.StatusCode)
	io.WriteString(writer, string(bytes))
}

func doRequest0(method, url string, body io.Reader,request *http.Request,writer http.ResponseWriter)  {
	var resp *http.Response
	var err error
	req, _  := http.NewRequest(method, url, body)
	req.Header = request.Header
	resp, err = httpClient.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	normalResponse(writer, resp)
}