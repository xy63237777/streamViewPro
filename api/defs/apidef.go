package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type UserSession struct {
	Username string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type User struct {
	Id int
	LoginName string
	Pwd string
}

type UserInfo struct {
	Id int `json:"id"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}