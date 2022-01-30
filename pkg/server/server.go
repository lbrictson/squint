package server

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"

	"github.com/lbrictson/squint/web"

	"github.com/labstack/echo/v4"
	"github.com/lbrictson/squint/pkg/store"
)

type Server struct {
	storage *store.Stores
}

type NewServerConfig struct {
	Storage *store.Stores
}

func New(config NewServerConfig) *Server {
	s := Server{storage: config.Storage}
	return &s
}

func (s Server) Run() {
	e := echo.New()
	// Use templates from embedded storage
	t := &Template{
		templates: template.Must(template.ParseFS(web.Assets, "templates/*.tmpl")),
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
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
