package handlers

import (
	"context"
	"katea_blog/db"
	"katea_blog/models"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Novel ...
func Novel(c echo.Context) error {

	cc := c.(*CustomContext)

	n := cc.Param("novel")
	nName, _ := url.PathUnescape(n)

	chapters := []models.Chapter{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	novel := new(models.Novel)
	cursor := db.C(novel).FindOne(
		ctx,
		bson.M{"name": nName},
		options.FindOne(),
	)
	cursor.Decode(novel)

	cc.Set("novel", novel)

	cur, err := db.C(new(models.Chapter)).Find(
		ctx,
		bson.M{"novel": nName, "active": true},
		options.Find().SetSort(bson.M{"_id": 1}),
	)

	if err == nil {
		cur.All(ctx, &chapters)
	}

	cc.Set("chapters", chapters)

	return cc.Render(http.StatusOK, "base.novel.html", cc)
}

//Chapter ...
func Chapter(c echo.Context) error {
	cc := c.(*CustomContext)

	idHex := cc.Param("id")
	id, _ := primitive.ObjectIDFromHex(idHex)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chapter := new(models.Chapter)

	cursor := db.C(chapter).FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	)

	cursor.Decode(chapter)

	cc.Set("chapter", chapter)

	n := cc.Param("novel")
	nName, _ := url.PathUnescape(n)

	prevChapter := new(models.Chapter)
	prevCur, _ := db.C(prevChapter).Find(
		ctx,
		bson.M{"active": true, "novel": nName, "_id": bson.M{"$lt": chapter.ID}},
		options.Find().SetSort(bson.M{"_id": -1}).SetLimit(1),
	)

	if prevCur.Next(ctx) {
		prevCur.Decode(prevChapter)
	}

	cc.Set("prev_chapter", prevChapter)

	nextChapter := new(models.Chapter)
	nextCur, _ := db.C(chapter).Find(
		ctx,
		bson.M{"active": true, "novel": nName, "_id": bson.M{"$gt": chapter.ID}},
		options.Find().SetSort(bson.M{"_id": 1}).SetLimit(1),
	)

	if nextCur.Next(ctx) {
		nextCur.Decode(nextChapter)
	}

	cc.Set("next_chapter", nextChapter)

	// random quote
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

	return cc.Render(http.StatusOK, "base.chapter.html", cc)

}
