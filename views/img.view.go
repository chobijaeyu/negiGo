package views

import (
	"context"
	"io"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

type NegiImg struct{}

var (
	bucketname string = os.Getenv("DB_NAME_negi") + "img"
)

func AddGoodsImg(c *gin.Context) {
	ctx := context.Background()
	// projectID := "YOUR_PROJECT_ID"

	file, _, err := c.Request.FormFile("upload")
	if err != nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	// filename := header.Filename

	filename := c.Param("name")
	client, err := storage.NewClient(ctx)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	writer := client.Bucket(bucketname).Object(filename).NewWriter(ctx)
	writer.ContentType = "image/png"

	if _, err := io.Copy(writer, file); err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := writer.Close(); err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	statusCode := http.StatusCreated
	c.JSON(statusCode, gin.H{
		"name": filename,
	})
}

func DeleteGoodsImg(c *gin.Context) {
	ctx := context.Background()
	filename := c.Query("filename")

	client, err := storage.NewClient(ctx)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	o := client.Bucket(bucketname).Object(filename)
	if err := o.Delete(ctx); err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	statusCode := http.StatusNoContent
	c.JSON(statusCode, gin.H{
		"name": filename,
	})
}
