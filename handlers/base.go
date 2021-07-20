package handlers

import (
	"strconv"

	"github.com/arion-dsh/jvmao"
)

type param struct {
	Active   bool   `query:"active, omitempty"`
	PageNo   int64  `query:"page_no"`
	PageSize int64  `query:"page_size"`
	Q        string `query:"q"`
}

func bindParams(c jvmao.Context) *param {
	p := new(param)
	// c.Bind(p)
	if p.PageNo < 1 {
		p.PageNo = 1
	}

	if p.PageSize < 1 || p.PageSize > 100 {
		p.PageSize = 10
	}

	return p
}

// CustomContext ...
type CustomContext struct {
	jvmao.Context
}

//URL ...
func (c *CustomContext) path() string {
	return c.Request().URL.Path
}

// PrevURL get the prev page url
func (c *CustomContext) PrevURL() string {
	q := c.Query()
	p, _ := strconv.Atoi(q.Get("page_no"))

	if p <= 1 {
		return c.path() + "?" + q.Encode()

	}
	q.Set("page_no", strconv.Itoa(p-1))
	return c.path() + "?" + q.Encode()
}

//HasPrevPage ...
func (c *CustomContext) HasPrevPage() bool {
	p, _ := strconv.Atoi(c.Query().Get("page_no"))

	if p <= 1 {
		return false
	}

	return true
}

//NextURL get the next page url
func (c *CustomContext) NextURL() string {
	q := c.Query()
	p, _ := strconv.Atoi(q.Get("page_no"))

	if p < 1 {
		p = 1
	}
	q.Set("page_no", strconv.Itoa(p+1))
	return c.path() + "?" + q.Encode()
}
