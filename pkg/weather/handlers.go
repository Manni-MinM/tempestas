package weather

import (
    "io"
    "net/http"

    "github.com/buger/jsonparser"
)

func CityTempHandler(ws WeatherService) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        if req.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        body, err := io.ReadAll(req.Body)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            panic(err)
        }

        city, err := jsonparser.GetString(body, "city")
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            panic(err)
        }

        cityTemp, err := ws.CityTemp(city)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            panic(err)
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(cityTemp))
    })
}

