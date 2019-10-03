package taskrunner

import (
	"errors"
	"log"
	"os"
	"streamViewPro/scheduler/dbops"
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	strings, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Println("Video clear dispatcher error: ", err)
		return err
	}

	if len(strings) <= 0 {
		return errors.New("All tasks finished")
	}
	
	for _, id := range strings {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
	forloop:
		for ; ;  {
			select {
			case vid := <- dc :
				go doClearVideo(vid,errMap)
			default:
				break forloop
			}
		}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}

func clearVideoForFile(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Println("Delete video error For File : ", err)
		return err
	}
	return nil
}

func clearVideoForDB(vid string) error {
	if err := dbops.DelVideoDeletionRecord(vid); err != nil {
		log.Println("Error -> Delete video error For DB : ",err)
		return err
	}
	return nil
}

func doClearVideo(obj interface{}, errMap *sync.Map)  {
	if err := clearVideoForFile(obj.(string)); err == nil {
		errMap.Store(obj, err)
		err = clearVideoForDB(obj.(string))
		if err != nil {
			errMap.Store(obj.(string) + "db",err)
		}
	}
}
