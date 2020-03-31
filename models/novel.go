package models

import (
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Novel ...
type Novel struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" form:"name"`
	Description string             `bson:"descrition" form:"description"`
}

// DBMate ..
func (n *Novel) DBMate() map[string]string {
	return map[string]string{
		"dbName": "katea_blog",
		"cName":  "novels",
	}
}

// Bind ...
func (n *Novel) Bind(c echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(n, c); err == echo.ErrUnsupportedMediaType {
		return err
	}

	if n.ID.IsZero() {
		n.ID = primitive.NewObjectID()
	}
	return nil
}

//Chapter ...
type Chapter struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title" form:"title"`
	CreateTime time.Time          `bson:"create_time" from:"create_time"`
	UpdateTime time.Time          `bson:"update_time"`
	Content    string             `bson:"content" form:"content"`
	Active     bool               `bson:"active" form:"active"`
	Novel      string             `bson:"novel" form:"novel"`
}

// DBMate ..
func (cp *Chapter) DBMate() map[string]string {
	return map[string]string{
		"dbName": "katea_blog",
		"cName":  "chapters",
	}
}

// Bind ...
func (cp *Chapter) Bind(c echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(cp, c); err != nil {
		return err
	}
	ct := c.FormValue("create_time")

	cp.CreateTime, _ = time.Parse("2006-01-02 15:04:05", ct)

	if cp.ID.IsZero() {
		cp.ID = primitive.NewObjectID()
	}

	// init CreateTime
	if cp.CreateTime.IsZero() {
		cp.CreateTime = time.Now()
	}
	cp.UpdateTime = time.Now()
	return nil
}
