# XM Rest Api

> This is an interview showcase work and not intended for production usage

## Overview
XM Golang Exercise - v21.0.0

## Config
```shell
+----------------------------------+----------+----------------+-------------------------+
|               ENV                |   Type   |    Default     |       Description       |
+----------------------------------+----------+----------------+-------------------------+
| PORT                             | int      | 9090           | Service Port            |
| METRICS_PORT                     | int      | 7000           | Metrics Port            |
| KAFKA_PRODUCER_BOOTSTRAP_SERVERS | []string | localhost:9092 | Kafka Bootstrap Servers |
| DATABASE_URL                     | string   | root:@/test    | Database Connect Url    |
+----------------------------------+----------+----------------+-------------------------+
```
## Quick start
Build the binary
```shell
go build -ldflags="-s -w" -o xm-rest-api main.go
```
Run the Binary
```shell
./xm-rest-api
```

## Testing
Start Services on docker-compose, This might take 20-30min for \
the first time until you pull all the required dependencies.
```shell
docker-compose up
```

Create Company
```shell
curl --location --request PUT 'http://localhost:9090/company' \
--header 'X-FORWARDED-FOR: 31.153.207.255,8.8.8.8' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"ABC2",
    "code":"123123",
    "country":"sri lanka",
    "website":"",
    "phone":123123
}'
```

Update Company
```shell
curl --location --request POST 'http://localhost:9090/company' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"ABC2",
    "code":"123123",
    "country":"sri lanka2",
    "website":"234234",
    "phone":123123
}'
```

Delete Company
```shell
curl --location --request DELETE 'http://localhost:9090/company?code=123123' \
--header 'X-FORWARDED-FOR: 31.153.207.255'
```

Get Company(s)
```shell
curl --location --request GET 'http://localhost:9090/company?code=123123&name=ABC2'
```

## Improvements
1. Add support to Multiple Filters on the same type 
```shell
ex:- 
  GET /company?code=123123&code=454545
```
2. Data Field Validation 
3. Multiple Source Data write failure handling
4. Topic Automatic creation if not exist
5. Database data indexing
