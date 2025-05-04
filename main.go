package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/kaczmarekdaniel/go-project/internal/app"
	"github.com/kaczmarekdaniel/go-project/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("running ...")

	defer app.DB.Close()
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
