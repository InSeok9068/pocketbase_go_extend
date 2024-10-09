package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed views
var viewsFS embed.FS

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/public/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		e.Router.GET("/", func(c echo.Context) error {
			f, err := viewsFS.Open("views/index.html")
			if err != nil {
				return err
			}
			defer f.Close()

			html, err := io.ReadAll(f) // ioutil.ReadAll 대신 io.ReadAll 사용
			if err != nil {
				return err
			}

			return c.HTML(http.StatusOK, string(html))
		})

		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")
			log.Println(name)
			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
