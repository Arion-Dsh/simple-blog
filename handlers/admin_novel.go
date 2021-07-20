package handlers

import (
	"context"
	"net/http"
	"net/url"

	"katea_blog/db"
	"katea_blog/models"

	"github.com/arion-dsh/jvmao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//AdminNovels get the Novel
func AdminNovels(c jvmao.Context) error {
	cc := c.(*CustomContext)

	p := bindParams(cc)

	novels := []models.Novel{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := options.Find().SetLimit(p.PageSize).
		SetSkip(p.PageSize * (p.PageNo - 1)).
		SetSort(bson.M{"_id": -1})

	cur, err := db.C(new(models.Novel)).Find(
		ctx,
		bson.M{},
		opt,
	)

	if err == nil {
		cur.All(ctx, &novels)
	}

	cc.Set("novels", novels)

	return cc.Render(http.StatusOK, "admin/base.novels.html", cc)
}

//AdminNovelEdit ...
func AdminNovelEdit(c jvmao.Context) error {
	cc := c.(*CustomContext)

	novel := new(models.Novel)

	idHex := cc.ParamValue("id")
	id, _ := primitive.ObjectIDFromHex(idHex)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor := db.C(novel).FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	)

	cursor.Decode(novel)

	oNovelName := novel.Name

	novel.Bind(cc)

	var err error
	if len(idHex) > 0 {
		_, err = db.UpdateOne(novel, bson.M{"_id": novel.ID})
	} else {
		_, err = db.InsertOne(novel)
	}
	url := cc.Request().URL.String()
	if err != nil {
		return cc.Redirect(http.StatusSeeOther, url)
	}

	if len(oNovelName) > 0 && oNovelName != novel.Name {
		db.C(new(models.Chapter)).UpdateMany(
			ctx,
			bson.M{"novel": oNovelName},
			bson.M{"$set": bson.M{"novel": novel.Name}},
		)
	}

	return cc.Redirect(http.StatusSeeOther, url)
}

//AdminNovel ...
func AdminNovel(c jvmao.Context) error {
	cc := c.(*CustomContext)

	idHex := cc.ParamValue("id")

	id, _ := primitive.ObjectIDFromHex(idHex)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	novel := new(models.Novel)

	cursor := db.C(novel).FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	)

	cursor.Decode(novel)

	cc.Set("novel", novel)

	return cc.Render(http.StatusOK, "admin/base.novel.html", cc)

}

//AdminChapters ...
func AdminChapters(c jvmao.Context) error {
	cc := c.(*CustomContext)

	p := bindParams(cc)

	n := cc.ParamValue("novel")
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

	opt := options.Find().SetLimit(p.PageSize).
		SetSkip(p.PageSize * (p.PageNo - 1)).
		SetSort(bson.M{"_id": -1})

	cur, err := db.C(new(models.Chapter)).Find(
		ctx,
		bson.M{"novel": nName},
		opt,
	)

	if err == nil {
		cur.All(ctx, &chapters)
	}

	cc.Set("chapters", chapters)

	return cc.Render(http.StatusOK, "admin/base.chapters.html", cc)
}

//AdminChapter ...
func AdminChapter(c jvmao.Context) error {
	cc := c.(*CustomContext)

	idHex := cc.ParamValue("id")

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

	return cc.Render(http.StatusOK, "admin/base.chapter.html", cc)

}

//AdminChapterEdit ...
func AdminChapterEdit(c jvmao.Context) error {
	cc := c.(*CustomContext)

	n := cc.ParamValue("novel")
	nName, _ := url.PathUnescape(n)

	idHex := cc.ParamValue("id")
	id, _ := primitive.ObjectIDFromHex(idHex)

	chapter := new(models.Chapter)
	if !id.IsZero() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cursor := db.C(chapter).FindOne(
			ctx,
			bson.M{"_id": id},
			options.FindOne(),
		)

		cursor.Decode(chapter)

	}
	chapter.Bind(cc)

	chapter.Novel = nName

	if !id.IsZero() {
		db.UpdateOne(chapter, bson.M{"_id": chapter.ID})
	} else {
		db.InsertOne(chapter)
	}

	return cc.Redirect(http.StatusSeeOther, cc.Reverse("admin_chapters", nName))
}
