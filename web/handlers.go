package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)


type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func apiHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//这里对错误的实现并没有分层 如果爬下来想写好的话 就像上面的分下层吧
	if request.Method != http.MethodPost {
		bytes, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(writer, string(bytes))
		return
	}
	res, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(writer, string(re))
	}
	doRequest(apiBody,writer, request)

}

func userHomeHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cname, err1 := request.Cookie("username")
	_, err2 := request.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(writer, request, "/", http.StatusFound)
		return
	}

	fname := request.FormValue("username")
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name:cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}
	file, err := template.ParseFiles("./templates/userhome.html")
	if err != nil {
		log.Println("Parsing, template home.html error : ", err)
		return
	}
	file.Execute(writer,p)
	return
}

func homeHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cname, err1 := request.Cookie("username")
	session, err2 := request.Cookie("session")
	if err1 != nil || err2 != nil {
		loginOrRegistry(writer)
		return
	}
	if len(cname.Value) != 0 && len(session.Value) != nil {
		http.Redirect(writer,request,"/userhome",http.StatusFound)
		return
	}
}

func proxyHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//这里应该写到配置里
	u := "http://127.0.0.1:9000/"
	parse, _ := url.Parse(u)
	proxy := httputil.NewSingleHostReverseProxy(parse)
	proxy.ServeHTTP(writer, request)
}

func loginOrRegistry(writer http.ResponseWriter)  {
	p := &HomePage{Name : "FOUR SEASONS"}
	file, err := template.ParseFiles("./template/home.html")
	if err != nil {
		log.Println("Parsing template home.html error: ", err)
		return
	}
	file.Execute(writer,p)
}