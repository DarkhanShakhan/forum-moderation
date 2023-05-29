package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/DarkhanShakhan/forum-moderation/internal/config"
	httpServer "github.com/DarkhanShakhan/forum-moderation/internal/ports/http"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/comments"
	commentsS "github.com/DarkhanShakhan/forum-moderation/internal/services/comments"

	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/categories"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/posts"
	categoriesS "github.com/DarkhanShakhan/forum-moderation/internal/services/categories"
	postsS "github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

type Application struct {
	logger     *log.Logger
	db         *sql.DB
	httpServer *http.Server
	config     *config.Config
}

func InitApp() *Application {
	app := &Application{}

	app.loadConfig()
	log.Println("loaded config")

	if err := app.initLogger(); err != nil {
		log.Fatal("init logger: ", err)
	}
	log.Println("logger initialized")

	if err := app.initDB(); err != nil {
		log.Fatal("init db: ", err)
	}
	log.Println("db initialized")

	if err := app.mirgateDB(); err != nil {
		log.Fatal("migrate db: ", err)
	}
	log.Println("db migrated succesfully")

	app.initHTTPServer()
	log.Println("server initialized")

	return app
}

func (a *Application) initLogger() error {
	// TODO: init logger

	return nil
}

func (a *Application) initDB() error {
	var err error
	a.db, err = sql.Open("sqlite3", a.config.SqliteDBName)
	if err != nil {
		return err
	}
	return a.db.Ping()
}

func (a *Application) mirgateDB() error {
	stmts, err := os.ReadFile("./migrations/init.sql")
	if err != nil {
		return err
	}
	_, err = a.db.Exec(string(stmts))
	return err
}

func (a *Application) initHTTPServer() {
	postsRepository := posts.New(a.db)
	categoriesRepository := categories.New(a.db)
	commentsRepository := comments.New(a.db)
	postsService := postsS.New(postsRepository, categoriesRepository)
	categoriesService := categoriesS.New(categoriesRepository)
	commentsService := commentsS.New(commentsRepository)
	a.httpServer = httpServer.NewServer(a.config, postsService, categoriesService, commentsService)
}

func (a *Application) Start() error {
	log.Printf("server is starting on %s\n", a.httpServer.Addr)
	return a.httpServer.ListenAndServe()
}

func (a *Application) CloseDB() error {
	return a.db.Close()
}
func (a *Application) loadConfig() {
	a.config = config.LoadConfig()
}
