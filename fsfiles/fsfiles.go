package handlers

import (
	"fmt"
	"net/http"

	"katea_blog/db"

	"github.com/arion-dsh/jvmao"
)

// SaveFsFile save file to mongodb
func SaveFsFile(c jvmao.Context) error {

	var err error

	file, err := c.FormFile("file")
	if err != nil {
		return c.Error(err)
	}

	fileID, err := db.SaveFsFile("katea_blog", file)
	if err != nil {
		return c.Error(err)
	}

	return c.Json(http.StatusOK, map[string]interface{}{"file_id": fileID})

}

// GetFsFile get fs file
func GetFsFile(c jvmao.Context) error {

	fsfile, err := db.GetFsFile("katea_blog", c.ParamValue("fid"))
	if err != nil {
		return c.Error(err)
	}

	c.Response().Header().Set("content-disposition", fmt.Sprintf("filename=%s", fsfile.Filename))

	return c.Blob(http.StatusOK, fsfile.Metadata["content-type"], fsfile.Bytes())

}
