package weather

import (
	"fmt"
	"io"
	"net/http"

	"tempestas/pkg/database"
)

type httpWeather struct{
    db      database.Database
    apiKey  string
}

func NewHTTP(db database.Database, apiKey string) WeatherService {
    return &httpWeather{db, apiKey}
}

func (w *httpWeather) CityTemp(city string) (string, error) {
    cityTemp, err := w.db.Get(city)
    if err == nil {
        return cityTemp, nil
    }

    url := fmt.Sprintf("https://api.api-ninjas.com/v1/weather?city=%v", city)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    req.Header.Set("X-Api-Key", w.apiKey)

    c := &http.Client{}
    resp, err := c.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    cityTemp = string(body)

    _, err = w.db.Set(city, cityTemp)
    if err != nil {
        return "", err
    }

    return cityTemp, nil
}

