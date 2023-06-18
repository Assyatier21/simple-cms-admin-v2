package helper

import (
	"log"
	"time"
)

var (
	jakartaLoc, _  = time.LoadLocation("Asia/Jakarta")
	TimeNowJakarta = time.Now().In(jakartaLoc)
)

func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
