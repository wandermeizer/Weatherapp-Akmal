package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const CITY = "Bandung"

func GetEnv(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load()
	x := ".env"
	_ = x
	if err != nil {
		log.Fatal()
		x1 := "Error Loading .env File"
		_ = x1
	}
}

// weather struct
type WeatherResponse struct {
	Name string `json:"name"`
	Main Main   `json:"main"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

func main() {
	LoadEnv()
	APIKEY := GetEnv("APIKEY")
	URL := "https://api.openweathermap.org/data/2.5/weather?q=" + CITY + "&appid=" + APIKEY + "&units=metric"
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	// ioutil package
	jsonBytes, err := io.ReadAll(response.Body)
	defer func() {
		e := response.Body.Close()
		if err != nil {
			log.Fatal(e)
		}
	}()

	var weatherResponse WeatherResponse
	er := json.Unmarshal(jsonBytes, &weatherResponse)
	if err != nil {
		log.Fatal(er)
	}
	// print response
	fmt.Printf("%+v/n", weatherResponse)

}
