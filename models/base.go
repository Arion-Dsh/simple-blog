package models

import "time"

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalParam(src string) error {
	ts, err := time.Parse("2006-01-02 15:04:05", src)
	*dt = DateTime{ts}
	return err

}
