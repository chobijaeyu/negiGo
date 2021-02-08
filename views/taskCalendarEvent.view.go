package views

import (
	"negigo/models"
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, Taskcal)
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

	c.JSON(http.StatusOK, Taskcal)
}

func (tcv TaskCalEvent) DeteleTaskCalEvent(c *gin.Context) {
	Taskcal := models.TaskCalEvent{}

	_uid64, _ := strconv.ParseUint(c.Param("eventid"), 10, 32)
	Taskcal.ID = uint(_uid64)

	if err := Taskcal.DeleteTaskCalEvent(); err != nil {
		c.String(http.StatusBadRequest, "Create taskCalEvent delete err:", err)
		return
	}
	c.String(http.StatusNoContent, "")
}

func (tcv TaskCalEvent) GetAllTaskCalEvent(c *gin.Context) {
	tc := models.TaskCalEvent{}
	if negifieldID, isexit := c.GetQuery("nfID"); isexit {
		cs, err := tc.GetTaskCalEventsByQuery("nfID", negifieldID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, cs)
		return
	}

	/////////
	cs, err := tc.GetAllTaskCalEvents()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cs)
}

func (tcv TaskCalEvent) GetTaskEventsByQuery(c *gin.Context) {
	tc := models.TaskCalEvent{}
	if c.ShouldBind(&tc) == nil {
		cvs, err := tc.GetTaskCalEventsByQuery("nfID", strconv.Itoa(int(tc.NegiFieldID)))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, cvs)
	}
}
