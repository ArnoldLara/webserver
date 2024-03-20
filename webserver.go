// forms.go
package main

import (
    "html/template"
    "net/http"
	"fmt"

    "io/ioutil"
    "log"
    "os"
    "encoding/json"
)

//Estructura para almacenar los datos que serán mostrados en tabla al usuario
type CitiesInfo struct {
    Success bool
	City1 string
    City2 string
    TempC1 string
    TempC2 string
    PressureC1 string
    PressureC2 string
    HumidityC1 string
    HumidityC2 string
}
//Estructura usada para leer la inforamción de la consulta al API
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

//Obtener coordenadas para consultar API
func get_coord(city string)(float64, float64){
    if city=="Bogota" { 
        lat:=4.6534649
        lon:=-74.0836453 
        return lat,lon
    }else if city=="Madrid" { 
        lat:=40.4167047
        lon:=-3.7035825
        return lat,lon
    }else if city=="Barcelona" { 
        lat:=41.3828939
        lon:=2.1774322
        return lat,lon
    }else if city=="Francia" { 
        lat:=46.2337295
        lon:=5.3538775
        return lat,lon
    }else if city=="Bruselas" { 
        lat:=50.8465573
        lon:=4.351697
        return lat,lon
    }else if city=="Holanda" { 
        lat:=52.55270795
        lon:=4.831577956975391
        return lat,lon
    }

 return 0.0,0.0

}
//Consulta de API mediante coordenadas
//API usada en el proyecto: https://openweathermap.org/
func get_weather(lat float64,lon float64)(Data){
    str_lat := fmt.Sprintf("%f", lat)
    str_lon := fmt.Sprintf("%f", lon)
    //Por mejorar, guarda API key como variable de entorno
    api_key:="3ed929f54ee60b0da8b99029d4107246"
    response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat="+str_lat+"&lon="+str_lon+"&units=metric&appid="+api_key)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
    

    var city Data
    json.Unmarshal(responseData, &city)

    fmt.Println(city.Base)

    //fmt.Println(city.Value)
    return city
}



func main() {
    tmpl := template.Must(template.ParseFiles("./static/index.html"))
    var api_city1 Data
    var api_city2 Data

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }
        cities_info := CitiesInfo{
			Success: true,
            City1: r.FormValue("city1"),
            City2: r.FormValue("city2"),
        }

        //*******************************************************************
        //Obtener coordenadas de ciudad 1
        lat_c1,lon_c1 := get_coord(cities_info.City1)
        //Obtener coordenadas de ciudad 2
        lat_c2,lon_c2 := get_coord(cities_info.City2)


        //*********************************************************************
        //Obtener informacion del clima Ciudad 1
        api_city1=get_weather(lat_c1,lon_c1)
        //Obtener informacion del clima Ciudad 1
        api_city2=get_weather(lat_c2,lon_c2)
        //*********************************************************************

        //*********************************************************************
        //Organizacion información que se va a mostrar al usuario
        cities_info.TempC1=fmt.Sprintf("%.2f", api_city1.Value.Temp)
        cities_info.TempC2=fmt.Sprintf("%.2f", api_city2.Value.Temp)
        cities_info.PressureC1=fmt.Sprintf("%.2f", api_city1.Value.Pressure)
        cities_info.PressureC2=fmt.Sprintf("%.2f", api_city2.Value.Pressure)
        cities_info.HumidityC1=fmt.Sprintf("%.2f", api_city1.Value.Humidity)
        cities_info.HumidityC2=fmt.Sprintf("%.2f", api_city2.Value.Humidity)

        //*********************************************************************
   
        //Ejecutar el template HTML y pasando la info al usuario
        tmpl.Execute(w, cities_info)
    })

    http.ListenAndServe(":8080", nil)
}