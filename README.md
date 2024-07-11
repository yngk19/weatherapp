# Weather Forecast API

Сервис предоставляющий API для получения прогнозов погоды по городам.

Используемые технологии:

PostgreSQL (хранилище данных)
Docker (запуск сервиса)
Gin (веб-фреймворк)
pgx (драйвер работы с PostgreSQL)
cron (шедулер задач)
golang-migrate (для миграций)

Сервис написан по Clean Architecture

## Запуск
1. Склонировать репо
2. Создать файл .env в директории проекта и заполнить, также предварительно получить api_key на сайте openweathermap.org и указать в .env. В файле config/local.yaml можно конфигурировать http сервер
3. Выполнить в терминале:
   ```
   ~$ go mod download 
   ~$ make build
   ~$ make run 
   ``` 

## Спецификация API

### GET /cities 

Возвращает список городов, для которых доступен прогноз погоды. Формат ответа такой:

```
{
    "cities": [
        {
            "ID": 1,
            "Name": "Ussurijsk",
            "State": "Primorsky Krai"
        },
     ]
}
```

### GET /cities/:id 

Возвращает полный прогноз погоды для конкретного города. Формат ответа:

```
{
    "status": "success"
    "city_id": 1,
    "country": "RU",
    "lat": 43.7972447,
    "lon": 131.9520752,
    "name": "Ussurijsk",
    "state": "Primorsky Krai",
    "forecasts": [
        {
            "ID": 146,
            "Temperature": 295.57,
            "Date": "2024-07-16T00:00:00Z",
            "DetailInfo": [
	.......
}

```

### GET /cities/:id?date=2024-07-11

Возвращает полный прогноз погоды для конкретного города в конкретный день. Формат ответа:
  

```
{
    "status": "success"
    "city_id": 1,
    "country": "RU",
    "lat": 43.7972447,
    "lon": 131.9520752,
    "name": "Ussurijsk",
    "state": "Primorsky Krai",
    "2024-07-12": {
        "ID": 34,
        "Temperature": 292.33,
        "Date": "2024-07-12T00:00:00Z",
        "DetailInfo": [
            {
                "dt": 1720742400,
                "main": {
                    "temp": 292.33,
       .......
}
```

### GET /cities/:id?date=2024-07-11&time=03:00:00

Возвращает полный прогноз погоды для конкретного города в конкретное время. Формат ответа:

```
{
    "status": "success"
    "city_id": 1,
    "country": "RU",
    "lat": 43.7972447,
    "lon": 131.9520752,
    "name": "Ussurijsk",
    "state": "Primorsky Krai",
    {
    "2024-07-12 12:00:00": {
        "dt": 1720785600,
        "main": {
            "temp": 290.57,
            "feels_like": 290.74,
            "temp_min": 290.57,
            "temp_max": 290.57,	
	......
}
```

### GET /cities/:id/short

Возвращает короткий прогноз погоды для выбранного города. Формат ответа:

```
{
    "status": "success"
    "city_id": 1,
    "country": "RU",
    "lat": 43.7972447,
    "lon": 131.9520752,
    "name": "Ussurijsk",
    "state": "Primorsky Krai",
    "short_forecast": {
        "Country": "RU",
        "Name": "Ussurijsk",
        "Lat": 43.7972447,
        "Lon": 131.9520752,
        "AverageTemperature": 293.7980041503906,
        "Dates": [
            "2024-07-16",
            "2024-07-15",
            "2024-07-14",
            "2024-07-13",
            "2024-07-12"
        ]
    },
}
```




