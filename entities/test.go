package entities

import (
	"database/sql"
	"fmt"
	"healthcheck/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func MapperServiceTest(data *sql.Rows) []model.ListServiceRequest {
	var services []model.ListServiceRequest
	for data.Next() {
		var svc model.ListServiceRequest
		err := data.Scan(&svc.ServiceId, &svc.ServiceName, &svc.State, &svc.StartTime, &svc.EndTime, &svc.UpTime, &svc.Node.NodeId, &svc.ServiceMetadata)
		if err != nil {
			log.Fatal(err)
		}
		services = append(services, svc)
	}

	return services
}

func GetAllServiceTest(db *sql.DB) gin.HandlerFunc {
	var services []model.ListServiceRequest
	var nodes []model.ListNodeRequest
	dataSvc, err := db.Query(`SELECT * FROM services;`)
	if err != nil {
		log.Fatal(err)
	}
	services = MapperServiceTest(dataSvc)

	for i := 0; i < len(services); i++ {
		dataNo, err := db.Query(`SELECT * FROM nodes WHERE node_id = $1;`, services[i].Node.NodeId)
		if err != nil {
			log.Fatal(err)
		}
		nodes = MapperNodeTest(dataNo)
		services[i].Node = nodes[0]
	}

	return func(ctx *gin.Context) {

		ctx.IndentedJSON(http.StatusOK, services)
	}
}

func GetServiceByIdTest(db *sql.DB) gin.HandlerFunc {
	var services []model.ListServiceRequest
	var nodes []model.ListNodeRequest

	return func(ctx *gin.Context) {
		serviceId := ctx.Param("service_id")
		dataSvc, err := db.Query(`SELECT * FROM services WHERE service_id = $1;`, serviceId)
		if err != nil {
			log.Fatal(err)
		}
		services = MapperServiceTest(dataSvc)

		dataNo, err := db.Query(`SELECT * FROM nodes WHERE node_id = $1;`, services[0].Node.NodeId)
		if err != nil {
			log.Fatal(err)
		}
		nodes = MapperNodeTest(dataNo)
		services[0].Node = nodes[0]

		ctx.IndentedJSON(http.StatusOK, services)
	}
}

func ServiceRegisterTest(db *sql.DB, ser model.AddServiceRequest) (int64, bool) {
	var serviceId int64
	no := model.AddNodeRequest{
		NodeName:     ser.NodeName,
		NodeIp:       ser.NodeIp,
		NodeMetadata: ser.NodeMetadata,
	}
	// Register node info
	nodeId := NodeRegisterTest(db, no)
	// check exist
	dataService, err := db.Query(`SELECT * FROM services WHERE service_name = $1 and node_id = $2;`, ser.ServiceName, nodeId)
	if err != nil {
		log.Fatal(err)
	}
	services := MapperServiceTest(dataService)
	dataService.Close()
	if services != nil {
		serviceId = services[0].ServiceId
		log.Println("Service has already register")
		return serviceId, true
	} else {
		startTime := time.Now()
		upTime := "0h0m"
		serId, err := db.Query(`INSERT INTO services(service_name, status, start_time, uptime, node_id, service_metadata) VALUES ($1, $2, $3, $4, $5, $6) returning service_id`,
			ser.ServiceName, ser.State, startTime, upTime, nodeId, ser.ServiceMetadata)
		if err != nil {
			log.Fatal(err)
		}
		serId.Scan(&serviceId)
		serId.Close()
		return serviceId, false
	}

}
func ServiceUpdateInfoTest(db *sql.DB) gin.HandlerFunc {
	var requestBody model.AddServiceRequest
	// secretToken := os.Getenv("TOKEN")
	secretToken := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDAwMDAwMDAsImlzcyI6ImlzIiwic3ViIjoiNTE1MSJ9.X9znAANseGWiMsP6rih3OBYd8Id_jaIYEyO0uefviUNtx3l_DwV7e_t8eTjJz5BMCn0lq"

	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[2]

		if err := ctx.BindJSON(&requestBody); err != nil {
			log.Fatal(err)
		}

		if reqToken != secretToken {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

		} else {
			serviceId, serviceEx := ServiceRegisterTest(db, requestBody)
			if !serviceEx {
				ctx.IndentedJSON(http.StatusOK, "Service do not exist! Service Register successfully!")
			} else {
				dataService, err := db.Query(`SELECT * FROM services WHERE service_id = $1;`, serviceId)
				if err != nil {
					log.Fatal(err)
				}
				services := MapperServiceTest(dataService)
				dataService.Close()

				if services[0].State == "Down" {
					startTime := time.Now()
					upTime := "0h0m"
					_, err = db.Query(`UPDATE services SET status = 'Up', uptime = $1, start_time = $2 WHERE service_id = $3;`, upTime, startTime, serviceId)
				} else {
					endTime := time.Now()
					sTime, err := time.Parse(time.RFC3339Nano, services[0].StartTime)
					d := endTime.Sub(sTime)
					hours := int(d.Hours())
					minutes := int(d.Minutes()) % 60
					upTime := fmt.Sprintf("%dh%dm", hours, minutes)
					_, err = db.Query(`UPDATE services SET uptime = $1 , service_metadata = $2 WHERE service_id = $3;`,
						upTime, requestBody.ServiceMetadata, serviceId)
					if err != nil {
						log.Fatal(err)
					}
				}
				ctx.IndentedJSON(http.StatusOK, "Update info successfully!")
			}
		}
	}
}
