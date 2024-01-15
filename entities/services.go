package entities

import (
	"database/sql"
	"fmt"
	"healthcheck/model"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Mapper(data *sql.Rows) []model.ListRequest {
	var svcs []model.ListRequest
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

func ServiceRegister(db *sql.DB, svc model.AddRequest) string {
	var rsp []model.ListRequest
	// check exist
	rsp = ServiceState(db, svc.ServiceName)

	if rsp == nil {
		startTime := time.Now()
		uptime := "0h0m"
		_, err := db.Query(`INSERT INTO services(service_name, start_time, status, uptime, metadata) VALUES ($1, $2, $3, $4, $5)`,
			svc.ServiceName, startTime, svc.State, uptime, svc.Metadata)
		if err != nil {
			log.Fatal(err)
		}
		rsp = ServiceState(db, svc.ServiceName)
		return "Regist successfully!"
	} else {
		return "Service exist "
	}
}

func ServiceUpdateState(db *sql.DB) gin.HandlerFunc {
	var rsp []model.ListRequest
	secretToken := os.Getenv("TOKEN")
	return func(c *gin.Context) {
		service := c.Param("service")
		bearerToken := c.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]
		rsp = ServiceState(db, service)
		if rsp != nil {
			// ServiceRegister(db, )
		}
		if reqToken != secretToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

		} else {

		}

		var requestBody model.AddRequest
		if err := c.BindJSON(&requestBody); err != nil {
			log.Fatal(err)
		}
		// check exist

		if rsp != nil && reqToken == secretToken {
			if rsp[0].State == "Down" {
				startTime := time.Now()
				_, err := db.Query(`UPDATE services SET status = 'Up', start_time = $1 WHERE service_name = $2;`, startTime, service)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				EndTime := time.Now()
				stime, err := time.Parse(time.RFC3339Nano, rsp[0].StartTime)
				d := EndTime.Sub(stime)
				hours := int(d.Hours())
				minutes := int(d.Minutes()) % 60
				uptime := fmt.Sprintf("%dh%dm", hours, minutes)
				_, err = db.Query(`UPDATE services SET uptime = $1 WHERE service_name = $2;`, uptime, service)
				if err != nil {
					log.Fatal(err)
				}
			}
			// return "Update Status OK!"
			c.IndentedJSON(http.StatusOK, "Update Status OK!")
		} else if rsp != nil && reqToken == secretToken {

		} else {
			// return "service hasn't been registerd"
			c.IndentedJSON(http.StatusOK, "service hasn't been registerd")
		}

	}
}

func ServiceState(db *sql.DB, svc string) []model.ListRequest {
	service, err := db.Query(`SELECT * FROM services WHERE service_name = $1;`, svc)
	if err != nil {
		log.Fatal(err)
	}
	var serviceList []model.ListRequest
	serviceList = Mapper(service)
	defer service.Close()

	return serviceList
}

func ListAllService(db *sql.DB) []model.ListRequest {
	service, err := db.Query("SELECT * FROM services;")
	if err != nil {
		log.Fatal(err)
	}
	var serviceList []model.ListRequest
	serviceList = Mapper(service)
	defer service.Close()

	return serviceList
}

func TagDownService(db *sql.DB, svc_name string) []model.ListRequest {
	_, err := db.Query(`UPDATE services SET status = 'Down', uptime = '0h0m' WHERE service_name = $1;`, svc_name)
	if err != nil {
		log.Fatal(err)
	}
	var serviceList []model.ListRequest
	serviceList = ServiceState(db, svc_name)
	return serviceList
}
