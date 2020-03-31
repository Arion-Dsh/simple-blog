package models

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Quote ...
type Quote struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Author  string             `bson:"author" form:"author"`
	Content string             `bson:"content" form:"content"`
	IsDel   bool               `bson:"is_del"`
}

// DBMate ..
func (q *Quote) DBMate() map[string]string {
	return map[string]string{
		"dbName": "katea_blog",
		"cName":  "quotes",
	}
}

// Bind ...
func (q *Quote) Bind(c echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(q, c); err == echo.ErrUnsupportedMediaType {
		return err
	}

	if q.ID.IsZero() {
		q.ID = primitive.NewObjectID()
	}
	return nil
}
