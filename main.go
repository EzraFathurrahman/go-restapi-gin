package main

import (
	productcontroller "github.com/EzraFathurrahman/go-restapi-gin/controllers/productController"
	"github.com/EzraFathurrahman/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PATCH("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()

}
