package main

import (
	"log"
	"os"

	"tempestas/internal/config"
	"tempestas/internal/middleware"
	"tempestas/internal/server"
	"tempestas/pkg/database"
	"tempestas/pkg/weather"
)

func main() {
    conf, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    db, err := database.NewRedis(conf.Redis)
    if err != nil {
        log.Fatal(err)
    }

    apiConf := conf.WeatherAPI

    logger := middleware.NewLogger(log.New(os.Stdout, "", log.LstdFlags))

    service := weather.NewHTTP(db, apiConf.WeatherApiKey)
    cityTempHandler := logger(weather.CityTempHandler(service))

    router := server.NewRouter()
    router.AddRoute("/city/temp", cityTempHandler.ServeHTTP)

    err = server.ListenAndServe(conf.Server, router)
    if err != nil {
        log.Fatal(err)
    }
}

