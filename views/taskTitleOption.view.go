package views

import (
	"negigo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NegiTaskTitleOption struct{}

func (tov NegiTaskTitleOption) CreateTaskTitleOption(c *gin.Context) {
	tasktitleoption := models.TaskTitleOption{}

	if err := c.Bind(&tasktitleoption); err != nil {
		c.String(http.StatusBadRequest, "Create task title option bind err : %v", err)
		return
	}

	if err := tasktitleoption.CreateTaskTitleOption(); err != nil {
		c.JSON(http.StatusInternalServerError, "Create task title option db err : "+err.Error())
	}

	c.JSON(http.StatusCreated, tasktitleoption)
}

func (tov NegiTaskTitleOption) UpdateTaskTitleOption(c *gin.Context) {
	tasktitleoption := models.TaskTitleOption{}

	if err := c.Bind(&tasktitleoption); err != nil {
		c.String(http.StatusBadRequest, "Update Negi task title option bind err : %v", err)
		return
	}

	if err := tasktitleoption.UpdatesTaskTitleOption(); err != nil {
		c.String(http.StatusInternalServerError, "Update Negi task title option db err : %v", err)
		return
	}

	c.JSON(http.StatusOK, tasktitleoption)
}

func (tov NegiTaskTitleOption) DeleteTaskTitleOption(c *gin.Context) {
	tasktitleoption := models.TaskTitleOption{}

	negitasktitleoptionid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, "Detele task title opiton no id param : %v", err)
		return
	}

	tasktitleoption.ID = uint(negitasktitleoptionid)

	if err := tasktitleoption.DeleteTaskTitleOption(); err != nil {
		c.String(http.StatusInternalServerError, "Delete task title option db err : %v", err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func (tov NegiTaskTitleOption) GetAllTaskTitleOption(c *gin.Context) {
	tasktitleoption := models.TaskTitleOption{}

	titleoptions, err := tasktitleoption.GetAllTaskTitleOption()
	if err != nil {
		c.String(http.StatusInternalServerError, "Get all task title option err : %v", err)
		return
	}

	c.JSON(http.StatusOK, titleoptions)
}
