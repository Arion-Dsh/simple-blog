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

//Home ...
func Home(c echo.Context) error {

	cc := c.(*CustomContext)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	find := options.Find().SetLimit(5).SetSort(bson.M{"create_time": -1})
	cur, _ := db.C(new(models.Article)).Find(
		ctx,
		bson.M{"active": true, "is_del": false, "category": "zh-hans"},
		find,
	)

	articles := []models.Article{}

	cur.All(ctx, &articles)

	page := new(models.Article)

	cursor := db.C(page).FindOne(
		ctx,
		bson.M{"slug": "home"},
		options.FindOne(),
	)

	if cursor.Err() == nil {
		cursor.Decode(page)
	}

	quote := new(models.Quote)
	cr, _ := db.C(quote).Aggregate(
		ctx,
		bson.A{
			bson.M{"$sample": bson.M{"size": 1}},
		},
		options.Aggregate(),
	)

	if cr.Next(ctx) {
		cr.Decode(quote)
	}

	cc.Set("page", page)
	cc.Set("articles", articles)
	cc.Set("quote", quote)

	return c.Render(http.StatusOK, "base.home.html", cc)

}
