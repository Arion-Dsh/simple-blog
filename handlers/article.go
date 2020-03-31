package handlers

import (
	"context"
	"net/http"

	"katea_blog/db"
	"katea_blog/models"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ArticlesGet get the articles
func ArticlesGet(c echo.Context) error {

	cc := c.(*CustomContext)

	p := bindParams(cc)

	cid := cc.Param("cid")

	articles := []models.Article{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := options.Find().SetLimit(p.PageSize).
		SetSkip(p.PageSize * (p.PageNo - 1)).
		SetSort(bson.M{"create_time": -1})

	cur, err := db.C(new(models.Article)).Find(
		ctx,
		bson.M{"category": cid, "active": true, "is_del": false},
		opt,
	)

	if err == nil {
		cur.All(ctx, &articles)
	}

	cc.Set("articles", articles)
	return c.Render(http.StatusOK, "base.articles.html", cc)

}

//ArticleGet get Article for id
func ArticleGet(c echo.Context) error {
	cc := c.(*CustomContext)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	article := new(models.Article)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor := db.C(article).FindOne(
		ctx,
		bson.M{"_id": id, "category": cc.Param("cid")},
		options.FindOne(),
	)

	if cursor.Err() == nil {
		cursor.Decode(article)
	}

	cc.Set("article", article)

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

	cc.Set("quote", quote)

	return c.Render(http.StatusOK, "base.article.html", cc)
}
