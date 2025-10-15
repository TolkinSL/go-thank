package main

import (
	// "encoding/json"
	"fmt"
	"time"
)

type Weather struct {
	City string `json:"city"`
	Temperature float64 `json:"temperature"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WeatherStorage struct {
	weathers map[string]*Weather
}

type WeatherStorageOld struct {
	weathers map[string]Weather
}

func main() {
	t := Weather{City: "Tula", Temperature: 23.6, UpdatedAt: time.Now()}	
	m := Weather{City: "Moscow", Temperature: 21.2, UpdatedAt: time.Now()}
	
	ws := WeatherStorage{
		weathers: make(map[string]*Weather),
	}

	ws.weathers["Tula"] = &t
	ws.weathers["Moscow"] = &m

	wso := WeatherStorageOld {
		weathers: make(map[string]Weather),
	}

	wso.weathers["Tula"] = t
	wso.weathers["Moscow"] = m

	fmt.Printf("%#v", wso.weathers["Tula"])
}