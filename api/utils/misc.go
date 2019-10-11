package utils

import (
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

func NewUUID() (string, error) {
	uuids, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return  uuids.String() ,nil
}

func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano() / 1000000000,64))
	return ts
}