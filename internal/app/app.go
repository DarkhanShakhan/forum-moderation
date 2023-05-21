package app

import (
	"database/sql"
	"log"
	"net/http"

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
	if err := app.initLogger(); err != nil {
		log.Fatal("init logger:", err)
	}
	if err := app.initDB(); err != nil {
		log.Fatal("init db:", err)
	}
	app.initHTTPServer()
	return app
}

func (a *Application) initLogger() error {
	// TODO: init logger
	return nil
}

func (a *Application) initDB() error {
	// TODO: init DB
	return nil
}

func (a *Application) initHTTPServer() {
	postsRepository := posts.New(a.db)
	categoriesRepository := categories.New(a.db)
	postsService := postsS.New(postsRepository, categoriesRepository)
	a.httpServer = httpServer.NewServer(a.config, postsService)
}

func (a *Application) Start() error {
	return a.httpServer.ListenAndServe()
}

func (a *Application) loadConfig() {
	a.config = config.LoadConfig()
}
