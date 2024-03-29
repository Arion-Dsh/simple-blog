package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"katea_blog/db"
	"katea_blog/models"

	"github.com/arion-dsh/jvmao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//AdminArticles get all Articles
func AdminArticles(c jvmao.Context) error {
	cc := c.(*CustomContext)

	p := bindParams(cc)

	articles := []models.Article{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := options.Find().SetLimit(p.PageSize).
		SetSkip(p.PageSize * (p.PageNo - 1)).
		SetSort(bson.M{"_id": -1})

	cur, err := db.C(new(models.Article)).Find(
		ctx,
		bson.M{"is_del": false},
		opt,
	)
	defer cur.Close(ctx)

	if err == nil {
		err = cur.All(ctx, &articles)
	}

	cc.Set("articles", articles)

	return cc.Render(http.StatusOK, "admin/base.articles.html", cc)
}

//AdminArticle ...
func AdminArticle(c jvmao.Context) error {
	cc := c.(*CustomContext)

	id, _ := primitive.ObjectIDFromHex(cc.ParamValue("id"))

	article := new(models.Article)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor := db.C(article).FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	)

	if cursor.Err() == nil {
		cursor.Decode(article)
	}

	cc.Set("article", article)

	return cc.Render(http.StatusOK, "admin/base.article.html", cc)

}

//AdminArticlePost ...
func AdminArticlePost(c jvmao.Context) error {

	cc := c.(*CustomContext)

	art := new(models.Article)

	art.Bind(cc)

	db.InsertOne(art)

	return cc.Redirect(http.StatusSeeOther, cc.Reverse("admin_article", art.ID.Hex()))

}

// AdminArticleEdit ...
func AdminArticleEdit(c jvmao.Context) error {

	cc := c.(*CustomContext)

	idHex := cc.ParamValue("id")
	id, _ := primitive.ObjectIDFromHex(idHex)

	a := new(models.Article)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor := db.C(a).FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	)

	if cursor.Err() == nil {
		cursor.Decode(a)
	}

	if cc.FormValue("is_update_file") != "1" && cc.FormValue("is_del_file") != "1" {
		fmt.Println("----------")

		if err := a.Bind(cc); err != nil {
			log.Print(err.Error())
		}

	}
	//add img
	if cc.FormValue("is_update_file") == "1" {

		file, _ := c.FormFile("file")

		fileID, err := db.SaveFsFile("katea_blog", file)
		if err == nil {
			a.AddImg(fileID.Hex())
		}
		a.UpdateTime = models.DateTime(time.Now())
	}

	//del img
	if cc.FormValue("is_del_file") == "1" {
		imgID := cc.FormValue("img_id")
		err := db.DelFsFile("katea_blog", imgID)
		if err == nil {
			a.DelImg(imgID)
		}
		a.UpdateTime = models.DateTime(time.Now())
	}

	db.UpdateOne(a, bson.M{"_id": id})

	return cc.Redirect(http.StatusSeeOther, cc.Reverse("admin_article", idHex))
}
