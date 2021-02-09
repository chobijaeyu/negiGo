package main

import (
	"negigo/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(os.Getenv("GIN_ENV"))
	r := routers.Setup()
	r.Run()
}
