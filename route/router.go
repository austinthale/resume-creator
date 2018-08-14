package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/austinthale/resume-creator/api"
	"github.com/austinthale/resume-creator/handler"
	myMw "github.com/austinthale/resume-creator/middleware"
	"github.com/austinthale/resume-creator/db"
)

func Init() *echo.Echo {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Unauthenticated route
	e.GET("/", api.Accessible())

	// Custom Middleware
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	e.Use(myMw.TransactionHandler(db.Init()))

	e.File("/", "index.html")
	//e.File("/", "src/github.com/austinthale/resume-creator/index.html") // using to serve a static file that will contain our VueJS client code.
	e.Static("/static", "static")	// using to serve all files contained in public folder, and must be accessed
	// through the /static folder ("localhost:1323/static/resume.css")


	// Routes => handlers
	g := e.Group("/api")
	{
		g.GET("/resumejson/:id", api.GetResumeInfo())
		g.POST("/resumejson/:id", api.SetResumeInfo())
		g.POST("/login", api.Login())
		//g.Use(middleware.JWT([]byte("secret")))
		g.GET("", api.Restricted())
	}
	return e
}
