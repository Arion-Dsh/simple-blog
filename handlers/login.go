package handlers

import (
	"context"
	"net/http"
	"time"

	"katea_blog/db"
	"katea_blog/models"

	"github.com/arion-dsh/jvmao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//AdminLoginGet ...
func AdminLoginGet(c jvmao.Context) error {
	cc := c.(*CustomContext)

	return c.Render(http.StatusOK, "admin/no_header_base.login.html", cc)
}

//AdminLoginPost ...
func AdminLoginPost(c jvmao.Context) error {
	cc := c.(*CustomContext)

	email := c.FormValue("email")
	passwd := c.FormValue("passwd")

	msg := make(map[string]string)
	msg["err"] = "email 或 密码错误"

	if len(email) < 1 || len(passwd) < 1 {
		cc.Set("msg", msg)
		return c.Render(http.StatusOK, "admin/no_header_base.login.html", cc)
	}

	u := new(models.User)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	count, err := db.C(u).CountDocuments(
		ctx,
		bson.M{},
		options.Count(),
	)

	if count == 0 && err == nil {
		u.Email = email
		u.ID = primitive.NewObjectID()
		u.SetPasswd(passwd)

		db.InsertOne(u)
	}

	cursor := db.C(u).FindOne(
		ctx,
		bson.M{"email": email},
		options.FindOne(),
	)

	if cursor.Err() == nil {
		cursor.Decode(u)
	} else {

		cc.Set("msg", msg)

		return c.Render(http.StatusOK, "admin/no_header_base.login.html", cc)
	}

	if !u.ValidPasswd(passwd) {

		cc.Set("msg", msg)

		return c.Render(http.StatusOK, "admin/no_header_base.login.html", cc)

	}

	cookie := new(http.Cookie)
	cookie.Name = "commander"
	cookie.Value = u.ID.Hex()
	t := time.Now()
	cookie.Expires = t.Add(3 * time.Hour)
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, c.Reverse("admin_home"))

}

//AdminLogOut ...
func AdminLogOut(c jvmao.Context) error {

	next := c.QueryValue("next")
	if len(next) == 0 {
		next = "/"
	}
	cookie, err := c.Cookie("commander")

	if err != nil {
		return c.Redirect(http.StatusSeeOther, next)
	}
	cookie.Expires = time.Now()
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, next)
}
