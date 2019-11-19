package main

import (
	"fizzbuzz/app"
	"fizzbuzz/handler"

	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	go func() {
		log.Printf("Debug listener started on port %d", cfg.DebugPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.DebugPort), nil); err != nil {
			log.Printf("Debug listener failed: %s", err.Error())
		}
	}()

	log.Printf("started v2 (pid=%d port=%d)", os.Getpid(), cfg.Port)

	app, err := app.NewApp(*cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Routes
	handler.RegisterRoutes(app)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), app)
	if err != nil {
		panic(err.Error())
	}
}
