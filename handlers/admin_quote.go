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

//AdminQuotesGet get the Quotes
func AdminQuotesGet(c echo.Context) error {
	cc := c.(*CustomContext)

	p := bindParams(cc)

	quotes := []models.Quote{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := options.Find().SetLimit(p.PageSize).
		SetSkip(p.PageSize * (p.PageNo - 1)).
		SetSort(bson.M{"_id": -1})

	cur, err := db.C(new(models.Quote)).Find(
		ctx,
		bson.M{"is_del": false},
		opt,
	)

	if err == nil {
		cur.All(ctx, &quotes)
	}

	cc.Set("quotes", quotes)

	return cc.Render(http.StatusOK, "admin/base.quotes.html", cc)
}

//AdminQuotePost ...
func AdminQuotePost(c echo.Context) error {
	cc := c.(*CustomContext)
	q := new(models.Quote)

	qid := cc.Param("qid")

	if len(qid) > 0 {
		id, _ := primitive.ObjectIDFromHex(qid)
		q.ID = id
	}

	q.Bind(cc)

	if len(qid) > 0 {
		db.UpdateOne(q, bson.M{"_id": q.ID})
	} else {
		db.InsertOne(q)
	}

	return cc.Redirect(http.StatusSeeOther, cc.URL())
}

//AdminQuoteEdit ...
func AdminQuoteEdit(c echo.Context) error {
	cc := c.(*CustomContext)

	qid := cc.Param("qid")

	id, _ := primitive.ObjectIDFromHex(qid)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	q := new(models.Quote)

	cursor := db.C(q).FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	)

	cursor.Decode(q)

	cc.Set("quote", q)

	return cc.Render(http.StatusOK, "admin/base.quote.html", cc)

}
