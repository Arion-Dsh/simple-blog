package utils

import (
	"html/template"
	"io"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/russross/blackfriday"
)

//Template ...
type Template struct {
	templates map[string]*template.Template
	funcs     template.FuncMap
}

//Render the template name will be "layout.templateName" eg. "base.books.html"
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	//need format layout and the templateName
	//"path/layout.templateName"
	fullpath := strings.SplitN(name, "/", 2)
	tmplPath := fullpath[0]
	if len(fullpath) == 2 {
		tmplPath = fullpath[1]
	}
	names := strings.SplitN(tmplPath, ".", 2)

	if len(names) != 2 {
		panic("The template " + name + " is wrong format.")
	}

	tmpName := names[1]
	if len(fullpath) == 2 {
		tmpName = fullpath[0] + "/" + tmpName
	}

	tmpl, ok := t.templates[tmpName]

	if !ok {
		panic("The template " + name + " does not exist.")
	}
	return tmpl.ExecuteTemplate(w, names[0], data)
}

// AddTemp ...
func (t *Template) AddTemp(path, prefix string) error {

	if t.templates == nil {
		t.templates = make(map[string]*template.Template)
	}
	layouts, err := filepath.Glob(path + "/layouts/*.html")
	if err != nil {
		return err
	}

	includes, err := filepath.Glob(path + "/*.html")
	if err != nil {
		return err
	}

	for _, include := range includes {
		files := append(layouts, include)
		n := filepath.Base(include)
		if len(prefix) > 0 {
			n = prefix + "/" + n
		}
		t.templates[n] = template.Must(template.New(n).Funcs(t.funcs).ParseFiles(files...))
	}
	return nil

}

//AddFunc ...
func (t *Template) AddFunc(name string, f interface{}) {
	if t.funcs == nil {
		t.funcs = template.FuncMap{
			"safe":      safeHTML,
			"date_time": dateTimeHTML,
			"markdown":  markdownHTML,
		}
	}
	t.funcs[name] = f
}

func parseImg(str string) string {
	sr := strings.TrimSuffix(strings.TrimPrefix(str, "[["), "]]")
	rs := strings.Split(sr, "|")
	r := "<div class=\"cont-img pos-r\">" +
		"<div class=\"cont-img-dec pos-a\"><p>" +
		rs[0] + "</p></div><img src=\"" +
		rs[2] + "\" alt=\"" +
		rs[1] + "\" class=\"pure-img bdr-s\"></div>"

	return r
}

func markdownHTML(s string) template.HTML {
	re := regexp.MustCompile(`\[\[(.+?)\]\]`)
	ns := re.ReplaceAllStringFunc(s, parseImg)
	t := blackfriday.MarkdownCommon([]byte(ns))
	return template.HTML(string(t))
}

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func dateTimeHTML(t *time.Time) template.HTML {
	return template.HTML(t.Format("02 Jan 2006"))
}
