package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) error {
	stmt, err := dbConn.Prepare("insert into video_del_rec values (?)")
	defer stmt.Close()
	if err != nil {
		log.Println("Error -> AddVideoDeletionRecord error stmt error : ", err)
		return err
	}

	_, err = stmt.Exec(vid)
	if err != nil {
		log.Println("Error -> AddVideoDeletionRecord error stmt Exec error : ", err)
		return err
	}
	return nil
}