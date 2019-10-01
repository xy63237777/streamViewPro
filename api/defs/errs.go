package defs

import "log"

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400,
		Error:Err{Error:"Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC: 401,
		Error:Err{Error:"User authentication failed", ErrorCode: "002"}}
	ErrorDBError = ErrorResponse{HttpSC:500, Error:Err{Error:"DB ops failed", ErrorCode:"003"}}
	ErrorInternalFaults = ErrorResponse{HttpSC:500, Error:Err{Error:"Internal service error", ErrorCode:"004"}}
)

func CheckErrorForExit(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckErrorForExitOfMsg(err error, msg ...string)  {
	if err != nil {
		log.Fatalln("发生了错误: ",msg,"\n", err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErrorOfMsg(err error, msg ...string) {
	if err != nil {
		log.Println("发生了错误: ",msg,"\n", err)
	}
}