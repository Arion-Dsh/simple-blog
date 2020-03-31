package handlers

import (
	"strconv"

	"github.com/labstack/echo"
)

type param struct {
	Active   bool   `query:"active, omitempty"`
	PageNo   int64  `query:"page_no"`
	PageSize int64  `query:"page_size"`
	Q        string `query:"q"`
}

func bindParams(c echo.Context) *param {
	p := new(param)
	c.Bind(p)
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
	echo.Context
}

//URL ...
func (c *CustomContext) URL() string {
	return c.Request().URL.Path
}

// PrevURL get the prev page url
func (c *CustomContext) PrevURL() string {
	params := c.QueryParams()
	p, _ := strconv.Atoi(params.Get("page_no"))

	if p <= 1 {
		return c.URL() + "?" + params.Encode()
	}
	params.Set("page_no", strconv.Itoa(p-1))
	return c.URL() + "?" + params.Encode()
}

//HasPrevPage ...
func (c *CustomContext) HasPrevPage() bool {
	p, _ := strconv.Atoi(c.QueryParam("page_no"))

	if p <= 1 {
		return false
	}

	return true
}

//NextURL get the next page url
func (c *CustomContext) NextURL() string {
	params := c.QueryParams()
	p, _ := strconv.Atoi(params.Get("page_no"))

	if p < 1 {
		p = 1
	}
	params.Set("page_no", strconv.Itoa(p+1))
	return c.URL() + "?" + params.Encode()
}
