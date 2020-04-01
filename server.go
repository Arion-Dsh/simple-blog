// main

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	// need init utils first
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	"katea_blog/conf"
	"katea_blog/db"
	fsfileHandler "katea_blog/fsfiles"
	"katea_blog/handlers"
	"katea_blog/utils"
)

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

//Validate .....
func (c *CustomValidator) Validate(i interface{}) error {
	return c.validator.Struct(i)
}

var branch = flag.String("branch", "deploy", "the branch of project")

func main() {

	// parse the flag
	flag.Parse()
	// init config
	conf.SetCfg(utils.GetExPath()+"/conf", *branch)

	//wx
	// utils.NewWXToken("wxe473a50da2ee4873", "1c7d54d7ddc19be96541973d0d30dee1")

	// db
	db.ConnectMongo("mongodb://localhost:27017")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")
	e.File("/favicon.ico", "images/favicon.ico")

	// csrf
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
		CookieName:  "_csrf",
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	//template
	tempPath := utils.GetExPath() + "/templates"
	t := new(utils.Template)
	t.AddFunc("reverse", e.Reverse)
	t.AddTemp(tempPath+"/admin", "admin")
	t.AddTemp(tempPath+"/frontend", "")

	e.Renderer = t

	// e.POST("/auth/wxapp", API.GetWeAppToken)

	// route => fsfile
	e.POST("/files", fsfileHandler.SaveFsFile).Name = "update_fsfile"
	e.GET("/files/:fid", fsfileHandler.GetFsFile).Name = "get_fsfile"

	//front end
	e.Use(handlers.HandlerMiddleware)
	e.GET("/", handlers.Home)
	e.GET("/:cid/:id", handlers.ArticleGet).Name = "article"
	e.GET("/:cid", handlers.ArticlesGet).Name = "articles"
	e.GET("/novels/:novel", handlers.Novel).Name = "novel"
	e.GET("/novels/:novel/:id", handlers.Chapter).Name = "chapter"

	//admin

	loginURL, _ := conf.Config.String("server.loginURL")
	e.GET(loginURL, handlers.AdminLoginGet)
	e.POST(loginURL, handlers.AdminLoginPost)

	admin := e.Group("/admin")
	admin.Use(handlers.VerifyAuth)
	admin.GET("", handlers.AdminHome).Name = "admin_home"
	admin.POST("", handlers.AdminQuotePost)
	admin.GET("/logout", handlers.AdminLogOut).Name = "admin_logout"
	admin.GET("/quotes", handlers.AdminQuotesGet).Name = "admin_quotes"
	admin.POST("/quotes", handlers.AdminQuotePost).Name = "admin_quotes_add"
	admin.GET("/quotes/:qid", handlers.AdminQuoteEdit).Name = "admin_quote"
	admin.POST("/quotes/:qid", handlers.AdminQuotePost).Name = "admin_quote_post"
	admin.GET("/articles", handlers.AdminArticles).Name = "admin_articles"
	admin.GET("/articles/:id", handlers.AdminArticle).Name = "admin_article"
	admin.POST("/articles/:id", handlers.AdminArticleEdit).Name = "admin_article_edit"
	admin.GET("/articles/new", handlers.AdminArticle).Name = "admin_article_new"
	admin.POST("/articles/new", handlers.AdminArticlePost).Name = "admin_article_new_post"

	admin.GET("/novels", handlers.AdminNovels).Name = "admin_novels"
	admin.POST("/novels", handlers.AdminNovelEdit)
	admin.GET("/novels/:id", handlers.AdminNovel).Name = "admin_novel"
	admin.POST("/novels/:id", handlers.AdminNovelEdit)
	admin.GET("/novels/:novel/chapters", handlers.AdminChapters).Name = "admin_chapters"
	admin.GET("/novels/:novel/chapters/new", handlers.AdminChapter).Name = "admin_chapter_new"
	admin.POST("/novels/:novel/chapters/new", handlers.AdminChapterEdit)
	admin.GET("/novels/:novel/chapters/:id", handlers.AdminChapter).Name = "admin_chapter"
	admin.POST("/novels/:novel/chapters/:id", handlers.AdminChapterEdit).Name = "admin_chapter_edit"
	//debug
	debug, _ := conf.Config.Bool("server.debug")
	if debug {
		e.Debug = true
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		}))

	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		}))

	}
	// Start server
	port, _ := conf.Config.String("server.port")
	go func() {
		if err := e.Start(port); err != nil {
			e.Logger.Fatal(err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)

	}
}
