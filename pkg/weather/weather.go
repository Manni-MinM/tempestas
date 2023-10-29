package weather

type WeatherService interface {
    CityTemp(string) (string, error)
}

