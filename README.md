
# Welcome To Go Shop Service 
Go Shop Service is a simple API for some actions, imagine you have a big shop, customers come and take sales, meanwhile
some purchases got delays on delivery system, we have designed a small system that customers can report a delay, 
an agent can handle a delay request and assign new time to it based on an external api,
finally you can make a report request to see vendors how has delays based on minutes within last 7 days

## Software Architecture
This is an attempt to implement a clean architecture, and some other design patterns such as adapter, singleton, composition in combination
## Requirements/Dependencies
- Docker
- Docker-compose
- golang:1.18-alpine docker image
- mysql docker image
##  Getting Started
we have simple makefile in root of our project, currently makefile works on Windows, other systems will be added soon
`make clean` 
will do everything for you to come up and running

## API Request
|      URL          |HTTP Method|Discription|
|----------------|-------------------------------|-----------------------------|
|`/v1/transfers`|`POST`            |`Make card to card transaction with Given json data   `         |
|`/v1/report`|`GET`            |`Gets you the Report of last transactions within 10 minutes last`            |
|`/v1/health`|`ALL METHODS`|`Checks App is up and Runnig`|


## Test Endpoints API using Curl
you can test APIs with curl or post man as follows
 ##### Request
```
curl -i --request POST 'http://localhost:8000/v1/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from_card_number": "5022291302421266",
    "to_card_number": "5041721005782710"
    "transaction_amount": 123456
}
```
##### Response
```
{
    "ReceiverMsg": {
        "Msg": "\n        واریز\n    5041721005782710\n        به مبلغ\n    123456\n        مانده\n    823456",
        "To": "+989033934262"
    },
    "SenderMsg": {
        "Msg": "\n        برداشت از\n    5022291302421266\n        به مبلغ\n    123456\n        مانده\n    26044",
        "To": "+989123993699"
    },
    "status": {
        "Message": "Done"
    }
}
```

it also supports Persian  Or Arabic Numbers like below but of course its and error because of wrong card number pattern !!!

 ##### Report Request
 this API will give you the time each vendor has delay time based on minutes
```curl -i --request GET 'http://localhost:8000/v1/report' ```
you can give optional parameter as t like
```curl -i --request GET 'http://localhost:8000/v1/report?t=1000' ```
it will change the default last 10 minutes time of query to 1000
 ##### Report Response 
 ```
 [

{
"CustomerID": "1003",
"TransactionId": "1006",
"CardIdFrom": "4003",
"CardIdTo": "4007",
"Amount": 123456,
"TransactionType": 0,
"TransactionTime": "2023-04-02 07:36:54",
"Index": 1
},

{
"CustomerID": "1004",
"TransactionId": "1006",
"CardIdFrom": "4003",
"CardIdTo": "4007",
"Amount": 123456,
"TransactionType": 0,
"TransactionTime": "2023-04-02 07:36:54",
"Index": 1
}
]
```

## Code Status
still fixing bugs for v1

## Author
Copyright © 2020 [Omid-h70](https://github.com/omid-h70). This project is MIT licensed. its free to EveryOne,

Thanks
