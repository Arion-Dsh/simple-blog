package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type DateTime time.Time

func (dt DateTime) String() string {
	return time.Time(dt).String()
}

func (dt DateTime) IsZero() bool {
	return time.Time(dt).IsZero()
}

func (dt DateTime) Format(layout string) string {
	return time.Time(dt).Format(layout)
}

func (dt *DateTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	t := time.Time(*dt)
	return bson.MarshalValue(t)
}

func (dt *DateTime) UnmarshalBSONValue(t bsontype.Type, value []byte) error {
	if t != bsontype.DateTime {
		return fmt.Errorf("invalid bson value type '%s'", t.String())
	}
	s, _, ok := bsoncore.ReadTime(value)
	if !ok {
		return fmt.Errorf("invalid bson string value")
	}
	*dt = DateTime(s)
	return nil
}

func (dt *DateTime) UnmarshalBind(src interface{}) error {
	t, _ := src.(string)
	ts, err := time.Parse("2006-01-02 15:04:05", t)
	if ts.IsZero() {
		ts = time.Now()
	}
	*dt = DateTime(ts)
	return err
}
