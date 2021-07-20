// main

package main

import (
	"context"
	"embed"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/arion-dsh/jvmao"
	"github.com/arion-dsh/jvmao/middleware"

	"gopkg.in/go-playground/validator.v9"

	"katea_blog/conf"
	"katea_blog/db"
	fs "katea_blog/fsfiles"
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

//go:embed templates/admin/* templates/admin/layouts/*
//go:embed templates/frontend/* templates/frontend/layouts/*
var tmpl embed.FS

func main() {

	// parse the flag
	flag.Parse()
	// init config
	conf.SetCfg(utils.GetExPath()+"/conf/config.yaml", *branch)

	//wx
	// utils.NewWXToken("wxe473a50da2ee4873", "1c7d54d7ddc19be96541973d0d30dee1")

	// db
	db.ConnectMongo("mongodb://localhost:27017")

	// Echo instance
	jm := jvmao.New()

	// Middleware
	jm.Use(middleware.Logger())
	jm.Use(middleware.Recover())

	jm.Static("static", "/static/")

	t := utils.NewTemplate(tmpl)
	t.AddFunc("reverse", jm.Reverse)
	t.AddTmpl("", "templates/frontend")
	t.AddTmpl("admin", "templates/admin")
	jm.SetRenderer(t)

	// e.POST("/auth/wxapp", API.GetWeAppToken)

	// route => fsfile
	jm.POST("update_fsfile", "/files", fs.SaveFsFile)
	jm.GET("get_fsfile", "/files/:fid", fs.GetFsFile)

	//front end
	jm.Use(handlers.HandlerMiddleware)
	jm.GET("home", "/", handlers.Home)
	jm.GET("article", "/:cid/:id", handlers.ArticleGet)
	jm.GET("articles", "/:cid", handlers.ArticlesGet)
	jm.GET("novel", "/novels/:novel", handlers.Novel)
	jm.GET("chapter", "/novels/:novel/:id", handlers.Chapter)

	//admin

	loginURL, _ := conf.Config.String("server.loginURL")
	jm.GET("login", loginURL, handlers.AdminLoginGet)
	jm.POST("login_post", loginURL, handlers.AdminLoginPost)

	admin := jm.Group("/admin")
	admin.Use(handlers.VerifyAuth)
	admin.GET("admin_home", "", handlers.AdminHome)
	admin.POST("", "", handlers.AdminQuotePost)
	admin.GET("admin_logout", "/logout", handlers.AdminLogOut)
	admin.GET("admin_quotes", "/quotes", handlers.AdminQuotesGet)
	admin.POST("admin_quotes_add", "/quotes", handlers.AdminQuotePost)
	admin.GET("admin_quote", "/quotes/:qid", handlers.AdminQuoteEdit)
	admin.POST("admin_quote_post", "/quotes/:qid", handlers.AdminQuotePost)
	admin.GET("admin_articles", "/articles", handlers.AdminArticles)
	admin.GET("admin_article", "/articles/:id", handlers.AdminArticle)
	admin.POST("admin_article_edit", "/articles/:id", handlers.AdminArticleEdit)
	admin.GET("admin_article_new", "/articles/new", handlers.AdminArticle)
	admin.POST("admin_article_new_post", "/articles/new", handlers.AdminArticlePost)

	admin.GET("admin_novels", "/novels", handlers.AdminNovels)
	admin.POST("", "/novels", handlers.AdminNovelEdit)
	admin.GET("admin_novel", "/novels/:id", handlers.AdminNovel)
	admin.POST("admin_novel_edit", "/novels/:id", handlers.AdminNovelEdit)
	admin.GET("admin_chapters", "/novels/:novel/chapters", handlers.AdminChapters)
	admin.GET("admin_chapter_new", "/novels/:novel/chapters/new", handlers.AdminChapter)
	admin.POST("", "/novels/:novel/chapters/new", handlers.AdminChapterEdit)
	admin.GET("admin_chapter", "/novels/:novel/chapters/:id", handlers.AdminChapter)
	admin.POST("admin_chapter_edit", "/novels/:novel/chapters/:id", handlers.AdminChapterEdit)

	// Start server
	port, _ := conf.Config.String("server.port")

	go func() {
		if err := jm.Start(port); err != nil {
			jm.Logger.Fatal(err.Error())
		}
	}()

	/* // Wait for interrupt signal to gracefully shutdown the server with */
	/* // a timeout of 5 seconds. */
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := jm.Shutdown(ctx); err != nil {
		jm.Logger.Fatal(err.Error())

	}
}
