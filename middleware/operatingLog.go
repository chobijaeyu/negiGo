package middleware

import (
	"fmt"
	"negigo/models"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	post   = "追加"
	put    = "修正"
	delete = "削除"

	calevdetail            = "タスク"
	fielddetail            = "圃場"
	imgdetail              = "画像"
	membersdetail          = "メンバー"
	titleoptiondetail      = "タスクタイトル"
	seriestaskoptiondetail = "標準タスク"
)

func OperatingLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		ol := models.OperatingLog{}
		ol.Who, ol.Did, ol.What = parseOperate(c)
		if ol.Did != "" {
			ol.When = time.Now().Format(time.RFC3339)
			ol.LoggingOperating()
		}
		c.Next()
	}
}

func parseOperate(c *gin.Context) (who, did, what string) {
	//logging who did what
	_who, _ := c.Get("username")
	who = fmt.Sprintf("%v", _who)

	switch c.Request.Method {
	case "POST":
		did = post
	case "PUT":
		did = put
	case "DELETE":
		did = delete
	}

	var calevre = regexp.MustCompile(`(?m)negicalevent`)
	var fieldre = regexp.MustCompile(`(?m)negifield`)
	var imgre = regexp.MustCompile(`(?m)img`)
	var membersre = regexp.MustCompile(`(?m)members`)
	var titleoptionre = regexp.MustCompile(`(?m)negicustomtasktitleoption`)
	var seriesoptionre = regexp.MustCompile(`(?m)negicustomseriestaskoption`)

	switch {
	case calevre.Match([]byte(c.Request.URL.Path)):
		what = calevdetail
	case fieldre.Match([]byte(c.Request.URL.Path)):
		what = fielddetail
	case imgre.Match([]byte(c.Request.URL.Path)):
		what = imgdetail
	case membersre.Match([]byte(c.Request.URL.Path)):
		what = membersdetail
	case titleoptionre.Match([]byte(c.Request.URL.Path)):
		what = titleoptiondetail
	case seriesoptionre.Match([]byte(c.Request.URL.Path)):
		what = seriestaskoptiondetail
	}

	return
}
