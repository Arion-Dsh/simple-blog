package models

import (
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Article ...
type Article struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title" form:"title"`
	Slug       string             `bson:"slug" form:"slug"`
	CreateTime time.Time          `bson:"create_time" form:"-"`
	UpdateTime time.Time          `bson:"update_time" form:"-"`
	Content    string             `bson:"content" form:"content"`
	Images     []string           `bson:"images" form:"-"`
	Active     bool               `bson:"active" form:"active"`
	Category   string             `bson:"category" form:"category"`
	IsDel      bool               `bson:"is_del" form:"-"`
}

// DBMate ..
func (a *Article) DBMate() map[string]string {
	return map[string]string{
		"dbName": "katea_blog",
		"cName":  "articles",
	}
}

// Bind ...
func (a *Article) Bind(c echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(a, c); err != nil {
		println("Error in custom", err.Error())
		return err
	}
	ct := c.FormValue("create_time")

	a.CreateTime, _ = time.Parse("2006-01-02 15:04:05", ct)

	if a.Images == nil {
		a.Images = []string{}
	}

	if a.ID.IsZero() {
		a.ID = primitive.NewObjectID()
	}

	// init CreateTime
	if a.CreateTime.IsZero() {
		a.CreateTime = time.Now()
	}
	a.UpdateTime = time.Now()
	return nil
}

//AddImg ...
func (a *Article) AddImg(img string) {

	ok := false
	for _, i := range a.Images {
		if i == img {
			ok = true
			break
		}
	}

	if !ok {
		a.Images = append(a.Images, img)
	}

}

//DelImg ...
func (a *Article) DelImg(img string) {
	for i, v := range a.Images {
		if v == img {
			a.Images = append(a.Images[:i], a.Images[i+1:]...)
			break
		}
	}
}
