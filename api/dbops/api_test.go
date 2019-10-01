package dbops

import (
	"fmt"
	"strconv"
	"streamViewPro/api/utils"
	"testing"
	"time"
)

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

func TestVideoWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo",testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func TestComments(t *testing.T)  {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 15
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	comments, err := ListComments(vid, from, to)
	if err != nil {
		checkErrorForTest(t, err, "Error of testListComments: %v")
	}
	fmt.Println("打印数组")
	fmt.Println(len(comments))
	for _, ele := range comments {
		fmt.Println(ele.Content)
	}
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := AddNewComments(vid, aid, content)
	checkErrorForTest(t, err, "Error of testAddComments: %v")
}

func TestAddUser(t *testing.T)  {
	err := AddUserCredential("FOUR_SEASONS", "123456")
	checkErrorForTest(t,err,"Error of AddUser: %v")
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("FOUR_SEASONS", "123456")
	checkErrorForTest(t,err,"Error of AddUser: %v")
}

func testGetUser(t *testing.T)  {
	pwd, e := GetUserCredential("FOUR_SEASONS")
	checkErrorForTest(t,e,"Error of GetUser: %v")
	fmt.Println(pwd)
}

var tempId string

func testDeleteUser(t *testing.T)  {
	err := DeleteUser("FOUR_SEASONS", "123456")
	if err != nil {
		fmt.Println("错误")
	}
	checkErrorForTest(t,err,"Error of DeleteUser: %v")
}


func testAddNewVideo(t *testing.T) {
	info, e := AddNewVideo(1, "my-video")
	checkErrorForTest(t,e,"Error of testAddNewVideo: %v")
	tempId = info.Id
}

func testGetVideoInfo(t *testing.T) {
	res, e := GetVideoInfo(tempId)
	checkErrorForTest(t, e, "Error of testGetVideoInfo : %v")
	fmt.Println(res)
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempId)
	checkErrorForTest(t,err,"Error of testDeleteVideoInfo : %v")
}

func testReGetVideoInfo(t *testing.T) {
	info, err := GetVideoInfo(tempId)
	if err != nil || info != nil {
		checkErrorForTest(t,err,"Error of testReGetVideoInfo : %v")
	}
}

func testReGetUser(t *testing.T)  {
	pwd, e := GetUserCredential("FOUR_SEASONS")
	checkErrorForTest(t,e,"Error of ReGetUser: %v")

	if pwd != "" {
		checkErrorForTest(t,nil,"Deleting user test failed")
	}
}

func checkErrorForTest(t *testing.T,err error,format string)  {
	if err != nil {
		t.Errorf(format,err)
	}
}



func TestDefer(t *testing.T)  {
	fmt.Println("TestDefer invoke")
	defer func() {
		fmt.Println("defer invoke...")
	}()
	for i := 0; i < 3; i++ {
		if i == 2 {
			return
		}
	}
}

func TestUUID(t *testing.T)  {
	s, e := utils.NewUUID()
	if err != nil {
		fmt.Println(e)
	}
	fmt.Println(s)
}

