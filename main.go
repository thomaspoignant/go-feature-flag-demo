package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"html/template"
	"io"
	"net/http"
	"time"
)

var users = make(map[string]ffuser.User, 2500)

func main() {
	configFile := flag.String("configFile", "./flags.yaml", "flags.yaml")
	flag.Parse()

	_ = ffclient.Init(ffclient.Config{
		PollingInterval: 1 * time.Second,
		Context:         context.Background(),
		Retriever: &ffclient.FileRetriever{
			Path: *configFile,
		},
	})

	// Instantiate a template registry and register all html files inside the view folder

	e := echo.New()
	e.HideBanner = true
	e.Static("/js", "js")
	e.Static("/css", "css")
	e.Renderer = &TemplateRegistry{templates: template.Must(template.ParseGlob("view/*.html"))}

	// init ids
	for i := 0; i < 2500; i++ {
		id := uuid.New()
		u := ffuser.NewUser(id.String())
		users[fmt.Sprintf("user%d", i)] = u
	}

	e.GET("/", apiHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func apiHandler(c echo.Context) error {
	mapToRender := make(map[string]string, 2500)
	for k, user := range users {
		color, _ := ffclient.StringVariation("color-box", user, "grey")
		mapToRender[k] = color
	}
	return c.Render(http.StatusOK, "template.html", mapToRender)
}
