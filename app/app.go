package app

import (
	"GoConcurrency-Bootcamp-2022/router"
	"os"
)

func Start() {
	r := router.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
