package entities

import (
	"database/sql"
	"healthcheck/model"
	"log"
)

func MapperTest(data *sql.Rows) []model.ListServiceRequest {
	var svcs []model.ListServiceRequest
	for data.Next() {
		var a model.ListRequest
		err := data.Scan(&a.Id, &a.ServiceName, &a.State, &a.StartTime, &a.EndTime, &a.UpTime, &a.Metadata)
		if err != nil {
			log.Fatal(err)
		}
		svcs = append(svcs, a)
	}

	return svcs
}

func ListAllServiceTest(db *sql.DB) {

}
