### Weather Service API

A Simple weather service api takes lat, long coordinates as query params.
Use openweathermap service, and analyze the response.

# Run code

```
go run cmd/web/main.go
```

# Example request

```
http://localhost:3000/weather?lat=39.607909701969696&lon=-105.2044216193248

- United States
http://localhost:3000/weather?lat=42.536457&lon=-70.985786

- Europe
http://localhost:3000/weather?lat=54.5260&lon=15.2551
```
