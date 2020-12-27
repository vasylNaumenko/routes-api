# Routes API

Microservice demo of REST API that takes the source and a list of destinations
an returns a list of routes between source and each destination. Both source and
destination are defined as a pair of latitude and longitude. The returned list of routes
is sorted by driving time and distance (if time is equal).

list of make available targets
```
make help
```

Running:

if you have GO installed:
``` 
make
```

using docker compose:
```
make run-compose
```

#
#### Request example:
```
GET http://localhost:3001/routes?rc=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219&dst=13.428855,52.523239
```
Success:
```
200 OK application/json; charset=utf-8
{
    "source": "13.388798,52.517033",
    "routes": [
        {
            "destination": "13.428554,52.523239",
            "duration": 117.6,
            "distance": 950.3
        },
        {
            "destination": "13.388798,52.517033",
            "duration": 251.5,
            "distance": 1884.8
        },
        {
            "destination": "13.397631,52.529430",
            "duration": 372.2,
            "distance": 2946.1
        }
    ]
}
```

error:
```
400 Bad Request application/json; charset=utf-8
{
    "status": "error",
    "message": "{error message}"
}

where is {error message} is one of:
- no src parameter
- at least one dst parameter must exists
```
```
500 Internal Server Error application/json; charset=utf-8
{
    "status": "error",
    "message": "{error message}"
}

where is {error message} is an error from third-party router service
```
#####health check
```
GET http://localhost:3001/status
```
should respond http status ``200`` 
and valid json
``
{
"Status": "ok"
}
``