package session

import (
	"streamViewPro/api/dbops"
	"streamViewPro/api/defs"
	"streamViewPro/api/utils"
	"sync"
	"time"
)


var sessionsMap *sync.Map

func init()  {
	sessionsMap = &sync.Map{}
}

func deleteExpiredSession(sid string) {
	sessionsMap.Delete(sid)
	err := dbops.DeleteSession(sid)
	if err != nil {
		defs.CheckErrorOfMsg(err, "deleteExpiredSession invoke...")
	}
}

func nowMilli() int64 {
	return time.Now().UnixNano()/1000000
}

func loadSessionsFromDB()  {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		defs.CheckErrorOfMsg(err, "loadSessionsFromDB invoke...")
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionsMap.Store(key,ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowMilli()
	ttl := ct + 30 * 60 *1000 //Severside session valid time: 30 min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	err := dbops.InsertSession(id, ttl, un)
	if err != nil {
		defs.CheckErrorOfMsg(err, "GenerateNewSessionId invoke ... ")
		return ""
	} else {
		sessionsMap.Store(id, ss)
	}
	return id
}

func IsSessionExpired(sid string) (string, bool)  {
	value, ok := sessionsMap.Load(sid)
	if ok{
		ct := nowMilli()
		if value.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return value.(*defs.SimpleSession).Username, false
	}
	return "", true
}