package models

import (
	"time"

	"github.com/arion-dsh/jvmao"
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
func (n *Novel) Bind(c jvmao.Context) error {
	if err := c.BindForm(n); err != nil {
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
	CreateTime DateTime           `bson:"create_time" form:"create_time"`
	UpdateTime DateTime           `bson:"update_time"`
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
func (cp *Chapter) Bind(c jvmao.Context) error {

	if err := c.BindForm(cp); err != nil {
		return err
	}

	if cp.ID.IsZero() {
		cp.ID = primitive.NewObjectID()
	}

	cp.UpdateTime = DateTime(time.Now())
	return nil
}
