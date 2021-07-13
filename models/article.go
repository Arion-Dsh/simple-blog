package models

import (
	"time"

	"github.com/arion-dsh/jvmao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Article ...
type Article struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title" form:"title"`
	Slug       string             `bson:"slug" form:"slug"`
	CreateTime DateTime           `bson:"create_time" form:"create_time"`
	UpdateTime DateTime           `bson:"update_time" form:"-"`
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
func (a *Article) Bind(c jvmao.Context) error {
	if err := c.BindForm(a); err != nil {
		return err
	}

	if a.Images == nil {
		a.Images = []string{}
	}

	if a.ID.IsZero() {
		a.ID = primitive.NewObjectID()
	}

	a.UpdateTime = DateTime(time.Now())

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
