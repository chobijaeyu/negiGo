package views

import (
	"negigo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NegiField struct{}

func (tcv NegiField) CreateNegiField(c *gin.Context) {
	negiField := models.NegiField{}

	if err := c.Bind(&negiField); err != nil {
		c.String(http.StatusBadRequest, "Create NegiField bind err:", err)
		return
	}

	if err := negiField.CreateNegiField(); err != nil {
		c.String(http.StatusBadRequest, "Create NegiField create err:", err)
		return
	}

	c.JSON(http.StatusCreated, negiField)
}

func (tcv NegiField) UpdateNegiField(c *gin.Context) {
	negiField := models.NegiField{}

	if err := c.Bind(&negiField); err != nil {
		c.String(http.StatusBadRequest, "Updates NegiField bind err:", err)
		return
	}

	if err := negiField.UpdatesNegiField(); err != nil {
		c.String(http.StatusBadRequest, "Updates NegiField update err:", err)
		return
	}

	c.JSON(http.StatusOK, negiField)
}

func (tcv NegiField) DeteleNegiField(c *gin.Context) {
	negiField := models.NegiField{}

	if err := c.Bind(&negiField); err != nil {
		c.String(http.StatusBadRequest, "Delete NegiField bind err:", err)
		return
	}

	if err := negiField.DeleteNegiField(); err != nil {
		c.String(http.StatusBadRequest, "Delete NegiField delete err:", err)
		return
	}

}

func (tcv NegiField) GetNegiField(c *gin.Context) {
	negiField := models.NegiField{}

	negifieldid, err := strconv.ParseUint(c.Param("negifieldid"), 10, 32)
	if err != nil {
		return
	}

	negiField.ID = uint(negifieldid)
	if err := negiField.GetOneNegiField(); err != nil {
		return
	}

	c.JSON(http.StatusOK, negiField)
}

func (tcv NegiField) GetAllNeigFields(c *gin.Context) {
	nf := models.NegiField{}
	fs, err := nf.GetAllNegiField()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, fs)
}

//todo return and log every err
