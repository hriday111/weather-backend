# Weather Forecast Backend (written in GoLang)
This is a project written as part of an application for internship. 
This repository deals with just the backend part of my project, which is written in GoLang

This project leverages:
- Open-meteo api for weather data
- SQLite for chaching
- Goroutines for background refresh
- high scalablity of the project using a modular structre.
- centralized constant variables stored in internal/config/config.go



## Features
- Feathes 7- day weather forecast at lat-lon
- Calculates estimated solar enegry production per day
- Computes weekly summary (avg pressure[hPa], sun exposure time [hours], min/max temp [celcius], rainy/ non rainy day)\
- Polish/ English support
- Auto background cache refresh


## TODO

- Set up GitHub projects for this project and keep track of issues and todos
- Add a feature so that an api call can be made to this backend by just city name instead of co-ordinates
- Spell check this city name using Levenshtein distance. 

## Directory structre explainations
The main code of the backend is stored in main.go
in internal/ are self created helper libraries that make it easier to update, scale and troubleshoot different elements. 

internal/api is for routing http requests to functions, enabling CORS and then handling requests

internal/config stores configuration constant variables

internal/model is used to define the model of json recieved from open meteo api and model of json request to push/publish for the frontend

internal/service handles and process the conversion of json recieved and to be published

internal/util is for other utilities such as a small translation dictionary
## Docker build /setup format

```git clone https://github.com/hriday111/weather-backend.git``` 

``` cd weather-backend```

```docker build -t weather-backend . ```

```docker run -d -p 8080:8080 -v $(pwd)/data:/app/data weather-backend ```

## Non docker Usage format
go run main.go 

## Api Call structure
http://localhost:8080/<summary/forecast>?lat=<valid latitude>&lon=<valid longitude>&lang=<pl, en> (default is en)
