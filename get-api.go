package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

type Data struct {
    Base string     `json:"base"`
    Value weather   `json:"main"`
}

type weather struct {
    Temp float32        `json:"temp"`
    Temp_min float32    `json:"temp_min"`
    Temp_max float32    `json:"temp_max"`
    Pressure float32    `json:"pressure"`
    Humidity float32    `json:"humidity"`

}

func main() {
    response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=51.50&lon=-0.12&units=metric&appid=3ed929f54ee60b0da8b99029d4107246")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
    

    var city1 Data
    json.Unmarshal(responseData, &city1)

    fmt.Println(city1.Base)

    fmt.Println(city1.Value)

}
