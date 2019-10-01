package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"streamViewPro/api/defs"
	"streamViewPro/api/utils"
	"time"
)


func AddUserCredential(loginName string, pwd string) error  {
	stmt, e := dbConn.Prepare("insert into users (login_name, pwd) values(?,?)")
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, err := stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error)  {
	stmt, e := dbConn.Prepare("select pwd from users where login_name = ?")
	defer stmt.Close()
	if e != nil {
		defs.CheckErrorOfMsg(e,"GetUserCredential invoke... ",loginName)
		return "", e
	}
	var pwd string
	err := stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "",err
	}

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmt, e := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	defer stmt.Close()
	if e != nil{
		defs.CheckErrorOfMsg(e,"DeleteUser invoke ",loginName, " pwd : ",pwd)
		return e
	}
	_, err := stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error)  {
	//create uuid
	uuid, err := utils.NewUUID()
	if err != nil{
		return nil, err
	}
	t := time.Now()
	//固定的格式 M D y, HH:MM:SS Go的时间原点
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmt, err := dbConn.Prepare(`insert into video_info 
	    	(id, author_id, name, display_ctime) values (?,?,?,?)`)
	defer stmt.Close()
	if err != nil{
		return nil, err
	}
	_, err = stmt.Exec(uuid, aid, name, ctime)
	if err != nil{
		return nil, err
	}
	res := &defs.VideoInfo{Id: uuid, AuthorId:aid, Name:name, DisplayCtime:ctime}
	return res,nil
}

func GetVideoInfo(uuid string) (*defs.VideoInfo, error)  {
	stmt, err := dbConn.Prepare(`select author_id, name, display_ctime from video_info where id = ?`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	var aid int
	var dct string
	var name string

	err = stmt.QueryRow(uuid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.VideoInfo{Id: uuid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(uuid string) error  {
	stmt, err := dbConn.Prepare("DELETE FROM video_info where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid)
	return err
}

func AddNewComments(vid string, aid int, content string) error  {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmt, err := dbConn.Prepare(`insert into comments (id, video_id, author_id, content, time) values (?,?,?,?
                                                                      ,from_unixtime(?))`)
	if err != nil {
		return err
	}
	temp, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	_, err = stmt.Exec(id, vid, aid, content,temp)
	defer stmt.Close()
	return err 
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error)  {
	fmt.Println(vid, " ", from, "  " , to)
	stmt, err := dbConn.Prepare(`select comments.id, users.login_name,comments.content 
			from comments inner join users on comments.author_id = users.id
			where comments.video_id = ? and comments.time > from_unixtime(?) and comments.time <= from_unixtime(?)
			`)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	var res []*defs.Comment
	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res,err
		}
		temp := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, temp)
		fmt.Println(res)
	}
	return res, nil
}