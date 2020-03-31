package handlers

import (
	"context"
	"net/http"

	"katea_blog/db"
	"katea_blog/models"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//AdminHome ...
func AdminHome(c echo.Context) error {

	cc := c.(*CustomContext)

	articles := []models.Article{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := options.Find().SetLimit(5).
		SetSort(bson.M{"_id": -1})

	cur, err := db.C(new(models.Article)).Find(
		ctx,
		bson.M{"is_del": false},
		opt,
	)

	if err == nil {
		cur.All(ctx, &articles)
	}

	cc.Set("articles", articles)

	return cc.Render(http.StatusOK, "admin/base.home.html", cc)

}
