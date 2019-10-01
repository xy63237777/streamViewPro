package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"streamViewPro/api/defs"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error  {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmt, err := dbConn.Prepare("insert into sessions(session_id, TTL, login_name) VALUES (?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error)  {
	ss := &defs.SimpleSession{}
	stmt, err := dbConn.Prepare("select TTL, login_name from sessions where session_id = ?")
	defer stmt.Close()
	if err != nil {
		return nil,err
	}
	var ttl  string
	var uname string
	err = stmt.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl,10,64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error)  {
	m := &sync.Map{}
	stmt, err := dbConn.Prepare("select * from sessions")
	defer stmt.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Println("%s", err)
		return nil, err
	}

	for  rows.Next()  {
		var id  string
		var ttlstr string
		var login_name string
		if er := rows.Scan(&id, &ttlstr, &login_name); er !=nil {
			log.Printf("retrive sessions error: %s \n", er)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr,10,64); err !=nil {
			ss := &defs.SimpleSession{Username:login_name, TTL:ttl}
			m.Store(id, ss)
			log.Printf(" session id : %s, ttl: %d \n",id, ss.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error  {
	stmt, err := dbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		log.Panicf("%s\n", err)
		return err
	}

	if _, err := stmt.Query(sid); err != nil {
		return err
	}
	return nil
}

