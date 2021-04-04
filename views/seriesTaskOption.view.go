package views

import (
	"negigo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NegiSeriesTaskOption struct{}

func (stov NegiSeriesTaskOption) CreateSeriesTaskOption(c *gin.Context) {
	seriesTaskOption := models.SeriesTaskOption{}

	if err := c.Bind(&seriesTaskOption); err != nil {
		c.String(http.StatusBadRequest, "Create task title option bind err : %v", err)
		return
	}

	if err := seriesTaskOption.CreateSeriesTaskOption(); err != nil {
		c.JSON(http.StatusInternalServerError, "Create task title option db err : "+err.Error())
	}

	c.JSON(http.StatusCreated, seriesTaskOption)
}

func (stov NegiSeriesTaskOption) UpdateseriesTaskOption(c *gin.Context) {
	seriesTaskOption := models.SeriesTaskOption{}

	if err := c.Bind(&seriesTaskOption); err != nil {
		c.String(http.StatusBadRequest, "Update Negi task title option bind err : %v", err)
		return
	}

	if err := seriesTaskOption.UpdateSeriesTaskOption(); err != nil {
		c.String(http.StatusInternalServerError, "Update Negi task title option db err : %v", err)
		return
	}

	c.JSON(http.StatusOK, seriesTaskOption)
}

func (stov NegiSeriesTaskOption) DeleteseriesTaskOption(c *gin.Context) {
	seriesTaskOption := models.SeriesTaskOption{}

	negiseriesTaskOptionid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, "Detele task title opiton no id param : %v", err)
		return
	}

	seriesTaskOption.ID = uint(negiseriesTaskOptionid)

	if err := seriesTaskOption.DeleteSeriesTaskOption(); err != nil {
		c.String(http.StatusInternalServerError, "Delete task title option db err : %v", err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func (stov NegiSeriesTaskOption) GetAllseriesTaskOption(c *gin.Context) {
	seriesTaskOption := models.SeriesTaskOption{}

	titleoptions, err := seriesTaskOption.GetAllSeriesTaskOptions()
	if err != nil {
		c.String(http.StatusInternalServerError, "Get all task title option err : %v", err)
		return
	}

	c.JSON(http.StatusOK, titleoptions)
}
