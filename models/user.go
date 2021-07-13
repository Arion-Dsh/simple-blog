package models

import (
	"katea_blog/utils"

	"github.com/arion-dsh/jvmao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ...
type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Email      string             `bson:"email" json:"email"`
	Passwd     string             `bson:"passwd"`
	Name       string             `bson:"name" json:"name"`
	AvatarID   primitive.ObjectID `bson:"avatar,omitempty" json:"-"`
	Avatar     string             `bson:"-"json:"avatar"`
	URLAvatar  string             `bson:"url_avatar" json:"-"`
	Level      int                `bson:"level" json:"-"`
	Active     bool               `bson:"active" json:"-" form:"active"`
	Socials    []*UserSocial      `bson:"socials" json:"-"`
	QRCode     string             `bson:"-" json:"qr_code"`
	UpdateTime DateTime           `bson:"update_time" json:"update_time"`
	CreateTime DateTime           `bson:"create_time" json:"create_time"`
}

// UserSocial ...
type UserSocial struct {
	Source  string `bson:"source" json:"="`
	OpenID  string `bson:"open_id" json:"open_id"`
	UnionID string `bson:"union_id" json:"union_id"`
	Token   string `bson:"token" json:"token"`
}

// DBMate ..
func (u *User) DBMate() map[string]string {
	return map[string]string{
		"dbName": "katea_blog",
		"cName":  "users",
	}
}

func (u *User) signed() *utils.Signed {
	return utils.NewSigned(u.ID.Hex())
}

// ValidPasswd ...
func (u *User) ValidPasswd(p string) bool {
	if len(u.Passwd) == 0 {
		return false
	}
	pw := u.signed().AESDecode(string(u.Passwd))
	if pw == p {
		return true
	}
	return false
}

//SetPasswd ...
func (u *User) SetPasswd(pw string) {
	u.Passwd = u.signed().AESEncode(pw)
}

// Bind ...
func (u *User) Bind(c jvmao.Context) error {

	c.BindForm(u)

	if u.ID.IsZero() {
		u.ID = primitive.NewObjectID()
	}

	pw := c.FormValue("passwd")
	if len(pw) > 0 && !u.ValidPasswd(pw) {

		u.Passwd = u.signed().AESEncode(pw)
	}

	// init CreateTime
	if u.Socials == nil {
		u.Socials = []*UserSocial{}
	}
	return nil
}
