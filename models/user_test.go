package models

import (
	"context"
	"testing"
	"time"

	"katea_blog/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserToken(t *testing.T) {
	// init config
	// conf.SetCfg("deploy")
	// secretKey, err := conf.Config.String("secretkey")
	// if err != nil {
	// panic("secretkey not in config file")
	// }

	// init db
	db.ConnectMongo("mongodb://localhost:27017")

	user := new(User)
	user.ID = primitive.NewObjectID()
	user.Level = 0
	user.Active = true
	// user.Socials = []*UserSocial{new(UserSocial)}
	user.Name = "3"
	uid, _ := primitive.ObjectIDFromHex("5cfe2386a87e5c5a47120169")

	user.Firends = []primitive.ObjectID{uid}

	data, _ := bson.Marshal(user)

	u := bson.M{}

	bson.Unmarshal(data, &u)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.C(user).InsertOne(ctx, u)

}
