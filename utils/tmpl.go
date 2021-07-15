package utils

import (
	"embed"
	"html/template"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/arion-dsh/jvmao"
	"github.com/russross/blackfriday"
)

func NewTemplate(fs embed.FS) Template {
	return &tmpl{
		fs:        fs,
		templates: make(map[string]*template.Template),
		funcs: template.FuncMap{
			"safe":      safeHTML,
			"date_time": dateTimeHTML,
			"markdown":  markdownHTML,
		},
	}
}

type Template interface {
	Render(w io.Writer, name string, data interface{}, c jvmao.Context) error
	AddTmpl(prefix, templatesDir string) error
	AddFunc(name string, f interface{})
}

//Template ...
type tmpl struct {
	fs        embed.FS
	templates map[string]*template.Template
	funcs     template.FuncMap
}

//Render the template name will be "layout.templateName" eg. "base.books.html"
func (t *tmpl) Render(w io.Writer, name string, data interface{}, c jvmao.Context) error {
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
	return tmpl.Funcs(t.funcs).ExecuteTemplate(w, names[0], data)
}

// AddTmpl add template
func (t *tmpl) AddTmpl(prefix, templatesDir string) error {

	tmplFiles, err := t.fs.ReadDir(templatesDir)
	if err != nil {
		return err
	}
	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}
		n := tmpl.Name()
		if len(prefix) > 0 {
			n = prefix + "/" + n
		}
		pt := template.New(n).Funcs(t.funcs)
		pt, err := pt.ParseFS(t.fs, templatesDir+"/"+tmpl.Name(), templatesDir+"/layouts/*")
		if err != nil {
			return err
		}
		t.templates[n] = pt
	}
	return nil

}

//AddFunc must before Addtmpl
func (t *tmpl) AddFunc(name string, f interface{}) {
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
