package handlers

import (
	"context"
	"net/http"

	"katea_blog/conf"
	"katea_blog/db"
	"katea_blog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/arion-dsh/jvmao"
)

//VerifyAuth ...
func VerifyAuth(next jvmao.HandlerFunc) jvmao.HandlerFunc {
	return func(c jvmao.Context) error {
		_, err := c.Cookie("commander")

		loginURL, _ := conf.Config.String("server.loginURL")
		loginURL = loginURL + "?next=" + c.Request().URL.Path

		if err != nil {
			return c.Redirect(http.StatusSeeOther, loginURL)
		}

		return next(c)

	}
}

// HandlerMiddleware ...
func HandlerMiddleware(next jvmao.HandlerFunc) jvmao.HandlerFunc {
	return func(c jvmao.Context) error {
		cc := &CustomContext{c}
		cookie, err := cc.Cookie("commander")
		if err == nil {
			user := new(models.User)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			uid, _ := primitive.ObjectIDFromHex(cookie.Value)
			cursor := db.C(user).FindOne(ctx, bson.M{"_id": uid}, options.FindOne())
			if err := cursor.Decode(user); err != nil {
				cc.Set("user", user)
			}
		}

		return next(cc)
	}
}
