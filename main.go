package main

import (
	_ "github.com/lib/pq"
	"os"
	"github.com/austinthale/resume-creator/route"
)

func main() {
	router := route.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	// Start server
	router.Logger.Fatal(router.Start(":"+port))
}