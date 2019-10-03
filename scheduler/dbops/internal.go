package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ReadVideoDeletionRecord(count int) ([]string, error)  {
	stmt, err := dbConn.Prepare("SELECT video_id from video_del_rec limit ?")
	defer stmt.Close()

	var ids []string

	if err != nil {
		return ids, err
	}

	rows, err := stmt.Query(count)
	if err != nil{
		log.Println("Error -> Query VideoDeletionRecord error: ", err)
		return ids, err
	}

	for ; rows.Next();  {
		var id  string
		if err = rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmt, err := dbConn.Prepare("DELETE FROM video_del_rec where video_id = ?")
	defer stmt.Close()
	if err != nil{
		return err
	}
	_, err = stmt.Exec(vid)
	if err != nil {
		log.Println("Error -> Delete VideoDeletionRecord error: ", err)
		return err
	}
	return nil
}