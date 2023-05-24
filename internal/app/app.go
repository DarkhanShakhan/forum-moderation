package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/DarkhanShakhan/forum-moderation/internal/config"
	httpServer "github.com/DarkhanShakhan/forum-moderation/internal/ports/http"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/categories"
	"github.com/DarkhanShakhan/forum-moderation/internal/repositories/posts"
	postsS "github.com/DarkhanShakhan/forum-moderation/internal/services/posts"
)

type Application struct {
	logger     *log.Logger
	db         *sql.DB
	httpServer *http.Server
	config     *config.Config
}

func InitApp() *Application {
	// TODO: read envs
	// TODO: init repo
	app := &Application{}
	app.loadConfig()
	log.Println("loaded config")
	if err := app.initLogger(); err != nil {
		log.Fatal("init logger:", err)
	}
	log.Println("logger initialized")
	if err := app.initDB(); err != nil {
		log.Fatal("init db:", err)
	}
	log.Println("db initialized")
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
	return nil
}

func (a *Application) initHTTPServer() {
	postsRepository := posts.New(a.db)
	categoriesRepository := categories.New(a.db)
	postsService := postsS.New(postsRepository, categoriesRepository)
	a.httpServer = httpServer.NewServer(a.config, postsService)
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
