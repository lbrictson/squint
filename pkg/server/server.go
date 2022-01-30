package server

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/lbrictson/squint/web"

	"github.com/labstack/echo/v4"
	"github.com/lbrictson/squint/pkg/store"
)

type Server struct {
	storage *store.Stores
	log     *logrus.Logger
}

type NewServerConfig struct {
	Storage *store.Stores
	Log     *logrus.Logger
}

func New(config NewServerConfig) *Server {
	s := Server{
		storage: config.Storage,
		log:     config.Log,
	}
	return &s
}

func (s Server) Run() {
	e := echo.New()
	// Use templates from embedded storage
	t := &Template{
		templates: template.Must(template.ParseFS(web.Assets, "templates/*.tmpl")),
		log:       s.log,
	}
	e.Renderer = t
	embeddedFiles := web.Assets
	// Just the static directory
	fSys, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		panic(err)
	}
	// Serve all files in the directory
	assetHandler := http.FileServer(http.FS(fSys))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	frontend := e.Group("")
	frontend.Use(s.requestID)
	admin := e.Group("/admin")
	admin.Use(s.requestID)
	admin.GET("", s.adminHomeView)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", "5882")))
}

type Template struct {
	templates *template.Template
	log       *logrus.Logger
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		t.log.Errorf("unable to render template %v %v", name, err)
	}
	return err
}
