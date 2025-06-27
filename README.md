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


## setup format

<pre> ```bash git clone https://github.com/hriday111/weather-backend.git docker build -t weather-backend . docker run -p 8080:8080 -v $(pwd)/data:/app/data weather-backend ``` </pre>
## Usage format
go run main.go 
and then 
http://localhost:8080/<summary/forecast>?lat=<valid latitude>&lon=<valid longitude>&lang=<pl, en> (default is en)

