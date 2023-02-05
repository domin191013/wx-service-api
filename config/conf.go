package config

const APIID = "24572ec02742e943b625df83a904b2ca"

const GetWeatherUrl = "https://api.openweathermap.org/data/2.5/onecall?lat=%v&lon=%v&exclude=minutely,hourly,daily&units=imperial&appid=%v"

const Hot = 80.0

const Cold = 20.0

var Temps = [3]string{"Cold", "Comfortable", "Hot"}
