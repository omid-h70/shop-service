
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
| URL                |HTTP Method| Description                                                       |
|--------------------|-------------------------------|-------------------------------------------------------------------|
| `/v1/dealy_report` |`POST`            | `Make a Delay report Based On vendor_id and order_id   `          |

every order or test orders are submitted by default expiration time of 1 (one) minute, if you quickly send delay report request the 
response will be as below

Request:
```
{
	"vendor_id" : "1001",
	"order_id" : "2001"
}
```
Response:
```
{
    "message": "Delay Time Report Is Invalid - Try After Delivery Time Reached",
    "result": {
        "OrderId": 2001,
        "VendorId": 1001,
        "CreatedAt": "2023-07-25 12:08:29",
        "DeliveryTime": "2023-07-25 12:09:29",
        "OrderStatus": "1"
    }
}
```
And if Successful You can get below Response, you can get report submission count as `ReportCount`

```
{
    "message": "Done",
    "result": {
        "DelayOrderId": 5001,
        "OrderId": 2001,
        "VendorId": 1001,
        "AgentId": 0,
        "CreatedAt": "2023-07-25 12:14:07",
        "UpdatedAt": "2023-07-25 12:14:07",
        "DelayReportStatus": "",
        "ReportCount": 1
    }
}
```


| URL             |HTTP Method| Description                                    |
|-----------------|-------------------------------|------------------------------------------------|
| `/v1/set_agent` |`POST`            | `Sets a Free Agent to an Open Delay Report   ` |

it will assign an order based on timestamp sorting, and it actually acts like queue
first report that got submitted, gets assigned first

Request:
```
{
	"agent_id" : "4001",
}
```
Response:
```
{
    "message": "Done",
    "result": {
        "DelayOrderId": 5001,
        "OrderId": 2001,
        "VendorId": 1001,
        "AgentId": 4001,
        "ReportCount": 1,
        "CreatedAt": "2023-07-25 12:14:07",
        "UpdatedAt": "2023-07-25 12:14:07"
    }
}
```

| URL             |HTTP Method| Description                                 |
|-----------------|-------------------------------|---------------------------------------------|
| `/v1/v1/handle_delayed_order` |`POST`            | `Handle a Delay Request And Set Agent Free` |

Handle a Delay Report And Sets Image Free

| URL             | HTTP Method | Description                                     |
|-----------------|-------------|-------------------------------------------------|
| `/v1/get_all_delay_reports` | `GET`       | `Gets Delay Report Of a Vendor Based On Minute` |

Gets All Delay Report Of a Vendor Based On Minute, within 7 Days

Response
```
{
    "message": "Done",
    "result": [
        {
            "ID": "1001",
            "DelayTime": "9"
        }
    ]
}
```

| URL          |HTTP Method| Description                                 |
|--------------|-------------------------------|---------------------------------------------|
| `/v1/health` |`ALL METHODS`            | `Checks App is up and Runnig` |

Response
```
{
    "message": "Done",
    "result": [
        {
            "ID": "1001",
            "DelayTime": "9"
        }
    ]
}
```

## Test Endpoints API using Curl
you can test APIs with curl or postman 


## Code Status
still fixing bugs for v1

## Author
Copyright © 2020 [Omid-h70](https://github.com/omid-h70). This project is MIT licensed. its free to EveryOne,

Thanks
