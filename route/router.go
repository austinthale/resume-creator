package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/austinthale/resume-creator/api"
)

func Init() *echo.Echo {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "index.html")
	//e.File("/", "src/github.com/austinthale/resume-creator/index.html") // using to serve a static file that will contain our VueJS client code.
	e.Static("/static", "static")	// using to serve all files contained in public folder, and must be accessed
	// through the /static folder ("localhost:1323/static/resume.css")


	// Routes => handlers
	g := e.Group("/api")
	{
		g.GET("/resumejson", api.DisplayInfo())
		g.POST("/resumejson", api.SaveInfo())
	}

	e.GET("/test", api.DisplayJSON())

	return e
}
