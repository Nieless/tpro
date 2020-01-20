# tpro

### Build and run word-counter-service
```bash

export WORD_COUNTER_SERVICE_HTTP_PORT=8001
go build -o ./bin/wordcounter -i ./cmd/wordcounter/ && ./bin/wordcounter

GetWordCounter
* serves the html file
* URL : /crawling
* METHID: GET

PostWordCounter
* counts the word and return map
* URL : /crawling
* METHOD: POST

```


### Build and run datetime-updater-service
```bash

export DATE_TIME_SERVICE_HTTP_PORT=8002
go build -o ./bin/datetime -i ./cmd/datetime/ && ./bin/datetime


GetDateTime
* serves the html file
* URL : /datetime
* METHID: GET

PostDateTime
* updates the datetime and return it
* URL : /datetime
* METHOD: POST


```


### Build and run excel-column-finder-service
```bash

export WORD_COUNTER_SERVICE_HTTP_PORT=8003
go build -o ./bin/excelcolumnfinder -i ./cmd/excelcolumnfinder/ && ./bin/excelcolumnfinder

GetExcelColumnFinder
* GetExcelColumnFinder serves the html file
* URL : /ecf
* METHOD: GET


```

### Build and run perlin-image-noise-service
```bash

export PERLIN_IMAGE_NOISE_CALC_SERVICE_HTTP_PORT=8004
go build -o ./bin/perlinimage -i ./cmd/perlinimage/ && ./bin/perlinimage

GetPerlinImageNoiseCalcHandler
* serves the html file
* URL : /noisecalc
* METHID: GET

PerlinImageNoiseCalcHandler
* returns array of noises of image coordinates
* URL : /noisecalc
* METHOD: POST

```

### Build and run mongo-crud-service
```bash

export MONGO_CONNECT_STRING="mongodb://localhost/test"
export MONGO_CRUD_SERVICE_HTTP_PORT=8005
go build -o ./bin/http -i ./cmd/http/ && ./bin/http


GetUsers
* serves the html file
* URL : /users
* METHID: GET
* REQUEST BODY : NONE
* RESPOSE BODY : 
     [
        {
          "id": "5e258b8ee206a569f010c103",
          "name": "neel",
          "age": 26
        },
        {...}
     ]


CreateUser
* returns array of noises of image coordinates
* URL : /users
* METHOD: POST
* REQUEST BODY :
      {
        "name": "neel",
        "age": 26
      }

* RESPOSE BODY : 
      {
        "id": "5e258b8ee206a569f010c103",
        "name": "neel",
        "age": 26
      }


UpdateUser
* serves the html file
* URL : /users/{id}
* METHID: PUT
* REQUEST BODY :
      {
        "name": "Neel",
        "age": 25
      }

* RESPOSE BODY : 
      {
        "id": "5e258b8ee206a569f010c103",
        "name": "Neel",
        "age": 25
      }



DeleteUser
* returns array of noises of image coordinates
* URL : /users/{id}
* METHOD: DELETE
* REQUEST BODY : NONE
* RESPOSE BODY : "deleted"

GetUser
* serves the html file
* URL : /users/{id}
* METHID: GET
* REQUEST BODY : NONE
* RESPOSE BODY : 
      {
        "id": "5e258b8ee206a569f010c103",
        "name": "neel",
        "age": 26
      }

```

