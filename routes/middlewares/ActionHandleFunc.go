package middlewares

import (
	"database/sql"
	"healthcheck/entities"
	"healthcheck/model"
	"log"
	"time"
)

func convertTimeToTimestamp(date string) int {
	//Check the documentation on Go for the const variables!
	//They need to be exactly as they are shown in the documentation to be read correctly!
	// format := "2006-01-02T15:04:05.000-0700"

	// t, err := time.Parse(format, date)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(t.Unix())
	// 	return int(t.Unix())
	// }
	// a := 2 * time.Hour
	t, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(t)
	durationInMins := int(duration.Minutes())

	// fmt.Println(t.String(), err)
	// if durationInMins < 20 {
	// 	fmt.Println(int(duration.Minutes()), err)
	// }
	return durationInMins
}
func ActionHandleFunc(db *sql.DB) {
	var serviceList []model.ListRequest
	serviceList = entities.ListAllService(db)
	for i := 0; i < len(serviceList); i++ {
		if serviceList[i].State == "Up" && convertTimeToTimestamp(serviceList[i].EndTime) > 3 {
			entities.TagDownService(db, serviceList[i].ServiceName)
			// log.Println()
		}
	}

}
