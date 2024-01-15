package main

import (
	"fmt"
	"log"

	"healthcheck/routes"
	"healthcheck/routes/middlewares"
	"healthcheck/shared/database"

	cron "gopkg.in/robfig/cron.v2"

	_ "github.com/lib/pq"
)

func main() {
	// middlewares.generateJWT()
	db, err := database.OpenDB()
	if err != nil {
		log.Printf("Databse Connection Error", err)
	}

	c := cron.New()
	c.AddFunc("@every 0h2m0s", func() {
		middlewares.ActionHandleFunc(db)
		fmt.Println("start check")
	})
	c.Start()
	r := routes.Config(db)
	r.Run(":8080")
}
