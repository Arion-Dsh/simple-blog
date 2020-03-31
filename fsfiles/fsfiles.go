package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"katea_blog/db"
)

// SaveFsFile save file to mongodb
func SaveFsFile(c echo.Context) error {

	var err error

	file, err := c.FormFile("file")
	if err != nil {
		log.Panic(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fileID, _ := db.SaveFsFile("katea_blog", file)

	return c.JSON(http.StatusOK, map[string]interface{}{"file_id": fileID})

}

// GetFsFile get fs file
func GetFsFile(c echo.Context) error {

	log.Print(c.Param("fid"))
	fsfile, err := db.GetFsFile("katea_blog", c.Param("fid"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "404")
	}

	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("filename=%s", fsfile.Filename))
	c.Logger().Warn(fsfile.Metadata)

	return c.Blob(http.StatusOK, fsfile.Metadata["content-type"], fsfile.Bytes())

}
