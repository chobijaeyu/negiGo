package views

import (
	"negigo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskCalEvent struct{}

func (tcv TaskCalEvent) CreateTaskCalEvent(c *gin.Context) {
	Taskcal := models.TaskCalEvent{}

	if err := c.Bind(&Taskcal); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent bind err:", err)
		return
	}

	if err := Taskcal.CreateTaskCalEvent(); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent create err:", err)
		return
	}
}

func (tcv TaskCalEvent) UpdateTaskCalEvent(c *gin.Context) {
	Taskcal := models.TaskCalEvent{}

	if err := c.Bind(&Taskcal); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent bind err:", err)
		return
	}

	if err := Taskcal.UpdateTaskCalEvent(); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent update err:", err)
		return
	}
}

func (tcv TaskCalEvent) DeteleTaskCalEvent(c *gin.Context) {
	Taskcal := models.TaskCalEvent{}

	if err := c.Bind(&Taskcal); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent bind err:", err)
		return
	}

	if err := Taskcal.DeleteTaskCalEvent(); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent delete err:", err)
		return
	}
}
