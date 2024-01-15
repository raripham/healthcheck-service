package routes

import (
	"database/sql"
	"log"
	"net/http"

	middlewares "healthcheck/routes/middlewares"

	"github.com/gin-gonic/gin"

	"healthcheck/entities"
	"healthcheck/model"

	"github.com/gin-gonic/contrib/static"
)

func Config(db *sql.DB) *gin.Engine {

	r := gin.Default()
	r.Use(gin.Recovery())
	// r.Use(middlewares.AccessHandleFunc())
	// frontend := r.Group("ui")
	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	backend := r.Group("api")
	{
		backend.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})

		backend.GET("/:service", func(c *gin.Context) {
			service := c.Param("service")
			bearerToken := c.Request.Header.Get("Authorization")
			if middlewares.AuthorizeJWT(bearerToken) {
				c.IndentedJSON(http.StatusOK, entities.ServiceState(db, service))
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized",
				})
			}
			// c.IndentedJSON(http.StatusOK, entities.ServiceState(db, service))
		})

		backend.GET("", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, entities.ListAllService(db))
		})

		backend.POST("", func(c *gin.Context) {
			var requestBody model.AddRequest
			if err := c.BindJSON(&requestBody); err != nil {
				log.Fatal(err)
			}
			c.IndentedJSON(http.StatusOK, entities.ServiceRegister(db, requestBody))
		})

		backend.GET("/:service/up", entities.ServiceUpdateState(db))

		// backend.GET("/:service/down", func(c *gin.Context) {
		// 	service := c.Param("service")
		// 	c.IndentedJSON(http.StatusOK, entities.TagDownService(db, service))
		// })
	}
	return r
}
